package internal

import (
	"reflect"

	gamePb "server/msg/gamepb"

	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
)

func init() {
	handleMsg(&gamePb.ReqLogin{}, login)
	handleMsg(&gamePb.ReqRegister{}, register)
	handleMsg(&gamePb.ReqUserInfo{}, userInfo)
}

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func login(args []interface{}) {
	var name, pwd string
	if data, ok := args[0].(*gamePb.ReqLogin); ok {
		name = data.GetName()
		pwd = data.GetPassword()
	}

	err, user := UserLogin(name, pwd)

	result := new(gamePb.RpsAuthor)

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

		// time.Sleep(time.Second)
		pbU := GetUserInfo(user.uid, 0)
		a.WriteMsg(pbU)

	} else {
		log.Error("login Handler - login error, agent is nil")
	}
}

func register(args []interface{}) {
	var name, pwd string
	if data, ok := args[0].(*gamePb.ReqRegister); ok {
		name = data.GetName()
		pwd = data.GetPassword()
	}

	result := new(gamePb.RpsAuthor)

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

		pbU := GetUserInfo(user.uid, 0)
		a.WriteMsg(pbU)
	} else {
		log.Error("login Handler - register error, agent is nil")
	}

}

func userInfo(args []interface{}) {
	var uid int64
	var fields int32
	if data, ok := args[0].(*gamePb.ReqUserInfo); ok {
		uid = data.GetUid()
		fields = data.GetFields()
	}
	pbU := GetUserInfo(uid, fields)
	log.Debug("loign handler - userInfo : data = %v", pbU)
	if a, ok := args[1].(gate.Agent); ok {
		a.WriteMsg(pbU)
	} else {
		log.Error("login Handler - userInfo error, agent is nil")
	}
}
