package server

import (
	"Data_Bank/fabric-manager/common/connector"
	"Data_Bank/fabric-manager/server/routers"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

type ServerConfig struct {
	Addr           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

func NewServer(cfg *ServerConfig) *FabricServer {
	router := routers.InitRouter()
	server := &http.Server{
		Addr:           cfg.Addr,
		Handler:        router,
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
		MaxHeaderBytes: cfg.MaxHeaderBytes,
	}
	return &FabricServer{
		server: server,
	}
}

type FabricServer struct {
	server *http.Server
}

func (s *FabricServer) Start() error {
	logrus.Info("fabric-manager server started")
	return s.server.ListenAndServe()
}

func (s *FabricServer) WaitforShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	logrus.Info("quit/terminate signal received, shutting down server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// shutdown gin server
	if err := s.server.Shutdown(ctx); err != nil {
		logrus.Fatal("Server shutdown with error:", err)
	}
	// shutdown mysql connection
	if err := connector.DB.Close(); err != nil {
		logrus.Fatalf("Fail to shutdown mysql connection, %v \n", err)
	}
	// shutdown etcd connection
	if err := connector.ETCD.Close(); err != nil {
		logrus.Fatalf("Fail to shutdown etcd connection, %v \n", err)
	}
	logrus.Info("Server exited")
}
