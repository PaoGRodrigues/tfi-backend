package main

import (
	"github.com/PaoGRodrigues/tfi-backend/app/api"
	"github.com/gin-gonic/gin"
)

func main() {

	iniatializer := api.Initializer{}

	api := &api.Api{
		DeviceHandler: iniatializer.InitializeDeviceDependencies(),
		Engine:        gin.Default(),
	}

	api.MapURLToPing()
	api.MapGetDevicesURL()

	api.Run(":8080")
}
