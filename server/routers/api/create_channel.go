package api

import (
	"Data_Bank/fabric-manager/common/connector"
	"Data_Bank/fabric-manager/common/message"
	"Data_Bank/fabric-manager/common/mockv2"
	"Data_Bank/fabric-manager/common/utils"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CreateChannel(c *gin.Context) {
	var data message.CreateChannelMsg
	if err := c.BindJSON(&data); err != nil {
		logrus.Errorf("Fail to parse json object, error:%+v", err)
	}
	data.Orgs = append(data.Orgs, data.InitOrg)
	for _, org := range data.Orgs {
		if mockv2.GetOrgInfo(org) == nil {
			c.JSON(200, gin.H{
				"message": fmt.Sprintf("%s doesn't exist", org),
				"error":   "org doesn't exist",
			})
			return
		}
	}
	bytes, err := json.Marshal(data)
	if err != nil {
		logrus.Errorf("json marshal error:%v", err)
		c.JSON(200, gin.H{
			"message": "json marshal error",
			"error":   err,
		})
	}
	// send create channel command to the peer0 of init org of the consortium
	peer := mockv2.GetOrgInfo(data.InitOrg).GetPeerInfo(0)
	action := utils.CreateChannel
	key := strings.Join([]string{utils.AgentService, peer.IP, action}, "/")
	err = connector.ETCD.Put(key, string(bytes), 5*time.Second)
	if err != nil {
		logrus.Errorf("etcd put key:%s failed with error: %v", key, err)
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Fail to create channel:%s", data.Channel),
			"error":   err,
		})
	} else {
		logrus.Infof("etcd put key:%s succeeded", key)
	}
}
