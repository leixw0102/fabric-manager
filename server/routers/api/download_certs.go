package api

import (
	"Data_Bank/fabric-manager/common/connector"
	"Data_Bank/fabric-manager/common/utils"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func DownloadCmd(c *gin.Context) {
	type certsInfo struct {
		CertsServer string `json:"certs_server"`
	}

	var ci certsInfo
	if err := c.BindJSON(&ci); err != nil {
		logrus.Errorf("Fail to parse json object, error:%+v", err)
		return
	}
	key := fmt.Sprintf(utils.CertDownloadPath, utils.CertDownload)
	err := connector.ETCD.Put(key, ci.CertsServer, 4*time.Second)
	if err != nil {
		logrus.Errorf("etcd put failed with error: %v", err)
		c.JSON(200, gin.H{
			"message": "cmd certs push to etcd failed",
			"error":   err,
		})
	} else {
		logrus.Infof("etcd put succeeded")
	}
}
