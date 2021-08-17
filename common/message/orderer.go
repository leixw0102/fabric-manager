package message

import (
	"Data_Bank/fabric-manager/common/utils"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/sirupsen/logrus"
)

type OrdererInfo struct {
	Name   string
	Domain string
	IP     string
	Port   int
	Type   string // etcdraft
	Cert   string // tls cert path
}

// Capitalize the first character of a string
func ToPascalNaming(name string) string {
	if len(name) == 0 {
		return name
	}
	res := strings.ToUpper(string(name[0])) + name[1:]
	return res
}

func (o *OrdererInfo) Address() string {
	return fmt.Sprintf("%s:%d", o.Domain, o.Port)
}

func (o *OrdererInfo) GetOrgDomain() string {
	idx := strings.Index(o.Domain, o.Name)
	if idx > -1 {
		idx += len(o.Name)
		return o.Domain[idx+1:]
	}
	return ""
}

func (o *OrdererInfo) GetDockerPath() string {
	orgDomain := o.GetOrgDomain()
	if orgDomain == "" {
		logrus.Errorf("Fail to get org domain for peer:%s", o.Domain)
		return ""
	}
	return filepath.Join(utils.BlockchainRoot, "organizations", orgDomain, "docker", o.Domain+".yaml")
}

func (o *OrdererInfo) GetOrgName() string {
	return ToPascalNaming(o.Name)
}
