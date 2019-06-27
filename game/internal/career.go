package internal

import (
	"server/game/modules"

	"github.com/name5566/leaf/log"
)

// 职业基本数据
// 工资收入,总收入,支出-税金,支出-住房抵押贷欽,支出-教育贷款,支出-购车贷款,支出-俏用卡支出,支出-额外支出,支出-其他支出,支出-每个孩子支出,支出-总支出,月现金流,储蓄,负债-住房抵押贷款,负债-教育贷款,负债-购车贷款,负债-侑用卡,负债-额外负愤
var CarrerData = [12][18]int{
	{1600, 1600, 280, 200, 0, 80, 60, 50, 300, 70, 970, 630, 560, 20000, 0, 4000, 2000, 1000},
	{2000, 2000, 360, 300, 0, 60, 60, 50, 450, 110, 1280, 720, 400, 30000, 0, 3000, 2000, 1000},
	{2500, 2500, 460, 400, 0, 80, 60, 50, 570, 140, 1620, 880, 750, 40000, 0, 4000, 2000, 1000},
	{2500, 2500, 460, 400, 0, 80, 60, 50, 570, 140, 1620, 880, 710, 40000, 0, 4000, 2000, 1000},
	{3000, 3000, 580, 400, 0, 100, 60, 50, 690, 160, 1880, 1120, 520, 40000, 0, 5000, 2000, 1000},
	{3100, 3100, 600, 400, 30, 100, 90, 50, 710, 170, 1980, 1120, 480, 40000, 6000, 5000, 3000, 1000},
	{3300, 3300, 630, 500, 60, 100, 90, 50, 760, 180, 2190, 1110, 400, 50000, 12000, 5000, 3000, 1000},
	{4600, 4600, 910, 700, 60, 120, 90, 50, 1000, 240, 2930, 1670, 400, 75000, 12000, 1000, 4000, 1000},
	{4900, 4900, 1050, 750, 60, 140, 120, 50, 1090, 250, 3260, 1640, 400, 75000, 12000, 1000, 4000, 1000},
	{7500, 7500, 1830, 1100, 390, 220, 180, 50, 1650, 380, 5420, 2080, 400, 110000, 78000, 11000, 6000, 1000},
	{9500, 9500, 2350, 1330, 0, 300, 660, 50, 2210, 480, 6900, 2600, 400, 133000, 0, 15000, 22000, 1000},
	{13200, 13200, 3420, 1900, 750, 380, 270, 50, 2880, 640, 9650, 3550, 400, 190000, 150000, 19000, 9000, 1000},
}

var (
	nameData = [12]string{"保安", "修车师傅", "公交司机", "企业秘书", "警察", "医院护士", "中学教师", "企业经理", "工程师", "律师", "飞机机长", "主任医生"}
)

const (
	CA_MW   = 0  // 门卫
	CA_JXS  = 1  // 机械师
	CA_KCSJ = 2  // 卡车司机
	CA_MS   = 3  // 秘书
	CA_JG   = 4  // 警官
	CA_HS   = 5  // 护士
	CA_JS   = 6  // 教师
	CA_JL   = 7  // 经理
	CA_GCS  = 8  // 工程师
	CA_LS   = 9  // 律师
	CA_FJS  = 10 // 飞机驾驶员
	CA_YS   = 11 // 医生
)

// 报表详情
const (
	CA_TP_NONE = 1
	CA_TP_ASZZ = 2  //资产-纸质
	CA_TP_ASFC = 3  //资产-房产
	CA_TP_ASQY = 4  //资产-企业
	CA_TP_ASOT = 5  //资产-其他
	CA_TP_LBFC = 10 //负债-房产
	CA_TP_LBQY = 11 //负债-企业
	CA_TP_LBOT = 12 //负债-其他
	CA_TP_INZZ = 20 //收入-纸质
	CA_TP_INFC = 21 //收入-房产
	CA_TP_INOT = 22 //收入-其他
)

type Career struct {
	cid         int        // 职业ID
	name        string     // 职业名称
	money       int        // 存款
	wege        int        // 工资
	extraIn     int        // 被动收入
	cflow       int        // 现金流
	childNum    int        // 孩子个数
	childCost   int        // 每个孩子支付
	totalCost   int        // 总支付
	liability   []TabGroup // 负债详情
	assets      []TabGroup // 资产详情
	expenditure []TabGroup // 支出详情详情
	income      []TabGroup // 收入详情
	cards       []*modules.RMCard
}

// 资产
type TabGroup struct {
	mtype int    // 类型
	title string // 名称
	value int    // 总额
	extra int    // 附加字段
}

// 创建一个职业实例
func (c Career) Create(cid int) {
	if cid < 0 || cid > 11 {
		log.Error("career.Create error! cid is wrong number. ")
		return
	}

	data := CarrerData[cid]
	c.cid = cid
	c.name = nameData[cid]
	c.money = data[12]
	c.wege = data[0]
	c.extraIn = 0
	c.cflow = data[11]
	c.childNum = 0
	c.childCost = data[9]
	c.totalCost = data[10]

	// 负债
	tab := new(TabGroup)
	tab.mtype = CA_TP_NONE
	tab.title = "住宅抵押贷款"
	tab.value = data[13]
	tab.extra = 0
	c.liability[0] = *tab

	tab = new(TabGroup)
	tab.mtype = CA_TP_NONE
	tab.title = "教育贷款"
	tab.value = data[14]
	tab.extra = 0
	c.liability[1] = *tab

	tab = new(TabGroup)
	tab.mtype = CA_TP_NONE
	tab.title = "购车贷款"
	tab.value = data[15]
	tab.extra = 0
	c.liability[2] = *tab

	tab = new(TabGroup)
	tab.mtype = CA_TP_NONE
	tab.title = "信用卡"
	tab.value = data[16]
	tab.extra = 0
	c.liability[3] = *tab

	tab = new(TabGroup)
	tab.mtype = CA_TP_NONE
	tab.title = "额外负债"
	tab.value = data[17]
	tab.extra = 0
	c.liability[4] = *tab

	// 支出
	tab = new(TabGroup)
	tab.mtype = CA_TP_NONE
	tab.title = "税金"
	tab.value = data[2]
	tab.extra = 0
	c.expenditure[0] = *tab

	tab = new(TabGroup)
	tab.mtype = CA_TP_NONE
	tab.title = "住房抵押贷款/房租"
	tab.value = data[3]
	tab.extra = 0
	c.expenditure[1] = *tab

	tab = new(TabGroup)
	tab.mtype = CA_TP_NONE
	tab.title = "教育贷款"
	tab.value = data[4]
	tab.extra = 0
	c.expenditure[2] = *tab

	tab = new(TabGroup)
	tab.mtype = CA_TP_NONE
	tab.title = "购车贷款"
	tab.value = data[5]
	tab.extra = 0
	c.expenditure[3] = *tab

	tab = new(TabGroup)
	tab.mtype = CA_TP_NONE
	tab.title = "信用卡支出"
	tab.value = data[6]
	tab.extra = 0
	c.expenditure[4] = *tab

	tab = new(TabGroup)
	tab.mtype = CA_TP_NONE
	tab.title = "额外支出"
	tab.value = data[7]
	tab.extra = 0
	c.expenditure[5] = *tab

	tab = new(TabGroup)
	tab.mtype = CA_TP_NONE
	tab.title = "其他支出"
	tab.value = data[8]
	tab.extra = 0
	c.expenditure[6] = *tab
}
