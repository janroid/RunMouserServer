package gate

import (
	"server/login"
	"server/msg"
	gamePb "server/msg/gamepb"
)

func init() {
	msg.Processor.SetRouter(&gamePb.UserLogin{}, login.ChanRPC)
	msg.Processor.SetRouter(&gamePb.UserRegister{}, login.ChanRPC)
}
