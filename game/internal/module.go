package internal

import (
	"server/base"
	"server/game/modules"

	"github.com/name5566/leaf/module"

	"github.com/name5566/leaf/log"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func init() {
	log.Debug("game.module - init")

	modules.InitData()

}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
}

func (m *Module) OnDestroy() {

}
