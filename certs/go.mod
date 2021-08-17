module Data_Bank/fabric-manager/certs

go 1.16

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace Data_Bank/fabric-manager/common => ../common

require (
	Data_Bank/fabric-manager/common v0.0.0-00010101000000-000000000000
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/gin-gonic/gin v1.7.2
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/pflag v1.0.5
	go.etcd.io/etcd v3.3.25+incompatible
)
