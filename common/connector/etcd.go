package connector

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"
)

var ETCD *EtcdConnector

type EtcdConfig struct {
	Address     string
	DialTimeout time.Duration
}

// TODO: Add connector heartbeat
func NewETCD(config clientv3.Config) *EtcdConnector {
	etcd, err := clientv3.New(config)
	if err != nil {
		logrus.Panicf("Fail to connect to etcd clsuter,error:%v", err)
	} else {
		logrus.Info("connected to etcd")
	}
	return &EtcdConnector{
		etcd: etcd,
	}
}

type EtcdConnector struct {
	etcd *clientv3.Client
}

func (c *EtcdConnector) Put(key, value string, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	resp, err := c.etcd.Put(ctx, key, value)
	cancel()
	if err != nil {
		logrus.Errorf("Put key: %s with value: %s, resp:%v failed, error: %v", key, value, resp, err)
	} else {
		logrus.Infof("Put key: %s with value: %s, resp:%v succeeded", key, value, resp)
	}
	return err
}

func (c *EtcdConnector) Watch(key string) clientv3.WatchChan {
	return c.etcd.Watch(context.Background(), key)
}

func (c *EtcdConnector) WatchPrefix(prefix string) clientv3.WatchChan {
	return c.etcd.Watch(context.Background(), prefix, clientv3.WithPrefix())
}

func (c *EtcdConnector) Close() error {
	return c.etcd.Close()
}
