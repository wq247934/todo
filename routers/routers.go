package routers

import (
	"github.com/gin-gonic/gin"
	"todo/controller"
)

func SetRouters() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("v1")
	{
		v1.GET("/todos/:owner_id", controller.Index)
	}
	return r
}
