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

func CreateConsortium(c *gin.Context) {
	var data message.CreateConsortiumMsg
	if err := c.BindJSON(&data); err != nil {
		logrus.Errorf("Fail to parse json object, error:%+v", err)
		return
	}
	// sent create consortium command to the orderer0 of initiate org
	orgs := make([]*message.OrgInfo, 0)
	initOrg := mockv2.GetOrgInfo(data.InitiateOrg)
	orgs = append(orgs, initOrg) // add init org to consortium org list
	for _, orgName := range data.OrgDomains {
		org := mockv2.GetOrgInfo(orgName)
		if org != nil {
			orgs = append(orgs, org)
		}
	}
	orderer := initOrg.GetOrdererInfo(data.Orderer) // get first orderer of init org, this
	if orderer == nil {
		c.JSON(404, gin.H{
			"message": "database query error",
			"error":   fmt.Sprintf("%s not found in database", data.Orderer),
		})
		return
	}
	action := utils.CreateConsortium
	bytes, err := json.Marshal(&message.ConsortiumInfo{
		Name: data.Name,
		Orgs: orgs,
	})
	if err != nil {
		logrus.Errorf("json marshal error:%v", err)
		c.JSON(200, gin.H{
			"message": "json marshal error",
			"error":   err,
		})
		return
	}
	key := strings.Join([]string{utils.AgentService, orderer.IP, action}, "/")
	err = connector.ETCD.Put(key, string(bytes), 5*time.Second)
	if err != nil {
		logrus.Errorf("etcd put failed with error: %v", err)
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Fail to create consortium:%s", data.Name),
			"error":   err,
		})
	} else {
		logrus.Infof("etcd put key: %s succeeded", key)
	}
}
