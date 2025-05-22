package model

// AppPayRefundMarketingRule 交易退款营销规则
type AppPayRefundMarketingRule struct {
	RefundMarketingMerchantId string  `json:"refundMarketingMerchantNo"`        // 退款营销商编
	RefundMarketingAmount     float64 `json:"refundMarketingAmount"`            // 营销金额
	RefundCouponMerchantNo    string  `json:"refundCouponMerchantNo,omitempty"` // 卡券商编
}

// AppPayRefundSplitRule 订单退款分账规则
type AppPayRefundSplitRule struct {
	MerchantId             string  `json:"merchantNo"`                       // 分账商户号
	RefundAmount           float64 `json:"refundAmount"`                     // 分账金额
	SplitBillMerchantEmail string  `json:"splitBillMerchantEmail,omitempty"` // 分账商户的邮箱
	SplitBillOrderNum      string  `json:"splitBillOrderNum,omitempty"`      // 分账平台流水号(交易系统生成)
	SplitBillRequestNo     string  `json:"splitBillRequestNo,omitempty"`     // 分账订单号(商户上送的原分账订单号)
	RefundStatus           string  `json:"refundStatus,omitempty"`           // 分账状态SUCCESS-成功 FAILED-失败
	WithPayAdvanceTrx      string  `json:"withPayAdvanceTrx,omitempty"`      // 是否垫付true:是;空或false:否
}

// AppPayOrderCloseReq 交易订单关闭接口
type AppPayOrderCloseReq struct {
	InterfaceName string `json:"interfaceName"`        // 接口类型,固定为:AppPayClose
	MerchantId    string `json:"-"`                    // 商户编号
	PayOrderNo    string `json:"payOrderNo,omitempty"` // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一
}

// AppPayOrderCloseRes 交易订单关闭接口
type AppPayOrderCloseRes struct {
	InterfaceName    string `json:"interfaceName"`              // 接口类型,固定为:AppPayClose
	MerchantId       string `json:"merchantId"`                 // 商户编号,支付系统分配的商户号
	PayOrderNo       string `json:"payOrderNo,omitempty"`       // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一
	PayChannelNumber string `json:"payChannelNumber,omitempty"` // 上游订单号,上游交易请求订单号
}

// AppPayOrderCancelReq 交易订单撤销接口
type AppPayOrderCancelReq struct {
	InterfaceName string `json:"interfaceName"`        // 接口类型,固定为:AppPayCancel
	MerchantId    string `json:"-"`                    // 商户编号
	PayOrderNo    string `json:"payOrderNo,omitempty"` // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一
}

// AppPayOrderCancelRes 交易订单撤销接口
type AppPayOrderCancelRes struct {
	InterfaceName    string `json:"interfaceName"`              // 接口类型,固定为:AppPayCancel
	MerchantId       string `json:"merchantId"`                 // 商户编号,支付系统分配的商户号
	PayOrderNo       string `json:"payOrderNo,omitempty"`       // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一
	PayChannelNumber string `json:"payChannelNumber,omitempty"` // 上游订单号,上游交易请求订单号
	Recall           string `json:"recall,omitempty"`           // 重试标识:Y/N,是否需要继续调用撤销,Y- 需要,N-不需要;上游通道撤销未返回明确结果时可以重试调用
}

// AppPayOrderRefundReq 交易订单退款接口
type AppPayOrderRefundReq struct {
	InterfaceName    string  `json:"interfaceName"`         // 交易类型,固定值:AppPayRefund
	MerchantId       string  `json:"-"`                     // 商户编号
	PayOrderNo       string  `json:"payOrderNo,omitempty"`  // 交易请求订单号,要退款的交易订单的商户请求订单号与”上游订单号”二选一
	PayChannelNumber string  `json:"payChannelNumber"`      // 上游请求订单号,要退款的交易订单的上游订单号与”交易请求订单号”二选一
	RefundOrderNo    string  `json:"refundOrderNo"`         // 退款订单号,商户退款请求订单号，商户号下唯一
	RefundAmount     float64 `json:"refundAmount"`          // 退款金额,以元为单位,最小金额为0.01
	RefundDesc       string  `json:"refundDesc,omitempty"`  // 退款原因/备注,会在下发给用户的退款账单消息中体现退款原因
	NotifyUrl        string  `json:"notifyUrl,omitempty"`   // 通知回调地址,异步接收智付退款结果通知的回调地址,通知url必须为外网可访问的url,不能携带参数。
	AcqAddnData      string  `json:"acqAddnData,omitempty"` // 银联二维码平台收款方附加数据,收款方附加数据退款请求,上送参加银联二维码通道营销商品维度信息;银联二维码平台参加单品营销的订单发生退款时必须上送
	// Deprecated: 请勿直接赋值,应调用SplitRules添加
	RefundSplitRuleJson string                   `json:"refundSplitRules,omitempty"` // 分账退款规则JSON串,需使用DES进行对称加密
	RefundSplitRules    []*AppPayRefundSplitRule `json:"-"`                          // 分账退款规则
}

// AppPayOrderRefundRes 交易订单退款接口
type AppPayOrderRefundRes struct {
	InterfaceName       string  `json:"interfaceName"`          // 交易类型,固定值:AppPayRefund
	MerchantId          string  `json:"merchantId"`             // 商户编号,支付系统分配的商户号
	PayOrderNo          string  `json:"payOrderNo"`             // 交易请求订单号,要退款的交易订单的商户请求订单号
	RefundChannelNumber string  `json:"refundChannelNumber"`    // 上游退款订单号
	RefundOrderNo       string  `json:"refundOrderNo"`          // 退款订单号,商户退款请求订单号，商户号下唯一
	RefundAmount        float64 `json:"refundAmount,omitempty"` // 退款金额,以元为单位,最小金额为0.01
	Currency            string  `json:"currency,omitempty"`     // 币种类型,CNY:人民币
}

// AppPayOrderRefundQueryReq 交易订单退款查询接口
type AppPayOrderRefundQueryReq struct {
	InterfaceName       string `json:"interfaceName"`                 // 交易类型,固定值:AppPayRefundQuery
	MerchantId          string `json:"-"`                             // 商户编号
	RefundOrderNo       string `json:"refundOrderNo"`                 // 退款订单号,商户退款请求订单号，商户号下唯一
	RefundChannelNumber string `json:"refundChannelNumber,omitempty"` // 上游退款请求订单号
}

// AppPayOrderRefundQueryRes 交易订单退款查询接口
type AppPayOrderRefundQueryRes struct {
	InterfaceName            string                     `json:"interfaceName"`                     // 交易类型,固定值:AppPayRefundQuery
	MerchantId               string                     `json:"merchantId"`                        // 商户编号,支付系统分配的商户号
	PaymentType              string                     `json:"paymentType,omitempty"`             // 支付客户端类型,constants.PaymentType
	PaymentMethods           string                     `json:"paymentMethods"`                    // 支付类型,constants.PaymentMethods
	PayOrderNo               string                     `json:"payOrderNo"`                        // 交易请求订单号,要退款的交易订单的商户请求订单号
	RefundOrderNo            string                     `json:"refundOrderNo"`                     // 退款订单号,商户退款请求订单号，商户号下唯一
	RefundChannelNumber      string                     `json:"refundChannelNumber"`               // 上游退款订单号
	OrderStatus              string                     `json:"orderStatus,omitempty"`             // 订单状态,INIT:已接收,DOING:处理中,SUCCESS:成功,FAIL:失败,CLOSE:关闭,CANCEL:撤销
	RefundAmount             float64                    `json:"refundAmount,omitempty"`            // 退款金额,以元为单位,最小金额为0.01
	Currency                 string                     `json:"currency,omitempty"`                // 币种类型,CNY:人民币
	RefundOrderCompleteDate  string                     `json:"refundOrderCompleteDate,omitempty"` // 退款完成时间,格式:yyyy-MM-dd HH:mm:ss
	RefundChannelOrderNo     string                     `json:"refundChannelOrderNum,omitempty"`   // 上游退款支付订单号
	RefundOrderAmount        float64                    `json:"refundOrderAmount,omitempty"`       // 退款面额,以元为单位,最小金额为0.01
	RefundUserAmount         float64                    `json:"refundUserAmount,omitempty"`        // 用户实际退款到账金额,以元为单位,最小金额为0.01
	RefundFee                float64                    `json:"refundFee,omitempty"`               // 退还手续费,单位:元
	RefundFeeAccountAmt      float64                    `json:"refundFeeAccountAmt,omitempty"`     // 退还平台商补贴的手续费,单位:元
	RefundReceiverFee        float64                    `json:"refundReceiverFee,omitempty"`       // 退还实收手续费,单位:元
	RefundOfflineFee         float64                    `json:"refundOfflineFee,omitempty"`        // 退还后收手续费,单位:元
	RefundPromotionDetail    string                     `json:"refundPromotionDetail,omitempty"`   // 微信/支付宝等卡券优惠详情串
	RefundDesc               string                     `json:"refundDesc,omitempty"`              // 退款原因/备注,若商户传入,会在下发给用户的退款账单消息中体现退款原因
	RetReasonDesc            string                     `json:"retReasonDesc,omitempty"`           // 通道上游业务结果描述
	UpAddData                string                     `json:"upAddData,omitempty"`               // 银联二维码平台付款方附加数据
	RefundSplitRulesJson     string                     `json:"refundSplitRules,omitempty"`        // 分账退款规则JSON串,需使用DES进行对称加密
	RefundSplitRules         []*AppPayRefundSplitRule   `json:"-"`                                 // 分账退款规则
	RefundMarketingRulesJson string                     `json:"refundMarketingRules,omitempty"`    // 退款营销规则串,营销规则,JSON格式字符串,des加密返回,
	RefundMarketingRules     *AppPayRefundMarketingRule `json:"-"`                                 // 退款营销规则
}

// OrderRefundResultNotifyReqBody 订单退款结果异步通知
type OrderRefundResultNotifyReqBody struct {
	MerchantId               string                     `json:"merchantId"`                        // 商户编号,支付系统分配的商户号
	PaymentType              string                     `json:"paymentType,omitempty"`             // 支付客户端类型,constants.PaymentType
	PaymentMethods           string                     `json:"paymentMethods"`                    // 支付类型,constants.PaymentMethods
	PayOrderNo               string                     `json:"payOrderNo,omitempty"`              // 交易请求订单号,要退款的交易订单的商户请求订单号
	RefundOrderNo            string                     `json:"refundOrderNo"`                     // 退款订单号,商户退款请求订单号，商户号下唯一
	RefundChannelNumber      string                     `json:"refundChannelNumber,omitempty"`     // 上游退款订单号
	RefundOrderStatus        string                     `json:"refundOrderStatus"`                 // 订单状态,INIT:已接收,DOING:处理中,SUCCESS:成功,FAIL:失败,CLOSE:关闭,CANCEL:撤销
	RefundAmount             float64                    `json:"refundAmount,omitempty"`            // 退款金额,以元为单位,最小金额为0.01
	Currency                 string                     `json:"currency,omitempty"`                // 币种类型,CNY:人民币
	RefundOrderCompleteDate  string                     `json:"refundOrderCompleteDate,omitempty"` // 退款完成时间,格式:yyyy-MM-dd HH:mm:ss
	RefundNotifyDate         string                     `json:"refundNotifyDate,omitempty"`        // 通知时间,格式:yyyy-MM-dd HH:mm:ss
	RefundChannelOrderNo     string                     `json:"refundChannelOrderNum,omitempty"`   // 上游退款支付订单号
	RefundOrderAmount        float64                    `json:"refundOrderAmount,omitempty"`       // 退款面额,以元为单位,最小金额为0.01
	RefundUserAmount         float64                    `json:"refundUserAmount,omitempty"`        // 用户实际退款到账金额,以元为单位,最小金额为0.01
	RefundFee                float64                    `json:"refundFee,omitempty"`               // 退还手续费,单位:元
	RefundFeeAccountAmt      float64                    `json:"refundFeeAccountAmt,omitempty"`     // 退还平台商补贴的手续费,单位:元
	RefundReceiverFee        float64                    `json:"refundReceiverFee,omitempty"`       // 退还实收手续费,单位:元
	RefundOfflineFee         float64                    `json:"refundOfflineFee,omitempty"`        // 退还后收手续费,单位:元
	RefundPromotionDetail    string                     `json:"refundPromotionDetail,omitempty"`   // 微信/支付宝等卡券优惠详情串
	RefundDesc               string                     `json:"refundDesc,omitempty"`              // 退款原因/备注,若商户传入,会在下发给用户的退款账单消息中体现退款原因
	RetReasonDesc            string                     `json:"retReasonDesc,omitempty"`           // 通道上游业务结果描述
	UpAddData                string                     `json:"upAddData,omitempty"`               // 银联二维码平台付款方附加数据
	RefundSplitRulesJson     string                     `json:"refundSplitRules,omitempty"`        // 分账退款规则JSON串,需使用DES进行对称加密
	RefundSplitRules         []*AppPayRefundSplitRule   `json:"-"`                                 // 分账退款规则
	RefundMarketingRulesJson string                     `json:"refundMarketingRules,omitempty"`    // 退款营销规则串,营销规则,JSON格式字符串,des加密返回,
	RefundMarketingRules     *AppPayRefundMarketingRule `json:"-"`                                 // 退款营销规则

}
type OrderRefundResultNotifyReq = NotifyReq[OrderRefundResultNotifyReqBody]
