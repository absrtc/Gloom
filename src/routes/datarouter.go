package routes

import (
	"github.com/gin-gonic/gin"
)

type DataRouter struct {
	Status string `json:"status"`
	Code int `json:"number"`
}

func HandleDataRouter(c *gin.Context) {
	c.JSON(200, DataRouter{
		Status: "OK",
		Code: 200,
	})
}