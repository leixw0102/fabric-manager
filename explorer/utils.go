package main

import (
	"Data_Bank/fabric-manager/common/mock"
	"Data_Bank/fabric-manager/common/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

func DownloadCertsFromDB(orgName, downloadPath string) error {
	// get org's admin private key and cert
	org := mock.GetOrgInfo(orgName)
	admin := org.Admin.Name
	pk := GetPriviateKey(admin, orgName)
	cert := GetSignCert(admin, orgName)
	tlscacert := GetTLSCACert(org.Peers[0].Name, orgName)
	// write admin private key to local filesystem
	pkPath := filepath.Join(downloadPath, "orgazniations", orgName, "admin", "private_key")
	if err := ioutil.WriteFile(pkPath, []byte(pk), os.ModePerm); err != nil {
		logrus.Errorf("Fail to download %s admin private key, error:%v", orgName, err)
		return err
	}
	// write admin sign certs to local filesystem
	certPath := filepath.Join(downloadPath, "organizations", orgName, "admin", "cert.pem")
	if err := ioutil.WriteFile(certPath, []byte(cert), os.ModePerm); err != nil {
		logrus.Errorf("Fail to download %s admin sign cert, error:%v", orgName, err)
		return err
	}
	// write peer tls cacerts to local filesystem
	peerTLSCertPath := filepath.Join(downloadPath, "organizations", orgName, "peer", "tlsca.crt")
	if err := ioutil.WriteFile(peerTLSCertPath, []byte(tlscacert), os.ModePerm); err != nil {
		logrus.Errorf("Fail to download %s peer tls ca cert, error:%v", orgName, err)
		return err
	}

	return nil
}

func GetPriviateKey(name, org string) string {
	return ""
}

func GetSignCert(name, org string) string {
	return ""
}

func GetTLSCACert(name, org string) string {
	return ""
}

func NewExplorerService(extraHosts []string, consortium, config, profile, crypto, output string) error {
	var explorerTemp = map[string]interface{}{
		"version": "2.1",
		"volumes": map[string]interface{}{
			"pgdata":      nil,
			"walletstore": nil,
		},
		"services": map[string]interface{}{
			"explorerdb.mynetwork.com": map[string]interface{}{
				"image":          "hyperledger/explorer-db:1.1.4",
				"container_name": "explorerdb.mynetwork.com",
				"hostname":       "explorerdb.mynetwork.com",
				"environment": []string{
					"DATABASE_DATABASE=fabricexplorer",
					"DATABASE_USERNAME=hppoc",
					"DATABASE_PASSWORD=password",
				},
				"healthcheck": map[string]interface{}{
					"test":     "pg_isready -h localhost -p 5432 -q -U postgres",
					"interval": "30s",
					"timeout":  "10s",
					"retries":  5,
				},
				"extra_hosts": extraHosts,
				"volumes": []string{
					"pgdata:/var/lib/postgresql/data",
				},
			},
			"explorer.mynetwork.com": map[string]interface{}{
				"image":          "hyperledger/explorer:1.1.4",
				"container_name": "explorer.mynetwork.com",
				"hostname":       "explorer.mynetwork.com",
				"depends_on": map[string]interface{}{
					"explorerdb.mynetwork.com": map[string]interface{}{
						"condition": "service_healthy",
					},
				},
				"environment": []string{
					"DATABASE_HOST=explorerdb.mynetwork.com",
					"DATABASE_DATABASE=fabricexplorer",
					"DATABASE_USERNAME=hppoc",
					"DATABASE_PASSWD=password",
					"LOG_LEVEL_APP=debug",
					"LOG_LEVEL_DB=debug",
					"LOG_LEVEL_CONSOLE=info",
					"LOG_CONSOLE_STDOUT=true",
					"DISCOVERY_AS_LOCALHOST=false",
				},
				"volumes": []string{
					config + ":/opt/explorer/app/platform/fabric/config.json",
					profile + ":/opt/explorer/app/platform/fabric/" + consortium + ".json",
					"walletstore:/opt/explorer/wallet",
					fmt.Sprintf("%s:/tmp/crypto", crypto),
				},
				"extra_hosts": extraHosts,
				"ports": []string{
					"8080:8080",
				},
			},
		},
	}
	bytes := utils.ToYaml(explorerTemp)
	filename := filepath.Join(output, "docker-compose.yaml")
	return ioutil.WriteFile(filename, bytes, os.ModePerm)
}

func NewExplorerConfig(consortium, profilePath, configPath string) error {
	configTemp := map[string]interface{}{
		"network-configs": map[string]interface{}{
			"test-network": map[string]string{
				"name":    consortium + " " + "Network",
				"profile": profilePath,
			},
		},
		"license": "Apache-2.0",
	}
	bytes, _ := json.Marshal(configTemp)
	return ioutil.WriteFile(configPath, bytes, os.ModePerm)
}

func NewExplorerProfile(fileAbsPath, adminPrivateKey, adminCert, peerCert, peerUrl string) error {
	profileTemp := map[string]interface{}{
		"name":    "test-network",
		"version": "1.0.0",
		"client": map[string]interface{}{
			"tlsEnable": true,
			"adminCredential": map[string]string{
				"id":       "exploreradmin",
				"password": "exploreradminpw",
			},
			"enableAuthentication": true,
			"organization":         "Org1MSP",
			"connection": map[string]interface{}{
				"timeout": map[string]interface{}{
					"peer": map[string]string{
						"endorser": "300",
					},
					"orderer": "300",
				},
			},
		},
		"channels": map[string]interface{}{
			"mychannel": map[string]interface{}{
				"peers": map[string]interface{}{
					"peer0.org1.example.com": map[string]interface{}{},
				},
			},
		},
		"organizations": map[string]interface{}{
			"Org1MSP": map[string]interface{}{
				"mspid": "Org1MSP",
				"adminPrivateKey": map[string]string{
					"path": adminPrivateKey,
				},
				"peers": []string{"peer0.org1.example.com"},
				"signedCert": map[string]string{
					"path": adminCert,
				},
			},
		},
		"peers": map[string]interface{}{
			"peer0.org1.example.com": map[string]interface{}{
				"tlsCACerts": map[string]string{
					"path": peerCert,
				},
				"url": "grpcs://" + peerUrl,
			},
		},
	}
	bytes, _ := json.Marshal(profileTemp)
	return ioutil.WriteFile(fileAbsPath, bytes, os.ModePerm)
}
