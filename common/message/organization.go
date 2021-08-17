package message

type CreateOrgMsg struct {
	OrgDomain  string     `json:"org" binding:"required"`
	CaName     string     `json:"ca" binding:"required"`
	CaAddress  string     `json:"ca_addr" binding:"required"`
	Identities []Identity `json:"identities" binding:"required"`
}

type Identity struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	IdType   string `json:"idtype" binding:"required"`
	IP       string `json:"ip" binding:"required"`
}

type CreateIdentityMsg struct {
	OrgDomain string   `json:"org" binding:"required"`
	CaName    string   `json:"ca" binding:"required"`
	CaAddress string   `json:"ca_addr" binding:"required"`
	Identity  Identity `json:"identity" binding:"required"`
}

type OrgInfo struct {
	Name     string
	Domain   string
	Orderers []*OrdererInfo
	Peers    []*PeerInfo
	Users    []*UserInfo
	Admin    *UserInfo
}

func (o *OrgInfo) GetPeerInfo(i int) *PeerInfo {
	if len(o.Peers) <= i {
		return nil
	}
	return o.Peers[i]
}

func (o *OrgInfo) GetOrdererInfo(ordererName string) *OrdererInfo {
	for _, orderer := range o.Orderers {
		if orderer.Name == ordererName {
			return orderer
		}
	}
	return nil
}

func (o *OrgInfo) GetAdminInfo() *UserInfo {
	return o.Admin
}
