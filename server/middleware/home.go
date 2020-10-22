package middleware

import (
	"Generalkhun/go-todo-server/models"
	"context"
	"encoding/json"
	"log"
	"net/http"

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

		_ = json.NewDecoder(c.Request.Body).Decode(&userdb)
		// check if username is alreadyused
		err := checkAlreadyused(&userdb)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusPermanentRedirect, gin.H{"message": "This Username is already used, please try again"})
				return
			}
		}
		// add this user to database
		addOneUser(userdb)
		c.Redirect(http.StatusPermanentRedirect, "/")

	}

}

func checkAlreadyused(userdb *models.UsersDB) error {
	var result bson.M
	err := collection.FindOne(context.TODO(), bson.D{{"Username", userdb.Username}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil

}
