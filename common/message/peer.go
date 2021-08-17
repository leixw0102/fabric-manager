package message

import (
	"Data_Bank/fabric-manager/common/utils"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

type PeerInfo struct {
	Name   string
	Domain string
	IP     string
	Port   int
}

func (p *PeerInfo) Address() string {
	return fmt.Sprintf("%s:%d", p.IP, p.Port)
}

func (p *PeerInfo) GetOrgDomain() string {
	idx := strings.Index(p.Domain, p.Name)
	if idx > -1 {
		idx += len(p.Name)
		return p.Domain[idx+1:]
	}
	return ""
}

func (p *PeerInfo) GetOrgName() string {
	return "Org1"
}

func (p *PeerInfo) GetDockerPath() string {
	orgDomain := p.GetOrgDomain()
	if orgDomain == "" {
		logrus.Errorf("Fail to get org domain for peer:%s", p.Domain)
		return ""
	}
	return filepath.Join(utils.BlockchainRoot, "organizations", orgDomain, "docker", p.Domain+".yaml")
}
