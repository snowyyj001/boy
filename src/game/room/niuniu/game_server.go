//牛牛服务
package niuniu

import (
	"github.com/snowyyj001/loumiao"
	"github.com/snowyyj001/loumiao/gorpc"
	"github.com/snowyyj001/loumiao/log"
)

func GetRoomId(gameId int, idx int) int {
	return gameId*1000 + idx
}

var (
	This *GameServer
)

type GameServer struct {
	gorpc.GoRoutineLogic

	GameId   int
	GameRule string
	RoomNum  int
	Rooms    map[int]*Room
	Players  map[int]*Room
}

func (self *GameServer) DoInit() {
	log.Infof("%s DoInit", self.Name)
	This = self
	loumiao.RegisterNetHandler(self, "C_R_SitDown", handlerSitDown)

}

func (self *GameServer) DoRegsiter() {
	self.Register("joinRoom", self.handlerJoinRoom)
}

func (self *GameServer) DoStart() {
	log.Infof("%s DoStart", self.Name)

	self.Players = make(map[int]*Room)
	self.Rooms = make(map[int]*Room)
	for i := 0; i < self.RoomNum; i++ {
		roomid := GetRoomId(self.GameId, i+1)
		self.Rooms[roomid] = new(Room)
		self.Rooms[roomid].doStart(roomid)
	}

	self.RunTicker(1000, self.Update)
}

func (self *GameServer) DoDestory() {
	log.Info("niuniu GameServer destory")
}

func (self *GameServer) Update(igo gorpc.IGoRoutine, data interface{}) interface{} {
	dt := data.(int64)
	for _, room := range self.Rooms {
		room.update(dt)
	}
	return nil
}
