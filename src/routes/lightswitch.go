package routes

import (
	"github.com/gin-gonic/gin"
)

func LightSwitch(c *gin.Context) {
	c.JSON(200, gin.H{
		"serviceInstanceId": "fortnite",
		"status": "UP",
		"message": "Fortnite is online",
		"maintenanceUri": nil,
		"overrideCatalogIds": []string{"a7f138b2e51945ffbfdacc1af0541053"},
		"allowedActions": []string{},
		"banned": false,
		"launcherInfoDTO": gin.H{
		"appName": "Fortnite",
		"catalogItemId": "4fe75bbc5a674f4f9b356b5c90567da5",
		"namespace": "fn",
		},
	})
}

func LightSwitchBulk(c *gin.Context) {
	c.JSON(200, []gin.H{
		{
			"serviceInstanceId": "kairos",
			"status": "DOWN",
			"message": "Farewell. Thank you for all the memories.",
			"maintenanceUri": nil,
			"overrideCatalogIds": []string{"a7f138b2e51945ffbfdacc1af0541053"},
			"allowedActions": []any{},
			"banned": false,
			"launcherInfoDTO": nil,
		  },
		  {
			"serviceInstanceId": "fortnite",
			"status": "UP",
			"message": "Fortnite is online",
			"maintenanceUri": nil,
			"overrideCatalogIds": []string{"a7f138b2e51945ffbfdacc1af0541053"},
			"allowedActions": []any{},
			"banned": false,
			"launcherInfoDTO": gin.H{
			  "appName": "Fortnite",
			  "catalogItemId": "4fe75bbc5a674f4f9b356b5c90567da5",
			  "namespace": "fn",
			},
		  },
	})
}