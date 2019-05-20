package internal

import (
	"reflect"
	"server/msg"

	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	handler(&msg.Hello{}, handleHello)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHello(args []interface{}) {
	m := args[0].(*msg.Hello)

	a := args[1].(gate.Agent)

	log.Debug("hello %v", m.GetName())

	a.WriteMsg(&msg.Hello{
		Name: proto.String("client"),
	})

}
