// 卡片
package modules

import (
	"encoding/json"

	"github.com/name5566/leaf/log"
)

// 属性组
const (
	ATTR_TP_XMM  = 1 // 小买卖
	ATTR_TP_DMM  = 2 // 大买卖
	ATTR_TP_SCFY = 3 // 市场风云
	ATTR_TP_YW   = 4 // 意外开支
)

// 卡片类型组
const (
	GP_TP_GB     = 1  //袁大头
	GP_TP_HJ     = 2  //黄金
	GP_TP_ZZ2    = 3  //2室1厅
	GP_TP_ZZ3    = 4  //3室2厅
	GP_TP_PA     = 5  //pa银行存单
	GP_TP_ZSY    = 6  //优先股中石油
	GP_TP_ALBB   = 7  // 股票阿里巴巴
	GP_TP_XMKJ   = 8  // 股票xiaomi
	GP_TP_TX     = 9  // 股票腾讯
	GP_TP_CF200  = 10 // 基金
	GP_TP_LAND10 = 11 //10亩土地
	GP_TP_LAND20 = 12 // 20亩土地
	GP_TP_WD     = 13 // 网店
	GP_TP_YXGS   = 14 // 游戏公司
	GP_TP_GY2    = 15 // 2室公寓
	GP_TP_GY4    = 16 // 4室公寓
	GP_TP_GY8    = 17 //8室公寓
	GP_TP_GY12   = 18 //12套公寓
	GP_TP_GY24   = 19 // 24套公寓
	GP_TP_60     = 20 // 60套组合公寓
	GP_TP_ZDH    = 21 //自动化企业
	GP_TP_HHR    = 22 //合伙人公司
	GP_TP_SC     = 23 //商场
	GP_TP_XC     = 24 // 洗车店
	GP_TP_ZC     = 25 // 早餐店
	GP_TP_HB     = 26 // 汉堡店
	GP_TP_XH     = 50 // 小孩
	GP_TP_DK     = 51 // 高息贷款
	GP_TP_XYK    = 52 // 信用卡
	GP_TP_XGK    = 53 // 效果卡
	GP_TP_KQ     = 54 // 减少钱
	GP_TP_SGGD   = 60 // 按固定价格收购
	GP_TP_SGC    = 61 // 按价格基数*数量 收购
	GP_TP_SGADD  = 62 // 按差价收购
	GP_TP_XJGB   = 63 // 现金流改变
	GP_TP_SGFB   = 64 //按原价格翻倍收购
)

type RMCard struct {
	Cid          int    `bson:"cid"`          // 唯一ID
	AttrType     int    `bson:"attrType"`     // 属性组
	GroupType    int    `bson:"groupType"`    // 类型
	Title        string `bson:"title"`        // 标题
	Desc         string `bson:"desc"`         // 描述
	Value        int    `bson:"value"`        // 价值
	Prange       string `bson:"prange"`       // 价值区间
	Payment      int    `bson:"payment"`      // 首付
	Loan         int    `bson:"loan"`         // 贷款
	Cflow        int    `bson:"cflow"`        // 现金流
	Abbreviation string `bson:"abbreviation"` // 缩写，简称，使用于报表中
	LinkID       []int  `bson:"linkID"`       // 关联卡片类型，类型数组
}

func (c *RMCard) InitData() {
	log.Debug("game.modules.card -- InitData")
	var jsonData = []byte(JSON_DATA_XMM)

	var xmmData = make([]RMCard, 0)

	err := json.Unmarshal(jsonData, &xmmData)

	if err != nil {
		log.Error("card.go - init error: json err = %v", err)

		return
	}

	log.Debug("id = %v", xmmData)

	for _, v := range xmmData {
		log.Debug("id = %v", v.Cid)
	}

}
