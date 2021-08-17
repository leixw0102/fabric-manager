package command

import (
	"Data_Bank/fabric-manager/agent/config"
	"Data_Bank/fabric-manager/common/message"
	"Data_Bank/fabric-manager/common/utils"
	"fmt"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

var OrgAdminEnrolled map[string]bool

func init() {
	OrgAdminEnrolled = make(map[string]bool)
}

func Enroll(id message.Identity, domain, caAddress, caName, tlsCertPathAbs, mspDirAbs string, isTLS bool) error {
	cmdStr := ""
	if isTLS {
		logrus.Infof("enroll %s tls \n", id.Name)
		cmdStr = fmt.Sprintf("fabric-ca-client enroll -d -u https://%s:%s@%s --caname %s --csr.hosts %s --tls.certfiles %s --enrollment.profile tls -M %s", id.Name, id.Password, caAddress, caName, domain, tlsCertPathAbs, mspDirAbs)
	} else {
		logrus.Infof("enroll %s msp \n", id.Name)
		cmdStr = fmt.Sprintf("fabric-ca-client enroll -d -u https://%s:%s@%s --caname %s --csr.hosts %s --tls.certfiles %s -M %s", id.Name, id.Password, caAddress, caName, domain, tlsCertPathAbs, mspDirAbs)
	}
	return utils.ExecLocalCommand(cmdStr)
}

func EnrollAdmin(caAddress, caName, tlsCertPathAbs, homeDir string) error {
	cmdStr := fmt.Sprintf("fabric-ca-client enroll -u https://admin:adminpw@%s --caname %s --tls.certfiles %s --home %s", caAddress, caName, tlsCertPathAbs, homeDir)
	return utils.ExecLocalCommand(cmdStr)
}

// fabric-ca-client register --caname ca-org1 --id.name peer0 --id.secret peer0pw --id.type peer --tls.certfiles "${PWD}/organizations/fabric-ca/org1/tls-cert.pem"
func Register(id message.Identity, caName, tlsCertPathAbs, homeDir string) {
	logrus.Infof("register %s", id.Name)
	cmdStr := fmt.Sprintf("fabric-ca-client register --caname %s --id.name %s --id.secret %s --id.type %s --tls.certfiles %s --home %s", caName, id.Name, id.Password, id.IdType, tlsCertPathAbs, homeDir)
	utils.ExecLocalCommand(cmdStr)
}

func CreateIdentity(orgDomain, caName, caAddress string, identity message.Identity) error {
	orgRoot := filepath.Join(utils.BlockchainRoot, "organizations", orgDomain, "crypto")
	// generate org level crypto materials
	orgMspDir := filepath.Join(orgRoot, "msp")
	orgMspTlsDir := filepath.Join(orgMspDir, "tlscacerts")
	orgTlsDir := filepath.Join(orgRoot, "tlsca")
	orgCaDir := filepath.Join(orgRoot, "ca")
	utils.MkdirIfNotExists(orgMspDir)
	utils.MkdirIfNotExists(orgMspTlsDir)
	utils.MkdirIfNotExists(orgTlsDir)
	utils.MkdirIfNotExists(orgCaDir)

	// Org admin only enroll once
	if _, ok := OrgAdminEnrolled[orgDomain]; !ok {
		EnrollAdmin(caAddress, caName, utils.CACertPath, orgRoot)
		OrgAdminEnrolled[orgDomain] = true
	}

	// generate node ou config
	config.GenNodeOUConfig(caAddress, caName, filepath.Join(orgMspDir, "config.yaml"))
	// generate peer level crypto materials
	logrus.Infof("Creating identity:%s", identity.Name)
	Register(identity, caName, utils.CACertPath, orgRoot)
	role := "users"
	tlsDir := ""
	mspDir := filepath.Join(orgRoot, role, identity.Name+"@"+orgDomain, "msp")
	// generate tls certs for peer and orderer
	if identity.IdType == "peer" || identity.IdType == "orderer" {
		role = identity.IdType + "s" // peers, orderers
		domain := identity.Name + "." + orgDomain
		tlsDir = filepath.Join(orgRoot, role, domain, "tls")
		utils.MkdirIfNotExists(tlsDir)
		mspDir = filepath.Join(orgRoot, role, domain, "msp")
		Enroll(identity, domain, caAddress, caName, utils.CACertPath, tlsDir, true)
	}
	// generate msp for all identities
	utils.MkdirIfNotExists(mspDir)
	Enroll(identity, "", caAddress, caName, utils.CACertPath, mspDir, false)
	// generate nodeou config for each msp
	config.GenNodeOUConfig(caAddress, caName, filepath.Join(mspDir, "config.yaml"))

	if tlsDir != "" {
		// copy tls crypto materails to peer/orderer's tls dir, this is for easily mount of peer/orderer's tls crytpo material to docker volume
		utils.Copy(filepath.Join(tlsDir, "tlscacerts", "*"), filepath.Join(tlsDir, "ca.crt"))
		utils.Copy(filepath.Join(tlsDir, "signcerts", "*"), filepath.Join(tlsDir, "server.crt"))
		utils.Copy(filepath.Join(tlsDir, "keystore", "*"), filepath.Join(tlsDir, "server.key"))

		// for orderer, orderer's msp also need to have tlscacerts dir
		if identity.IdType == "orderer" {
			utils.MkdirIfNotExists(filepath.Join(mspDir, "tlscacerts"))
			utils.Copy(filepath.Join(tlsDir, "tlscacerts", "*"), filepath.Join(mspDir, "tlscacerts"))
		}
		// copy tls and ca certs to org msp and tls
		utils.Copy(filepath.Join(tlsDir, "tlscacerts", "*"), filepath.Join(orgMspTlsDir, "ca.crt"))
		utils.Copy(filepath.Join(tlsDir, "tlscacerts", "*"), filepath.Join(orgTlsDir, fmt.Sprintf("tlsca.%s-cert.pem", orgDomain)))
		utils.Copy(filepath.Join(mspDir, "cacerts", "*"), filepath.Join(orgCaDir, fmt.Sprintf("ca.%s-cert.pem", orgDomain)))
	}

	return nil
}
