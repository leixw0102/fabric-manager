package monitor

import (
	"Data_Bank/fabric-manager/common/message"
	"time"

	"github.com/sirupsen/logrus"
)

var agentMonitor AgentMonitor

func NewAgentMonitor() AgentMonitor {
	return AgentMonitor{
		agents: make(map[string]*message.AgentHeartbeat),
	}
}

type AgentMonitor struct {
	agents map[string]*message.AgentHeartbeat
}

func (m *AgentMonitor) OnPing(id string) {
	now := time.Now()
	if a, ok := m.agents[id]; ok {
		a.Heartbeat = now
	} else {
		m.agents[id] = &message.AgentHeartbeat{AgentID: id, Heartbeat: now}
	}
}

func (m *AgentMonitor) MonitorAgentHeartbeat(timeout time.Duration) {
	for {
		now := time.Now()
		for id, a := range m.agents {
			if a.Heartbeat.Add(timeout).Before(now) {
				logrus.Errorf("agent:%s lost connection", id)
				delete(m.agents, id)
			}
		}
		time.Sleep(timeout)
	}
}

func (m *AgentMonitor) HasIP(IP string) bool {
	for _, a := range m.agents {
		if a.IP == IP {
			return true
		}
	}
	return false
}
