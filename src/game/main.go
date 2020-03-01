package main

import (
	"game/config"

	"github.com/snowyyj001/loumiao"
	"github.com/snowyyj001/loumiao/gate"
	"github.com/snowyyj001/loumiao/log"

	lconf "github.com/snowyyj001/loumiao/config"
)

func init() {
	lconf.NET_PROTOCOL = config.NET_PROTOCOL
	lconf.NET_WEBSOCKET = config.NET_WEBSOCKET
	lconf.NET_GATE_PORT = config.NET_GATE_PORT
	lconf.NET_BE_CHILD = config.NET_BE_CHILD
}

func main() {
	log.Info("server run!")

	//gate start, do some other init
	gate := new(gate.GateServer)
	if config.NET_BE_CHILD {
		for _, cfg := range config.ServerCfg.ServerNodes {
			gate.BuildRpc(cfg.Ip, cfg.Port, cfg.Id, cfg.Name)
		}
	}
	loumiao.Prepare(gate, "GateServer", false)

	loumiao.Run()
}
