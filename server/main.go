package main

import (
	"Data_Bank/fabric-manager/common/connector"
	"Data_Bank/fabric-manager/common/mock"
	"Data_Bank/fabric-manager/common/mockv2"
	"Data_Bank/fabric-manager/common/utils"
	"Data_Bank/fabric-manager/server/server"
	"fmt"
	"net/http"
	"strings"
	"time"

	"Data_Bank/fabric-manager/server/monitor"

	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.etcd.io/etcd/clientv3"
)

var agentMonitor monitor.AgentMonitor

func main() {
	agentMonitor = monitor.NewAgentMonitor()
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		logrus.Panicf("Fail to read config, error:%v", err)
	}
	// InitMysql()
	InitEtcd()
	go agentMonitor.MonitorAgentHeartbeat(5 * time.Second)
	go ListenToEtcd()
	InitServer()
}

func InitMysql() {
	config := viper.GetStringMapString("mysql")
	// Init mysql connector
	logrus.Infoln("Initializing mysql connector ...")
	logrus.Infoln("mysql username:", config["username"])
	logrus.Infoln("mysql password:", config["password"])
	logrus.Infoln("mysql host:", config["host"])
	logrus.Infoln("mysql db:", config["db"])
	dbConfig := &connector.DBConfig{
		Username: config["username"],
		Password: config["password"],
		Hostname: config["host"],
		DBName:   config["db"],
	}
	connector.DB = connector.NewMysql(dbConfig) // init global mysql connector instance
}

func InitServer() {
	// Init server
	logrus.Info("Initializing server ...")
	port := viper.GetInt("server.port")
	readTimeout := viper.GetDuration("server.read_timeout") * time.Second
	writeTimeout := viper.GetDuration("server.write_timeout") * time.Second
	maxHeaderBytes := viper.GetInt("server.max_header_bytes")
	logrus.Infoln("server port:", port)
	logrus.Infoln("server read timeout:", readTimeout)
	logrus.Infoln("server write timeout:", writeTimeout)
	logrus.Infoln("server max header bytes:", maxHeaderBytes)
	serverConfig := &server.ServerConfig{
		Addr:           fmt.Sprintf(":%d", port),
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	s := server.NewServer(serverConfig)
	go func() {
		if err := s.Start(); err != nil && err != http.ErrServerClosed {
			logrus.Fatal("server crashed with error:", err)
		}
	}()
	// Graceful exit
	s.WaitforShutdown()
}

func InitEtcd() {
	// Init etcd connector
	endpoints := viper.GetStringSlice("etcd.endpoints")
	timeout := viper.GetDuration("etcd.timeout") * time.Second
	logrus.Info("Initializing etcd connector ...")
	logrus.Infof("etcd endpoints:%v", endpoints)
	logrus.Infoln("etcd dial timeout:", timeout)
	etcdConfig := clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: timeout * time.Second,
	}
	connector.ETCD = connector.NewETCD(etcdConfig) // init global etcd connector
}

func ListenToEtcd() {
	// Setup etcd watch channel
	ch := connector.ETCD.WatchPrefix("fabric-manager/server")
	for resp := range ch {
		for _, event := range resp.Events {
			if event.Type == mvccpb.PUT {
				key := event.Kv.Key
				cmd := utils.GetCmd(key)
				value := utils.GetParams(event.Kv.Value)
				switch cmd {
				case "heartbeat":
					agentMonitor.OnPing(value)
				case "reportAgentInfo":
					logrus.Infof("reportAgentInfo event received, value is %v", value)
					mock.AgentIP = strings.Split(value, ",")[0]
					mockv2.AgentIP = strings.Split(value, ",")[0]
				default:
					logrus.Errorf("message type not known, key:%s, value:%v", key, value)
				}
			}
		}
	}
}
