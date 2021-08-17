package config

import (
	"Data_Bank/fabric-manager/common/utils"
	"fmt"
	"path/filepath"
)

func NewPeerService(consortium, orgDomain, peerDomain, mspID string, ports, hosts []string) []byte {
	consortiumRoot := filepath.Join(utils.BlockchainRoot, "consortiums", consortium)
	cryptoRoot := filepath.Join(utils.BlockchainRoot, "organizations", orgDomain, "crypto")
	var peer = map[string]interface{}{
		"version": "2",
		"services": map[string]interface{}{
			peerDomain: map[string]interface{}{
				"container_name": peerDomain,
				"image":          "hyperledger/fabric-peer:2.2.0",
				"working_dir":    "/opt/gopath/src/github.com/hyperledger/fabric/peer",
				"command":        "peer node start",
				"volumes": []string{
					"/var/run/:/host/var/run/",
					fmt.Sprintf("%s/organizations/%s/crypto/peers/%s/msp:/etc/hyperledger/fabric/msp", utils.BlockchainRoot, orgDomain, peerDomain),
					fmt.Sprintf("%s/organizations/%s/crypto/peers/%s/tls:/etc/hyperledger/fabric/tls", utils.BlockchainRoot, orgDomain, peerDomain),
				},
				"ports":       ports,
				"extra_hosts": hosts,
				"environment": []string{
					"CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock",
					fmt.Sprintf("CORE_PEER_ID=%s", peerDomain),
					fmt.Sprintf("CORE_PEER_ADDRESS=%s:7051", peerDomain),
					"CORE_PEER_LISTENADDRESS=0.0.0.0:7051",
					"CORE_PEER_CHAINCODELISTENADDRESS=0.0.0.0:7052",
					"CORE_VM_DOCKER_HOSTCONFIG_NETWORKMODE=peers_default",
					fmt.Sprintf("CORE_PEER_CHAINCODEADDRESS=%s:7052", peerDomain),
					fmt.Sprintf("CORE_PEER_GOSSIP_EXTERNALENDPOINT=%s:7051", peerDomain),
					fmt.Sprintf("CORE_PEER_GOSSIP_BOOTSTRAP=%s:7051", peerDomain),
					fmt.Sprintf("CORE_PEER_LOCALMSPID=%s", mspID),
					"FABRIC_LOGGING_SPEC=INFO",
					"CORE_PEER_TLS_ENABLED=true",
					"CORE_PEER_GOSSIP_USELEADERELECTION=true",
					"CORE_PEER_GOSSIP_ORGLEADER=false",
					"CORE_PEER_PROFILE_ENABLED=true",
					"CORE_PEER_TLS_CERT_FILE=/etc/hyperledger/fabric/tls/server.crt",
					"CORE_PEER_TLS_KEY_FILE=/etc/hyperledger/fabric/tls/server.key",
					"CORE_PEER_TLS_ROOTCERT_FILE=/etc/hyperledger/fabric/tls/ca.crt",
					"CORE_CHAINCODE_EXECUTETIMEOUT=300s",
				},
			},
			"cli": map[string]interface{}{
				"container_name": "cli",
				"image":          "hyperledger/fabric-tools:2.2.0",
				"tty":            true,
				"stdin_open":     true,
				"environment": []string{
					"GOPROXY=https://goproxy.io,direct",
					"GO111MODULE=on",
					"GOPATH=/opt/gopath",
					"CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock",
					"FABRIC_LOGGING_SPEC=DEBUG",
					"CORE_PEER_ID=cli",
					fmt.Sprintf("CORE_PEER_ADDRESS=%s:7051", peerDomain),
					fmt.Sprintf("CORE_PEER_LOCALMSPID=%s", mspID),
					"CORE_PEER_TLS_ENABLED=true",
					fmt.Sprintf("CORE_PEER_TLS_CERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peers/%s/tls/server.crt", peerDomain),
					fmt.Sprintf("CORE_PEER_TLS_KEY_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peers/%s/tls/server.key", peerDomain),
					fmt.Sprintf("CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peers/%s/tls/ca.crt", peerDomain),
					fmt.Sprintf("CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/users/Admin@%s/msp", orgDomain),
				},
				"working_dir": "/opt/gopath/src/github.com/hyperledger/fabric/peer",
				"command":     "/bin/bash",
				"volumes": []string{
					"/var/run/:/host/var/run/",
					fmt.Sprintf("%s/chaincode/go/:/opt/gopath/src/github.com/hyperledger/chaincode/go", consortiumRoot),
					fmt.Sprintf("%s:/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/", cryptoRoot),
					fmt.Sprintf("%s/channel-artifacts:/opt/gopath/src/github.com/hyperledger/fabric/peer/channel-artifacts", consortiumRoot),
					fmt.Sprintf("%s/chaincode:/opt/gopath/src/github.com/hyperledger/fabric/peer/chaincode", consortiumRoot),
				},
				"depends_on": []string{
					peerDomain,
				},
				"extra_hosts": hosts,
			},
		},
	}
	return utils.ToYaml(peer)
}

func NewOrdererService(orgDomain, ordererDomain, mspID, genesisBlockPath string, ports, hosts []string) []byte {
	mspPath := filepath.Join(utils.BlockchainRoot, "organizations", orgDomain, "crypto", "orderers", ordererDomain, "msp")
	tlsPath := filepath.Join(utils.BlockchainRoot, "organizations", orgDomain, "crypto", "orderers", ordererDomain, "tls")
	var orderer = map[string]interface{}{
		"version": "2",
		"services": map[string]interface{}{
			ordererDomain: map[string]interface{}{
				"container_name": ordererDomain,
				"image":          "hyperledger/fabric-orderer:2.2.0",
				"working_dir":    "/opt/gopath/src/github.com/hyperledger/fabric",
				"command":        "orderer",
				"volumes": []string{
					fmt.Sprintf("%s:/var/hyperledger/orderer/orderer.genesis.block", genesisBlockPath),
					fmt.Sprintf("%s:/var/hyperledger/orderer/msp", mspPath),
					fmt.Sprintf("%s:/var/hyperledger/orderer/tls", tlsPath),
				},
				"ports":       ports,
				"extra_hosts": hosts,
				"environment": []string{
					"FABRIC_LOGGING_SPEC=DEBUG",
					"ORDERER_GENERAL_LISTENADDRESS=0.0.0.0",
					"ORDERER_GENERAL_BOOTSTRAPMETHOD=file",
					"ORDERER_GENERAL_BOOTSTRAPFILE=/var/hyperledger/orderer/orderer.genesis.block",
					fmt.Sprintf("ORDERER_GENERAL_LOCALMSPID=%s", mspID),
					"ORDERER_GENERAL_LOCALMSPDIR=/var/hyperledger/orderer/msp",
					"ORDERER_GENERAL_TLS_ENABLED=true",
					"ORDERER_GENERAL_TLS_PRIVATEKEY=/var/hyperledger/orderer/tls/server.key",
					"ORDERER_GENERAL_TLS_CERTIFICATE=/var/hyperledger/orderer/tls/server.crt",
					"ORDERER_GENERAL_TLS_ROOTCAS=[/var/hyperledger/orderer/tls/ca.crt]",
					"ORDERER_GENERAL_CLUSTER_CLIENTCERTIFICATE=/var/hyperledger/orderer/tls/server.crt",
					"ORDERER_GENERAL_CLUSTER_CLIENTPRIVATEKEY=/var/hyperledger/orderer/tls/server.key",
					"ORDERER_GENERAL_CLUSTER_ROOTCAS=[/var/hyperledger/orderer/tls/ca.crt]",
				},
			},
		},
	}
	return utils.ToYaml(&orderer)
}
