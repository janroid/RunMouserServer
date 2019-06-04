package internal

import (
	"math/rand"

	"server/data"

	"github.com/name5566/leaf/log"
	"gopkg.in/mgo.v2/bson"
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

var (
	userList map[int64]*User
	uidIndex int64

	dbName  = "runmouse"
	tabAcc  = "admin"
	tabUser = "user"
)

// 用户信息
type User struct {
	uid        int64  // 用户ID
	exp        int64  // 经验
	name       string // 用户名
	icon       int    // 头像
	acName     string // 用户名
	acPwd      string // 用户密码
	playCount  int64  // 总玩牌局数
	playWin    int64  // 赢的次数
	playOut    int64  // 淘汰次数
	playCreate int64  // 创建房间次数
	honor      int64  // 荣誉
	money      int64  // 游戏币
	gold       int64  // 代币
	title      int    // 头衔
	status     int    // 账号状态
}

// 返回的错误信息
type ErrCode struct {
	Code   int
	ErrMsg error
}

// 账户信息
type DBAcc struct {
	ID     int64  `bson:"uid"`
	ACNAME string `bson:"name"`
	PWD    string `bson:"password"`
}

type DBUser struct {
	ID     int64  `bson:"uid"`
	EXP    int64  `bson:"exp"`
	NAME   string `bson:"name"`
	IC     int    `bson:"icon"`
	TOTAL  int64  `bson:"playcount"`
	WIN    int64  `bson:"playwin"`
	OUT    int64  `bson:"playout"`
	CREATE int64  `bson:"create"`
	HONOR  int64  `bson:"honor"`
	MONEY  int64  `bson:"money"`
	GLOD   int64  `bson:"gold"`
	TITLE  int    `bson:"title"`
	STATYS int    `bson:"status"`
}

func init() {
	userList = make(map[int64]*User)

	var users []DBAcc
	err := data.FindAll(dbName, tabAcc, &users)

	if err != nil {
		log.Error("******************************* err = %v", err)
		return
	}

	uidIndex = 0
	for _, u := range users {
		if uidIndex < u.ID {
			uidIndex = u.ID
		}
	}
}

// 根据ID获取用户信息
func GetUserByUID(uid int64) (*User, int) {
	for _, value := range userList {
		if value.uid == uid {
			return value, 0
		}
	}

	u := findDBUser(uid)

	if u == nil {
		return nil, -1
	}

	user := assemblyUser(uid, nil, u)
	addUserList(user)

	return user, 0
}

// 用户名登陆
func LoginByName(name string) *User {
	for _, user := range userList {
		if user.acName == name {
			return user
		}
	}

	acc := findDBAcc(name, -1)

	if acc == nil {
		return nil
	}

	user := assemblyUser(acc.ID, acc, nil)

	addUserList(user)

	return user
}

// 创建新用户
func createUser(name string, pwd string) *User {
	if len(name) <= 0 || len(pwd) <= 0 {
		return nil
	}

	user := new(User)

	initUserInfo(user)
	user.name = name
	user.acName = name
	user.acPwd = pwd

	userList[user.uid] = user

	updateUser(user, true, true, 2)

	return user
}

func initUserInfo(u *User) {
	u.uid = allocUID()
	u.exp = 0
	u.icon = rand.Intn(11)
	u.playCount = 0
	u.playWin = 0
	u.playOut = 0
	u.playCreate = 0
	u.honor = 0
	u.money = 1000
	u.gold = 0
	u.title = 1
	u.status = STRUCT_NONE
}

func allocUID() int64 {
	uidIndex += uidIndex

	return uidIndex
}

// 从数据库中查找账号
func findDBAcc(name string, id int64) *DBAcc {
	acc := new(DBAcc)

	var err error
	if id <= 0 {
		err = data.Find(dbName, tabAcc, bson.M{"name": name}, &acc)
	} else {
		err = data.Find(dbName, tabAcc, bson.M{"uid": id}, &acc)
	}

	if err != nil {
		log.Error("user.go - findDBAcc error : %v", err)
		return nil
	}

	return acc
}

// 从数据库中查找用户
func findDBUser(id int64) *DBUser {
	u := new(DBUser)

	err := data.Find(dbName, tabUser, bson.M{"uid": id}, &u)

	if err != nil {
		log.Error("user.go - findDBUser error : %v", err)

		return nil
	}

	return u
}

// 组合user
func assemblyUser(id int64, acc *DBAcc, u *DBUser) *User {
	if acc == nil {
		acc = findDBAcc("", id)
	}

	if u == nil {
		u = findDBUser(id)
	}

	if acc == nil || u == nil {
		return nil
	}

	user := new(User)

	user.uid = id
	user.exp = u.EXP
	user.name = u.NAME
	user.icon = u.IC
	user.acName = acc.ACNAME
	user.acPwd = acc.PWD
	user.playCount = u.TOTAL
	user.playWin = u.WIN
	user.playOut = u.OUT
	user.playCreate = u.CREATE
	user.honor = u.HONOR
	user.money = u.MONEY
	user.gold = u.GLOD
	user.title = u.TITLE
	user.status = u.STATYS

	return user
}

// 添加user到缓存列表
func addUserList(user *User) {
	if user == nil {
		return
	}

	if user.uid <= 0 {
		return
	}

	userList[user.uid] = user
}

// 更新账户表
func updateAccDB(acc *DBAcc) {
	if acc == nil {
		return
	}

	data.Update(dbName, tabAcc, bson.M{"uid": acc.ID}, acc)

}

// 更新用户数据库表
func updateUserDB(user *DBUser) {
	if user == nil {
		return
	}

	data.Update(dbName, tabUser, bson.M{"uid": user.ID}, user)
}

// 新增账户表数据
func insertAccDB(acc *DBAcc) {
	if acc == nil {
		return
	}

	data.Insert(dbName, tabAcc, acc)
}

// 新增用户表数据
func insertUserDB(user *DBUser) {
	if user == nil {
		return
	}

	data.Insert(dbName, tabUser, user)
}

// 根据用户对象更新账号或用户表, upOadd : 1；表示更新，2：表示添加
func updateUser(user *User, isAcc bool, isUser bool, upOadd int) {
	if user == nil {
		return
	}

	if isAcc {
		acc := new(DBAcc)
		acc.ID = user.uid
		acc.ACNAME = user.acName
		acc.PWD = user.acPwd

		if upOadd == 1 {
			updateAccDB(acc)
		} else if upOadd == 2 {
			insertAccDB(acc)
		}

	}

	if isUser {
		dbu := new(DBUser)
		dbu.ID = user.uid
		dbu.EXP = user.exp
		dbu.NAME = user.name
		dbu.IC = user.icon
		dbu.TOTAL = user.playCount
		dbu.WIN = user.playWin
		dbu.OUT = user.playOut
		dbu.CREATE = user.playCreate
		dbu.HONOR = user.honor
		dbu.MONEY = user.money
		dbu.GLOD = user.gold
		dbu.TITLE = user.title
		dbu.STATYS = user.status

		if upOadd == 1 {
			updateUserDB(dbu)
		} else if upOadd == 2 {
			insertUserDB(dbu)
		}

	}
}
