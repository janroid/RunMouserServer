package internal

import (
	"server/game/modules"
)

const (
	U_STA_NONE  = 0  // 无状态
	U_STA_READY = 1  // 游戏开始前准备状态
	U_STA_OPE   = 2  // 玩家操作状态
	U_STA_MSE   = 10 // 内圈状态
	U_STA_CAT   = 11 // 外圈状态
	U_STA_LJ    = 12 // 失业等待状态
)

type GUser struct {
	uid    int             // 用户ID
	name   string          // 用户名
	icon   int             // 头像
	career *modules.Career // 职业
	status int             //操作状态
}

func (u *GUser) getStatus() int {
	return u.career.status
}
