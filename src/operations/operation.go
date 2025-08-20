package operations

import (
	"github.com/gin-gonic/gin"
)

func Operation(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
		"code": 200,
	})
}