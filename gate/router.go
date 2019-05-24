package gate

import (
	"server/game"
	"server/msg"
	gamePb "server/msg/gamepb"
)

func init() {
	msg.Processor.SetRouter(&gamePb.UserLogin{}, game.ChanRPC)
}
