package api

import (
	"Data_Bank/fabric-manager/common/connector"
	"Data_Bank/fabric-manager/common/message"
	"Data_Bank/fabric-manager/common/mock"
	"Data_Bank/fabric-manager/common/utils"
	"encoding/json"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func StartConsortium(c *gin.Context) {
	var data message.StartConsortiumMsg
	if err := c.BindJSON(&data); err != nil {
		logrus.Errorf("Fail to parse json object, error:%+v", err)
	}
	// send start command to each peer and orderer in each org
	for _, org := range mock.GetConsoritumInfo(data.Name).Orgs {
		// start orderers
		for _, orderer := range org.Orderers {
			logrus.Infof("Sending start command to %s", orderer.Domain)
			action := utils.StartOrderer
			key := strings.Join([]string{utils.AgentService, orderer.IP, action}, "/")
			msg := message.StartServiceMsg{Consortium: data.Name, Domain: orderer.Domain}
			bytes, _ := json.Marshal(msg)
			if err := connector.ETCD.Put(key, string(bytes), 5*time.Second); err != nil {
				logrus.Errorf("Fail to put message %q to etcd with path:%s")
			}
		}
		// start peers
		for _, peer := range org.Peers {
			logrus.Infof("Sending start command to %s", peer.Domain)
			action := utils.StartPeer
			key := strings.Join([]string{utils.AgentService, peer.IP, action}, "/")
			msg := message.StartServiceMsg{Consortium: data.Name, Domain: peer.Domain}
			bytes, _ := json.Marshal(msg)
			if err := connector.ETCD.Put(key, string(bytes), 5*time.Second); err != nil {
				logrus.Errorf("Fail to put message %q to etcd with path:%s")
			}
		}
	}
}
