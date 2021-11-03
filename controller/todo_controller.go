package controller

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"todo/common"
	"todo/repositories"
	"todo/service"
)

var todoService = service.NewTodoService(repositories.NewTodoRepository(common.NewMysqlConn()))

func Index(c *gin.Context) {
	ownerID, _ := strconv.Atoi(c.Param("owner_id"))
	todos := todoService.FindByOwnerID(int64(ownerID))
	c.JSON(200, gin.H{
		"result": todos,
	})
}
