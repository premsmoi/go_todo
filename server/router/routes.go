package router

import (
	"github.com/gin-gonic/gin"
)

// func *Engine GroupRouterLongin() {
	
// }

func *gin.Engine GroupRouterTask( *gin.RouterGroup) {
	router.Group("/v1"){
		v1.POST("api/")
	}

	

