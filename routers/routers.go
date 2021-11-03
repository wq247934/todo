package routers

import "github.com/gin-gonic/gin"

func SetRouters() *gin.Engine {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, "HHHHHHH")

	})
	return r
}
