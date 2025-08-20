package routes

import (
	"github.com/gin-gonic/gin"
)

func WaitingroomAPI(c  *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
		"code": 200,
	})
}