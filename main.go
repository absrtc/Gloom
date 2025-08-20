package main

import (
	"fmt"
	"os"

	"github.com/absrtc/gloom/src/operations"
	"github.com/absrtc/gloom/src/routes"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Gloom struct {
		Port int `yaml:"port"`
	} `yaml:"gloom"`
}

func main() {
	var config Config
	gin := gin.Default()

	data, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Println("error while reading config:", err)
		return
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Println("error while reading config:", err)
		return
	}
	
	port := fmt.Sprintf("%v", config.Gloom.Port)
	gin.Run(port)
}

func HandleSetupServerRoutes(r *gin.Engine) {
	fn := r.Group("/fortnite/api")
	{
		fn.GET("/storefront/v2/keychain", routes.HandleRegisterKeyChain)
		fn.GET("/storefront/v2/catalog", routes.CatalogAPI)
		fn.GET("/calendar/v1/timeline", routes.HandleTimeline)
		fn.GET("/version", routes.HandleVersion)
		fn.GET("/versioncheck/*any", routes.HandleNoUpdate)
		fn.GET("/v2/versioncheck/*any", routes.HandleNoUpdate)

		fn.POST("/game/v2/profile/:accountId/client/:operation", operations.Operation)
	}

	eula := r.Group("/eulatracking")
	{
		eula.GET("/shared/agreements/fn", routes.SharedAgreements)
		eula.GET("/public/agreements/fn/account/:accountId", routes.SharedAgreements_Account)
	}

	datarouter := r.Group("/datarouter/api")
	{
		datarouter.POST("/v1/public/data", routes.HandleDataRouter)
	}

	content := r.Group("/content/api")
	{
		content.GET("/pages/:key", routes.HandleContentPages)
	}

	Waitingroom := r.Group("/waitingroom/api")
	{
		Waitingroom.GET("/waitingroom", routes.WaitingroomAPI)
	}

	LightSwitch := r.Group("/lightswitch/api/")
	{
		LightSwitch.GET("/service/:serviceId/status", routes.LightSwitch)
		LightSwitch.GET("/service/bulk/status", routes.LightSwitchBulk)
	}

}