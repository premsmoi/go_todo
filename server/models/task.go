package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//ToDoList is struct of todolist read form POST request
type ToDoList struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Task     string             `json:"task,omitempty"`
	Status   bool               `json:"status,omitempty"`
	Username string
}
