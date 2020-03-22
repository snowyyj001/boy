package main

import (
	"game/config"
	"game/gate"

	"github.com/snowyyj001/loumiao"
	"github.com/snowyyj001/loumiao/log"

	lconf "github.com/snowyyj001/loumiao/config"
)

func init() {
	lconf.NET_PROTOCOL = config.NET_PROTOCOL
	lconf.NET_WEBSOCKET = config.NET_WEBSOCKET
	lconf.NET_GATE_PORT = config.NET_GATE_PORT
	lconf.NET_GATE_IP = config.NET_GATE_IP
	lconf.NET_BE_CHILD = config.NET_BE_CHILD
}

func main() {
	log.Info("server run!")

	gate.StartGate()

	loumiao.Run()
}
