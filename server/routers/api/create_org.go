package api

import (
	"Data_Bank/fabric-manager/common/connector"
	"Data_Bank/fabric-manager/common/message"
	"Data_Bank/fabric-manager/common/utils"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type PeerInfo struct {
	Name string
}

type OrgInfo struct {
	Name   string
	Domain string
	Peers  []PeerInfo
}

func CreateOrg(c *gin.Context) {
	// forward request to etcd
	var data message.CreateOrgMsg
	if err := c.BindJSON(&data); err != nil {
		logrus.Errorf("Fail to parse json object, error:%+v", err)
	}
	for _, identity := range data.Identities {
		ip := identity.IP
		action := utils.CreateIdentity
		key := strings.Join([]string{utils.AgentService, ip, action}, "/")
		msg := message.CreateIdentityMsg{
			OrgDomain: data.OrgDomain,
			CaName:    data.CaName,
			CaAddress: data.CaAddress,
			Identity:  identity,
		}
		bytes, err := json.Marshal(msg)
		if err != nil {
			logrus.Errorf("json marshal error:%+v", err)
		}
		err = connector.ETCD.Put(key, string(bytes), 5*time.Second)
		if err != nil {
			logrus.Errorf("etcd put failed with error: %v", err)
			c.JSON(200, gin.H{
				"message": fmt.Sprintf("Failt to generate identity:%s", identity.Name),
				"error":   err,
			})
		} else {
			logrus.Infof("etcd put succeeded")
		}
	}
}
