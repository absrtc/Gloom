package routes

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func CatalogAPI(c *gin.Context) {
	data, err := os.ReadFile("./static/catalog.json")
	if err != nil {
		fmt.Println("error while reading catalog.json:", err)
		return
	}

	var catalog interface{}
	err = json.Unmarshal(data, &catalog)
	if err != nil {
		fmt.Println("error while unmarshaling catalog.json:", err)
		return
	}

	c.JSON(200, catalog)

}