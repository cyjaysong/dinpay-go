package model

// DelayedSplitReq 订单延迟分账
type DelayedSplitReq struct {
	InterfaceName string                     `json:"interfaceName"`        // 接口名称,固定值:delaySplitting
	MerchantId    string                     `json:"-"`                    // 商户编号
	DelayOrderNo  string                     `json:"delayOrderNo"`         // 分账请求订单号,同一商户号下唯一订单号
	PayOrderNo    string                     `json:"payOrderNo"`           // 原交易订单号,原收款交易商户请求流水号
	PayOrderType  string                     `json:"payOrderType"`         // 原交易产品类型,APPPAY,枚举详见附录
	SplitRules    []*AppPayPreOrderSplitRule `json:"splitRules,omitempty"` // 分账规则JSON字符串
}

// DelayedSplitRes 订单延迟分账
type DelayedSplitRes struct {
	DelayOrderNo    string `json:"delayOrderNo"`    // 分账请求订单号,同一商户号下唯一订单号
	SplittingStatus string `json:"splittingStatus"` // 分账状态,constants.SplittingStatus
}

// DelayedSplitQueryReq 订单延迟分账查询
type DelayedSplitQueryReq struct {
	InterfaceName string `json:"interfaceName"` // 接口名称,固定值:delaySplittingQuery
	MerchantId    string `json:"-"`             // 商户编号
	DelayOrderNo  string `json:"delayOrderNo"`  // 分账请求订单号,同一商户号下唯一订单号
}

// DelayedSplitQueryRes 订单延迟分账查询
type DelayedSplitQueryRes struct {
	DelayOrderNo string                     `json:"delayOrderNo"`           // 分账请求订单号,同一商户号下唯一订单号
	DelayStatus  string                     `json:"delayStatus"`            // 分账状态,constants.SplittingStatus
	StatusDesc   string                     `json:"statusDesc"`             // 交易状态描述
	CompleteTime string                     `json:"completeTime,omitempty"` // 交易完成时间,yyyy-MM-dd HH:mm:ss
	SplitRules   []*AppPayPreOrderSplitRule `json:"splitRules,omitempty"`   // 分账结果
}

// DelayedSplitBackReq 订单延迟分账退回
type DelayedSplitBackReq struct {
	InterfaceName    string                   `json:"interfaceName"`    // 接口名称,固定值:delaySplittingRfund
	MerchantId       string                   `json:"-"`                // 商户编号
	RefundOrderNo    string                   `json:"refundOrderNo"`    // 商户请求分账退回订单号,同一商户号下唯一订单号
	DelayOrderNo     string                   `json:"delayOrderNo"`     // 分账请求订单号,同一商户号下唯一订单号
	RefundSplitRules []*AppPayRefundSplitRule `json:"refundSplitRules"` // 分账规则JSON串
}

// DelayedSplitBackRes 订单延迟分账退回
type DelayedSplitBackRes struct {
	RefundOrderNo string `json:"refundOrderNo"` // 商户请求分账退回订单号,同一商户号下唯一订单号
	RefundStatus  string `json:"refundStatus"`  // 分账状态,constants.SplittingStatus
}

// DelayedSplitBackQueryReq 订单延迟分账退回查询
type DelayedSplitBackQueryReq struct {
	InterfaceName string `json:"interfaceName"` // 接口名称,固定值:delaySplittingRfundQuery
	MerchantId    string `json:"-"`             // 商户编号
	RefundOrderNo string `json:"refundOrderNo"` // 商户请求分账退回订单号,同一商户号下唯一订单号
}

// DelayedSplitBackQueryRes 订单延迟分账退回查询
type DelayedSplitBackQueryRes struct {
	RefundOrderNo    string                   `json:"refundOrderNo"`          // 商户请求分账退回订单号,同一商户号下唯一订单号
	RefundStatus     string                   `json:"refundStatus"`           // 分账状态,constants.SplittingStatus
	StatusDesc       string                   `json:"statusDesc,omitempty"`   // 状态说明
	CompleteTime     string                   `json:"completeTime,omitempty"` // 完成时间
	RefundSplitRules []*AppPayRefundSplitRule `json:"refundSplitRules"`       // 分账退回明细,分账规则JSON串
}
