package command

import (
	"Data_Bank/fabric-manager/agent/config"
	"Data_Bank/fabric-manager/common/message"
	"Data_Bank/fabric-manager/common/mock"
	"Data_Bank/fabric-manager/common/utils"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func CreateChannel(consortium, channel string, orgNames []string) error {
	// create configtx.yaml for channel
	orgs := make([]*message.OrgInfo, 0)
	for _, org := range orgNames {
		orgs = append(orgs, mock.GetOrgInfo(org))
	}
	channelCfg := config.GenChannelConfigtx(consortium, channel, orgs)
	bytes, err := yaml.Marshal(channelCfg)
	if err != nil {
		logrus.Errorf("json marshal error:%v", err)
		return err
	}
	consortiumRoot := filepath.Join(utils.BlockchainRoot, "consortiums", consortium)
	utils.MkdirIfNotExists(consortiumRoot)
	err = ioutil.WriteFile(filepath.Join(consortiumRoot, "configtx.yaml"), bytes, os.ModePerm)
	if err != nil {
		logrus.Errorf("write file error:%v", err)
		return err
	}
	logrus.Infof("Generating channel config tx for channel:%s", channel)
	txPath := filepath.Join(consortiumRoot, "channel-artifacts", channel+".tx")
	cmdStr := fmt.Sprintf("configtxgen -profile %s -channelID %s -outputCreateChannelTx %s -configPath %s", channel, channel, txPath, consortiumRoot)
	utils.ExecLocalCommand(cmdStr)
	// rename channel config file to reserve configtx.yaml
	utils.RenameFile(consortiumRoot, "configtx.yaml", channel+".yaml")
	return nil
}
