package command

import (
	"Data_Bank/fabric-manager/agent/config"
	"Data_Bank/fabric-manager/common/message"
	"Data_Bank/fabric-manager/common/utils"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// 生成创世区块分两步（采用sdk的方法暂告失败）：
// 1. 生成configtx.yaml
// 2. 用configtxgen生成genesis.block
func CreateGenesisBlock(consortiumName string, orgInfos []*message.OrgInfo) {
	logrus.Info("Generating genesis config ...")
	genesisCfg := config.GenGenesisConfigtx(consortiumName, orgInfos)
	bytes, _ := yaml.Marshal(genesisCfg)
	consortiumRoot := filepath.Join(utils.BlockchainRoot, "consortiums", consortiumName)
	utils.MkdirIfNotExists(consortiumRoot)
	utils.MkdirIfNotExists(filepath.Join(consortiumRoot, "system-genesis-block"))
	err := ioutil.WriteFile(filepath.Join(consortiumRoot, "configtx.yaml"), bytes, os.ModePerm)
	if err != nil {
		logrus.Errorf("error:%v", err)
	}
	logrus.Info("Generating genesis block using genesis config ...")
	cmdStr := fmt.Sprintf("configtxgen -profile GenesisChannel -channelID system-channel -outputBlock %s/system-genesis-block/genesis.block -configPath %s", consortiumRoot, consortiumRoot)
	utils.ExecLocalCommand(cmdStr)
	// rename channel config file to reserve configtx.yaml
	utils.RenameFile(consortiumRoot, "configtx.yaml", "genesis.yaml")
}
