package internal

import (
	"strconv"
	"time"

	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/timer"
)

const (
	GYP_READY = 0 //准备阶段
	GYP_START = 1 //开始阶段
	GYP_END   = 2 //结束阶段
)

type GameServer struct {
	sid      int               //服务器ID
	uCount   int               //玩家人数
	maxCount int               //最大人数
	uidList  map[int]*GUser    //玩家列表
	round    int               //当前第几轮
	status   int               // 状态
	title    string            //房间名称
	time     int               //倒计时时间
	curID    int               //当前操作玩家
	timerDis *timer.Dispatcher // 定时器
}

var serverID = 10000

func allocID() int {
	serverID += 1

	return serverID
}

// 初始化房间
func (g *GameServer) create(num int) {
	if g.sid > 0 {
		log.Error("gameServer.ctor --- Error : server is init!")

		return
	}

	g.sid = allocID()
	if num <= 0 {
		num = 6
	}

	g.maxCount = num
	g.uCount = 0
	g.uidList = make([]*GUser, num)
	g.round = 0
	g.status = GYP_READY
	g.title = "百万富翁" + strconv.Itoa(g.sid) + "场"
	g.time = 0
	g.curID = 0

	g.timerDis = timer.NewDispatcher(10)
}

func (g *GameServer) start() int {
	for _, u := range g.uidList {
		if u.status != U_STA_READY {
			log.Error("gameServer.start -- Error : user is not ready !")
			return GYP_READY
		}
	}

	g.curID = 0

	g.startOperation(g.curID)

	return GYP_START
}

func (g *GameServer) startOperation(tid int) {
	user := g.uidList[tid]

	if user != nil {
		status := user.getStatus()
	}

}

func (g *GameServer) GetUserByID(id int) *GUser {
	for _, u := range g.uidList {
		if u.uid == id {
			return u
		}
	}

	return nil
}

func (g *GameServer) runTimer(d time.Duration, cb func()) *timer.Timer {
	t := g.timerDis.AfterFunc(d, cb)

	(<-g.timerDis.ChanTimer).Cb()

	return t
}

func (g *GameServer) stopTimer(t *timer.Timer) {
	t.Stop()
}
