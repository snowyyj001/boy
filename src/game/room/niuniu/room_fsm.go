package niuniu

import (
	"game/msg"

	"github.com/snowyyj001/loumiao/util"
)

func (self *Room) SetState(state FSMState) {
	self.nextState = state
}

func (self *Room) IsState(state FSMState) bool {
	//self.curState==self.nextState,状态机切换期间，不属于任何状态
	//这时不应接受任何输入
	if self.curState == state && self.nextState == self.curState {
		return true
	} else {
		return false
	}
}

func (self *Room) changeState() {
	switch self.nextState {
	case FSM_Idle:
		self.curFsmFunc[0] = self.onEnterIdle
		self.curFsmFunc[1] = self.onExecIdle
		self.curFsmFunc[2] = self.onExitIdle
	case FSM_Fapai:
		self.curFsmFunc[0] = self.onEnterFaPai
		self.curFsmFunc[1] = self.onExecFaPai
		self.curFsmFunc[2] = self.onExitFaPai
	case FSM_Bipai:
		self.curFsmFunc[0] = self.onEnterBiPai
		self.curFsmFunc[1] = self.onExecBiPai
		self.curFsmFunc[2] = self.onExitBiPai
	case FSM_Result:
		self.curFsmFunc[0] = self.onEnterResult
		self.curFsmFunc[1] = self.onExecResult
		self.curFsmFunc[2] = self.onExitResult
	default:

	}
}

func (self *Room) update(dt int64) {
	if self.curState == self.nextState {
		self.curFsmFunc[1](dt)
	} else {
		self.curFsmFunc[2](dt)
		self.changeState()
		self.curFsmFunc[0](dt)
		self.curFsmFunc[1](dt)
	}
}

func (self *Room) onEnterIdle(dt int64) {
	self.fsmTime = Time_Idle
	l_shuffle(self.Cards)
}

func (self *Room) onExecIdle(dt int64) {
	if self.playerLen >= MIN_GAMER {
		self.SetState(FSM_Fapai)
	}
}

func (self *Room) onExitIdle(dt int64) {

}

func (self *Room) onEnterFaPai(dt int64) {
	self.fsmTime = Time_FaPai
	start := 0

	for i := 0; i < self.playerLen; i++ {
		self.players[i].handCards = self.Cards[start : start+HANDCARD_NUM-1]
		start += HANDCARD_NUM - 1
		l_sort(self.players[i].handCards, false)

		req := &msg.NN_FaPai{}
		req.Seat = self.players[i].seat
		req.Cards = self.players[i].handCards
		self.SendClient(self.players[i], req)
	}

}

func (self *Room) onExecFaPai(dt int64) {
	self.fsmTime -= dt
	if self.fsmTime <= 0 {
		self.SetState(FSM_Bipai)
	}
}

func (self *Room) onExitFaPai(dt int64) {
	self.ResultCard = util.Random(10) + 1
	req := &msg.NN_LuckNumber{}
	req.Number = self.ResultCard
	self.SendAllAgents(req)
}

func (self *Room) onEnterBiPai(dt int64) {
	self.fsmTime = Time_BiPai
	for i := 0; i < self.playerLen; i++ {
		self.players[i].handCards[HANDCARD_NUM-1] = self.ResultCard
	}
}

func (self *Room) onExecBiPai(dt int64) {
	self.fsmTime -= dt
	if self.fsmTime <= 0 {
		self.SetState(FSM_Bipai)
	}
}

func (self *Room) onExitBiPai(dt int64) {

}

func (self *Room) onEnterResult(dt int64) {
	self.fsmTime = Time_Result
}

func (self *Room) onExecResult(dt int64) {
	self.fsmTime -= dt
	if self.fsmTime <= 0 {
		self.SetState(FSM_Idle)
	}
}

func (self *Room) onExitResult(dt int64) {

}
