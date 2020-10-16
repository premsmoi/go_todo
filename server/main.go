package main

import "github.com/gin-gonic/gin"

func main() {
	// create router
	router := gin.Default()
	//Login-router (coming soon)

	//router.Group("/test")
	//routerLogin := router.GroupRouterLongin()

	//task-router
	routerTask := router.Group("/v1")
	{
		routerTask.GET("/api/task", middleware.GetAllTask)
		routerTask.POST("/api/task", middleware.CreateTask)
		routerTask.PUT("/api/undoTask/", middleware.UndoTask)
		routerTask.DELETE("/api/deleteTask", middleware.DeleteTask)
		routerTask.DELETE("/api/deleteAllTask", middleware.DeleteAllTask)
	}

	//Run the server
	router.Run(":8080")

}
