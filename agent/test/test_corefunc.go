package test

import (
	"Data_Bank/fabric-manager/agent/command"
	"Data_Bank/fabric-manager/agent/config"
	"Data_Bank/fabric-manager/common/message"
	"Data_Bank/fabric-manager/common/mock"
	"Data_Bank/fabric-manager/common/utils"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func TestCreateOrg() {
	orgDomain := "org1.example.com"
	caAddress := "172.38.50.212:7054"
	caName := "ca-org1"
	identities := []message.Identity{
		{
			Name:     "peer0",
			Password: "peer0pw",
			IdType:   "peer",
		},
		{
			Name:     "org1admin",
			Password: "org1adminpw",
			IdType:   "admin",
		},
		{
			Name:     "user1",
			Password: "user1pw",
			IdType:   "client",
		},
		{
			Name:     "orderer0",
			Password: "orderer0pw",
			IdType:   "orderer",
		},
	}
	command.CreateIdentity(orgDomain, caName, caAddress, identities[0])
}

func TestCreatePeerDockerCompose() {
	consortium := "SampleConsortium"
	orgDomain := "org1.example.com"
	peerDomain := "peer0." + orgDomain
	mspID := "ORG1MSP"
	hosts := []string{fmt.Sprintf("%s:172.38.50.211", peerDomain)}
	ports := []string{"7051:7051"}
	outputRoot := filepath.Join(utils.BlockchainRoot, "organizations", orgDomain, "docker")
	content := config.NewPeerService(consortium, orgDomain, peerDomain, mspID, ports, hosts)
	filename := filepath.Join(outputRoot, peerDomain+".yaml")
	utils.MkdirIfNotExists(outputRoot)
	ioutil.WriteFile(filename, content, os.ModePerm)
}

func TestCreateOrdererDockerCompose() {
	orgDomain := "org1.example.com"
	ordererDomain := "orderer0." + orgDomain
	mspID := "ORG1MSP"
	hosts := []string{orgDomain + ":172.38.50.211"}
	ports := []string{"7050:7050"}
	genesisBlockPath := filepath.Join(utils.BlockchainRoot, "system-genesis-block", "genesis.block")
	content := config.NewOrdererService(orgDomain, ordererDomain, mspID, genesisBlockPath, ports, hosts)
	outputRoot := filepath.Join(utils.BlockchainRoot, "organizations", orgDomain, "docker")
	filename := filepath.Join(outputRoot, ordererDomain+".yaml")
	utils.MkdirIfNotExists(outputRoot)
	ioutil.WriteFile(filename, content, os.ModePerm)
}

func TestGenNodeOU() {
	caName := "ca-org1"
	caAddress := "172.38.50.212:7054"
	orgDomain := "org1.example.com"
	orgRoot := filepath.Join(utils.BlockchainRoot, "organizations", orgDomain)
	config.GenNodeOUConfig(caAddress, caName, filepath.Join(orgRoot, "crypto", "msp", "config.yaml"))
}

func TestCreateGenesisBlock() {
	consortiumName := "SampleConsortium"
	command.CreateGenesisBlock(consortiumName, []*message.OrgInfo{mock.GetOrgInfo("org1.example.com")})
}
