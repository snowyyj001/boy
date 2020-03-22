package msg

import (
	"github.com/snowyyj001/loumiao/message"
)

func init() {
	message.RegisterPacket(&C_A_Login{})
	message.RegisterPacket(&A_C_Login{})
}
