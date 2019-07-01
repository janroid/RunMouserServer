// 卡片
package modules

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/name5566/leaf/log"
)

// 属性组
const (
	ATTR_TP_XMM  = 1 // 小买卖
	ATTR_TP_DMM  = 2 // 大买卖
	ATTR_TP_SCFY = 3 // 市场风云
	ATTR_TP_YW   = 4 // 意外开支
	ATTR_TP_TDK  = 5 // 特定事件
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
	Count        int    `bson:"count"`        // 数量
}

type RMCards struct {
	XMMCards []*RMCard
	DMMCards []*RMCard
	SCCards  []*RMCard
	EWCards  []*RMCard
	XHCard   *RMCard // 小孩卡
	SYCard   *RMCard // 失业卡
	CSCard   *RMCard // 慈善卡
	YHCard   *RMCard // 银行结算

}

var CardsFactory = new(RMCards)

func InitData() {
	log.Debug("game.modules.card -- InitData")
	var jsonData = []byte(JSON_DATA_XMM)

	CardsFactory.XMMCards = make([]*RMCard, 0)

	err := json.Unmarshal(jsonData, &CardsFactory.XMMCards)

	if err != nil {
		log.Error("card.go - init error: json err = %v", err)

		return
	}

	jsonData = []byte(JSON_DATA_DMM)

	CardsFactory.DMMCards = make([]*RMCard, 0)

	err = json.Unmarshal(jsonData, &CardsFactory.DMMCards)

	if err != nil {
		log.Error("card.go - init error: json err = %v", err)

		return
	}

	jsonData = []byte(JSON_DATA_SCFY)

	CardsFactory.SCCards = make([]*RMCard, 0)

	err = json.Unmarshal(jsonData, &CardsFactory.SCCards)

	if err != nil {
		log.Error("card.go - init error: json err = %v", err)

		return
	}

	jsonData = []byte(JSON_DATA_EWZC)

	CardsFactory.EWCards = make([]*RMCard, 0)

	err = json.Unmarshal(jsonData, &CardsFactory.EWCards)

	if err != nil {
		log.Error("card.go - init error: json err = %v", err)

		return
	}

	h := new(RMCard)
	(*h).Cid = 501
	(*h).AttrType = 5
	(*h).GroupType = 63
	(*h).Title = "家里添丁"
	(*h).Desc = "特大喜讯，恭喜%s，迎来了他们的新宝宝！\\r\\n每月增加小孩花费180."
	(*h).Value = 0
	(*h).Prange = "0"
	(*h).Payment = 0
	(*h).Loan = 0
	(*h).Cflow = -180
	(*h).Abbreviation = "孩子支付"
	(*h).LinkID = []int{-1}
	(*h).Count = 1

	CardsFactory.XHCard = h

	h = new(RMCard)
	(*h).Cid = 502
	(*h).AttrType = 5
	(*h).GroupType = 53
	(*h).Title = "失业"
	(*h).Desc = "因公司经营不善，市场萎缩，公司无法支撑当前规模，你处于裁员名单中。\\r\\n你将减少一个%d现金，并停赛2轮。"
	(*h).Value = 0
	(*h).Prange = "0"
	(*h).Payment = 0
	(*h).Loan = 0
	(*h).Cflow = 0
	(*h).Abbreviation = ""
	(*h).LinkID = []int{-1}
	(*h).Count = 1

	CardsFactory.SYCard = h

	h = new(RMCard)
	(*h).Cid = 504
	(*h).AttrType = 5
	(*h).GroupType = 53
	(*h).Title = "慈善捐款"
	(*h).Desc = "你被要求参加一个慈善宴会，如果同意，则扣除%10现金(%d),后面3轮你将可以投振2颗骰子。"
	(*h).Value = 0
	(*h).Prange = "0"
	(*h).Payment = 0
	(*h).Loan = 0
	(*h).Cflow = 0
	(*h).Abbreviation = ""
	(*h).LinkID = []int{-1}
	(*h).Count = 1

	CardsFactory.CSCard = h

	h = new(RMCard)
	(*h).Cid = 503
	(*h).AttrType = 5
	(*h).GroupType = 53
	(*h).Title = "银行结算日"
	(*h).Desc = "现金增加当前现金流数量"
	(*h).Value = 0
	(*h).Prange = "0"
	(*h).Payment = 0
	(*h).Loan = 0
	(*h).Cflow = 0
	(*h).Abbreviation = ""
	(*h).LinkID = []int{-1}
	(*h).Count = 1

	CardsFactory.YHCard = h

}

func (c RMCards) getXMMCard() RMCard {
	rand.Seed(int64(time.Now().UnixNano()))
	n := rand.Intn(len(c.XMMCards))

	return *c.XMMCards[n]
}

func (c RMCards) getDMMCard() RMCard {
	rand.Seed(int64(time.Now().UnixNano()))
	n := rand.Intn(len(c.DMMCards))

	return *c.DMMCards[n]
}

func (c RMCards) getSCCard() RMCard {
	rand.Seed(int64(time.Now().UnixNano()))
	n := rand.Intn(len(c.SCCards))

	return *c.SCCards[n]
}

func (c RMCards) getEWCard() RMCard {
	rand.Seed(int64(time.Now().UnixNano()))
	n := rand.Intn(len(c.EWCards))

	return *c.EWCards[n]
}

func (c RMCards) getSYCard() RMCard {
	return *c.SYCard
}

func (c RMCards) getXHCard() RMCard {
	return *c.XHCard
}

func (c RMCards) getCSCard() RMCard {
	return *c.CSCard
}

func (c RMCards) getYHCard() RMCard {
	return *c.YHCard
}

// 创建一张信用卡卡片
func (c RMCards) createXYK(money float64) RMCard {
	h := new(RMCard)
	(*h).Cid = 504
	(*h).AttrType = 5
	(*h).GroupType = 52
	(*h).Title = "信用卡"
	(*h).Desc = "使用了信用卡消费"
	(*h).Value = 0
	(*h).Prange = "0"
	(*h).Payment = 0
	(*h).Loan = int(money)
	(*h).Cflow = int(money * -0.03)
	(*h).Abbreviation = "信用卡"
	(*h).LinkID = []int{-1}
	(*h).Count = 1

	return *h
}

//创建一张高息贷款卡片
func (c RMCards) createGXDK(money float64) RMCard {
	h := new(RMCard)
	(*h).Cid = 505
	(*h).AttrType = 5
	(*h).GroupType = 51
	(*h).Title = "高息贷款"
	(*h).Desc = "向银行借款"
	(*h).Value = 0
	(*h).Prange = "0"
	(*h).Payment = 0
	(*h).Loan = int(money)
	(*h).Cflow = int(money * -0.1)
	(*h).Abbreviation = "高息贷款"
	(*h).LinkID = []int{-1}
	(*h).Count = 1

	return *h
}
