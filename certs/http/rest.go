package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"Data_Bank/fabric-manager/common/utils"
)

func StartHTTP() *gin.Engine {
	r := gin.Default()
	r.GET("/v1/download/client/cert", DownloadClient)
	return r
}

const DownloadTar = "/opt/cert.tar.gz"

// DownloadClient download implementation
func DownloadClient(c *gin.Context) {
	if utils.Exists(DownloadTar) {
		c.Writer.WriteHeader(http.StatusOK)
		c.Header("Content-Disposition", "attachment; filename=cert.tar.gz")
		c.Header("Content-Type", "application/tar+gzip")
		// c.Header("Accept-Length", fmt.Sprintf("%d", len(content)))
		c.File(DownloadTar)
		return
	}
	c.JSON(500, "{'code':404}")
}
