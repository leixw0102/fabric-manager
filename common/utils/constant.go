package utils

// cert related
const (
	BlockchainRoot = "/root/fabric_networks" // the root dir for all fabric consortium chain, fixed for every agent
	// there's only 1 CA for all consortium chains and organizations, the ca admin username and password are fixed
	CAAdminUsername = "admin"
	CAAdminPassword = "adminpw"
	CACertPath      = BlockchainRoot + "/fabric-ca/root-cert/tls-cert.pem"
)

// action related
const (
	CreateIdentity   = "CreateIdentity"
	CreateCrypto     = "CreateCrypto"
	CreateConsortium = "CreateConsortium"
	CreateChannel    = "CreateChannel"
	StartNetwork     = "StartNetwork"
	StartPeer        = "StartPeer"
	StartOrderer     = "StartOrderer"
)

// service
const (
	AgentService  = "fabric-manager/agent"
	ServerService = "fabric-manager/server"
)

const (
	CertDownloadPath = "/fabric-manager/certs/%s"
)
const (
	CertCreateCmd = "CertCreate"
	CertDownload  = "CertDownload"
)
