package gate

import (
	"server/login"
	"server/msg"
	gamePb "server/msg/gamepb"
)

func init() {
	msg.Processor.SetRouter(&gamePb.ReqLogin{}, login.ChanRPC)
	msg.Processor.SetRouter(&gamePb.ReqRegister{}, login.ChanRPC)
	msg.Processor.SetRouter(&gamePb.ReqUserInfo{}, login.ChanRPC)
}
