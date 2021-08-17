package routers

import (
	"Data_Bank/fabric-manager/server/routers/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/createOrg", api.CreateOrg)
	r.POST("/createConsortium", api.CreateConsortium)
	r.POST("/startConsortium", api.StartConsortium)
	r.POST("/createChannel", api.CreateChannel)
	return r
}
