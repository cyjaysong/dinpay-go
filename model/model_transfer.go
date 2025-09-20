package model

// TransferOrderReq 商户代付下单
type TransferOrderReq struct {
	InterfaceName string  `json:"interfaceName"`           // 接口名称:固定值:accountPay
	MerchantId    string  `json:"-"`                       // 商户编号
	OrderNo       string  `json:"transferOrderNo"`         // 代付订单号,代付订单发起时的唯一订单号
	Amount        float64 `json:"transferAmount"`          // 代付金额,金额单位为元,最少值0.01
	CardNo        string  `json:"cardNo"`                  // 收款银行卡号
	CardName      string  `json:"cardName"`                // 收款银行账户名
	CardType      string  `json:"cardType"`                // 收款账户类型,B2B:对公,B2C:对私
	BankCode      string  `json:"bankCode"`                // 收款银行编码
	BankUnionCode string  `json:"bankUnionCode,omitempty"` // 银行联行号,对公或对私单笔超过5w联行号必填
	FeeUndertaker string  `json:"feeUndertaker"`           // 手续费承担方,PAYER:付款方承担;RECEIVER:收款方承担
	NotifyUrl     string  `json:"notifyUrl"`               // 通知地址,商户用于接收代付结果通知的地址,notifyUrl接收通知成功后必须返回:"success";
	Urgency       bool    `json:"urgency"`                 // 固定值;true:加急
	TransferDesc  string  `json:"transferDesc"`            // 代付备注;通用备注:结算、工资奖金、供应链款项、交易退款、公共事业费用
}

// TransferOrderRes 商户代付下单
type TransferOrderRes struct {
	OrderNo       string `json:"transferOrderNo"`         // 代付订单发起时的代付订单号
	ChannelNumber string `json:"channelNumber,omitempty"` // 平台流水号,虚拟账户支付唯一平台流水号
}

// TransferQueryReq 商户代付订单查询
type TransferQueryReq struct {
	InterfaceName string `json:"interfaceName"`   // 接口名称:固定值:accountPayQuery
	MerchantId    string `json:"-"`               // 商户编号
	OrderNo       string `json:"transferOrderNo"` // 代付订单号,代付订单发起时的代付订单号
}

// TransferQueryRes 商户代付订单查询
type TransferQueryRes struct {
	OrderNo        string  `json:"transferOrderNo"`          // 代付订单发起时的代付订单号
	Amount         float64 `json:"transferAmount,omitempty"` // 代付金额,金额单位为元,最少值0.01
	TransferStatus string  `json:"transferStatus,omitempty"` // 代付状态,RECEIVE:已接收;INIT:初始化;DOING:处理中;SUCCESS:成功;FAIL:失败
	ChannelNumber  string  `json:"channelNumber,omitempty"`  // 平台流水号,虚拟账户支付唯一平台流水号
	ChannelMsg     string  `json:"channelMsg,omitempty"`     // 代付返回信息
	CreateDate     string  `json:"createDate,omitempty"`     // 订单创建时间,创建时间:格式为:yyyy-MM-dd HH:mm:ss
	CompleteDate   string  `json:"completeDate,omitempty"`   // 订单完成时间,完成时间:格式为:yyyy-MM-dd HH:mm:ss
}

// TransferNotifyReqBody 商户代付异步通知Body
type TransferNotifyReqBody struct {
	InterfaceName  string  `json:"interfaceName"`            // 接口名称:固定值:Transfer
	OrderNo        string  `json:"transferOrderNo"`          // 代付订单发起时的代付订单号
	Amount         float64 `json:"transferAmount,omitempty"` // 代付金额,金额单位为元,最少值0.01
	TransferStatus string  `json:"transferStatus"`           // 代付状态,RECEIVE:已接收;INIT:初始化;DOING:处理中;SUCCESS:成功;FAIL:失败
	ChannelNumber  string  `json:"channelNumber"`            // 平台流水号,虚拟账户支付唯一平台流水号
	ChannelMsg     string  `json:"channelMsg"`               // 代付返回信息
	CreateDate     string  `json:"createDate"`               // 订单创建时间,创建时间:格式为:yyyy-MM-dd HH:mm:ss
	CompleteDate   string  `json:"completeDate"`             // 订单完成时间,完成时间:格式为:yyyy-MM-dd HH:mm:ss
	NotifyType     string  `json:"notifyType"`               // 通知类型,ORDER_STATUS:普通通知;RETURN_REMITTANCE	:退汇通知
}
type TransferNotifyReq = NotifyReq[TransferNotifyReqBody]
