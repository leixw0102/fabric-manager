package test

import (
	"Data_Bank/fabric-manager/common/utils"
	"fmt"
)

func TestGetIP() {
	ip := utils.GetOutboundIP().String()
	fmt.Println(ip)
}

func TestCopy() {
	utils.Copy("/root/fabric_networks/organizations/peerOrganizations/org1.example.com/peers/peer3.org1.example.com/tls/signcerts/*", "/root/fabric_networks/organizations/peerOrganizations/org1.example.com/peers/peer3.org1.example.com/tls/server.crt")
}
