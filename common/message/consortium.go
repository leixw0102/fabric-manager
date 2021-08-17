package message

// front end message
type CreateConsortiumMsg struct {
	Name        string   `json:"name" binding:"required"`
	OrgDomains  []string `json:"orgs" binding:"required"`     // all org domains that're included in the consortium
	InitiateOrg string   `json:"init_org" binding:"required"` //domain
	Orderer     string   `json:"orderer" binding:"required"`  // orderer name that
}

type StartConsortiumMsg struct {
	Name string `json:"name" binding:"required"`
}

type StartServiceMsg struct {
	Consortium string
	Domain     string
}

// mysql queried message
type ConsortiumInfo struct {
	Name string
	Orgs []*OrgInfo
}
