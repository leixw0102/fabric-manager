package command

import (
	"Data_Bank/fabric-manager/agent/config"
	"Data_Bank/fabric-manager/common/mock"
	"Data_Bank/fabric-manager/common/utils"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func StartContainer(dockerComposePath string) error {
	cmd := fmt.Sprintf("docker-compose -f %s up -d", dockerComposePath)
	return utils.ExecLocalCommand(cmd)
}

func StartOrdererService(consortium, domain string) error {
	ordererInfo := mock.GetOrdererInfo(domain)
	orgDomain := ordererInfo.GetOrgDomain()
	dockerComposeRoot := filepath.Join(utils.BlockchainRoot, "organizations", orgDomain, "docker")
	utils.MkdirIfNotExists(dockerComposeRoot)
	dockerComposePath := filepath.Join(dockerComposeRoot, domain+".yaml")
	orgMSP := ordererInfo.GetOrgName() + "MSP"
	ports := []string{fmt.Sprintf("%d:%d", ordererInfo.Port, ordererInfo.Port)}
	hosts := mock.GetConsortiumHosts(consortium)
	consortiumRoot := filepath.Join(utils.BlockchainRoot, "consortiums", consortium)
	genesisBlockPath := filepath.Join(consortiumRoot, "system-genesis-block", "genesis.block")
	data := config.NewOrdererService(orgDomain, domain, orgMSP, genesisBlockPath, ports, hosts)
	if err := ioutil.WriteFile(dockerComposePath, data, os.ModePerm); err != nil {
		logrus.Errorf("Fail to generate %s, error:%v", dockerComposePath, err)
	}
	return StartContainer(dockerComposePath)
}

func StartPeerService(consortium, domain string) error {
	peerInfo := mock.GetPeerInfo(domain)
	orgDomain := peerInfo.GetOrgDomain()
	dockerComposeRoot := filepath.Join(utils.BlockchainRoot, "organizations", orgDomain, "docker")
	utils.MkdirIfNotExists(dockerComposeRoot)
	dockerComposePath := filepath.Join(dockerComposeRoot, domain+".yaml")
	orgMSP := peerInfo.GetOrgName() + "MSP"
	ports := []string{fmt.Sprintf("%d:%d", peerInfo.Port, peerInfo.Port)}
	hosts := mock.GetConsortiumHosts(consortium)
	data := config.NewPeerService(consortium, orgDomain, domain, orgMSP, ports, hosts)
	if err := ioutil.WriteFile(dockerComposePath, data, os.ModePerm); err != nil {
		logrus.Errorf("Fail to generate %s, error:%v", dockerComposePath, err)
	}
	return StartContainer(dockerComposePath)
}
