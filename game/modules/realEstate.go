// 房地产
package modules

// 玩家持有房地产
type RealEstate struct {
	id        int     // 牌的ID
	name      string  // 名称
	buyDesc   string  // 求购描述
	sellDesc  string  // 待售描述
	minValue  int     // 最低卖价值
	maxValue  int     // 最高卖价值
	minCost   int     // 最低买成本
	maxCost   int     // 最高卖成本
	minCF     int     // 最小现金流
	maxCF     int     // 最大现金流
	buyCost   int     // 收购价
	sellValue int     // 售价
	cFlow     int     // 每月现金流
	downPay   int     // 首付
	loan      int     // 贷款
	income    float32 // 收益率

}
