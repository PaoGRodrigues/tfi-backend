package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *Api) StoreHosts(c *gin.Context) {
	err := api.StoreHostsUseCase.StoreHosts()
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"data": "error"})
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Header("Access-Control-Allow-Origin", "*") //There is a vuln here, that's only for testing purpose.
	c.Header("Access-Control-Allow-Methods", "POST")
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}
