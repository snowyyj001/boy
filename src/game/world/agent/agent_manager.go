package agent

import (
	"log"
	"sync"
)

type AgentMgr struct {
	agents     map[int]*Agent
	acc_id_Map map[string]int
	sid_id_Map map[int]int
}

var (
	inst *AgentMgr
	once sync.Once
)

func GetAgentMgr() *AgentMgr {
	once.Do(func() {
		inst = &AgentMgr{}
		inst.agents = make(map[int]*Agent)
		inst.acc_id_Map = make(map[string]int)
		inst.sid_id_Map = make(map[int]int)
	})
	return inst
}

func (self *AgentMgr) AddAgent(agent *Agent) {
	if self.agents[agent.ID] != nil {
		log.Fatalf("AgentMgr.AddAgent error,[%d] has already added", agent.ID)
		return
	}
	if agent.ID <= 0 || agent.ClientId <= 0 || agent.Account == "" {
		log.Fatalf("AgentMgr.AddAgent error,[%d][%s][%s] not leagl", agent.ID, agent.ClientId, agent.Account)
		return
	}
	self.agents[agent.ID] = agent
	self.acc_id_Map[agent.Account] = agent.ID
	self.sid_id_Map[agent.ClientId] = agent.ID
}

func (self AgentMgr) RemoveAgent(userid int) {
	agent := self.GetAgent(userid)
	if agent == nil {
		return
	}
	delete(self.acc_id_Map, agent.Account)
	delete(self.sid_id_Map, agent.ClientId)
	delete(self.agents, userid)
}

func (self *AgentMgr) GetAgent(userid int) *Agent {
	return self.agents[userid]
}

func (self *AgentMgr) GetAgentByAccount(accName string) *Agent {
	userid, has := self.acc_id_Map[accName]
	if has {
		return self.agents[userid]
	}
	return nil
}

func (self *AgentMgr) GetAgentByServerId(clientId int) *Agent {
	userid, has := self.sid_id_Map[clientId]
	if has {
		return self.agents[userid]
	}
	return nil
}
