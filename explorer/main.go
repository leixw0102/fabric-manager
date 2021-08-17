package main

import (
	"Data_Bank/fabric-manager/common/connector"
	"Data_Bank/fabric-manager/common/mock"
	"Data_Bank/fabric-manager/common/utils"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

var consortium, orgName, downloadPath, db, username, password, host string

func main() {
	pflag.StringVar(&consortium, "consortium", "SampleConsortium", "consortium name")
	pflag.StringVar(&orgName, "orgname", "org1", "the name of an org")
	pflag.StringVar(&downloadPath, "download", "/root/explorer", "root for downloading org admin cert, private key and create config.json and test-network.json")
	pflag.StringVar(&db, "db", "defaultdb", "mysql db name that cli needs to connect to")
	pflag.StringVar(&username, "username", "root", "username of mysql")
	pflag.StringVar(&password, "password", "ehl1234", "password of mysql")
	pflag.StringVar(&host, "host", "192.168.133.130:3306", "ip/domain:port of mysql")
	help := pflag.BoolP("help", "h", false, "print help message")
	pflag.Parse()
	if *help {
		pflag.PrintDefaults()
		return
	}
	// connect to mysql
	dbConfig := &connector.DBConfig{
		Username: username,
		Password: password,
		Hostname: host,
		DBName:   db,
	}
	connector.DB = connector.NewMysql(dbConfig)
	// download org certs from mysql
	if err := DownloadCertsFromDB(orgName, downloadPath); err != nil {
		logrus.Error("Fail to download cert")
		// os.Exit(-1)
	}
	// construct test-network.json
	cryptoAbs := filepath.Join(downloadPath, "organizations", orgName)
	cryptoDockerAbs := "/tmp/crypto"
	adminPrivateKey := filepath.Join(cryptoDockerAbs, "admin", "private_key")
	adminCert := filepath.Join(cryptoDockerAbs, "admin", "cert.pem")
	peerCert := filepath.Join(cryptoDockerAbs, "peer", "tlsca.crt")
	peer := mock.GetOrgInfo(orgName).GetPeerInfo(0)
	peerUrl := fmt.Sprintf("%s:%d", peer.Domain, peer.Port)
	profileAbs := filepath.Join(downloadPath, consortium+".json")
	profileDockerAbs := filepath.Join("/opt/explorer/app/platform/fabric/", consortium+".json")

	if err := NewExplorerProfile(profileAbs, adminPrivateKey, adminCert, peerCert, peerUrl); err != nil {
		logrus.Errorf("Fail to create %s.json, error:%v", consortium, err)
		os.Exit(-1)
	}
	// construct config.json
	configAbs := filepath.Join(downloadPath, "config.json")
	if err := NewExplorerConfig(consortium, profileDockerAbs, configAbs); err != nil {
		logrus.Errorf("Fail to create config.json, error:%v", err)
		os.Exit(-1)
	}

	// construct docker-compose.yaml
	extraHosts := mock.GetConsortiumHosts(consortium)
	if err := NewExplorerService(extraHosts, consortium, configAbs, profileAbs, cryptoAbs, downloadPath); err != nil {
		logrus.Errorln("Fail to create explorer docker-compose.yaml")
		os.Exit(-1)
	}

	// run docker-compose to start docker container
	dockerComposePath := filepath.Join(downloadPath, "docker-compose.yaml")
	cmdStr := fmt.Sprintf("docker-compose -f %s up -d", dockerComposePath)
	if err := utils.ExecLocalCommand(cmdStr); err != nil {
		logrus.Errorf("Fail to execute docker-compose up command, error:%v", err)
		os.Exit(-1)
	}
	logrus.Info("Explorer started successfully.")
}
