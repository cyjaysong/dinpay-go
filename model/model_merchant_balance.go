package model

// MerchantBalanceQueryReq 商户余额查询接口
type MerchantBalanceQueryReq struct {
	InterfaceName string `json:"interfaceName"` // 交易类型 商户账户查询接口
	MerchantId    string `json:"merchantId"`    // 商户编号 商户号由系统分配。注:平台商子商户模式下，查询子商户余额时值为子商户号
}

// MerchantBalanceQueryRes 商户余额查询接口
type MerchantBalanceQueryRes struct {
	MerchantId       string  `json:"merchantId"`              // 商户编号 商户号由系统分配。注：平台商子商户模式下，查询子商户余额时值为子商户号
	AccountStatus    string  `json:"accountStatus"`           // 账户状态 AVAILABLE("正常",1), FROZEN("冻结",2), FREEZE_DEBIT("冻结出帐",3), FREEZE_CREDIT("冻结入账",4), CANCELLED("已注销",5);
	Currency         string  `json:"currency"`                // 币种
	CreateDate       string  `json:"createDate"`              // 创建时间 格式:yyyy-MM-dd HH:mm:ss
	MerchantBalance  float64 `json:"merchantBalance,string"`  // 账户余额 订单金额，以元为单位，最小金额为0.01；
	D0Balance        float64 `json:"d0Balance,string"`        // D0余额 订单金额，以元为单位，最小金额为0.01
	D1Balance        float64 `json:"d1Balance,string"`        // D1余额 已经过了一个自然日的资金余额
	T1Balance        float64 `json:"t1Balance,string"`        // T1余额 订单金额，以元为单位，最小金额为0.01
	RechargeBalance  float64 `json:"rechargeBalance,string"`  // 充值金额 充值金额，根据结算周期清0（T1则次工作日清0，D1则次日清0）
	FrozenBalance    float64 `json:"frozenBalance,string"`    // 冻结余额 订单金额，以元为单位，最小金额为0.01； 因触犯风控规则而进行冻结的金额；
	SettlableBalance float64 `json:"settlableBalance,string"` // 可结算金额 商户可出款金额；
	// SettlableBalance 计算公式：
	// 可结算金额（D0结算周期）= 账户余额 – 冻结余额；
	// 可结算金额（D1结算周期）= D1余额+充值金额-冻结余额；
	// 可结算金额（T1结算周期）= T1余额+充值金额-冻结余额；
	Desc string `json:"desc,omitempty"` // 备注 备注信息
}
