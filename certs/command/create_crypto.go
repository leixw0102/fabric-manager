package command

import (
	"Data_Bank/fabric-manager/common/utils"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func CreateCrypto(params map[string]map[string]string) {
	// generate crypto.yaml
	cryptoDir := filepath.Join(utils.BlockchainRoot, "organizations", "crypto")
	utils.MkdirIfNotExists(cryptoDir)
	configPath := filepath.Join(cryptoDir, "crypto.yaml")
	bytes := CreateCryptoConfig(params)
	if err := ioutil.WriteFile(configPath, bytes, os.ModePerm); err != nil {
		logrus.Errorf("WriteFile error:%v", err)
	}
	// execute cryptogen command
	outputPath := filepath.Join(cryptoDir, "organizations")
	cmdStr := fmt.Sprintf("cryptogen generate --config %s --output %s", configPath, outputPath)
	utils.ExecLocalCommand(cmdStr)
	// add tar.gz
	_, _, err := utils.ExecuteCommandFile("./cert.sh")
	fmt.Println(err)
}

// create crypto.yaml
func CreateCryptoConfig(params map[string]map[string]string) []byte {
	cfg := map[string]interface{}{
		"PeerOrgs": []map[string]interface{}{
			{
				"Name":          params["PeerOrgs"]["Name"],
				"Domain":        params["PeerOrgs"]["Domain"],
				"EnableNodeOUs": true,
				"Template": map[string]interface{}{
					"Count": 1,
					"Sans":  []string{"localhost"},
				},
				"Users": map[string]interface{}{
					"Count": 1,
				},
			},
		},
		"OrdererOrgs": []map[string]interface{}{
			{
				"Name":          params["OrdererOrgs"]["Name"],
				"Domain":        params["OrdererOrgs"]["Domain"],
				"EnableNodeOUs": true,
				"Specs": []map[string]interface{}{
					{
						"Hostname": "orderer",
						"SANS":     []string{"localhost"},
					},
				},
			},
		},
	}
	return utils.ToYaml(cfg)
}
