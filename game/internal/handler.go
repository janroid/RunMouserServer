package internal

import (
	"reflect"
	gamePb "server/msg/gamepb"

	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	handler(&gamePb.UserLogin{}, handleLogin)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleLogin(args []interface{}) {
	m := args[0].(*gamePb.UserLogin)
	a := args[1].(gate.Agent)

	log.Release("handleLogin - type = %v, name = %v, pwd = %v", m.GetMtype(), m.GetName(), m.GetPassword())

	a.WriteMsg(&gamePb.LoginResult{
		Result: proto.Int32(0),
		Name:   proto.String(m.GetName()),
	})

}
