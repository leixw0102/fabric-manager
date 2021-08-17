package mock

import (
	"Data_Bank/fabric-manager/common/message"
	"Data_Bank/fabric-manager/common/utils"
	"path/filepath"
)

var AgentIP string

func GetOrgInfo(orgDomain string) *message.OrgInfo {
	return &message.OrgInfo{
		Name:   "Org1",
		Domain: "org1.example.com",
		Orderers: []*message.OrdererInfo{
			GetOrdererInfo("orderer.example.com"),
			// {
			// 	Name:   "orderer0",
			// 	Domain: "orderer0.org1.example.com",
			// 	IP:     AgentIP,
			// 	Port:   7050,
			// 	Type:   "etcdraft",
			// 	Cert:   fmt.Sprintf("%s/%s/%s/%s/%s/%s/%s/%s", utils.BlockchainRoot, "organizations", "org1.example.com", "crypto", "orderers", "orderer0.org1.example.com", "tls", "server.crt"),
			// 	// Cert:   filepath.Join(utils.BlockchainRoot, "organizations", "org1.example.com", "crypto", "orderers", "orderer0.org1.example.com", "tls", "server.crt"),
			// },
		},
		Peers: []*message.PeerInfo{
			{
				Name:   "peer0",
				Domain: "peer0.org1.example.com",
				IP:     AgentIP,
				Port:   7051,
			},
		},
		Users: []*message.UserInfo{{Name: "user1", Role: "user"}},
		Admin: &message.UserInfo{Name: "Admin", Role: "admin"},
	}
}

func GetConsoritumInfo(name string) message.ConsortiumInfo {
	return message.ConsortiumInfo{
		Name: "SampleConsortium",
		Orgs: []*message.OrgInfo{
			GetOrgInfo("org1.example.com"),
		},
	}
}

func GetPeerInfo(peerDomain string) *message.PeerInfo {
	return &message.PeerInfo{
		Name:   "peer0",
		Domain: "peer0.org1.example.com",
		IP:     AgentIP,
		Port:   7051,
	}
}

func GetOrdererInfo(ordererDomain string) *message.OrdererInfo {
	return &message.OrdererInfo{
		Name:   "orderer0",
		Domain: ordererDomain,
		IP:     AgentIP,
		Port:   7050,
		Type:   "etcdraft",
		Cert:   filepath.Join(utils.BlockchainRoot, "organizations", "org1.example.com", "crypto", "orderers", "orderer0.org1.example.com", "tls", "server.crt"),
	}
}

func GetOrdererInfoV2(ordererDomain string) *message.OrdererInfo {
	return &message.OrdererInfo{
		Name:   "orderer",
		Domain: ordererDomain,
		IP:     AgentIP,
		Port:   7050,
		Type:   "etcdraft",
		Cert:   filepath.Join(utils.BlockchainRoot, "organizations", "org1.example.com", "crypto", "orderers", "orderer0.org1.example.com", "tls", "server.crt"),
	}
}

func GetConsortiumHosts(consortium string) []string {
	return []string{"orderer0.org1.example.com:" + AgentIP, "peer0.org1.example.com:" + AgentIP}
}
