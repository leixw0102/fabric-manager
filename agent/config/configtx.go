package config

import (
	"Data_Bank/fabric-manager/common/message"
	"Data_Bank/fabric-manager/common/utils"
	"fmt"
	"path/filepath"
)

/*configtx.yaml 结构定义*/

// TxConfig consists of the structs used by the configtxgen tool.
type TxConfig struct {
	Profiles      map[string]*Profile        `yaml:"Profiles"`
	Organizations []*Organization            `yaml:"-"`
	Channel       Channel                    `yaml:"-"`
	Application   *Application               `yaml:"-"`
	Orderer       *Orderer                   `yaml:"-"`
	Capabilities  map[string]map[string]bool `yaml:"-"`
}

// Profile encodes orderer/application configuration combinations for the
// configtxgen tool.
type Profile struct {
	Consortium   string                 `yaml:"Consortium,omitempty"`
	Application  *Application           `yaml:"Application,omitempty"`
	Orderer      *Orderer               `yaml:"Orderer,omitempty"`
	Consortiums  map[string]*Consortium `yaml:"Consortiums,omitempty"`
	Capabilities map[string]bool        `yaml:"Capabilities,omitempty"`
	Policies     map[string]*Policy     `yaml:"Policies,omitempty"`
}

// Policy encodes a channel config policy
type Policy struct {
	Type string `yaml:"Type"`
	Rule string `yaml:"Rule"`
}

type Channel struct {
	Policies     map[string]*Policy `yaml:"Policies"`
	Capabilities map[string]bool    `yaml:"Capabilities"`
}

// Consortium represents a group of organizations which may create channels
// with each other
type Consortium struct {
	Organizations []*Organization `yaml:"Organizations"`
}

// Application encodes the application-level configuration needed in config
// transactions.
type Application struct {
	Organizations []*Organization    `yaml:"Organizations"`
	Capabilities  map[string]bool    `yaml:"Capabilities"`
	Policies      map[string]*Policy `yaml:"Policies"`
	ACLs          map[string]string  `yaml:"ACLs"`
}

// Organization encodes the organization-level configuration needed in
// config transactions.
type Organization struct {
	Name        string             `yaml:"Name"`
	ID          string             `yaml:"ID"`
	MSPDir      string             `yaml:"MSPDir"`
	Policies    map[string]*Policy `yaml:"Policies"`
	AnchorPeers []AnchorPeer       `yaml:"AnchorPeers,omitempty"`
}

// AnchorPeer encodes the necessary fields to identify an anchor peer.
type AnchorPeer struct {
	Host string `yaml:"Host"`
	Port int    `yaml:"Port"`
}

// Orderer contains configuration associated to a channel.
type Orderer struct {
	OrdererType   string             `yaml:"OrdererType"`
	Addresses     []string           `yaml:"Addresses"`
	BatchTimeout  string             `yaml:"BatchTimeout"`
	BatchSize     BatchSize          `yaml:"BatchSize"`
	EtcdRaft      EtcdRaft           `yaml:"EtcdRaft"`
	Organizations []*Organization    `yaml:"Organizations"`
	MaxChannels   uint64             `yaml:"MaxChannels"`
	Capabilities  map[string]bool    `yaml:"Capabilities"`
	Policies      map[string]*Policy `yaml:"Policies"`
}

// BatchSize contains configuration affecting the size of batches.
type BatchSize struct {
	MaxMessageCount   uint32 `yaml:"MaxMessageCount"`
	AbsoluteMaxBytes  string `yaml:"AbsoluteMaxBytes"`
	PreferredMaxBytes string `yaml:"PreferredMaxBytes"`
}

type Consenter struct {
	Host          string `yaml:"Host"`
	Port          int    `yaml:"Port"`
	ClientTLSCert string `yaml:"ClientTLSCert"`
	ServerTLSCert string `yaml:"ServerTLSCert"`
}
type EtcdRaft struct {
	Consenters []Consenter `yaml:"Consenters"`
}

/*如何生成configtx.yaml的函数*/
// 生成创世区块配置文件
func GenGenesisConfigtx(consortiumName string, orgInfos []*message.OrgInfo) *TxConfig {
	consenters := make([]Consenter, 0)
	addresses := make([]string, 0)

	ordererOrgs := make([]*message.OrgInfo, 0)
	peerOrgs := make([]*message.OrgInfo, 0)
	for _, org := range orgInfos {
		if org.Orderers != nil {
			ordererOrgs = append(ordererOrgs, org)
		}
		if org.Peers != nil {
			peerOrgs = append(peerOrgs, org)
		}
	}

	// form orderer section of configtx.yaml, consenter and address of orderers are needed
	for _, org := range ordererOrgs {
		for _, orderer := range org.Orderers {
			consenters = append(consenters, Consenter{
				Host:          orderer.Domain,
				Port:          orderer.Port,
				ClientTLSCert: orderer.Cert,
				ServerTLSCert: orderer.Cert,
			})
			addresses = append(addresses, orderer.Address())
		}
	}
	return &TxConfig{
		Profiles: map[string]*Profile{
			"GenesisChannel": {
				Orderer: GenOrdererSection(addresses, consenters, OrgInfoToOrgTx(ordererOrgs)),
				Consortiums: map[string]*Consortium{
					consortiumName: {
						Organizations: OrgInfoToOrgTx(peerOrgs),
					},
				},
				Policies: map[string]*Policy{
					"Readers": {Type: "ImplicitMeta", Rule: "ANY Readers"},
					"Writers": {Type: "ImplicitMeta", Rule: "ANY Writers"},
					"Admins":  {Type: "ImplicitMeta", Rule: "MAJORITY Admins"},
				},
				Capabilities: map[string]bool{"V2_0": true},
			},
		},
	}
}

func GenOrdererSection(ordererAddresses []string, consenters []Consenter, orgs []*Organization) *Orderer {
	return &Orderer{
		OrdererType:  "etcdraft",
		Addresses:    ordererAddresses,
		BatchTimeout: "2s",
		BatchSize: BatchSize{
			MaxMessageCount:   10,
			AbsoluteMaxBytes:  "99 MB",
			PreferredMaxBytes: "512 KB",
		},
		MaxChannels:  10,
		Capabilities: map[string]bool{"V2_0": true},
		Policies: map[string]*Policy{
			"Readers":         {Type: "ImplicitMeta", Rule: "ANY Readers"},
			"Writers":         {Type: "ImplicitMeta", Rule: "ANY Writers"},
			"Admins":          {Type: "ImplicitMeta", Rule: "MAJORITY Admins"},
			"BlockValidation": {Type: "ImplicitMeta", Rule: "ANY Writers"},
		},
		EtcdRaft: EtcdRaft{
			Consenters: consenters,
		},
		Organizations: orgs,
	}
}

func GenChannelConfigtx(consortiumName, channelName string, orgInfos []*message.OrgInfo) *TxConfig {
	return &TxConfig{
		Profiles: map[string]*Profile{
			channelName: {
				Consortium: consortiumName,
				Application: &Application{
					Organizations: OrgInfoToOrgTx(orgInfos),
					Capabilities:  map[string]bool{"V2_0": true},
					Policies: map[string]*Policy{
						"Readers":              {Type: "ImplicitMeta", Rule: "ANY Readers"},
						"Writers":              {Type: "ImplicitMeta", Rule: "ANY Writers"},
						"Admins":               {Type: "ImplicitMeta", Rule: "MAJORITY Admins"},
						"LifecycleEndorsement": {Type: "ImplicitMeta", Rule: "MAJORITY Endorsement"},
						"Endorsement":          {Type: "ImplicitMeta", Rule: "MAJORITY Endorsement"},
					},
				},
				Policies: map[string]*Policy{
					"Readers": {Type: "ImplicitMeta", Rule: "ANY Readers"},
					"Writers": {Type: "ImplicitMeta", Rule: "ANY Writers"},
					"Admins":  {Type: "ImplicitMeta", Rule: "MAJORITY Admins"},
				},
				Capabilities: map[string]bool{"V2_0": true},
			},
		},
	}
}

func OrgInfoToOrgTx(orgInfos []*message.OrgInfo) []*Organization {
	orgs := make([]*Organization, 0)
	for _, org := range orgInfos {
		MSPID := org.Name + "MSP"
		var mspDir string
		var policies map[string]*Policy
		if org.Orderers == nil {
			// peer org
			mspDir = filepath.Join(utils.BlockchainRoot, "organizations", "crypto", "organizations", "peerOrganizations", org.Domain, "msp")
			policies = map[string]*Policy{
				"Readers":     {Type: "Signature", Rule: fmt.Sprintf("OR('%s.admin','%s.peer','%s.client')", MSPID, MSPID, MSPID)},
				"Writers":     {Type: "Signature", Rule: fmt.Sprintf("OR('%s.admin','%s.client')", MSPID, MSPID)},
				"Admins":      {Type: "Signature", Rule: fmt.Sprintf("OR('%s.admin')", MSPID)},
				"Endorsement": {Type: "Signature", Rule: fmt.Sprintf("OR('%s.peer')", MSPID)},
			}
		} else {
			// orderer org
			mspDir = filepath.Join(utils.BlockchainRoot, "organizations", "crypto", "organizations", "ordererOrganizations", org.Domain, "msp")
			policies = map[string]*Policy{
				"Readers": {Type: "Signature", Rule: fmt.Sprintf("OR('%s.member')", MSPID)},
				"Writers": {Type: "Signature", Rule: fmt.Sprintf("OR('%s.member')", MSPID)},
				"Admins":  {Type: "Signature", Rule: fmt.Sprintf("OR('%s.admin')", MSPID)},
			}
		}
		orgs = append(orgs, &Organization{
			Name:     org.Name,
			ID:       MSPID,
			MSPDir:   mspDir,
			Policies: policies,
		})
	}
	return orgs
}
