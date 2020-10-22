//DB connect
package middleware

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://127.0.0.1:27017/?compressors=disabled&gssapiServiceName=mongodb"
const dbName = "todo"

//const collectionName = "todoTasks"

var collection *mongo.Collection

//IntiateMongoConn to database when app is start
func IntiateMongoConn() *mongo.Client {
	// Set the client options, specified database location by using ApplyURI
	clientOptions := options.Client().ApplyURI(connectionString)

	//Connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

func connectTodotasks(username, userpass_id string, client *mongo.Client) (string, err) {

	username, err = client.Database(dbName).Collection("todoTasks").FindOne(context.TODO(), bson.D{{"_id", userpass_id}})

	return collection

}

func getUserPass(username string) {
	// Set the client options, specified database location by using ApplyURI
	clientOptions := options.Client().ApplyURI(connectionString)

	//Connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection("UsersDB")

}
