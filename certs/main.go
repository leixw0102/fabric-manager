package main

import (
	"Data_Bank/fabric-manager/certs/command"
	agentHTTP "Data_Bank/fabric-manager/certs/http"
	"Data_Bank/fabric-manager/common/connector"
	"Data_Bank/fabric-manager/common/utils"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"go.etcd.io/etcd/clientv3"
)

var httpServer *http.Server

func ListenToServerRequest() {
	// Setup etcd watch key
	watchKey := "/fabric-manager/certs/" + utils.CertCreateCmd
	watchChan := connector.ETCD.Watch(watchKey)
	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			logrus.Infof("Event received! %s executed on %q with value %q", event.Type, event.Kv.Key, event.Kv.Value)
			// cmd := utils.GetCmd(event.Kv.Key)
			if event.Type == mvccpb.PUT {
				logrus.Info("create crypto action received")
				var cryptoConfig map[string]map[string]string
				if err := json.Unmarshal(event.Kv.Value, &cryptoConfig); err != nil {
					logrus.Errorf("Fail to parse json, error:%+v", err)
				}
				command.CreateCrypto(cryptoConfig)

				// logrus.Errorf("%s is not known", cmd)
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
func initHttpServer() {
	router := agentHTTP.StartHTTP()
	httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}

// global variables
var userID string
var etcdAddr string
var port string

func main() {
	pflag.StringVar(&userID, "userid", "", "user id returned by fabric manager server when sign up")
	pflag.StringVar(&etcdAddr, "etcd_addr", "localhost:2379", "etcd address, ip:port")
	pflag.StringVar(&port, "port", "8080", "input server listen port")
	pflag.Parse()

	InitEtcd()

	go ListenToServerRequest()

	go initHttpServer()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch,
		// kill -SIGINT XXXX or Ctrl+c
		os.Interrupt,
		syscall.SIGINT, // register that too, it should be ok
		// os.Kill  is equivalent with the syscall.Kill
		os.Kill,
		syscall.SIGKILL, // register that too, it should be ok
		// kill -SIGTERM XXXX
		syscall.SIGTERM,
	)
	select {
	case <-ch:
		println("shutdown...")
		shutdownHTTP()
		shutdownETCD()
	}

}

func shutdownETCD() {
	connector.ETCD.Close()
}
func shutdownHTTP() {
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
