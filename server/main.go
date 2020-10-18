package main

import (
	"./middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// create router
	router := gin.Default()

	// CORS
	router.Use(middleware.CORSMiddleware())

	//Login-router (coming soon)

	//router.Group("/test")
	//routerLogin := router.GroupRouterLongin()

	//task-router
	routerTask := router.Group("/task")
	{
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
