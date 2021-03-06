package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"Data_Bank/fabric-manager/agent/command"
	"Data_Bank/fabric-manager/common/connector"
	"Data_Bank/fabric-manager/common/message"
	"Data_Bank/fabric-manager/common/mock"
	"Data_Bank/fabric-manager/common/utils"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	bolt "go.etcd.io/bbolt"
	"go.etcd.io/etcd/clientv3"
)

func BOLT() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		logrus.Panicf("Fail to open db, error:%v", err)
	}
	defer db.Close()
	db.Update(func(tx *bolt.Tx) error {
		b, _ := tx.CreateBucket([]byte("myBucket"))
		return b.Put([]byte("foo"), []byte("bar"))
	})
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("myBucket"))
		v := b.Get([]byte("foo"))
		logrus.Infof("%s \n", v)
		return nil
	})
}

func ListenToServerRequest() {
	// Setup etcd watch key
	prefix := fmt.Sprintf("fabric-manager/agent/%s/", mock.AgentIP)
	watchChan := connector.ETCD.WatchPrefix(prefix)
	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			logrus.Infof("Event received! %s executed on %q with value %q", event.Type, event.Kv.Key, event.Kv.Value)
			cmd := utils.GetCmd(event.Kv.Key)
			switch cmd {
			case utils.CreateIdentity:
				logrus.Infoln("create org action received")
				data := message.CreateIdentityMsg{}
				if err := json.Unmarshal(event.Kv.Value, &data); err != nil {
					logrus.Errorf("Fail to parse json, error:%+v", err)
				}
				command.CreateIdentity(data.OrgDomain, data.CaName, data.CaAddress, data.Identity)
			case utils.CreateConsortium:
				logrus.Infoln("create consortium action received")
				data := message.ConsortiumInfo{}
				if err := json.Unmarshal(event.Kv.Value, &data); err != nil {
					logrus.Errorf("Fail to parse json, error:%+v", err)
				}
				command.CreateGenesisBlock(data.Name, data.Orgs)
			case utils.StartOrderer:
				logrus.Infoln("start network action received")
				msg := message.StartServiceMsg{}
				if err := json.Unmarshal(event.Kv.Value, &msg); err != nil {
					logrus.Errorf("json unmarshal error,%v", err)
				}
				command.StartOrdererService(msg.Consortium, msg.Domain)
			case utils.StartPeer:
				logrus.Infoln("start network action received")
				msg := message.StartServiceMsg{}
				if err := json.Unmarshal(event.Kv.Value, &msg); err != nil {
					logrus.Errorf("json unmarshal error,%v", err)
				}
				command.StartPeerService(msg.Consortium, msg.Domain)
			case utils.CreateChannel:
				msg := message.CreateChannelMsg{}
				if err := json.Unmarshal(event.Kv.Value, &msg); err != nil {
					logrus.Errorf("json unmarshal error,%v", err)
				}
				command.CreateChannel(msg.Consortium, msg.Channel, msg.Orgs)
			default:
				logrus.Errorf("%s is not known", cmd)
			}
		}
	}
}

func InitEtcd() {
	// Init etcd connector
	logrus.Info("Initializing etcd connector ...")
	etcdConfig := clientv3.Config{
		Endpoints:   []string{"http://" + etcdAddr},
		DialTimeout: 5 * time.Second,
	}
	connector.ETCD = connector.NewETCD(etcdConfig) // init global etcd connector
}

// global variables
var userID string
var agentID string
var serverChan string = "fabric-manager/server"
var debug bool
var etcdAddr string

func ReportAgentInfo() {
	mock.AgentIP = utils.GetOutboundIP().String()
	agentID = utils.GetUUID()
	connector.ETCD.Put(serverChan+"/"+"reportAgentInfo", mock.AgentIP+","+agentID+","+userID, 5*time.Second)
	// go func() {
	// 	for {
	// 		connector.ETCD.Put(serverChan+"/"+"heartbeat", agentID, 5*time.Second)
	// 		time.Sleep(5 * time.Second)
	// 	}
	// }()
}

func main() {
	pflag.StringVar(&userID, "userid", "", "user id returned by fabric manager server when sign up")
	pflag.BoolVar(&debug, "debug", false, "if run agent in debug mode")
	pflag.StringVar(&etcdAddr, "etcd_addr", "", "etcd address, ip:port")
	pflag.Parse()
	if etcdAddr == "" {
		logrus.Error("etcd_addr must be provided")
		os.Exit(1)
	}
	if debug {
		logrus.Info("Running in debug mode.")
		// test.TestCreateOrg()
		// logrus.Infoln("Creating orderer docker compose ...")
		// test.TestCreateOrdererDockerCompose()
		// logrus.Infoln("Creating genesis block ...")
		// // test.TestGenNodeOU()
		// test.TestCreateGenesisBlock()
	} else {
		InitEtcd()
		ReportAgentInfo()
		ListenToServerRequest()
	}
}
