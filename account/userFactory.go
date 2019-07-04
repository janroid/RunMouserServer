package account

import (
	"server/data"

	"github.com/name5566/leaf/log"
)

var (
	UserList map[int64]*User
	uidIndex int64
)

// 用户信息
type User struct {
	Uid        int64  // 用户ID
	Exp        int64  // 经验
	Name       string // 用户名
	Icon       int    // 头像
	AcName     string // 用户名
	AcPwd      string // 用户密码
	PlayCount  int64  // 总玩牌局数
	PlayWin    int64  // 赢的次数
	PlayOut    int64  // 淘汰次数
	PlayCreate int64  // 创建房间次数
	Honor      int64  // 荣誉
	Money      int64  // 游戏币
	Gold       int64  // 代币
	Title      int    // 头衔
	Status     int    // 账号状态
}

func init() {
	UserList = make(map[int64]*User)

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

func AllocUid() int64 {
	uidIndex += 1

	return uidIndex
}
