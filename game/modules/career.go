package modules

import (
	"github.com/name5566/leaf/log"

	"server/language"
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
	nameData = [12]string{"str_ca_ba", "str_ca_xcsf", "str_ca_gjsj", "str_ca_qyms", "str_ca_jc", "str_ca_yyhs", "str_ca_zxjs", "str_ca_jl", "str_ca_gcs", "str_ca_ls", "str_ca_jz", "str_ca_ys"}
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
	CA_TP_ASZZ = 11 //资产-纸质
	CA_TP_ASFC = 12 //资产-房产
	CA_TP_ASQY = 13 //资产-企业
	CA_TP_ASOT = 14 //资产-其他
	CA_TP_LBFC = 22 //负债-房产
	CA_TP_LBQY = 23 //负债-企业
	CA_TP_LBOT = 24 //负债-其他
	CA_TP_INZZ = 31 //收入-纸质
	CA_TP_INFC = 32 //收入-房产
	CA_TP_INQY = 33 //收入-企业
	CA_TP_INOT = 34 //收入-其他
	CA_TP_XYK  = 41 //支出-信用卡
	CA_TP_GXDK = 42 //支出-高息贷款
	CA_TP_XH   = 43 //支出-小孩
)

// 玩家状态
const (
	CA_SAT_NONE = 0 // 无状态
	CA_SAT_CS   = 1 // 慈善状态
	CA_SAT_SY   = 2 // 失业状态
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
	Status      int        // 状态
	liability   []TabGroup // 负债详情
	assets      []TabGroup // 资产详情
	expenditure []TabGroup // 支出详情详情
	income      []TabGroup // 收入详情
	cards       []*RMCard
}

// 资产
type TabGroup struct {
	cid   int    // 对应的卡片ID,职业初始资产，id默认为-1
	mtype int    // 类型
	title string // 名称
	value int    // 总额
	extra int    // 附加字段
}

// 创建一个职业实例
func (c *Career) Create(cid int) {
	if cid < 0 || cid > 11 {
		log.Error("career.Create error! cid is wrong number. ")
		return
	}

	data := CarrerData[cid]
	c.cid = cid
	c.name = language.Get(nameData[cid])
	c.money = data[12]
	c.wege = data[0]
	c.extraIn = 0
	c.cflow = data[11]
	c.childNum = 0
	c.childCost = data[9]
	c.totalCost = data[10]
	c.Status = CA_SAT_NONE

	// 负债
	tab := new(TabGroup)
	tab.mtype = CA_TP_NONE
	tab.title = language.Get("str_ca_dkzz")
	tab.value = data[13]
	tab.extra = 0
	tab.cid = -1
	c.liability[0] = *tab

	tab = new(TabGroup)
	tab.mtype = CA_TP_NONE
	tab.title = language.Get("str_ca_dkjy")
	tab.value = data[14]
	tab.extra = 0
	tab.cid = -1
	c.liability[1] = *tab

	tab = new(TabGroup)
	tab.mtype = CA_TP_NONE
	tab.title = language.Get("str_ca_dkgc")
	tab.value = data[15]
	tab.extra = 0
	tab.cid = -1
	c.liability[2] = *tab

	tab = new(TabGroup)
	tab.mtype = CA_TP_XYK
	tab.title = language.Get("str_ca_dkxy")
	tab.value = data[16]
	tab.extra = 0
	tab.cid = CA_ID_XYK
	c.liability[3] = *tab

	tab = new(TabGroup)
	tab.mtype = CA_TP_NONE
	tab.title = language.Get("str_ca_dkew")
	tab.value = data[17]
	tab.extra = 0
	tab.cid = -1
	c.liability[4] = *tab

	//高息贷款
	tab = new(TabGroup)
	tab.mtype = CA_TP_GXDK
	tab.title = language.Get("str_ca_dkgx")
	tab.value = 0
	tab.extra = 0
	tab.cid = CA_ID_GXDK
	c.liability[5] = *tab

	// 支出
	//税金
	tab = new(TabGroup)
	tab.mtype = CA_TP_NONE
	tab.title = language.Get("str_ca_zcsj")
	tab.value = data[2]
	tab.extra = 0
	tab.cid = -1
	c.expenditure[0] = *tab

	//住房支出
	tab = new(TabGroup)
	tab.mtype = CA_TP_NONE
	tab.title = language.Get("str_ca_zczz")
	tab.value = data[3]
	tab.extra = 0
	tab.cid = -1
	c.expenditure[1] = *tab

	//教育贷款
	tab = new(TabGroup)
	tab.mtype = CA_TP_NONE
	tab.title = language.Get("str_ca_zcjy")
	tab.value = data[4]
	tab.extra = 0
	tab.cid = -1
	c.expenditure[2] = *tab

	//购车贷款
	tab = new(TabGroup)
	tab.mtype = CA_TP_NONE
	tab.title = language.Get("str_ca_zcgc")
	tab.value = data[5]
	tab.extra = 0
	tab.cid = -1
	c.expenditure[3] = *tab

	//信用卡支出
	tab = new(TabGroup)
	tab.mtype = CA_TP_XYK
	tab.title = language.Get("str_ca_zcxy")
	tab.value = data[6]
	tab.extra = 0
	tab.cid = CA_ID_XYK
	c.expenditure[4] = *tab

	//额外支出
	tab = new(TabGroup)
	tab.mtype = CA_TP_NONE
	tab.title = language.Get("str_ca_zcew")
	tab.value = data[7]
	tab.extra = 0
	tab.cid = -1
	c.expenditure[5] = *tab

	//其他支出
	tab = new(TabGroup)
	tab.mtype = CA_TP_NONE
	tab.title = language.Get("str_ca_zcot")
	tab.value = data[8]
	tab.extra = 0
	tab.cid = -1
	c.expenditure[6] = *tab

	//高息贷款
	tab = new(TabGroup)
	tab.mtype = CA_TP_GXDK
	tab.title = language.Get("str_ca_gxdk")
	tab.value = 0
	tab.extra = 0
	tab.cid = CA_ID_GXDK
	c.expenditure[7] = *tab

}

// 添加一项负债详情
func (c *Career) AddLiability(tab TabGroup) {
	c.liability[len(c.liability)] = tab
}

// 删除负债，num表示数量
func (c *Career) DelLiability(id int) {
	for i := 0; i < len(c.liability); i++ {
		if c.liability[i].cid == id {
			c.liability = append(c.liability[:i], c.liability[i+1:]...)
		}
	}
}

// 修改负债数据
func (c *Career) ChangeLiability(id int, value int) {
	for i := 0; i < len(c.liability); i++ {
		if c.liability[i].cid == id {
			c.liability[i].value += value
			return
		}
	}
}

// 添加一项资产详情
func (c *Career) AddAssets(tab TabGroup) {
	if c.ChangeAssets(tab.cid, tab.extra) {
		return
	}

	c.assets[len(c.assets)] = tab
}

// 删除资产详情
func (c *Career) DelAssets(id int) {
	for i := 0; i < len(c.assets); i++ {
		if c.assets[i].cid == id {
			c.assets = append(c.assets[:i], c.assets[i+1:]...)
			return
		}
	}
}

// 修改持有资产的数量
func (c *Career) ChangeAssets(id int, count int) bool {
	if id >= 117 && id <= 141 {
		for i := 0; i < len(c.assets); i++ {
			if c.assets[i].cid == id {
				c.assets[i].extra += count
				return true
			}
		}
	}

	return false
}

// 支出详情,只有高息贷款和信用卡可以修改
func (c *Career) AddExpenditure(tab TabGroup) {
	c.totalCost -= tab.value
	c.changeCflow(tab.value)

	if tab.mtype == CA_TP_XH { // 小孩支出
		c.childNum += 1
	}

	for i := 0; i < len(c.expenditure); i++ {
		if c.expenditure[i].mtype == tab.mtype {
			c.expenditure[i].value += tab.value

			return
		}
	}

	c.expenditure[len(c.expenditure)] = tab
}

func (c *Career) DelExpenditure(mtype int, value int) {
	for i := 0; i < len(c.expenditure); i++ {
		if c.expenditure[i].mtype == mtype {
			c.totalCost += value
			c.changeCflow(-value)

			c.expenditure[i].value -= value

			if c.expenditure[i].value <= 0 {
				c.expenditure[i].value = 0
			}
			break
		}
	}
}

//收入详情
func (c *Career) AddIncome(tab TabGroup) {
	c.changeCflow(tab.value)
	c.extraIn += tab.value

	c.income[len(c.income)] = tab
}

func (c *Career) DelIncome(id int, value int) {
	c.changeCflow(-value)
	c.extraIn -= value

	for i := 0; i < len(c.income); i++ {
		if c.income[i].cid == id {
			c.income = append(c.income[:i], c.income[i+1:]...)

			break
		}
	}
}

// 股票类卡片删除,gtype 类型， num 数量
func (c *Career) DelStock(card *RMCard) {
	var num = card.Count

	for i := 0; i < len(c.cards); i++ {
		if c.cards[i].GroupType == card.GroupType {
			if c.cards[i].Count <= num {
				num -= c.cards[i].Count

				c.DelAssets(c.cards[i].Cid)
				c.DelIncome(c.cards[i].Cid, c.cards[i].Cflow)

				if num <= 0 {
					return
				}

			} else {
				c.cards[i].Count -= num

				c.ChangeAssets(card.Cid, -num)
			}
		}
	}

}

func (c *Career) AddCard(card *RMCard) {
	if card.AttrType == ATTR_TP_XMM || card.AttrType == ATTR_TP_DMM {
		c.processChance(card)
	} else if card.AttrType == ATTR_TP_SCFY {
		c.processMart(card)
	} else if card.AttrType == ATTR_TP_YW {
		c.processAccident(card)
	} else if card.AttrType == ATTR_TP_TDK {
		c.processOther(card)
	}
}

func (c *Career) DelCard(card *RMCard) {
	if card.Cid >= 117 && card.Cid <= 141 {
		c.DelStock(card)
	} else if card.AttrType == ATTR_TP_XMM || card.AttrType == ATTR_TP_DMM {
		c.DelLiability(card.Cid)
		c.DelAssets(card.Cid)
		c.DelIncome(card.Cid, card.Cflow)
	} else { // 额外开支
		if card.GroupType == GP_TP_DK { // 高息贷款
			c.DelExpenditure(CA_TP_GXDK, card.Cflow)
		} else {
			c.DelExpenditure(CA_TP_XYK, card.Cflow)
		}
	}
}

// 处理机会卡
func (c *Career) processChance(card *RMCard) {
	if card.GroupType == GP_TP_XGK {
		c.processEffect(card)
		return
	}

	// 负债表
	litab := c.newLiability(card)
	// 资产
	astab := c.newAssets(card)
	// 收入
	intab := c.newIncome(card)

	c.AddLiability(*litab)
	c.AddAssets(*astab)
	c.AddIncome(*intab)

	c.changeMoney(-card.Payment * card.Count)
	c.cards[len(c.cards)] = card
}

// 处理市场风云卡
func (c *Career) processMart(card *RMCard) {
	if card.GroupType == GP_TP_XGK {
		c.processEffect(card)
		return
	}

}

// 处理意外卡
func (c *Career) processAccident(card *RMCard) {

}

// 处理特定卡
func (c *Career) processOther(card *RMCard) {
	if card.Cid == CA_ID_XH { // 小孩卡
		card.Cflow = -c.childCost
		tap := c.newExpenditure(card)
		c.AddExpenditure(*tap)

	} else if card.Cid == CA_ID_SY { // 失业卡
		c.changeMoney(-c.wege)
		c.Status = CA_SAT_SY

	} else if card.Cid == CA_ID_FGZ { // 银行结算
		c.changeMoney(c.cflow)

	} else if card.Cid == CA_ID_CSJK { // 慈善捐款
		c.changeMoney(c.money / 10)
		c.Status = CA_SAT_CS
	} else if card.Cid == CA_ID_XYK { // 信用卡
		c.ChangeLiability(card.Cid, card.Loan)
		tap := c.newExpenditure(card)
		c.AddExpenditure(*tap)

	} else if card.Cid == CA_ID_GXDK { // 高息贷款
		c.ChangeLiability(card.Cid, card.Loan)
		tap := c.newExpenditure(card)
		c.AddExpenditure(*tap)
	}

}

// 处理效果卡
func (c *Career) processEffect(card *RMCard) {
	if card.Cid == 124 {
		for i := 0; i < len(c.cards); i++ {
			if c.cards[i].GroupType == GP_TP_ALBB {
				num := c.cards[i].Count
				c.cards[i].Count = num * 2

				c.ChangeAssets(c.cards[i].Cid, num)
			}
		}
	} else if card.Cid == 125 {
		for i := 0; i < len(c.cards); i++ {
			if c.cards[i].GroupType == GP_TP_ALBB {
				num := c.cards[i].Count
				c.cards[i].Count = num / 2

				c.ChangeAssets(c.cards[i].Cid, -num)
			}
		}
	} else if card.Cid == 136 {
		for i := 0; i < len(c.cards); i++ {
			if c.cards[i].GroupType == GP_TP_XMKJ {
				num := c.cards[i].Count
				c.cards[i].Count = num * 2

				c.ChangeAssets(c.cards[i].Cid, num)
			}
		}
	} else if card.Cid == 137 {
		for i := 0; i < len(c.cards); i++ {
			if c.cards[i].GroupType == GP_TP_XMKJ {
				num := c.cards[i].Count
				c.cards[i].Count = num / 2

				c.ChangeAssets(c.cards[i].Cid, -num)
			}
		}
	} else if card.Cid == 142 { // test 待完成

	} else if card.Cid == 145 {

	} else if card.Cid == 202 {

	} else if card.Cid == 203 {

	} else if card.Cid == 306 {
		var tmp []int

		for i := 0; i < len(c.cards); i++ {
			if c.cards[i].GroupType == GP_TP_ZZ3 {
				tmp[len(tmp)] = i
				c.DelCard(c.cards[i])
			}
		}

		for i := 0; i < len(tmp); i++ {
			n := tmp[i] - i
			c.cards = append(c.cards[:n], c.cards[n+1:]...)
		}
	}
}

// 创建一个负债
func (c *Career) newLiability(card *RMCard) *TabGroup {
	tab := new(TabGroup)
	tab.cid = card.Cid
	tab.title = card.Abbreviation
	tab.value = card.Loan
	tab.mtype = c.getTabGroupType(card.GroupType) + 20

	return tab
}

// 创建一个资产
func (c *Career) newAssets(card *RMCard) *TabGroup {
	tab := new(TabGroup)
	tab.cid = card.Cid
	tab.title = card.Abbreviation
	tab.value = card.Value
	t := c.getTabGroupType(card.GroupType)

	if t == 1 {
		tab.extra = card.Count
	} else if t == 2 {
		tab.extra = card.Payment
	} else if t == 3 {
		tab.extra = card.Payment
	} else if t == 4 {
		tab.extra = card.Payment
	}

	tab.mtype = t + 10

	return tab
}

// 创建一个收入
func (c *Career) newIncome(card *RMCard) *TabGroup {
	tab := new(TabGroup)
	tab.cid = card.Cid
	tab.title = card.Abbreviation
	tab.value = card.Cflow
	tab.mtype = c.getTabGroupType(card.Cid) + 30

	return tab
}

// 创建一个支出
func (c *Career) newExpenditure(card *RMCard) *TabGroup {
	tab := new(TabGroup)
	tab.cid = card.Cid
	tab.title = card.Abbreviation
	tab.value = card.Cflow
	if card.Cid == CA_ID_XH { // 小孩
		tab.mtype = CA_TP_XH
	} else if card.Cid == CA_ID_XYK { // 信用卡
		tab.mtype = CA_TP_XYK
	} else if card.Cid == CA_ID_GXDK { // 高息贷款
		tab.mtype = CA_TP_GXDK
	}

	return tab
}

func (c *Career) getTabGroupType(id int) int {
	if id >= GP_TP_PA && id <= GP_TP_CF200 {
		return 1
	} else if id == GP_TP_ZZ2 || id == GP_TP_ZZ3 {
		return 2
	} else if id >= GP_TP_GY2 && id <= GP_TP_60 {
		return 2
	} else if id == GP_TP_WD || id == GP_TP_YXGS {
		return 3
	} else if id >= GP_TP_ZDH && id <= GP_TP_HB {
		return 3
	} else {
		return 4
	}
}

// 处理玩家钱，破产或晋级
func (c *Career) changeMoney(money int) {
	c.money += money
}

// 处理玩家现金流，处理晋级
func (c *Career) changeCflow(money int) {
	c.cflow += money
}
