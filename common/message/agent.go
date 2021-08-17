package message

import "time"

type AgentHeartbeat struct {
	AgentID   string
	IP        string
	Heartbeat time.Time
}
