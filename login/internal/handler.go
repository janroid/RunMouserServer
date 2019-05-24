package internal

import (
	"reflect"

	"github.com/name5566/leaf/log"
)

func init() {
	log.Debug("Login.handler.init ")
	// handleMsg(&msg.UserLogin{}, handleLogin)
}

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleLogin(args []interface{}) {
	// m := args[0].(*msg.UserLogin)
	// a := args[1].(gate.Agent)

	// log.Debug("Test login %v ", m.LoginName)
	// for i := 1; i < 10; i++ {
	// 	a.WriteMsg(&msg.Hello{
	// 		Name: "client",
	// 	})

	// 	time.Sleep(time.Second)
	// }

}
