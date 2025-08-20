package routes

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func HandleRegisterKeyChain(c *gin.Context) {
	data, err := os.ReadFile("./static/keychain.json")
	if err != nil {
		fmt.Println("error!:", err)
		return
	}

	var keychain interface{}
	err = json.Unmarshal(data, &keychain)
	if err != nil {
		fmt.Println("error!:", err)
		return
	}

	c.JSON(200, keychain)
}