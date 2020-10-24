package middleware

import (
	"Generalkhun/go-todo-server/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

//PreSignin is handler function for route /
func PreSignin() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "This is Todo app, please login to start using the app")
	}
}

//Register is handler function for route /register
func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userdb models.UsersDB
		client := IntiateMongoConn() //init mongoDB connection
		_ = json.NewDecoder(c.Request.Body).Decode(&userdb)
		// check if username is alreadyused
		alreadyUsed, err := checkAlreadyused(&userdb, client)
		if err != nil {
			if alreadyUsed {
				c.String(http.StatusNotAcceptable, err.Error())

			} else {
				log.Fatal(err)
				c.String(http.StatusInternalServerError, err.Error())

			}
		}
		if !alreadyUsed {
			// add this user to database if username is not ever used
			addOneUser(&userdb, client)
			c.String(http.StatusOK, "Successfully registered, your username:"+userdb.Username+", your password:"+userdb.Password)

		}

	}

}

func checkAlreadyused(userdb *models.UsersDB, client *mongo.Client) (bool, error) {
	var result bson.M
	collection := client.Database("todo").Collection("UsersDB")
	condition := primitive.E{Key: "username", Value: userdb.Username}
	err := collection.FindOne(context.TODO(), bson.D{condition}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("Err no  doc ")
			return false, nil
		}
		log.Fatal(err)
		fmt.Println("Err other ")
		return false, err
	}
	fmt.Println("found doc")
	return true, errors.New("This Username is already used, please try again")
}

func addOneUser(userdb *models.UsersDB, client *mongo.Client) {
	collection := client.Database("todo").Collection("UsersDB")
	insertResult, err := collection.InsertOne(context.Background(), userdb)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a Single Record ", insertResult.InsertedID.(primitive.ObjectID).Hex())
}
