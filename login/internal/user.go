package internal

import (
	"math/rand"

	ACC "server/account"
)

// 登陆注册错误代码
const (
	LGOINSUCC    = 1000 // 登陆成功
	REGISTERSUCC = 1001 // 注册成功
	ERRNOUSER    = 1002 // 无法找到用户
	ERRHAVEUSER  = 1003 // 账号已存在
	ERRLOGIN     = 1004 // 用户名或密码错误
	ERRFORBID    = 1005 // 账户被封
	ERRNULL      = 1006 // 用户名或密码不能为空
	ERRSHORT     = 1007 // 用户名或密码长度太短
)

// 头像预定义
const (
	DEFICON_0  = 0
	DEFICON_1  = 1
	DEFICON_2  = 2
	DEFICON_3  = 3
	DEFICON_4  = 4
	DEFICON_5  = 5
	DEFICON_6  = 6
	DEFICON_7  = 7
	DEFICON_8  = 8
	DEFICON_9  = 9
	DEFICON_10 = 10
	DEFICON_11 = 11
)

// 位字段
const (
	FIELD_MONEY      = 1
	FIELD_EXP        = 2
	FIELD_NAME       = 4
	FIELD_ICON       = 8
	FIELD_ACNAME     = 16
	FIELD_ACPWD      = 32
	FIELD_PLAYCOUNT  = 64
	FIELD_PLAYWIN    = 128
	FIELD_PLAYOUT    = 256
	FIELD_PLAYCREATE = 512
	FIELD_HONOR      = 1024
	FIELD_GOLD       = 2048
	FIELD_TITLE      = 4096
	FIELD_STATUS     = 8192
)

// 玩家账户状态
const (
	STRUCT_NONE   = 0
	STRUCT_FORBID = 1
)

// 请求类型
const (
	TYPE_LOGIN    = 1
	TYPE_REGISTER = 2
)

// 返回的错误信息
type ErrCode struct {
	Code   int
	ErrMsg error
}

// 根据ID获取用户信息
func GetUserByUID(uid int64) (*ACC.User, int) {
	for _, value := range ACC.UserList {
		if value.Uid == uid {
			return value, 0
		}
	}

	u := ACC.FindDBUser(uid)

	if u == nil {
		return nil, -1
	}

	user := assemblyUser(uid, nil, u)
	addUserList(user)

	return user, 0
}

// 用户名登陆
func LoginByName(name string) *ACC.User {
	for _, user := range ACC.UserList {
		if user.AcName == name {
			return user
		}
	}

	acc := ACC.FindDBAcc(name, -1)

	if acc == nil {
		return nil
	}

	user := assemblyUser(acc.ID, acc, nil)

	addUserList(user)

	return user
}

// 创建新用户
func createUser(name string, pwd string) *ACC.User {
	if len(name) <= 0 || len(pwd) <= 0 {
		return nil
	}

	user := new(ACC.User)

	initUserInfo(user)
	user.Name = name
	user.AcName = name
	user.AcPwd = pwd

	ACC.UserList[user.Uid] = user

	updateUser(user, true, true, 2)

	return user
}

func initUserInfo(u *ACC.User) {
	u.Uid = ACC.AllocUid()
	u.Exp = 0
	u.Icon = rand.Intn(11)
	u.PlayCount = 0
	u.PlayWin = 0
	u.PlayOut = 0
	u.PlayCreate = 0
	u.Honor = 0
	u.Money = 1000
	u.Gold = 0
	u.Title = 1
	u.Status = STRUCT_NONE
}

// 组合user
func assemblyUser(id int64, acc *ACC.DBAcc, u *ACC.DBUser) *ACC.User {
	if acc == nil {
		acc = ACC.FindDBAcc("", id)
	}

	if u == nil {
		u = ACC.FindDBUser(id)
	}

	if acc == nil || u == nil {
		return nil
	}

	user := new(ACC.User)

	user.Uid = id
	user.Exp = u.EXP
	user.Name = u.NAME
	user.Icon = u.IC
	user.AcName = acc.ACNAME
	user.AcPwd = acc.PWD
	user.PlayCount = u.TOTAL
	user.PlayWin = u.WIN
	user.PlayOut = u.OUT
	user.PlayCreate = u.CREATE
	user.Honor = u.HONOR
	user.Money = u.MONEY
	user.Gold = u.GLOD
	user.Title = u.TITLE
	user.Status = u.STATYS

	return user
}

// 添加user到缓存列表
func addUserList(user *ACC.User) {
	if user == nil {
		return
	}

	if user.Uid <= 0 {
		return
	}

	ACC.UserList[user.Uid] = user
}

// 根据用户对象更新账号或用户表, upOadd : 1；表示更新，2：表示添加
func updateUser(user *ACC.User, isAcc bool, isUser bool, upOadd int) {
	if user == nil {
		return
	}

	if isAcc {
		acc := new(ACC.DBAcc)
		acc.ID = user.Uid
		acc.ACNAME = user.AcName
		acc.PWD = user.AcPwd

		if upOadd == 1 {
			ACC.UpdateAccDB(acc)
		} else if upOadd == 2 {
			ACC.InsertAccDB(acc)
		}

	}

	if isUser {
		dbu := new(ACC.DBUser)
		dbu.ID = user.Uid
		dbu.EXP = user.Exp
		dbu.NAME = user.Name
		dbu.IC = user.Icon
		dbu.TOTAL = user.PlayCount
		dbu.WIN = user.PlayWin
		dbu.OUT = user.PlayOut
		dbu.CREATE = user.PlayCreate
		dbu.HONOR = user.Honor
		dbu.MONEY = user.Money
		dbu.GLOD = user.Gold
		dbu.TITLE = user.Title
		dbu.STATYS = user.Status

		if upOadd == 1 {
			ACC.UpdateUserDB(dbu)
		} else if upOadd == 2 {
			ACC.InsertUserDB(dbu)
		}

	}
}
