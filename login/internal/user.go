package internal

import "math/rand"

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

type ErrCode struct {
	Code   int
	ErrMsg error
}

func init() {
	userList = make(map[int64]*User)
	uidIndex = 1001

	user := createUser("janroid", "123456")

	userList[user.uid] = user
}

func GetUserByUID(uid int64) (*User, int) {
	for _, value := range userList {
		if value.uid == uid {
			return value, 0
		}
	}

	return nil, -1
}

func GetUserByName(name string) *User {
	for _, user := range userList {
		if user.acName == name {
			return user
		}
	}
	return nil
}

func createUser(name string, pwd string) *User {
	user := new(User)

	initUserInfo(user)
	user.name = name
	user.acName = name
	user.acPwd = pwd

	userList[user.uid] = user

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
