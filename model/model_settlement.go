package model

// MerchantSettlementReq 商户结算接口
type MerchantSettlementReq struct {
	InterfaceName    string  `json:"interfaceName"`             // 固定值:MerchantSettlement
	MerchantId       string  `json:"-"`                         // 商户编号
	OrderNo          string  `json:"settlementOrderNo"`         // 商户订单号 商户系统内部订单号，要求40字符以内，同一商户号下订单号唯一
	SettlementAmount float64 `json:"settlementAmount"`          // 金额 金额单位为元
	SettlementDesc   string  `json:"settlementDesc,omitempty"`  // 结算备注
	TransferSummary  string  `json:"transferSummary,omitempty"` // 打款摘要
	NotifyUrl        string  `json:"notifyUrl,omitempty"`       // 结算结果通知地址
}

// MerchantSettlementRes 商户结算接口
type MerchantSettlementRes struct {
	OrderNo string `json:"settlementOrderNo"` // 商户订单号 商户系统内部订单号，要求40字符以内，同一商户号下订单号唯一
}

// MerchantSettlementQueryReq 商户结算查询接口
type MerchantSettlementQueryReq struct {
	InterfaceName string `json:"interfaceName"`               // 固定值:MerchantSettlementQuery
	MerchantId    string `json:"-"`                           // 商户编号
	OrderNo       string `json:"settlementOrderNo,omitempty"` // 商户订单号 商户系统内部订单号，要求40字符以内，同一商户号下订单号唯一
	SettleDate    string `json:"settleDate,omitempty"`        // 结算日期 通常都是查前一个工作日
}

// MerchantSettlementQueryRes 商户结算查询接口
type MerchantSettlementQueryRes struct {
	SettleRecords []MerchantSettlementRecord `json:"settleRecords,omitempty"` // 结算记录
}

type MerchantSettlementRecord struct {
	MerchantId            string  `json:"merchantId"`                             // 商户编号
	OrderNo               string  `json:"settlementOrderNo"`                      // 订单号 同一商户号下订单号唯一
	SettleDate            string  `json:"settleDate"`                             // 结算日期,yyyyMMdd
	SettlementAmount      float64 `json:"settlementAmount,string"`                // 结算金额
	SettleFee             float64 `json:"settleFee,string"`                       // 结算手续费
	D0Amount              float64 `json:"d0Amount,string,omitempty"`              // D0资金 加急资金
	D0AmountFee           float64 `json:"d0AmountFee,string,omitempty"`           // D0手续费 加急手续费
	IntradayAmount        float64 `json:"intradayAmount,string,omitempty"`        // Dr资金 当日加急资金
	IntradayFee           float64 `json:"intradayFee,string,omitempty"`           // Dr手续费 当日加急手续费
	NonCurrentDayD0Amount float64 `json:"nonCurrentDayD0Amount,string,omitempty"` // D0-Dr资金 非当日加急资金
	NonCurrentDayD0Fee    float64 `json:"nonCurrentDayD0Fee,string,omitempty"`    // D0-Dr手续费 非当日加急手续费
	SettleType            string  `json:"settleType"`                             // 结算类型 结算类型（T1/D1）
	SettlementStatus      string  `json:"settlementStatus"`                       // 结算状态,INIT已接收,DOING处理中,DONE成功,FAILED失败,MANUAL人工处理
	FailMsg               string  `json:"failMsg,omitempty"`                      // 失败原因
	CompletionTime        string  `json:"completionTime,omitempty"`               // 订单完成时间;格式: yyyy-MM-dd HH:mm:ss
}

// MerchantSettlementNotifyReqBody 商户结算异步通知Body
type MerchantSettlementNotifyReqBody struct {
	InterfaceName string                     `json:"interfaceName"` // 交易类型:固定值:MerchantSettlement
	SettleRecords []MerchantSettlementRecord `json:"settleRecords"` // 结算记录
}
type MerchantSettlementNotifyReq = NotifyReq[MerchantSettlementNotifyReqBody]
