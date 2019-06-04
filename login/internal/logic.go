package internal

import (
	gamePb "server/msg/gamepb"
)

var MINLEN = 6 //用户名密码最小长度

func UserLogin(name string, pwd string) (int, *User) {
	if len(name) < MINLEN {
		return ERRSHORT, nil
	}

	if len(pwd) < MINLEN {
		return ERRSHORT, nil
	}

	user := LoginByName(name)

	if user == nil {
		return ERRNOUSER, nil
	}

	if user.acPwd != pwd {
		return ERRLOGIN, nil
	}

	if user.status == STRUCT_FORBID {
		return ERRFORBID, nil
	}

	return LGOINSUCC, user
}

func UserRegister(name string, pwd string) (int, *User) {
	if len(name) < MINLEN {
		return ERRSHORT, nil
	}

	if len(pwd) < MINLEN {
		return ERRSHORT, nil
	}

	user := LoginByName(name)
	if user != nil {
		return ERRHAVEUSER, nil
	}

	user = createUser(name, pwd)

	return REGISTERSUCC, user

}

func GetUserInfo(uid int64, fields int32) *gamePb.RpsUserInfo {
	pbU := new(gamePb.RpsUserInfo)

	if uid <= 0 || fields < 0 {
		id := int64(-1)
		pbU.Uid = &id

		return pbU
	}

	user, _ := GetUserByUID(uid)
	if user == nil {
		id := int64(-1)
		pbU.Uid = &id

		return pbU
	}

	isAll := fields == 0

	pbU.Uid = &uid

	tmp := fields & FIELD_MONEY
	if isAll || tmp > 0 {
		pbU.Money = &user.money
	}

	tmp = fields & FIELD_EXP
	if isAll || tmp > 0 {
		pbU.Exp = &user.exp
	}

	tmp = fields & FIELD_NAME
	if isAll || tmp > 0 {
		pbU.Name = &user.name
	}

	tmp = fields & FIELD_ICON
	if isAll || tmp > 0 {
		ic := int32(user.icon)

		pbU.Icon = &ic
	}

	tmp = fields & FIELD_PLAYCOUNT
	if isAll || tmp > 0 {
		pbU.PlayCount = &user.playCount
	}

	tmp = fields & FIELD_PLAYWIN
	if isAll || tmp > 0 {
		pbU.PlayWin = &user.playWin
	}

	tmp = fields & FIELD_PLAYOUT
	if isAll || tmp > 0 {
		pbU.PlayOut = &user.playOut
	}

	tmp = fields & FIELD_PLAYCREATE
	if isAll || tmp > 0 {
		pbU.PlayCreate = &user.playCreate
	}

	tmp = fields & FIELD_HONOR
	if isAll || tmp > 0 {
		pbU.Honor = &user.honor
	}

	tmp = fields & FIELD_GOLD
	if isAll || tmp > 0 {
		pbU.Gold = &user.gold
	}

	tmp = fields & FIELD_TITLE
	if isAll || tmp > 0 {
		ti := int32(user.title)
		pbU.Title = &ti
	}

	tmp = fields & FIELD_STATUS
	if isAll || tmp > 0 {
		st := int32(user.status)
		pbU.Status = &st
	}

	return pbU
}
