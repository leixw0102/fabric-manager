package api

import (
	"Data_Bank/fabric-manager/common/connector"
	"Data_Bank/fabric-manager/common/mockv2"
	"Data_Bank/fabric-manager/common/utils"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CreateCrypto(c *gin.Context) {
	action := utils.CreateCrypto
	key := strings.Join([]string{utils.AgentService, mockv2.AgentIP, action}, "/")
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logrus.Errorf("Fail to read request body, error:%v", err)
		return
	}
	err = connector.ETCD.Put(key, string(bytes), 5*time.Second)
	if err != nil {
		logrus.Errorf("etcd put failed with error: %v", err)
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Fail to create crypto:%v", err),
			"error":   err,
		})
	} else {
		logrus.Infof("etcd put succeeded")
	}
}

func CreateCryptoV2(c *gin.Context) {
	// action := utils.CreateCrypto
	// key := strings.Join([]string{utils.AgentService, mock.AgentIP, action}, "/")
	key := fmt.Sprintf(utils.CertDownloadPath, utils.CertCreateCmd)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logrus.Errorf("Fail to read request body, error:%v", err)
		return
	}
	err = connector.ETCD.Put(key, string(bytes), 5*time.Second)
	if err != nil {
		logrus.Errorf("etcd put failed with error: %v", err)
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Fail to create crypto:%v", err),
			"error":   err,
		})
	} else {
		logrus.Infof("etcd put succeeded")
	}
}
