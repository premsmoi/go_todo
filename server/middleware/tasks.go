package middleware

import (
	"Generalkhun/go-todo-server/models"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CORSMiddleware (Cross-Origin Resource Sharing) middleware that used to handle Response Header
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-HTTP-Method-Override,Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Accept,accept, origin, Cache-Control, X-Requested-With, set-cookie")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,UPDATE,OPTIONS")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// GetAllTask get all the task route
func GetAllTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		// use username that pas through previous authentication process
		u, exist := c.Get("contextUsername")
		username := u.(string)

		if !exist {
			log.Fatal(errors.New("Context do not contains username, there are some problems"))
		}
		cur, err := connectTodotasks(username, IntiateMongoConn())
		if err != nil {
			log.Fatal(err)
		}

		var results []primitive.M
		for cur.Next(context.Background()) {
			var result bson.M
			e := cur.Decode(&result)
			if e != nil {
				log.Fatal(e)
			}
			results = append(results, result)

		}

		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}

		cur.Close(context.Background())
		werr := json.NewEncoder(c.Writer).Encode(results)
		if werr != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
		c.JSON(http.StatusOK, "?"+username)

	}

}

// CreateTask create task route
func CreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var task models.ToDoList
		_ = json.NewDecoder(c.Request.Body).Decode(&task)
		// insert Username gather from context
		u, exist := c.Get("contextUsername")
		username := u.(string)

		if !exist {
			log.Fatal(errors.New("Context do not contains username, there are some problems"))
		}
		task.Username = username
		insertOneTask(task)
		w := json.NewEncoder(c.Writer).Encode(task)
		c.JSON(http.StatusOK, w)
	}
}

// Insert one task in the DB
func insertOneTask(task models.ToDoList) {

	collection := IntiateMongoConn().Database(dbName).Collection("todoTasks")
	insertResult, err := collection.InsertOne(context.Background(), task)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a Single Record ", insertResult.InsertedID.(primitive.ObjectID).Hex())
}

// UndoTask undo the complete task route
func UndoTask() gin.HandlerFunc {

	return func(c *gin.Context) {
		// params := mux.Vars(r)
		// undoTask(params["id"])
		undoTask(c.Param("id"))
		json.NewEncoder(c.Writer).Encode(c.Param("id"))

	}
}

//task undo method, update task's status to false
func undoTask(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": false}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("modified count: ", result.ModifiedCount)
}

// DeleteTask delete one task route
func DeleteTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		deleteOneTask(c.Param("id"))
		json.NewEncoder(c.Writer).Encode(c.Param("id"))
		// json.NewEncoder(w).Encode("Task not found")

	}
}

// delete one task from the DB, delete by ID
func deleteOneTask(task string) {
	fmt.Println(task)
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	d, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
}

// DeleteAllTask delete all tasks route
func DeleteAllTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		count := deleteAllTask()
		json.NewEncoder(c.Writer).Encode(count)
	}

}

// delete all the tasks from the DB
func deleteAllTask() int64 {
	d, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted Document", d.DeletedCount)
	return d.DeletedCount
}

// Welcome function is a handler funciton to /welcome route
func Welcome() gin.HandlerFunc {
	return func(c *gin.Context) {

		// We can obtain the session token from the requests cookies, which come with every request
		cookie, err := c.Request.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				c.Writer.WriteHeader(http.StatusUnauthorized)
				return
			}
			c.Writer.WriteHeader(http.StatusInternalServerError)
			return

		}

		// Get the JWT string from the cookie
		tknStr := cookie.Value

		// Initialize a new instance of `Claims`
		claims := &models.Claims{}

		// Parse the JWT string and store the result in `claims`.
		// Note that we are passing the key in this method as well. This method will return an error
		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
		// or if the signature does not match
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return models.JwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.Writer.WriteHeader(http.StatusUnauthorized)
				return
			}
			c.Writer.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			c.Writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Finally, return the welcome message to the user, along with their
		// username given in the token
		c.Writer.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))

	}

}
