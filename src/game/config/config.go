package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//player
const (
	PLAYER_GOLD  int64 = 1000
	PLAYER_COIN  int64 = 0
	PLAYER_MONEY int64 = 0
)

//net
const (
	NET_PROTOCOL  = "JSON" //OR"PROTOBUF" 使用JSON协议
	NET_WEBSOCKET = true   //使用websocket
	NET_GATE_PORT = 6788   //网关监听端口
	NET_BE_CHILD  = true   //作为分布式子网节点
)

//dbinit
const (
	MYSQL_URI     = "117.51.136.136:3306"
	MYSQL_DBNAME  = "loumiao"
	MYSQL_ACCOUNT = "root"
	MYSQL_PASS    = "123456"
)

type NetNode struct {
	Id   int    `json:"id"`
	Ip   string `json:"ip"`
	Port int    `json:"port"`
	Name string `json:"string"`
}

type Server struct {
	ServerNodes []NetNode `json:"net"`
}

var ServerCfg Server

func init() {
	data, err := ioutil.ReadFile("config/cfg.json")
	if err != nil {
		fmt.Errorf("%v", err)
	}
	err = json.Unmarshal(data, &ServerCfg)
	if err != nil {
		fmt.Println("%v", err)
	}
}
