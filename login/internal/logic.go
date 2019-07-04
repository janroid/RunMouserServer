package internal

import (
	ACC "server/account"
	gamePb "server/msg/gamepb"
)

var MINLEN = 6 //用户名密码最小长度

func UserLogin(name string, pwd string) (int, *ACC.User) {
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

	if user.AcPwd != pwd {
		return ERRLOGIN, nil
	}

	if user.Status == STRUCT_FORBID {
		return ERRFORBID, nil
	}

	return LGOINSUCC, user
}

func UserRegister(name string, pwd string) (int, *ACC.User) {
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
		pbU.Money = &user.Money
	}

	tmp = fields & FIELD_EXP
	if isAll || tmp > 0 {
		pbU.Exp = &user.Exp
	}

	tmp = fields & FIELD_NAME
	if isAll || tmp > 0 {
		pbU.Name = &user.Name
	}

	tmp = fields & FIELD_ICON
	if isAll || tmp > 0 {
		ic := int32(user.Icon)

		pbU.Icon = &ic
	}

	tmp = fields & FIELD_PLAYCOUNT
	if isAll || tmp > 0 {
		pbU.PlayCount = &user.PlayCount
	}

	tmp = fields & FIELD_PLAYWIN
	if isAll || tmp > 0 {
		pbU.PlayWin = &user.PlayWin
	}

	tmp = fields & FIELD_PLAYOUT
	if isAll || tmp > 0 {
		pbU.PlayOut = &user.PlayOut
	}

	tmp = fields & FIELD_PLAYCREATE
	if isAll || tmp > 0 {
		pbU.PlayCreate = &user.PlayCreate
	}

	tmp = fields & FIELD_HONOR
	if isAll || tmp > 0 {
		pbU.Honor = &user.Honor
	}

	tmp = fields & FIELD_GOLD
	if isAll || tmp > 0 {
		pbU.Gold = &user.Gold
	}

	tmp = fields & FIELD_TITLE
	if isAll || tmp > 0 {
		ti := int32(user.Title)
		pbU.Title = &ti
	}

	tmp = fields & FIELD_STATUS
	if isAll || tmp > 0 {
		st := int32(user.Status)
		pbU.Status = &st
	}

	return pbU
}
