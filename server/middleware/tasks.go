package middleware

import (
	"Generalkhun/go-todo-server/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CORSMiddleware (Cross-Origin Resource Sharing) middleware that used to handle Response Header
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

// GetAllTask get all the task route
func GetAllTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		cur, err := collection.Find(context.Background(), bson.D{{}})
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
			// fmt.Println("cur..>", cur, "result", reflect.TypeOf(result), reflect.TypeOf(result["_id"]))
			results = append(results, result)

		}

		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}

		cur.Close(context.Background())
		w := json.NewEncoder(c.Writer).Encode(results)
		c.JSON(http.StatusOK, w)

	}

}

// CreateTask create task route
func CreateTask() gin.HandlerFunc {
	return func(c *gin.Context) {
		var task models.ToDoList
		_ = json.NewDecoder(c.Request.Body).Decode(&task)
		insertOneTask(task)
		w := json.NewEncoder(c.Writer).Encode(task)
		c.JSON(http.StatusOK, w)
	}
}

// Insert one task in the DB
func insertOneTask(task models.ToDoList) {
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
