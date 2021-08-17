module Data_Bank/fabric-manager/server

go 1.16

require (
	Data_Bank/fabric-manager/common v0.0.0-00010101000000-000000000000
	github.com/beego/beego/v2 v2.0.1
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/gin-gonic/gin v1.7.1
	github.com/go-playground/validator/v10 v10.5.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/btree v1.0.1 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/soheilhy/cmux v0.1.5 // indirect
	github.com/spf13/viper v1.8.0
	github.com/tmc/grpc-websocket-proxy v0.0.0-20201229170055-e5319fda7802 // indirect
	github.com/ugorji/go v1.2.5 // indirect
	go.etcd.io/etcd v3.3.25+incompatible
	golang.org/x/text v0.3.6 // indirect
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba // indirect
	sigs.k8s.io/yaml v1.2.0 // indirect
)

replace Data_Bank/fabric-manager/common => ../common

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
