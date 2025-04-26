package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *Api) BlockHost(c *gin.Context) {
	host := blockHostRequest{}
	if err := c.BindJSON(&host); err != nil {
		c.JSON(400, gin.H{"data": "error"})
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	blockedHost, err := api.BlockHostUseCase.Block(host.Host)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if blockedHost == nil {
		c.JSON(400, gin.H{"data": "Host not found"})
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Header("Access-Control-Allow-Origin", "*") //There is a vuln here, that's only for testing purpose.
	c.Header("Access-Control-Allow-Methods", "POST")
	c.JSON(http.StatusOK, gin.H{"message": "Host " + host.Host + "has been blocked"})
}
