package gate

import (
	"game/config"

	"github.com/snowyyj001/loumiao"
	lgate "github.com/snowyyj001/loumiao/gate"
	"github.com/snowyyj001/loumiao/network"
)

func RegisterRpc(watchdog *lgate.GateServer) {
	watchdog.RegisterRpc("C_S_Login", config.WORLD_NODE)

}

func StartGate() {
	watchdog := &lgate.GateServer{ServerType: network.CLIENT_CONNECT}
	RegisterRpc(watchdog)
	loumiao.Prepare(watchdog, "GateServer", false)
}
