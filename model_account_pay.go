package dinpay

// AccountPayOrderReq 账户支付下单
type AccountPayOrderReq struct {
	InterfaceName string `json:"interfaceName"` // 接口名称:固定值:accountPay
	OrderNo       string `json:"orderNo"`       // 商户订单号
	MerchantId    string `json:"merchantId"`    // 付款商编
	// Deprecated: 请勿直接赋值,使用 OrderParameter
	OrderParameters string                             `json:"orderParameters"` // 订单参数
	OrderParameter  AccountPayParametersOrderParameter `json:"-"`               // 订单参数
}

// AccountPayParametersOrderParameter 账户支付订单参数结构体
type AccountPayParametersOrderParameter struct {
	RecordedMerchantId string                `json:"recordedMerchantId"`       // 收款商编,与付款商户merchantId隶属同一服务商
	OrderType          string                `json:"orderType"`                // 订单类型,ALLOWANCE:活动补贴,需关联交易订单;TRANSFER:资金划拨,无需关联交易订单
	PayAmount          float64               `json:"payAmount"`                // 订单金额,订单金额:不能小于等于零,小数点后最多两位
	GoodsName          string                `json:"goodsName"`                // 商品名称,长度限制128
	OrderDesc          string                `json:"orderDesc,omitempty"`      // 订单备注,长度限制128
	PayProductType     string                `json:"payProductType,omitempty"` // 原订单产品类型,当orderType值为ALLOWANCE时,此字段必填;APPPAY:扫码
	PayOrderNo         string                `json:"payOrderNo,omitempty"`     // 原订单商户订单号,当orderType值为ALLOWANCE时,此字段必填
	IsGuarantee        string                `json:"isGuarantee,omitempty"`    // 担保交易标识,当orderType值为TRANSFER时,此字段可填写;true:需要担保。交易不会直接成功,资金临时冻结直到完成担保确认,false:不需要担保,交易直接完成
	SplitRules         []AccountPaySplitRule `json:"splitRules,omitempty"`     // 收款分账规则,当orderType值为TRANSFER时,此字段可填写;若用于支持多个收款方,此时请求的recordedMerchantId、payAmount字段值可不填写。
	FeeUndertaker      string                `json:"feeUndertaker,omitempty"`  // 手续费承担方向,PAYER:表示付款方承担;RECEIVER:收款方承担:如需平台商手续费账户承担,请联系智付人员后台配置
	NotifyUrl          string                `json:"notifyUrl,omitempty"`      // 后台通知地址,订单完成支付之后,异步通知商户支付结果;若此字段填写则值不能为 null
}

// AccountPaySplitRule 账户支付订单分账规则结构体
type AccountPaySplitRule struct {
	SplitBillMerchantId string  `json:"splitBillMerchantNo"` // 分账商户号
	SplitBillAmount     float64 `json:"splitBillAmount"`     // 分账金额
	SplitBillFee        float64 `json:"splitBillFee"`        // 分账手续费
}

// AccountPayOrderRes 账户支付下单
type AccountPayOrderRes struct {
	OrderNo       string `json:"orderNo,omitempty"`       // 商户订单号 120000000000001
	MerchantId    string `json:"merchantId,omitempty"`    // 商户编号 D10000000000002
	ChannelNumber string `json:"channelNumber,omitempty"` // 平台流水号,虚拟账户支付唯一平台流水号
	OrderStatus   string `json:"orderStatus,omitempty"`   // 订单状态
}

// AccountPayQueryReq 账户支付订单查询
type AccountPayQueryReq struct {
	InterfaceName string `json:"interfaceName"` // 接口名称:固定值:accountPayQuery
	OrderNo       string `json:"orderNo"`       // 商户订单号
	MerchantId    string `json:"merchantId"`    // 商户编号
}

// AccountPayQueryRes 账户支付订单查询
type AccountPayQueryRes struct {
	OrderNo          string  `json:"orderNo,omitempty"`          // 商户订单号
	MerchantId       string  `json:"merchantId,omitempty"`       // 商户编号
	PayAmount        float64 `json:"payAmount,string,omitempty"` // 订单金额 PS:接口返回的是字符串
	ChannelNumber    string  `json:"channelNumber,omitempty"`    // 平台流水号
	OrderStatus      string  `json:"orderStatus,omitempty"`      // 订单状态
	StatusIllustrate string  `json:"statusIllustrate,omitempty"` // 订单状态说明
}

// AccountPayGuaranteeConfirmReq 账户支付担保确认
type AccountPayGuaranteeConfirmReq struct {
	InterfaceName string `json:"interfaceName"` // 接口名称:固定值:guaranteeConfirm
	OrderNo       string `json:"orderNo"`       // 商户订单号
	MerchantId    string `json:"merchantId"`    // 商户编号
	OperateType   string `json:"operateType"`   // 操作类型,CONFIRM:确认支付;CANCEL:撤销
}

// AccountPayGuaranteeConfirmRes 账户支付担保确认
type AccountPayGuaranteeConfirmRes struct {
	OrderNo          string `json:"orderNo,omitempty"`          // 商户订单号
	MerchantId       string `json:"merchantId,omitempty"`       // 商户编号
	PayAmount        string `json:"payAmount,string,omitempty"` // 订单金额
	ChannelNumber    string `json:"channelNumber,omitempty"`    // 平台流水号 虚拟账户支付唯一平台流水号
	OrderStatus      string `json:"orderStatus,omitempty"`      // 订单状态 状态详见以下说明
	StatusIllustrate string `json:"statusIllustrate,omitempty"` // 订单状态说明 商户余额不足
}

// AccountPayTransferValidateCodeReq 资金划拨类账户支付获取验证码接口
type AccountPayTransferValidateCodeReq struct {
	InterfaceName string `json:"interfaceName"` // 接口名称:固定值:accountPayValidateCode
	OrderNo       string `json:"orderNo"`       // 商户订单号
	MerchantId    string `json:"merchantId"`    // 商户编号
}

// AccountPayTransferValidateCodeRes 资金划拨类账户支付获取验证码接口
type AccountPayTransferValidateCodeRes struct {
	OrderNo    string `json:"orderNo,omitempty"`    // 商户订单号
	MerchantId string `json:"merchantId,omitempty"` // 商户编号
}

// AccountPayTransferConfirmReq 资金划拨类账户支付确认接口
type AccountPayTransferConfirmReq struct {
	InterfaceName string `json:"interfaceName"` // 接口名称:固定值:accountPayConfirmPay
	OrderNo       string `json:"orderNo"`       // 商户订单号
	MerchantId    string `json:"merchantId"`    // 商户编号
	ValidateCode  string `json:"validateCode"`  // 手机短信验证码
}

// AccountPayTransferConfirmRes 资金划拨类账户支付确认接口
type AccountPayTransferConfirmRes struct {
	OrderNo       string `json:"orderNo,omitempty"`       // 商户订单号
	MerchantId    string `json:"merchantId,omitempty"`    // 商户编号
	ChannelNumber string `json:"channelNumber,omitempty"` // 虚拟账户支付通道流水号
	OrderStatus   string `json:"orderStatus,omitempty"`   // 订单状态
}

// AccountPayOrderRefundReq 账户支付退款接口
type AccountPayOrderRefundReq struct {
	InterfaceName string `json:"interfaceName"`       // 接口名称:固定值:accountPayRefund
	OrderNo       string `json:"orderNo"`             // 商户原交易订单号 原交易订单订单号
	MerchantId    string `json:"merchantId"`          // 退款发起方商编 商户编号 原交易订单付款方商编
	RefundOrderNo string `json:"refundOrderNo"`       // 退款订单号 商户自定义退款订单号，同一商户退款订单号唯一
	NotifyUrl     string `json:"notifyUrl,omitempty"` // 异步回调地址 订单终态时回调通知商户地址，商户需返回“SUCCESS”
	// Deprecated: 请勿直接赋值,使用 RefundDetails
	RefundDetail  string                            `json:"refundDetail"` // 退款详细参数 JSON 字符串，仅必须包含退款商户编号以及退款金额
	RefundDetails []AccountPayOrderRefundDetailInfo `json:"-"`            // 退款详细参数
}

// AccountPayOrderRefundDetailInfo 定义退款详细参数的结构
type AccountPayOrderRefundDetailInfo struct {
	MerchantId   string  `json:"merchantNo"`   // 退款商户号 原交易订单中入账的商户商编inMerchantNo
	RefundAmount float64 `json:"refundAmount"` // 退款金额 退款金额
}

// AccountPayOrderRefundRes 账户支付退款接口
type AccountPayOrderRefundRes struct {
	InterfaceName string  `json:"interfaceName"`           // 接口名称
	MerchantId    string  `json:"merchantId,omitempty"`    // 商户编号
	OrderNo       string  `json:"orderNo,omitempty"`       // 原交易订单号
	RefundOrderNo string  `json:"refundOrderNo,omitempty"` // 退款订单号
	RefundAmount  float64 `json:"refundAmount,omitempty"`  // 退款金额
	RefundStatus  string  `json:"refundStatus,omitempty"`  // 订单状态，退款状态以该字段为准
}

// AccountPayOrderRefundQueryReq 账户支付退款查询接口
type AccountPayOrderRefundQueryReq struct {
	InterfaceName string `json:"interfaceName"` // 接口名称:固定值:accountPayRefundQuery
	MerchantId    string `json:"merchantId"`    // 退款发起方商编 商户编号
	RefundOrderNo string `json:"refundOrderNo"` // 退款订单号 商户自定义退款订单号，同一商户退款订单号唯一
}

// AccountPayOrderRefundQueryRes 账户支付退款查询接口
type AccountPayOrderRefundQueryRes struct {
	InterfaceName string  `json:"interfaceName"`           // 接口名称
	MerchantId    string  `json:"merchantId,omitempty"`    // 商户编号
	RefundOrderNo string  `json:"refundOrderNo,omitempty"` // 退款订单号
	RefundStatus  string  `json:"refundStatus,omitempty"`  // 订单状态，退款状态以该字段为准
	RefundAmount  float64 `json:"refundAmount,omitempty"`  // 退款金额
}

// AccountPayOrderResultNotifyReqBody 账户支付回调通知
type AccountPayOrderResultNotifyReqBody struct {
	InterfaceName    string `json:"interfaceName"`              // 接口名称
	OrderNo          string `json:"orderNo"`                    // 商户订单号
	MerchantId       string `json:"merchantId"`                 // 商户编号
	PayAmount        string `json:"payAmount"`                  // 订单金额
	ChannelNumber    string `json:"channelNumber"`              // 平台流水号
	OrderStatus      string `json:"orderStatus"`                // 订单状态
	StatusIllustrate string `json:"statusIllustrate,omitempty"` // 订单状态说明
}
type AccountPayOrderResultNotifyReq = NotifyReq[AccountPayOrderResultNotifyReqBody]

// AccountPayOrderRefundNotifyReqBody 账户支付退款回调通知
type AccountPayOrderRefundNotifyReqBody struct {
	InterfaceName string  `json:"interfaceName"`                 // 接口名称
	MerchantId    string  `json:"merchantId"`                    // 商户编号
	RefundOrderNo string  `json:"refundOrderNo,omitempty"`       // 退款订单号
	RefundStatus  string  `json:"refundStatus,omitempty"`        // 订单状态，退款状态以该字段为准
	RefundAmount  float64 `json:"refundAmount,string,omitempty"` // 退款金额
}
type AccountPayOrderRefundNotifyReq = NotifyReq[AccountPayOrderRefundNotifyReqBody]
