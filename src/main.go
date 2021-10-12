package main

import "github.com/gin-gonic/gin"

func main() {
	api := &Api{
		Engine: gin.Default(),
	}

	api.MapURLToPing()

	api.Run(":8080")
}
