# Goreact todoapp
Toy Todo app using Go, mongoDB where a logged in user can access their todo list and use simple react app to display user interface

## server
Use Go with Gin (great web framework built on Golang:https://github.com/gin-gonic/gin) to be serverside. After recieved requests from client, server will perform tasks 
 1. create account => add username and password to mongoDB
 2. login => check username and password via matching JWT token
 3. create/delete/undo tasks => perform CRUD operations on MongoDB
 4. refresh token for users that interact with app and their token is not expired yet

## client 
Use react to connect API from server and display user's todolist data. 
After login, User can interact with their todolist including: Add a todo item, toggle a todo item as finished (or unfinished), Delete a todo item and Delete whole todolist 
