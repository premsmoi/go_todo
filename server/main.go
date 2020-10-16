package main

import "github.com/gin-gonic/gin"

func main() {
	// create router
	router := gin.Default()
	//Login-router (coming soon)

	router.Group("/test")
	//routerLogin := router.GroupRouterLongin()

	//task-router
	// routerTask := router.Group("/task"){
	// 	task.POST()
	// }

	//task
	router.Run(":8080")

}
