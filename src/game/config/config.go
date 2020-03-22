package config

//rpc node
const (
	WORLD_NODE = 1
)

//net
const (
	NET_PROTOCOL  = "JSON"      //OR"PROTOBUF" 使用JSON协议
	NET_WEBSOCKET = true        //使用websocket
	NET_GATE_IP   = "127.0.0.1" //网关监听ip
	NET_GATE_PORT = 6789        //网关监听端口
	NET_BE_CHILD  = true        //作为分布式子网节点
)
