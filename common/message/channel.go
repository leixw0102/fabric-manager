package message

type CreateChannelMsg struct {
	Consortium string   `json:"consortium" binding:"required"`
	Channel    string   `json:"channel" binding:"required"`
	Orgs       []string `json:"orgs" binding:"required"`
	InitOrg    string   `json:"init_org" binding:"required"`
}
