package internal

// 职业

type Career struct {
	cid         int          // 职业ID
	name        string       // 职业名称
	money       int          // 存款
	wege        int          // 工资
	wegeMul     int          // 工资基数
	childCost   int          // 每个孩子花费
	liability   *Liability   // 负债详情
	assets      *Assets      // 资产详情
	expenditure *Expenditure // 支出详情详情
	income      *Income      // 收入详情
}

// 负债
type Liability struct {
	homeLoan  int // 住房贷款
	eduLoan   int // 教育贷款
	carLoan   int // 购车贷款
	credit    int // 信用卡
	extraLoan int // 额外负债
	highLoan  int // 高息贷款
	//rlEstateDebt []*RealEstate // 房地产抵押贷款
	companyDebt []*Company // 企业负债
}

// 资产
type Assets struct {
}

//收入
type Income struct {
}

//支出
type Expenditure struct {
}

// 企业
type Company struct {
	id      int // 企业类型
	cost    int //成本
	downPay int //首付
	loan    int // 贷款
}
