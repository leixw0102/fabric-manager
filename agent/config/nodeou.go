package config

import (
	"Data_Bank/fabric-manager/common/utils"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func GenNodeOUConfig(caAddress, caName, where string) error {
	certName := fmt.Sprintf("cacerts/%s-%s", caAddress, caName)
	certName = strings.Replace(certName, ".", "-", -1)
	certName = strings.Replace(certName, ":", "-", -1)
	certName += ".pem"
	config := map[string]interface{}{
		"NodeOUs": map[string]interface{}{
			"Enable": true,
			"ClientOUIdentifier": map[string]interface{}{
				"Certificate":                  certName,
				"OrganizationalUnitIdentifier": "client",
			},
			"PeerOUIdentifier": map[string]interface{}{
				"Certificate":                  certName,
				"OrganizationalUnitIdentifier": "peer",
			},
			"AdminOUIdentifier": map[string]interface{}{
				"Certificate":                  certName,
				"OrganizationalUnitIdentifier": "admin",
			},
			"OrdererOUIdentifier": map[string]interface{}{
				"Certificate":                  certName,
				"OrganizationalUnitIdentifier": "orderer",
			},
		},
	}
	content := utils.ToYaml(config)
	return ioutil.WriteFile(where, content, os.ModePerm)
}
