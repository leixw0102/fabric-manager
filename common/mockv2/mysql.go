package mockv2

import (
	"Data_Bank/fabric-manager/common/message"
	"Data_Bank/fabric-manager/common/utils"
	"strings"

	"github.com/sirupsen/logrus"
)

var AgentIP string

func GetOrgInfo(orgDomain string) *message.OrgInfo {
	var res *message.OrgInfo
	if orgDomain == "org1.example.com" {
		// peer org
		res = &message.OrgInfo{
			Name:   "Org1",
			Domain: orgDomain,
			Peers: []*message.PeerInfo{
				GetPeerInfo("peer0.org1.example.com"),
			},
			Users: []*message.UserInfo{{Name: "User1", Role: "user"}},
			Admin: &message.UserInfo{Name: "Admin", Role: "admin"},
		}
	} else if orgDomain == "example.com" {
		// orderer org
		res = &message.OrgInfo{
			Name:   "Orderer",
			Domain: orgDomain,
			Orderers: []*message.OrdererInfo{
				GetOrdererInfo("orderer.example.com"),
			},
			Admin: &message.UserInfo{Name: "Admin", Role: "admin"},
		}
	} else {
		logrus.Errorf("org:%s not known", orgDomain)
	}
	return res
}

func GetConsoritumInfo(name string) message.ConsortiumInfo {
	return message.ConsortiumInfo{
		Name: name,
		Orgs: []*message.OrgInfo{
			GetOrgInfo("example.com"),
			GetOrgInfo("org1.example.com"),
		},
	}
}

func GetPeerInfo(peerDomain string) *message.PeerInfo {
	return &message.PeerInfo{
		Name:   "peer0",
		Domain: peerDomain,
		IP:     AgentIP,
		Port:   7051,
	}
}

func GetOrdererInfo(ordererDomain string) *message.OrdererInfo {
	return &message.OrdererInfo{
		Name:   "orderer",
		Domain: ordererDomain,
		IP:     AgentIP,
		Port:   7050,
		Type:   "etcdraft",
		Cert:   strings.Join([]string{utils.BlockchainRoot, "organizations", "crypto", "organizations", "ordererOrganizations", "example.com", "orderers", ordererDomain, "tls", "server.crt"}, "/"),
	}
}

func GetConsortiumHosts(consortium string) []string {
	return []string{"orderer.example.com:" + AgentIP, "peer0.org1.example.com:" + AgentIP}
}
