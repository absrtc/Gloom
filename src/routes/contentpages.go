package routes

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func HandleContentPages(c *gin.Context) {
	data, err := os.ReadFile("./static/c.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to read content pages file: " + err.Error(),
		})
		return
	}

	var contentPages interface{}
	if err := json.Unmarshal(data, &contentPages); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to parse JSON: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, contentPages)
}
