module Data_Bank/fabric-manager/agent

go 1.16

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace Data_Bank/fabric-manager/common => ../common

require (
	Data_Bank/fabric-manager/common v0.0.0-00010101000000-000000000000
	github.com/coreos/bbolt v1.3.2 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20190321100706-95778dfbb74e // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/pflag v1.0.3
	github.com/tmc/grpc-websocket-proxy v0.0.0-20190109142713-0ad062ec5ee5 // indirect
	go.etcd.io/bbolt v1.3.3
	go.etcd.io/etcd v3.3.25+incompatible
	google.golang.org/genproto v0.0.0-20191108220845-16a3f7862a1a // indirect
	gopkg.in/yaml.v2 v2.3.0
)
