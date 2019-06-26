package internal

const (
	U_STA_NONE  = 0  // 无状态
	U_STA_READY = 1  // 游戏开始前准备状态
	U_STA_OPE   = 2  // 玩家操作状态
	U_STA_MSE   = 10 // 内圈状态
	U_STA_CAT   = 11 // 外圈状态
	U_STA_LJ    = 12 // 失业等待状态
)

type GUser struct {
	uid      int64       // 用户ID
	name     string      // 用户名
	icon     int         // 头像
	caree    int         // 职业
	fsta     *FStatement //财务报表
	status   int         //操作状态
	Identity int         // 身份状态
}

type FStatement struct {
	money       int64      //现金
	cflow       int64      //现金流
	wege        int64      // 工资
	passin      int64      //被动收入
	expenditure int64      //支出
	assets      *Assets    // 资产
	liability   *Liability // 负债
}

// 玩家持有股票
type Stock struct {
	id    int   // 股票ID
	count int64 //数量
	cost  int   //成本
}
