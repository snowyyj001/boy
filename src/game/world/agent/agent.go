package agent

import (
	"game/dbmodel"

	"github.com/snowyyj001/loumiao/util"
)

type Agent struct {
	dbmodel.User
	ClientId int
	RoomId   int
}

func (self *Agent) OnLogin() {
	self.LoginTime = util.TimeStamp()
}

func (self *Agent) OnLoginOut() {

}

func (self *Agent) SendClient() {

}
