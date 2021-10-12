package main

import "github.com/gin-gonic/gin"

type Api struct {
	*gin.Engine
}

func (api *Api) MapURLToPing() {
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}
