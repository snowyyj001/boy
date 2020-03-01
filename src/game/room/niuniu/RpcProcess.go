package niuniu

import (
	"game/define"
	"game/world/agent"

	//"unsafe"

	"github.com/snowyyj001/loumiao/log"

	"github.com/snowyyj001/loumiao/gorpc"
)

func (self *GameServer) handlerJoinRoom(igo gorpc.IGoRoutine, data interface{}) interface{} {
	m := data.(gorpc.M)

	player := gorpc.SimpleGK(m, "user").(agent.Agent)
	roomid := gorpc.SimpleGK(m, "roomid").(int)

	err := 0
	if self.Rooms[roomid] != nil {
		err = self.Rooms[roomid].canJoinRoom(&player)
	} else {
		err = define.Err_Room_NoExist
	}

	if err == 0 {
		self.Rooms[roomid].joinRoom(&player)
		log.Debugf("玩家%d加入房间%d", player.ID, roomid)
		self.Players[player.ClientId] = self.Rooms[roomid]
	}
	return err
}

func handlerSitDown(igo gorpc.IGoRoutine, clientid int, data interface{}) interface{} {
	room := This.Players[clientid]
	room.sitDown(clientid)

	return nil
}
