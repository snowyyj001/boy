//小游戏管理服务
package room

import (
	"game/dbmodel"

	"game/room/niuniu"

	"github.com/snowyyj001/loumiao"
	"github.com/snowyyj001/loumiao/gorpc"
	"github.com/snowyyj001/loumiao/log"
)

var (
	This *RoomServer
)

type RoomServer struct {
	gorpc.GoRoutineLogic

	Rooms map[int]string
}

func (self *RoomServer) DoInit() {
	log.Info("RoomServer DoInit")
	This = self

	self.Rooms = make(map[int]string)
}

func (self *RoomServer) DoRegsiter() {

}

func (self *RoomServer) DoStart() {
	log.Info("RoomServer DoStart")
	cfgs := self.Call("DBServer", "getGameCfg", nil).([]dbmodel.GameCfg)
	for _, cfg := range cfgs {
		if cfg.ServiceName == "niuniu" {
			self.Rooms[cfg.GameId] = cfg.ServiceName
			gameServer := new(niuniu.GameServer)
			gameServer.GameId = cfg.GameId
			gameServer.GameRule = cfg.GameRule
			gameServer.RoomNum = cfg.RoomNumber
			loumiao.Prepare(gameServer, cfg.ServiceName, false)
		}
	}

}

func (self *RoomServer) DoDestory() {
	log.Info("RoomServer destory")
}
