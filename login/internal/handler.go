package internal

import (
	"reflect"

	gamePb "server/msg/gamepb"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	handleMsg(&gamePb.UserLogin{}, login)
	handleMsg(&gamePb.UserRegister{}, register)
}

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func login(args []interface{}) {
	var name, pwd string
	if data, ok := args[0].(*gamePb.UserLogin); ok {
		name = data.GetName()
		pwd = data.GetPassword()
	}

	err, user := UserLogin(name, pwd)

	result := new(gamePb.LoginResult)

	er := int32(err)

	result.Result = &er

	t := int32(TYPE_LOGIN)
	result.Type = &t

	if err == LGOINSUCC {
		result.Uid = &user.uid
		result.Money = &user.money
		icon := int32(user.icon)
		result.Icon = &icon
		result.Exp = &user.exp
		result.Name = &user.name
	}

	log.Debug("login Handler - login : result = %v ", result)

	if a, ok := args[1].(gate.Agent); ok {
		a.WriteMsg(result)
	} else {
		log.Error("login Handler - login error, agent is nil")
	}
}

func register(args []interface{}) {
	var name, pwd string
	if data, ok := args[0].(*gamePb.UserRegister); ok {
		name = data.GetName()
		pwd = data.GetPassword()
	}

	result := new(gamePb.LoginResult)

	err, user := UserRegister(name, pwd)
	e := int32(err)
	result.Result = &e

	t := int32(TYPE_REGISTER)
	result.Type = &t

	if err == REGISTERSUCC {
		result.Uid = &user.uid
		result.Money = &user.money
		icon := int32(user.icon)
		result.Icon = &icon
		result.Exp = &user.exp
		result.Name = &user.name
	}

	log.Debug("loign handler - register : data = %v", result)

	if a, ok := args[1].(gate.Agent); ok {
		a.WriteMsg(result)
	} else {
		log.Error("login Handler - register error, agent is nil")
	}

}
