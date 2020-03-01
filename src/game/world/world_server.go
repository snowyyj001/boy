// 游戏服务
package world

import (
	"github.com/snowyyj001/loumiao"
	"github.com/snowyyj001/loumiao/gorpc"
	"github.com/snowyyj001/loumiao/log"
)

var (
	This *WorldServer
)

type WorldServer struct {
	gorpc.GoRoutineLogic
}

//在这里注册网络消息
func (self *WorldServer) DoInit() {
	log.Info("WorldServer DoInit")
	This = self

	loumiao.RegisterNetHandler(self, "C_S_Login", handlerLogin)
	loumiao.RegisterNetHandler(self, "C_S_JoinRoom", handlerJoinRoom)
}

//在这里注册rpc消息
func (self *WorldServer) DoRegsiter() {
	//self.Register("handlerLogin", handlerLogin)
}

func (self *WorldServer) DoDestory() {
	log.Info("WorldServer destory")
}

func (self *WorldServer) Update() {

}
