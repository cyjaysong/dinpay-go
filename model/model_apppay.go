package model

// AppPayPreOrderMarketingRule 交易下单营销规则
type AppPayPreOrderMarketingRule struct {
	MarketingMerchantId string  `json:"marketingMerchantNo"`          // 营销商编
	MarketingAmount     float64 `json:"marketingAmount"`              // 营销金额
	MarketingOwnerType  string  `json:"marketingOwnerType,omitempty"` // 所属营销对象USER:用户,MERCHANT:商户,不填默认USER
	CouponMerchantNo    string  `json:"couponMerchantNo,omitempty"`   // 卡券商编
}

// AppPayPreOrderSplitRule 交易下单分账规则
type AppPayPreOrderSplitRule struct {
	SplitBillMerchantId  string  `json:"splitBillMerchantNo"`            // 分账商户号
	SplitBillAmount      float64 `json:"splitBillAmount"`                // 分账金额
	SplitBillFee         float64 `json:"splitBillFee,omitempty"`         // 分账手续费
	SplitBillOrderNum    string  `json:"splitBillOrderNum,omitempty"`    // 分账平台流水号(交易系统生成)
	SplitBillOrderStatus string  `json:"splitBillOrderStatus,omitempty"` // 分账状态SUCCESS-成功 FAILED-失败
	SplitBillRequestNo   string  `json:"splitBillRequestNo,omitempty"`   // 分账订单号(商户上送)
	SplitBillRetCode     string  `json:"splitBillRetCode,omitempty"`     // 分账错误码
	SplitBillRetMsg      string  `json:"splitBillRetMsg,omitempty"`      // 分账错误信息
}

// AppPayOrderQueryReq 交易订单查询接口
type AppPayOrderQueryReq struct {
	InterfaceName string `json:"interfaceName"`           // 接口类型,固定为:AppPayQuery
	MerchantId    string `json:"-"`                       // 商户编号
	OrderNo       string `json:"orderNo,omitempty"`       // 要查询的交易订单的商户请求订单号
	ChannelNumber string `json:"channelNumber,omitempty"` // 下单返回的上游订单号与”请求订单号”二选一
}

// AppPayOrderQueryRes 交易订单查询接口
type AppPayOrderQueryRes struct {
	InterfaceName           string                       `json:"interfaceName"`                     // 接口类型,固定为:AppPayQuery
	PaymentType             string                       `json:"paymentType,omitempty"`             // 支付客户端类型,constants.PaymentType
	PaymentMethods          string                       `json:"paymentMethods"`                    // 支付类型,constants.PaymentMethods
	MerchantId              string                       `json:"merchantId"`                        // 商户编号,支付系统分配的商户号
	OrderNo                 string                       `json:"orderNo"`                           // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一
	PayAmount               float64                      `json:"payAmount"`                         // 交易金额,以元为单位,最小金额为0.01
	Currency                string                       `json:"currency"`                          // 币种类型,CNY:人民币
	OrderStatus             string                       `json:"orderStatus,omitempty"`             // 订单状态,INIT:已接收,DOING:处理中,SUCCESS:成功,FAIL:失败,CLOSE:关闭,CANCEL:撤销
	RefundStatus            string                       `json:"refundStatus,omitempty"`            // 退款状态,PRE_REFUND:转入退款;PART_REFUND:部分退款;ALL_REFUND:全额退款;FAIL_REFUND:退款失败;NOT_YET:尚未退款,(转入退款,部分退款,全额退款,退款失败 均表示这 笔订单发起过退款)
	OrderDesc               string                       `json:"orderDesc,omitempty"`               // 订单备注信息,商户自定义上送,原样返回。
	ChannelNumber           string                       `json:"channelNumber,omitempty"`           // 上游请求订单号,智付交易订单号
	OutTransactionOrderId   string                       `json:"outTransactionOrderId,omitempty"`   // 上游交易单号,成功时有返回
	SubMerchantNo           string                       `json:"subMerchantNo,omitempty"`           // 上游子商户号(U/A/T)
	AppId                   string                       `json:"appid,omitempty"`                   // 商户公众号sub_appid
	OpenId                  string                       `json:"openid,omitempty"`                  // 微信用户openId
	SubOpenId               string                       `json:"subopenid,omitempty"`               // 微信子商户subopenid或支付宝子商户用户buyer_id
	BankType                string                       `json:"bankType,omitempty"`                // 用户付款银行,用户付款银行,成功时有返回
	OnlineCardType          string                       `json:"onlineCardType,omitempty"`          // 支付卡类型,EBIT(借记卡),CREDIT(贷记卡),UNKNOWN(未知),CFT(钱包零钱)
	CashFee                 float64                      `json:"cashFee,omitempty"`                 // 上游返回:现金支付金额(订单总金额-现金券金额=现金支付金额)
	CouponFee               float64                      `json:"couponFee,omitempty"`               // 上游返回:现金券金额:现金券金额
	OrderCreditAmount       float64                      `json:"orderCreditAmount,omitempty"`       // 订单入账面额(不扣手续费)
	MerchantCreditAmount    float64                      `json:"merchantCreditAmount,omitempty"`    // 商户实际入账发生额,收单商户实际入账发生额,成功时返回
	PaymentAmount           float64                      `json:"paymentAmount,omitempty"`           // 用户实际支付金额
	ChannelSettlementAmount float64                      `json:"channelSettlementAmount,omitempty"` // 渠道结算金额,成功时返回
	OrderFee                float64                      `json:"orderFee,omitempty"`                // 该笔交易产生的手续费,成功时返回
	MerchantFee             float64                      `json:"merchantFee,omitempty"`             // 订单商户交易的费率,成功时返回
	FeeAccountAmt           float64                      `json:"feeAccountAmt,omitempty"`           // 平台商补贴的手续费,单位:元,支付成功时返回
	ReceiverFee             float64                      `json:"receiverFee,omitempty"`             // 实收手续费,单位:元,支付成功时返回
	OfflineFee              float64                      `json:"offlineFee,omitempty"`              // 后收手续费,单位:元,支付成功时返回
	OrderPayDate            string                       `json:"orderPayDate,omitempty"`            // 订单完成时间,格式:yyyy-MM-dd HH:mm:ss
	ChargeFlag              string                       `json:"chargeFlag,omitempty"`              // 渠道支付宝费率活动标识,支付宝活动通道成功时有返回
	WxTradeType             string                       `json:"wxTradeType,omitempty"`             // 微信交易类型,微信成功时有返回
	UpAddData               string                       `json:"upAddData,omitempty"`               // 银联二维码返回的付款方附加数据,Base64编码
	ResvData                string                       `json:"resvData,omitempty"`                // 银联二维码返回的保留数据,Base64编码
	FundBillList            string                       `json:"fundBillList,omitempty"`            // 交易资金渠道和优惠信息-支付宝
	PromotionDetail         string                       `json:"promotionDetail,omitempty"`         // 上游返回的优惠详情,JSON字符串
	VoucherDetailList       string                       `json:"voucherDetailList,omitempty"`       // 支付宝优惠信息详情,JSON字符串
	MarketingRulesJson      string                       `json:"marketingRules,omitempty"`          // 营销规则,JSON格式字符串，des加密传输
	MarketingRules          *AppPayPreOrderMarketingRule `json:"-"`                                 // 营销规则
	SplitRulesJson          string                       `json:"splitRules,omitempty"`              // 分账规则及状态,响应分账结果规则以及对应状态
	SplitRules              []*AppPayPreOrderSplitRule   `json:"-"`                                 // 分账规则
}

// OrderPayResultNotifyReqBody 订单支付结果异步通知
type OrderPayResultNotifyReqBody struct {
	PaymentType             string                       `json:"paymentType"`                       // 支付客户端类型,constants.PaymentType
	PaymentMethods          string                       `json:"paymentMethods"`                    // 支付类型,constants.PaymentMethods
	MerchantId              string                       `json:"merchantId"`                        // 商户编号,支付系统分配的商户号
	OrderNo                 string                       `json:"orderNo"`                           // 请求订单号,商户请求订单号,要求50字符以内,同一商户号下订单号唯一
	PayAmount               float64                      `json:"payAmount"`                         // 交易金额，单位：元。最小金额为0.01
	Currency                string                       `json:"currency"`                          // 币种,CNY:人民币
	OrderStatus             string                       `json:"orderStatus"`                       // 订单状态,INIT:已接收,DOING:处理中,SUCCESS:成功,FAIL:失败,CLOSE:关闭,CANCEL:撤销
	RefundStatus            string                       `json:"refundStatus,omitempty"`            // 退款状态,PRE_REFUND:转入退款;PART_REFUND:部分退款;ALL_REFUND:全额退款;FAIL_REFUND:退款失败;NOT_YET:尚未退款,(转入退款,部分退款,全额退款,退款失败 均表示这 笔订单发起过退款)
	OrderDesc               string                       `json:"orderDesc,omitempty"`               // 订单备注,交易下单时上送值原样返回
	ChannelNumber           string                       `json:"channelNumber"`                     // 上游订单号,交易订单号
	OutTransactionOrderId   string                       `json:"outTransactionOrderId,omitempty"`   // 上游交易单号
	SubMerchantNo           string                       `json:"subMerchantNo,omitempty"`           // 上游子商户号(U/A/T)
	AppId                   string                       `json:"appid,omitempty"`                   // 商户公众号sub_appid
	OpenId                  string                       `json:"openid,omitempty"`                  // 微信用户openId
	SubOpenId               string                       `json:"subopenid,omitempty"`               // 微信子商户subopenid或支付宝子商户用户buyer_id
	BankType                string                       `json:"bankType,omitempty"`                // 用户付款银行,用户付款银行,成功时有返回
	OnlineCardType          string                       `json:"onlineCardType,omitempty"`          // 支付卡类型,EBIT(借记卡),CREDIT(贷记卡),UNKNOWN(未知),CFT(钱包零钱)
	CashFee                 float64                      `json:"cashFee,omitempty"`                 // 上游返回:现金支付金额(订单总金额-现金券金额=现金支付金额)
	CouponFee               float64                      `json:"couponFee,omitempty"`               // 上游返回:现金券金额:现金券金额
	OrderCreditAmount       float64                      `json:"orderCreditAmount,omitempty"`       // 订单入账面额(不扣手续费)
	MerchantCreditAmount    float64                      `json:"merchantCreditAmount,omitempty"`    // 商户实际入账发生额,该笔交易成功后收单商户实际入账发生额,成功时返回
	PaymentAmount           float64                      `json:"paymentAmount,omitempty"`           // 用户实际支付金额
	ChannelSettlementAmount float64                      `json:"channelSettlementAmount,omitempty"` // 渠道结算金额,成功时返回
	OrderFee                float64                      `json:"orderFee,omitempty"`                // 该笔交易产生的手续费,成功时返回
	MerchantFee             float64                      `json:"merchantFee,omitempty"`             // 订单商户交易的费率,成功时返回
	FeeAccountAmt           float64                      `json:"feeAccountAmt,omitempty"`           // 平台商补贴的手续费,单位:元,支付成功时返回
	ReceiverFee             float64                      `json:"receiverFee,omitempty"`             // 实收手续费,单位:元,支付成功时返回
	OfflineFee              float64                      `json:"offlineFee,omitempty"`              // 后收手续费,单位:元,支付成功时返回
	OrderPayDate            string                       `json:"orderPayDate,omitempty"`            // 订单完成时间,格式:yyyy-MM-dd HH:mm:ss
	Timestamp               string                       `json:"timestamp"`                         // 通知时间,精确到通知时间的毫秒数,yyyyMMddHHmmss
	ChargeFlag              string                       `json:"chargeFlag,omitempty"`              // 渠道支付宝费率活动标识,支付宝活动通道成功时有返回
	WxTradeType             string                       `json:"wxTradeType,omitempty"`             // 微信交易类型,微信成功时有返回
	UpAddData               string                       `json:"upAddData,omitempty"`               // 银联二维码返回的付款方附加数据,Base64编码
	ResvData                string                       `json:"resvData,omitempty"`                // 银联二维码返回的保留数据,Base64编码
	FundBillList            string                       `json:"fundBillList,omitempty"`            // 交易资金渠道和优惠信息-支付宝
	PromotionDetail         string                       `json:"promotionDetail,omitempty"`         // 上游返回的优惠详情,JSON字符串
	VoucherDetailList       string                       `json:"voucherDetailList,omitempty"`       // 支付宝优惠信息详情,JSON字符串
	MarketingRulesJson      string                       `json:"marketingRules,omitempty"`          // 营销规则,JSON格式字符串,des加密传输
	MarketingRule           *AppPayPreOrderMarketingRule `json:"-"`                                 // 营销规则
	SplitRulesJson          string                       `json:"splitRules,omitempty"`              // 响应分账结果规则以及对应状态
	SplitRules              []*AppPayPreOrderSplitRule   `json:"-"`                                 // 分账规则
}

type OrderPayResultNotifyReq = NotifyReq[OrderPayResultNotifyReqBody]
