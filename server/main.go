package main

import (
	"Generalkhun/go-todo-server/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// create router
	router := gin.Default()

	// CORS
	router.Use(middleware.CORSMiddleware())

	//home-router
	router.GET("/", middleware.PreSignin())
	router.POST("/register", middleware.Register())

	//auth-router
	routerAuth := router.Group("/auth")
	{

		routerAuth.POST("/signin", middleware.Signin())

		routerAuth.GET("/refresh", middleware.Refresh())
		routerAuth.GET("/logout", middleware.Logout())
	}

	//task-router
	routerTask := router.Group("/task")
	routerTask.Use(middleware.AuthRequired())
	{
		routerTask.GET("/welcome", middleware.Welcome())
		routerTask.GET("/", middleware.GetAllTask())
		routerTask.POST("/", middleware.CreateTask())
		routerTask.PUT("/:id", middleware.UndoTask())
		routerTask.DELETE("/:id", middleware.DeleteTask())

	}
	routerTask2 := router.Group("/task2")
	{
		routerTask2.DELETE("/deleteAllTask", middleware.DeleteAllTask())
	}

	//Run the server
	router.Run(":8080")

}
