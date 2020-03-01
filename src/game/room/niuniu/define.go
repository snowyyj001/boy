package niuniu

const (
	MAX_ROOM_RENSHU = 50 //房间最大人数
	MAX_SEAT        = 13 //座位数
	POKER_NUMBER    = 52 //扑克数量
	MIN_GAMER       = 2  //最小玩家数量
	HANDCARD_NUM    = 4  //手牌数量
)

type RoomState int

const (
	State_Idle RoomState = iota
	State_FaPai
	State_Ready
	State_Gameing
)

type FSMState int

const (
	FSM_Idle FSMState = iota
	FSM_Fapai
	FSM_Bipai
	FSM_Result
)

const (
	Time_Idle   = 20
	Time_FaPai  = 60
	Time_BiPai  = 30
	Time_Result = 10
)
