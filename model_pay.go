package dinpay

import (
	"github.com/bytedance/sonic"
)

// AppPayPreOrderMarketingRule 交易下单营销规则
type AppPayPreOrderMarketingRule struct {
	MarketingMerchantNo string  `json:"marketingMerchantNo"` // 营销商编
	MarketingAmount     float64 `json:"marketingAmount"`     // 营销金额
	CouponMerchantNo    string  `json:"couponMerchantNo"`    // 卡券商编
}

// AppPayRefundMarketingRule 交易退款营销规则
type AppPayRefundMarketingRule struct {
	RefundMarketingMerchantNo string  `json:"refundMarketingMerchantNo"` // 退款营销商编
	RefundMarketingAmount     float64 `json:"refundMarketingAmount"`     // 营销金额
	RefundCouponMerchantNo    string  `json:"refundCouponMerchantNo"`    // 卡券商编
}

// AppPayPreOrderSplitRule 交易下单分账规则
type AppPayPreOrderSplitRule struct {
	SplitBillMerchantNo string  `json:"splitBillMerchantNo"`
	SplitBillAmount     float64 `json:"splitBillAmount"`
}

// AppPayRefundSplitRule 订单退款分账规则
type AppPayRefundSplitRule struct {
	MerchantNo   string  `json:"merchantNo"`
	RefundAmount float64 `json:"refundAmount"`
}

// AppPayScanOrderReq 扫码/刷卡下单接口
type AppPayScanOrderReq struct {
	BizType        string  `json:"P1_bizType" sign:"1"`                   // 接口类型,固定为:AppPay
	OrderId        string  `json:"P2_orderId" sign:"1"`                   // 商户系统内部订单号，要求50字符以内，同一商户号下订单号唯一
	CustomerNumber string  `json:"P3_customerNumber" sign:"1"`            // 合利宝分配的商户号
	PayType        string  `json:"P4_payType" sign:"1"`                   // 支付类型,SWIPE:刷卡(被扫);SCAN:扫码(主扫)
	OrderAmount    float64 `json:"P5_orderAmount" sign:"1"`               // 订单金额,以元为单位，最小金额为0.01
	Currency       string  `json:"P6_currency" sign:"1"`                  // 币种类型,CNY:人民币
	AuthCode       string  `json:"P7_authcode" sign:"1"`                  // 授权码,主扫支付类型传入1,被扫传入
	AppType        string  `json:"P8_appType" sign:"1"`                   // 客户端类型,AppPayTypes
	NotifyUrl      string  `json:"P9_notifyUrl,omitempty" sign:"1"`       // 通知回调地址
	SuccessToUrl   string  `json:"P10_successToUrl,omitempty" sign:"1"`   // 成功跳转URL
	OrderIp        string  `json:"P11_orderIp" sign:"1"`                  // 商户IP
	GoodsName      string  `json:"P12_goodsName" sign:"1"`                // 商品名称,账单上显示
	GoodsDetail    string  `json:"P13_goodsDetail,omitempty" sign:"1"`    // 商品的优惠活动:单品优惠活动该字段必传,且必须按照规范上传,JSON格式
	Desc           string  `json:"P14_desc,omitempty" sign:"1"`           // 订单备注信息，原样返回
	AppId          string  `json:"P16_appId,omitempty" sign:"0"`          // 公众号appId
	LimitCreditPay string  `json:"P17_limitCreditPay,omitempty" sign:"0"` // 是否限制借贷记,1:禁用贷记卡,0:不限,2:针对支付宝禁用借记;不传以平台为准
	GoodsTag       string  `json:"P18_goodsTag,omitempty" sign:"0"`       // 商品标记,微信平台配置的商品标记，用于优惠券或者满减使用
	Guid           string  `json:"P19_guid,omitempty" sign:"0"`           // 微信进件时上送的唯一号,此值非必填,填写时必须是商户侧做微信进件时上送的唯一号.传错或不匹配的交易失败
	Identity       string  `json:"P21_identity,omitempty" sign:"0"`       // 实名参数,实名支付功能,用于公安和保险类商户使用
	// Deprecated: 请勿直接赋值,应调从MarketingRules添加
	MarketingRuleJson string `json:"P20_marketingRule,omitempty" sign:"0"` // 营销规则,JSON格式字符串,des加密传输
	// Deprecated: 请勿直接赋值,应调从SplitRules添加
	SplitRuleJson      string                       `json:"ruleJson,omitempty" sign:"0"`           // 分账规则,Json格式字符串;分账类型和分账规则串出现时须2个字段都要上送
	SplitBillType      string                       `json:"splitBillType,omitempty" sign:"0"`      // 分账类型,FIXED_AMOUNT:固定金额(默认,目前只支持固定金额);RATE:比率
	HbfqNum            string                       `json:"hbfqNum,omitempty" sign:"0"`            // 花呗分期数,目前仅支持3,6,12期
	DeviceInfo         string                       `json:"deviceInfo,omitempty" sign:"0"`         // 终端号
	StoreId            string                       `json:"storeId,omitempty" sign:"0"`            // 商户自定义的门店编码
	TimeExpire         string                       `json:"timeExpire,omitempty" sign:"0"`         // 超时时间,微信/银联传参规则:超时时间单位为秒,微信超时时间要大于等于60秒;支付宝传参规则:取值范围:1m～15d。m-分钟,h-小时,d-天,1c-当天(1c-当天的情况下,无论交易何时创建,都在0点关闭)。该参数数值不接受小数点,1.5h,可转换为90m。
	IndustryRefluxInfo string                       `json:"industryRefluxInfo,omitempty" sign:"0"` // 支付宝行业数据回流信息,json字符串格式;详见 https://myshangpu.yuque.com/org-wiki-myshangpu-sfbw9n/qnpgdn/phtfgreqpllk9i2m#oC6y0
	TermInfo           string                       `json:"termInfo,omitempty" sign:"0"`           // 银联终端信息,参加银联官方单品营销必送,json字符串格式,详见 https://myshangpu.yuque.com/org-wiki-myshangpu-sfbw9n/qnpgdn/phtfgreqpllk9i2m#DuXus
	OpenId             string                       `json:"openId,omitempty"`                      // 用户id
	AuthConfirmMode    string                       `json:"authConfirmMode,omitempty" sign:"0"`    // 预授权确认模式,预授权场景必填COMPLETE:转交易支付完成结束预授权;NOT_COMPLETE:转交易支付完成不结束预授权
	TerminalSysBindNo  string                       `json:"terminalSysBindNo,omitempty" sign:"0"`  // 终端绑定号,在平台已报备过的终端信息的绑定号(采集接口返回);线下场景时必填
	EncryptRandNum     string                       `json:"encryptRandNum,omitempty" sign:"0"`     // 加密随机因子,仅在被扫支付类交易报文中出现:若付款码为19位数字,则取后6位;若付款码码为EMV二维码,则取其tag57的卡号/token号的后6位
	SecretText         string                       `json:"secretText,omitempty" sign:"0"`         // 密文数据
	SceneInfo          string                       `json:"sceneInfo,omitempty" sign:"0"`          // 场景信息,该字段用于上报场景信息,目前支持上报实际门店信息; 详见 https://myshangpu.yuque.com/org-wiki-myshangpu-sfbw9n/qnpgdn/phtfgreqpllk9i2m#fxUGg
	EduSubject         string                       `json:"eduSubject,omitempty" sign:"0"`         // 学校名称、场景名称
	BusinessParams     string                       `json:"businessParams,omitempty" sign:"0"`     // 学校ID和场景ID
	MarketingRule      *AppPayPreOrderMarketingRule `json:"-" sign:"0"`                            // 营销规则
	SplitRules         []*AppPayPreOrderSplitRule   `json:"-" sign:"0"`                            // 分账规则
}

// AppPayScanOrderRes 扫码/刷卡下单接口
type AppPayScanOrderRes struct {
	BizType                 string                       `json:"rt1_bizType" sign:"1"`                          // 接口类型,固定为:AppPayPublic
	RetCode                 string                       `json:"rt2_retCode" sign:"1"`                          // 返回码
	RetMsg                  string                       `json:"rt3_retMsg,omitempty" sign:"0"`                 // 返回信息
	CustomerNumber          string                       `json:"rt4_customerNumber" sign:"1"`                   // 商户编号
	OrderId                 string                       `json:"rt5_orderId" sign:"1"`                          // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一
	SerialNumber            string                       `json:"rt6_serialNumber,omitempty" sign:"1"`           // 合利宝交易订单号
	PayType                 string                       `json:"rt7_payType" sign:"1"`                          // 支付类型,SWIPE:刷卡;SCAN:扫码
	Qrcode                  string                       `json:"rt8_qrcode,omitempty" sign:"1"`                 // 二维码信息,用于生成二维码图片，展示给用户进行扫码支付
	WapUrl                  string                       `json:"rt9_wapurl,omitempty" sign:"1"`                 // 支付地址,商户系统显示二维码页面地址(暂不用)
	OrderAmount             float64                      `json:"rt10_orderAmount,string" sign:"1"`              // 交易金额,以元为单位,最小金额为0.01
	Currency                string                       `json:"rt11_currency" sign:"1"`                        // 币种类型,CNY:人民币
	OpenId                  string                       `json:"rt12_openId,omitempty" sign:"0"`                // 用户openId
	OrderStatus             string                       `json:"rt13_orderStatus,omitempty" sign:"0"`           // 订单状态,刷卡支付时响应状态
	FundBillList            string                       `json:"rt14_fundBillList,omitempty" sign:"0"`          // 交易支付使用的资金渠道和优惠信息-支付宝
	ChannelRetCode          string                       `json:"rt15_channelRetCode,omitempty" sign:"0"`        // 上游返回码
	OutTransactionOrderId   string                       `json:"rt16_outTransactionOrderId,omitempty" sign:"0"` // 微信支付宝交易订单号
	BankType                string                       `json:"rt17_bankType,omitempty" sign:"0"`              // 用户付款银行
	SubOpenId               string                       `json:"rt18_subOpenId,omitempty" sign:"0"`             // subOpenId;微信子商户subOpenId.或支付宝子商户用户buyer_id
	OrderAttribute          string                       `json:"rt19_orderAttribute,omitempty" sign:"0"`        // 通道订单属性;UNDIRECT_DEFAULT:间连通道;DIRECT_CHANNEL 直连通道
	MarketingRuleJson       string                       `json:"rt20_marketingRule,omitempty" sign:"0"`         // 营销规则,JSON格式字符串，des加密传输
	PromotionDetail         string                       `json:"rt21_promotionDetail,omitempty" sign:"0"`       // 上游返回的优惠详情
	CreditAmount            string                       `json:"rt22_creditAmount,omitempty" sign:"0"`          // 入账面额(不扣手续费)
	PaymentAmount           string                       `json:"rt23_paymentAmount,omitempty" sign:"0"`         // 用户实际支付金额
	OrderCompleteDate       string                       `json:"rt24_orderCompleteDate,omitempty" sign:"0"`     // 订单完成时间,格式:yyyy-MM-dd HH:mm:ss
	AppPayType              string                       `json:"rt25_appPayType,omitempty" sign:"0"`            // 客户端类型,AppPayTypes
	AppId                   string                       `json:"rt26_appId,omitempty" sign:"0"`                 // 子商户公众号sub_appid
	SplitRuleJson           string                       `json:"ruleJson,omitempty" sign:"0"`                   // 分账规则及状态,响应分账结果规则以及对应状态
	ProductFee              string                       `json:"productFee,omitempty" sign:"0"`                 // 该笔交易产生的手续费,成功时返回
	ChannelSettlementAmount string                       `json:"channelSettlementAmount,omitempty" sign:"0"`    // 渠道结算金额,成功时返回
	RealCreditAmount        string                       `json:"realCreditAmount,omitempty" sign:"0"`           // 该笔交易成功后收单商户实际入账发生额,成功时返回
	TradeType               string                       `json:"tradeType,omitempty" sign:"0"`                  // 微信交易类型,刷卡成功时有返回
	ChargeFlag              string                       `json:"chargeFlag,omitempty" sign:"0"`                 // 渠道支付宝费率活动标识
	UpAddData               string                       `json:"upAddData,omitempty" sign:"0"`                  // 银联二维码返回的付款方附加数据,Base64编码
	ResvData                string                       `json:"resvData,omitempty" sign:"0"`                   // 银联二维码返回的保留数据,Base64编码的
	OnlineCardType          string                       `json:"onlineCardType,omitempty" sign:"0"`             // 支付卡类型,DEBIT(借记卡);CREDIT(贷记卡);UNKNOWN(未知);CFT(钱包零钱)
	SubMerchantNo           string                       `json:"subMerchantNo,omitempty" sign:"0"`              // 渠道子商户号(U/A/T)
	FeeRate                 string                       `json:"feeRate,omitempty" sign:"0"`                    // 商户交易的费率,被扫支付成功时返回
	FeeAccountAmt           string                       `json:"feeAccountAmt,omitempty" sign:"0"`              // 平台商补贴的手续费,单位:元,被扫支付成功时返回
	Sign                    string                       `json:"sign" sign:"0"`                                 // MD5 签名结果
	MarketingRule           *AppPayPreOrderMarketingRule `json:"-" sign:"0"`                                    // 营销规则
	SplitRules              []*AppPayPreOrderSplitRule   `json:"-" sign:"0"`                                    // 分账规则
}

// AppPayPublicPreOrderReq 公众号/JS/服务窗预下单接口
type AppPayPublicPreOrderReq struct {
	BizType        string  `json:"P1_bizType" sign:"1"`                   // 接口类型,固定为:AppPayPublic
	OrderId        string  `json:"P2_orderId" sign:"1"`                   // 商户订单号,要求50字符以内,同一商户号下订单号唯一
	CustomerNumber string  `json:"P3_customerNumber" sign:"1"`            // 合利宝分配的商户号
	PayType        string  `json:"P4_payType" sign:"1"`                   // 支付类型,公众号/JS/服务窗固定为:PUBLIC
	AppId          string  `json:"P5_appid" sign:"1"`                     // 微信支付填公众账号ID;支付宝填1;银联云闪付填1;合利付填应用编号appNo
	DeviceInfo     string  `json:"P6_deviceInfo,omitempty" sign:"1"`      // 可以为终端设备号(门店号或收银设备ID),PC网页或公众号内支付可以传"WEB"
	IsRaw          string  `json:"P7_isRaw,omitempty" sign:"1"`           // 是否原生态,1:是;0:否;不传默认是1
	OpenId         string  `json:"P8_openid" sign:"1"`                    // 微信用户关注商家公众号的openid,当P7_isRaw=0微信公众号非原生态时openid上送1即可
	OrderAmount    float64 `json:"P9_orderAmount" sign:"1"`               // 订单金额,以元为单位,最小金额为0.01
	Currency       string  `json:"P10_currency" sign:"1"`                 // 币种类型,人民币:CNY
	AppType        string  `json:"P11_appType" sign:"1"`                  // 客户端类型,AppPayType
	NotifyUrl      string  `json:"P12_notifyUrl,omitempty" sign:"1"`      // 通知回调地址,异步接收合利宝支付结果通知的回调地址,通知url必须为外网可访问
	SuccessToUrl   string  `json:"P13_successToUrl,omitempty" sign:"1"`   // 支付完成后,展示支付结果的页面地址;为原生态时该地址请忽略,需要商户自己在js页面控制跳转;当P7_isRaw=0微信公众号非原生态时支付完会跳转到该地址
	OrderIp        string  `json:"P14_orderIp" sign:"1"`                  // 下单机器IP
	GoodsName      string  `json:"P15_goodsName" sign:"1"`                // 商品名称
	GoodsDetail    string  `json:"P16_goodsDetail,omitempty" sign:"1"`    // 商品详情,商品的优惠活动:单品优惠活动该字段必传,且必须按照规范上传,JSON字符串格式
	LimitCreditPay string  `json:"P17_limitCreditPay,omitempty" sign:"1"` // 值为1,禁用信用卡;值为0不限,不传以平台为准
	Desc           string  `json:"P18_desc,omitempty" sign:"1"`           // 备注原样返回
	SubscribeAppId string  `json:"P19_subscribeAppId,omitempty" sign:"0"` // 关注appId,微信参数
	GoodsTag       string  `json:"P21_goodsTag,omitempty" sign:"0"`       // 商品标记,用于优惠券或者满减使用
	Guid           string  `json:"P22_guid,omitempty" sign:"0"`           // 微信进件时上送的唯一号,此值非必填,传递时必须是商户侧做微信进件时上送的唯一号,传错或不匹配会交易失败
	Identity       string  `json:"P24_identity,omitempty" sign:"0"`       // 实名参数,实名支付功能,用于公安和保险类商户使用
	// Deprecated: 请勿直接赋值,应调从MarketingRules添加
	MarketingRuleJson string `json:"P23_marketingRule,omitempty" sign:"0"` // 营销规则,JSON格式字符串,des加密传输
	// Deprecated: 请勿直接赋值,应调从SplitRules添加
	SplitRuleJson      string                       `json:"ruleJson,omitempty" sign:"0"`           // 分账规则,Json格式字符串;分账类型和分账规则串出现时须2个字段都要上送
	SplitBillType      string                       `json:"splitBillType,omitempty" sign:"0"`      // 分账类型,FIXED_AMOUNT:固定金额(默认,目前只支持固定金额);RATE:比率
	TimeExpire         string                       `json:"timeExpire,omitempty" sign:"0"`         // 超时时间,微信/银联传参规则:超时时间单位为秒,微信超时时间要大于等于60秒;支付宝传参规则:取值范围:1m～15d。m-分钟,h-小时,d-天,1c-当天(1c-当天的情况下,无论交易何时创建,都在0点关闭)。该参数数值不接受小数点,1.5h,可转换为90m。
	HbfqNum            string                       `json:"hbfqNum,omitempty" sign:"0"`            // 花呗分期数,目前仅支持3,6,12期
	IndustryRefluxInfo string                       `json:"industryRefluxInfo,omitempty" sign:"0"` // 支付宝行业数据回流信息,json字符串格式;详见 https://myshangpu.yuque.com/org-wiki-myshangpu-sfbw9n/qnpgdn/phtfgreqpllk9i2m#oC6y0
	FoodOrderType      string                       `json:"foodOrderType,omitempty" sign:"0"`      // (店内扫码点餐)支付宝点餐码,详见 https://myshangpu.yuque.com/org-wiki-myshangpu-sfbw9n/qnpgdn/phtfgreqpllk9i2m#JF10T
	TermInfo           string                       `json:"termInfo,omitempty" sign:"0"`           // 银联终端信息,参加银联官方单品营销必送,json字符串格式,详见 https://myshangpu.yuque.com/org-wiki-myshangpu-sfbw9n/qnpgdn/phtfgreqpllk9i2m#DuXus
	UserAgent          string                       `json:"userAgent,omitempty" sign:"0"`          // 银联用户代理,银联JS识别用户的客户端,不填默认是UnionPay/1.0 CloudPay;详见 https://myshangpu.yuque.com/org-wiki-myshangpu-sfbw9n/qnpgdn/phtfgreqpllk9i2m#ENYf3
	OpenIdType         string                       `json:"openIdType,omitempty" sign:"0"`         // userAuthCode(临时授权码)/openId(用户id),此参表示上送的P8_openid值的类型;不填默认是userAuthCode
	StoreId            string                       `json:"storeId,omitempty" sign:"0"`            // 商户门店编号,商户自定义的门店编码
	TerminalSysBindNo  string                       `json:"terminalSysBindNo,omitempty" sign:"0"`  // 终端绑定号,在平台已报备过的终端信息的绑定号
	SceneInfo          string                       `json:"sceneInfo,omitempty" sign:"0"`          // 场景信息,该字段用于上报场景信息,目前支持上报实际门店信息; 详见 https://myshangpu.yuque.com/org-wiki-myshangpu-sfbw9n/qnpgdn/phtfgreqpllk9i2m#fxUGg
	MarketingRule      *AppPayPreOrderMarketingRule `json:"-" sign:"0"`                            // 营销规则
	SplitRules         []*AppPayPreOrderSplitRule   `json:"-" sign:"0"`                            // 分账规则
}

// AppPayPublicPreOrderRes 公众号/JS/服务窗预下单接口
type AppPayPublicPreOrderRes struct {
	BizType              string  `json:"rt1_bizType" sign:"1"`                   // 接口类型,固定为:AppPayPublic
	RetCode              string  `json:"rt2_retCode" sign:"1"`                   // 返回码
	RetMsg               string  `json:"rt3_retMsg,omitempty" sign:"0"`          // 返回信息
	CustomerNumber       string  `json:"rt4_customerNumber" sign:"1"`            // 商户编号
	OrderId              string  `json:"rt5_orderId" sign:"1"`                   // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一
	SerialNumber         string  `json:"rt6_serialNumber,omitempty" sign:"1"`    // 合利宝交易订单号
	PayType              string  `json:"rt7_payType" sign:"1"`                   // 支付类型,PUBLIC:公众号/服务窗/JS支付
	AppId                string  `json:"rt8_appid,omitempty" sign:"1"`           // 公众账号ID
	TokenId              string  `json:"rt9_tokenId,omitempty" sign:"1"`         // 动态口令,预支付ID,用于后续接口调用中使用,预留参数
	PayInfo              string  `json:"rt10_payInfo" sign:"1"`                  // 原生态JS支付信息,is_raw为1时返回,json格式的字符串,作用于原生态JS支付时的参数;为小程序时返回的json串集成小程序JS接口时需要
	OrderAmount          float64 `json:"rt11_orderAmount,string" sign:"1"`       // 订单金额,以元为单位,最小金额为0.01
	Currency             string  `json:"rt12_currency" sign:"1"`                 // 币种类型,CNY:人民币
	ChannelRetCode       string  `json:"rt13_channelRetCode,omitempty" sign:"0"` // 上游返回码,失败时透传上游返回码,仅供参考,请以订单状态为准
	AppPayType           string  `json:"rt14_appPayType,omitempty" sign:"0"`     // 客户端类型,AppPayType
	ChannelSubMerchantNo string  `json:"subMerchantNo,omitempty" sign:"0"`       // 渠道子商户号(U/A/T)
	Sign                 string  `json:"sign" sign:"0"`                          // 签名
}

func (res *AppPayPublicPreOrderRes) GetPayInfo() (payInfo *AppPayPublicPreOrderPayInfo) {
	if len(res.PayInfo) > 0 {
		payInfo = new(AppPayPublicPreOrderPayInfo)
		_ = sonic.Unmarshal([]byte(res.PayInfo), payInfo)
	}
	return
}

// AppPayPublicPreOrderPayInfo 公众号/JS/服务窗预下单支付信息
type AppPayPublicPreOrderPayInfo struct {
	AppId     string `json:"appId"`     // AppId
	TimeStamp string `json:"timeStamp"` // 时间戳
	NonceStr  string `json:"nonceStr"`  // 随机字符串
	Package   string `json:"package"`   // 订单详情扩展字符串
	SignType  string `json:"signType"`  // 签名方式
	PaySign   string `json:"paySign"`   // 签名
	TradeNO   string `json:"tradeNO"`   // 支付宝交易单号
}

// AppPayAppletPreOrderReq 小程序预下单接口
type AppPayAppletPreOrderReq struct {
	BizType        string  `json:"P1_bizType" sign:"1"`                   // 接口类型,固定为:AppPayApplet
	OrderId        string  `json:"P2_orderId" sign:"1"`                   // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一
	CustomerNumber string  `json:"P3_customerNumber" sign:"1"`            // 合利宝分配的商户号
	PayType        string  `json:"P4_payType" sign:"1"`                   // 支付类型,小程序固定为:"APPLET"
	AppId          string  `json:"P5_appid" sign:"1"`                     // 支付宝买家账号ID,能获取买家账号账号ID就送,反之填1
	DeviceInfo     string  `json:"P6_deviceInfo,omitempty" sign:"1"`      // 终端设备号(门店号或收银设备ID),PC网页或服务窗内支付可以传"WEB"
	IsRaw          string  `json:"P7_isRaw,omitempty" sign:"1"`           // 是否原生态,1:是;0:否。不传默认是1
	OpenId         string  `json:"P8_openid" sign:"1"`                    // 用户id获取方法,详见6.2注意事项
	OrderAmount    float64 `json:"P9_orderAmount" sign:"1"`               // 订单金额,以元为单位,最小金额为0.01
	Currency       string  `json:"P10_currency" sign:"1"`                 // 人民币:CNY
	AppType        string  `json:"P11_appType" sign:"1"`                  // 客户端类型,AppPayType
	NotifyUrl      string  `json:"P12_notifyUrl,omitempty" sign:"1"`      // 异步接收合利宝支付结果通知的回调地址,通知url必须为外网可访问
	SuccessToUrl   string  `json:"P13_successToUrl,omitempty" sign:"1"`   // 支付完成后,展示支付结果的页面地址
	OrderIp        string  `json:"P14_orderIp" sign:"1"`                  // 下单机器IP
	GoodsName      string  `json:"P15_goodsName" sign:"1"`                // 商品名称,支付时账单显示
	GoodsDetail    string  `json:"P16_goodsDetail,omitempty" sign:"1"`    // 商品的优惠活动：单品优惠活动该字段必传,且必须按照规范上传,JSON字符串格式
	LimitCreditPay string  `json:"P17_limitCreditPay,omitempty" sign:"1"` // 1:禁用信用卡;0:不限,2:针对支付宝禁用借记;不传以平台为准
	Desc           string  `json:"P18_desc,omitempty" sign:"1"`           // 商户备注,原样返回
	GoodsTag       string  `json:"P21_goodsTag,omitempty" sign:"0"`       // 商品标记,支付宝平台配置的商品标记,用于优惠券或者满减使用
	// Deprecated: 请勿直接赋值,应调用MarketingRules添加
	MarketingRuleJson string `json:"P23_marketingRule,omitempty" sign:"0"` // 营销规则,JSON格式字符串,des加密传输
	// Deprecated: 请勿直接赋值,应调用SplitRules添加
	SplitRuleJson      string                       `json:"ruleJson,omitempty" sign:"0"`           // 分账规则,Json格式字符串,分账类型和分账规则串出现时须2个字段都要上送
	SplitBillType      string                       `json:"splitBillType,omitempty" sign:"0"`      // 分账类型,FIXED_AMOUNT:固定金额(默认,目前只支持固定金额);RATE:比率
	HbfqNum            string                       `json:"hbfqNum,omitempty" sign:"0"`            // 花呗分期数,目前仅支持3,6,12期
	TimeExpire         string                       `json:"timeExpire,omitempty" sign:"0"`         // 订单超时时间参数规则 https://myshangpu.yuque.com/org-wiki-myshangpu-sfbw9n/qnpgdn/phtfgreqpllk9i2m#P8kcu
	IndustryRefluxInfo string                       `json:"industryRefluxInfo,omitempty" sign:"0"` // 支付宝行业数据回流信息,JSON格式字符串
	FoodOrderType      string                       `json:"foodOrderType,omitempty" sign:"0"`      // (店内扫码点餐)支付宝点餐码详情
	StoreId            string                       `json:"storeId,omitempty" sign:"0"`            // 商户自定义的门店编码
	UserAgent          string                       `json:"userAgent,omitempty" sign:"0"`          // 银联JS识别用户的客户端,不填默认是UnionPay/1.0 CloudPay
	OpenIdType         string                       `json:"openIdType,omitempty" sign:"0"`         // userAuthCode(临时授权码)/openId(用户id),此参表示上送的P8_openid值的类型;不填默认是userAuthCode
	TerminalSysBind    string                       `json:"terminalSysBind,omitempty" sign:"0"`    // 在平台已报备过的终端信息的绑定号(采集接口返回)
	SceneInfo          string                       `json:"sceneInfo,omitempty" sign:"0"`          // 该字段用于上报实际门店信息,目前支持上报实际门店信息
	MarketingRule      *AppPayPreOrderMarketingRule `json:"-" sign:"0"`                            // 营销规则
	SplitRules         []*AppPayPreOrderSplitRule   `json:"-" sign:"0"`                            // 分账规则
}

// AppPayAppletPreOrderRes 小程序预下单接口
type AppPayAppletPreOrderRes struct {
	BizType              string  `json:"rt1_bizType" sign:"1"`                   // 接口类型,固定为:AppPayApplet
	RetCode              string  `json:"rt2_retCode" sign:"1"`                   // 参考附录：https://myshangpu.yuque.com/org-wiki-myshangpu-sfbw9n/qnpgdn/phtfgreqpllk9i2m#EkA5A
	RetMsg               string  `json:"rt3_retMsg,omitempty" sign:"0"`          // 返回信息,不用参与签名
	CustomerNumber       string  `json:"rt4_customerNumber" sign:"1"`            // 合利宝分配的商户号
	OrderId              string  `json:"rt5_orderId" sign:"1"`                   // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一
	SerialNumber         string  `json:"rt6_serialNumber,omitempty" sign:"1"`    // 合利宝交易订单号
	PayType              string  `json:"rt7_payType" sign:"1"`                   // 支付类型APPLET:小程序支付
	AppId                string  `json:"rt8_appid,omitempty" sign:"1"`           // app账号 Id
	TokenId              string  `json:"rt9_tokenId,omitempty" sign:"1"`         // 动态口令,预支付ID,用于后续接口调用中使用,预留参数
	PayInfo              string  `json:"rt10_payInfo" sign:"1"`                  // 原生态js 支付信息,返回,json 格式的字符串,唤起钱包支付控件{"tradeNO":"2016071121001004610284287324","status":"0"},
	OrderAmount          float64 `json:"rt11_orderAmount,string" sign:"1"`       // 交易金额,以元为单位,最小金额为0.01
	Currency             string  `json:"rt12_currency" sign:"1"`                 // 币种类型 CNY:人民币
	ChannelRetCode       string  `json:"rt13_channelRetCode,omitempty" sign:"0"` // 上游返回码,失败时透传上游返回码,仅供参考,请以订单状态为准
	AppPayType           string  `json:"rt14_appPayType,omitempty" sign:"0"`     // 客户端类型,AppPayType
	ChannelSubMerchantNo string  `json:"subMerchantNo,omitempty" sign:"0"`       // 渠道子商户号(U/A/T),不参与签名
	Sign                 string  `json:"sign" sign:"0"`                          // MD5 签名结果
}

func (res *AppPayAppletPreOrderRes) GetPayInfo() (payInfo *AppPayAppletPreOrderPayInfo) {
	if len(res.PayInfo) > 0 {
		payInfo = new(AppPayAppletPreOrderPayInfo)
		_ = sonic.Unmarshal([]byte(res.PayInfo), payInfo)
	}
	return
}

// AppPayAppletPreOrderPayInfo 小程序预下单支付信息
type AppPayAppletPreOrderPayInfo struct {
	TimeStamp string `json:"timeStamp"` // 时间戳
	NonceStr  string `json:"nonceStr"`  // 随机字符串
	Package   string `json:"package"`   // 订单详情扩展字符串
	SignType  string `json:"signType"`  // 签名方式
	PaySign   string `json:"paySign"`   // 签名
	TradeNO   string `json:"tradeNO"`   // 支付宝交易单号
}

// AppPayOrderCloseReq 交易订单关闭接口
type AppPayOrderCloseReq struct {
	BizType        string `json:"P1_bizType" sign:"1"`           // 接口类型,固定为:AppPayClose
	OrderId        string `json:"P2_orderId,omitempty" sign:"1"` // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一
	CustomerNumber string `json:"P3_customerNumber" sign:"1"`    // 合利宝分配的商户号
}

// AppPayOrderCloseRes 交易订单关闭接口
type AppPayOrderCloseRes struct {
	BizType        string `json:"rt1_bizType" sign:"1"`                // 小程序支付
	RetCode        string `json:"rt2_retCode" sign:"1"`                // 参考附录：https://myshangpu.yuque.com/org-wiki-myshangpu-sfbw9n/qnpgdn/phtfgreqpllk9i2m#EkA5A
	RetMsg         string `json:"rt3_retMsg,omitempty" sign:"0"`       // 返回信息
	CustomerNumber string `json:"rt4_customerNumber" sign:"1"`         // 合利宝分配的商户号
	OrderId        string `json:"rt5_orderId" sign:"1"`                // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一
	SerialNumber   string `json:"rt6_serialNumber,omitempty" sign:"1"` // 合利宝平台唯一流水号
	Sign           string `json:"sign" sign:"0"`                       // MD5 签名结果
}

// AppPayOrderCancelReq 交易订单撤销接口
type AppPayOrderCancelReq struct {
	BizType        string `json:"P1_bizType" sign:"1"`           // 接口类型,固定为:AppPayCancel
	OrderId        string `json:"P2_orderId,omitempty" sign:"1"` // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一
	CustomerNumber string `json:"P3_customerNumber" sign:"1"`    // 合利宝分配的商户号
}

// AppPayOrderCancelRes 交易订单撤销接口
type AppPayOrderCancelRes struct {
	BizType        string `json:"rt1_bizType" sign:"1"`                // 接口类型,固定为:AppPayCancel
	RetCode        string `json:"rt2_retCode" sign:"1"`                // 参考附录：https://myshangpu.yuque.com/org-wiki-myshangpu-sfbw9n/qnpgdn/phtfgreqpllk9i2m#EkA5A
	RetMsg         string `json:"rt3_retMsg,omitempty" sign:"0"`       // 返回信息
	CustomerNumber string `json:"rt4_customerNumber" sign:"1"`         // 合利宝分配的商户号
	OrderId        string `json:"rt5_orderId" sign:"1"`                // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一
	SerialNumber   string `json:"rt6_serialNumber,omitempty" sign:"1"` // 合利宝平台唯一流水号
	Recall         string `json:"recall,omitempty" sign:"0"`           // 重试标识:Y/N,是否需要继续调用撤销,Y- 需要,N-不需要;上游通道撤销未返回明确结果时可以重试调用
	Sign           string `json:"sign" sign:"0"`                       // MD5 签名结果
}

// AppPayOrderQueryReq 交易订单查询接口
type AppPayOrderQueryReq struct {
	BizType        string `json:"P1_bizType" sign:"1"`                // 订单查询接口,固定为:AppPayQuery
	OrderId        string `json:"P2_orderId,omitempty" sign:"1"`      // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一与” 合利宝订单流水号”二选一
	CustomerNumber string `json:"P3_customerNumber" sign:"1"`         // 合利宝分配的商户号
	SerialNumber   string `json:"P4_serialNumber,omitempty" sign:"0"` // 合利宝平台返回的订单流水号与”商户订单号”二选一
}

// AppPayOrderQueryRes 交易订单查询接口
type AppPayOrderQueryRes struct {
	BizType                 string                       `json:"rt1_bizType" sign:"1"`                          // 交易类型,固定为:AppPayQuery
	RetCode                 string                       `json:"rt2_retCode" sign:"1"`                          // 返回码
	RetMsg                  string                       `json:"rt3_retMsg,omitempty" sign:"0"`                 // 返回信息
	CustomerNumber          string                       `json:"rt4_customerNumber" sign:"1"`                   // 合利宝分配的商户号
	OrderId                 string                       `json:"rt5_orderId" sign:"1"`                          // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一
	SerialNumber            string                       `json:"rt6_serialNumber" sign:"1"`                     // 合利宝平台唯一流水号
	OrderStatus             string                       `json:"rt7_orderStatus" sign:"1"`                      // 订单状态,INIT:已接收,DOING:处理中,SUCCESS:成功,FAIL:失败,CLOSE:关闭,CANCEL:撤销
	OrderAmount             float64                      `json:"rt8_orderAmount,string" sign:"1"`               // 订单金额,以元为单位,最小金额为0.01
	Currency                string                       `json:"rt9_currency" sign:"1"`                         // 币种类型,CNY:人民币
	Desc                    string                       `json:"rt10_desc,omitempty" sign:"0"`                  // 订单描述字段
	OpenId                  string                       `json:"rt11_openId,omitempty" sign:"0"`                // 微信用户openId(暂不用)
	ChannelOrderNum         string                       `json:"rt12_channelOrderNum,omitempty" sign:"0"`       // 第三方平台订单号
	OrderCompleteDate       string                       `json:"rt13_orderCompleteDate,omitempty" sign:"0"`     // 订单完成时间,格式:yyyy-MM-dd HH:mm:ss
	CashFee                 string                       `json:"rt14_cashFee,omitempty" sign:"0"`               // 上游返回:现金支付金额(订单总金额-现金券金额=现金支付金额)
	CouponFee               string                       `json:"rt15_couponFee,omitempty" sign:"0"`             // 上游返回:现金券金额
	OnlineCardType          string                       `json:"rt16_onlineCardType,omitempty" sign:"0"`        // 支付卡类型,EBIT(借记卡),CREDIT(贷记卡),UNKNOWN(未知),CFT(钱包零钱)
	FundBillList            string                       `json:"rt17_fundBillList,omitempty" sign:"0"`          // 支付宝使用的资金渠道和优惠信息字段
	OutTransactionOrderId   string                       `json:"rt18_outTransactionOrderId,omitempty" sign:"0"` // 微信支付宝交易订单号,成功时有返回( 不参与签名)
	BankType                string                       `json:"rt19_bankType,omitempty" sign:"0"`              // 用户付款银行,用户付款银行,成功时有返回
	SubOpenId               string                       `json:"rt20_subOpenId,omitempty" sign:"0"`             // 微信子商户subOpenId字段
	OrderAttribute          string                       `json:"rt21_orderAttribute,omitempty" sign:"0"`        // 通道订单属性,标记通道订单属性,UNDIRECT_DEFAULT:间连通道,DIRECT_CHANNEL 直连通道
	MarketingRuleJson       string                       `json:"rt22_marketingRule,omitempty" sign:"0"`         // 营销规则,JSON格式字符串,des加密传输
	PromotionDetail         string                       `json:"rt23_promotionDetail,omitempty" sign:"0"`       // 上游返回的优惠详情
	CreditAmount            string                       `json:"rt24_creditAmount,omitempty" sign:"0"`          // 入账面额(不扣手续费)
	PaymentAmount           string                       `json:"rt25_paymentAmount,omitempty" sign:"0"`         // 用户实际支付金额
	AppPayType              string                       `json:"rt26_appPayType,omitempty" sign:"0"`            // 客户端,AppPayType
	PayType                 string                       `json:"rt27_payType,omitempty" sign:"0"`               // 支付类型,PayType
	AppId                   string                       `json:"rt28_appId,omitempty" sign:"0"`                 // 子商户公众号sub_appid字段
	RefundStatus            string                       `json:"rt29_refundStatus,omitempty" sign:"0"`          // 退款状态,PRE_REFUND:转入退款;PART_REFUND:部分退款;ALL_REFUND:全额退款;FAIL_REFUND:退款失败;NOT_YET:尚未退款,(转入退款,部分退款,全额退款,退款失败 均表示这 笔订单发起过退款)
	SubMerchantNo           string                       `json:"rt30_subMerchantNo,omitempty" sign:"0"`         // 二级商编
	SplitRuleJson           string                       `json:"ruleJson,omitempty" sign:"0"`                   // 分账规则及状态
	ProductFee              string                       `json:"productFee,omitempty" sign:"0"`                 // 该笔交易产生的手续费,成功时返回
	ChannelSettlementAmount string                       `json:"channelSettlementAmount,omitempty" sign:"0"`    // 渠道结算金额,成功时返回
	RealCreditAmount        string                       `json:"realCreditAmount,omitempty" sign:"0"`           // 该笔交易成功后收单商户实际入账发生额,成功时返回
	TradeType               string                       `json:"tradeType,omitempty" sign:"0"`                  // 微信交易类型,微信成功时有返回
	ChargeFlag              string                       `json:"chargeFlag,omitempty" sign:"0"`                 // 支付宝活动通道标识,支付宝活动通道成功时有返回
	UpAddData               string                       `json:"upAddData,omitempty" sign:"0"`                  // 银联二维码返回的付款方附加数据,Base64编码
	ResvData                string                       `json:"resvData,omitempty" sign:"0"`                   // 银联二维码返回的保留数据,Base64编码
	ChannelSubMerchantNo    string                       `json:"subMerchantNo,omitempty" sign:"0"`              // 渠道子商户号(U/A/T)
	FeeRate                 string                       `json:"feeRate,omitempty" sign:"0"`                    // 商户交易的费率,被扫支付成功时返回
	FeeAccountAmt           string                       `json:"feeAccountAmt,omitempty" sign:"0"`              // 平台商补贴的手续费,单位:元,被扫支付成功时返回
	Sign                    string                       `json:"sign" sign:"0"`                                 // 签名
	MarketingRule           *AppPayPreOrderMarketingRule `json:"-" sign:"0"`                                    // 营销规则
	SplitRules              []*AppPayPreOrderSplitRule   `json:"-" sign:"0"`                                    // 分账规则
}

// AppPayOrderRefundReq 交易订单退款接口
type AppPayOrderRefundReq struct {
	BizType           string  `json:"P1_bizType" sign:"1"`                     // 交易类型,固定为:AppPayRefund
	OrderId           string  `json:"P2_orderId,omitempty" sign:"1"`           // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一与”合利宝原订单号”二选一
	CustomerNumber    string  `json:"P3_customerNumber" sign:"1"`              // 合利宝分配的商户号
	RefundOrderId     string  `json:"P4_refundOrderId" sign:"1"`               // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一
	Amount            float64 `json:"P5_amount" sign:"1"`                      // 退款金额,以元为单位,最小金额为0.01
	CallbackUrl       string  `json:"P6_callbackUrl,omitempty" sign:"1"`       // 异步接收合利宝退款结果通知的回调地址,通知url必须为外网可访问的url,不能携带参数。
	Desc              string  `json:"P7_desc,omitempty" sign:"0"`              // 若商户传入,会在下发给用户的退款账单消息中体现退款原因
	OrderSerialNumber string  `json:"P8_orderSerialNumber,omitempty" sign:"0"` // 合利宝平台原订单与”商户订单号”二选一
	// Deprecated: 请勿直接赋值,应调用SplitRules添加
	SplitRuleJson string                   `json:"ruleJson,omitempty" sign:"0"`    // 分账退款规则JSON串,需使用DES进行对称加密
	AcqAddnData   string                   `json:"acqAddnData,omitempty" sign:"0"` // 银联二维码平台收款方附加数据,收款方附加数据退款请求,上送参加银联二维码通道营销商品维度信息;银联二维码平台参加单品营销的订单发生退款时必须上送
	SplitRules    []*AppPayRefundSplitRule `json:"-" sign:"0"`                     // 分账退款规则
}

// AppPayOrderRefundRes 交易订单退款接口
type AppPayOrderRefundRes struct {
	BizType        string `json:"rt1_bizType" sign:"1"`                // 交易类型,固定为:AppPayRefund
	RetCode        string `json:"rt2_retCode" sign:"1"`                // 返回码
	RetMsg         string `json:"rt3_retMsg,omitempty" sign:"0"`       // 返回信息
	CustomerNumber string `json:"rt4_customerNumber" sign:"1"`         // 合利宝分配的商户号
	OrderId        string `json:"rt5_orderId" sign:"1"`                // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一
	RefundOrderNum string `json:"rt6_refundOrderNum" sign:"1"`         // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一
	SerialNumber   string `json:"rt7_serialNumber,omitempty" sign:"1"` // 合利宝平台唯一流水号
	Amount         string `json:"rt8_amount,omitempty" sign:"1"`       // 退款金额,以元为单位,最小金额为0.01
	Currency       string `json:"rt9_currency,omitempty" sign:"1"`     // 币种类型,CNY:人民币
	Sign           string `json:"sign" sign:"0"`                       // 签名
}

// AppPayOrderRefundQueryReq 交易订单退款查询接口
type AppPayOrderRefundQueryReq struct {
	BizType        string `json:"P1_bizType" sign:"1"`        // 交易类型,固定为:AppPayRefundQuery
	RefundOrderId  string `json:"P2_refundOrderId" sign:"1"`  // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一与合利宝退款订单流水号二选一
	CustomerNumber string `json:"P3_customerNumber" sign:"1"` // 合利宝分配的商户号
	SerialNumber   string `json:"P4_serialNumber" sign:"0"`   // 合利宝平台返回的退款订单流水号与商户退款订单号二选一
}

// AppPayOrderRefundQueryRes 交易订单退款查询接口
type AppPayOrderRefundQueryRes struct {
	BizType                 string                     `json:"rt1_bizType" sign:"1"`                            // 交易类型,固定为:AppPayRefundQuery
	RetCode                 string                     `json:"rt2_retCode" sign:"1"`                            // 返回码
	RetMsg                  string                     `json:"rt3_retMsg,omitempty" sign:"0"`                   // 返回信息
	CustomerNumber          string                     `json:"rt4_customerNumber" sign:"1"`                     // 合利宝分配的商户号
	OrderId                 string                     `json:"rt5_orderId" sign:"1"`                            // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一
	RefundOrderNum          string                     `json:"rt6_refundOrderNum" sign:"1"`                     // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一
	SerialNumber            string                     `json:"rt7_serialNumber,omitempty" sign:"1"`             // 合利宝平台唯一流水号
	OrderStatus             string                     `json:"rt8_orderStatus" sign:"1"`                        // 订单状态,BEFORERECEIVE:等待处理,RECEIVE:接收成功,INIT:初始化,DOING:处理中,SUCCESS:成功,FAIL:失败,CLOSE:关闭
	Amount                  string                     `json:"rt9_amount" sign:"1"`                             // 退款金额,以元为单位,最小金额为0.01
	Currency                string                     `json:"rt10_currency,omitempty" sign:"1"`                // 币种类型,CNY:人民币
	RefundOrderCompleteDate string                     `json:"rt11_refundOrderCompleteDate,omitempty" sign:"0"` // 退款完成时间,格式:yyyy-MM-dd HH:mm:ss
	RefundChannelOrderNum   string                     `json:"rt12_refundChannelOrderNum,omitempty" sign:"0"`   // 第三方平台退款订单号
	Desc                    string                     `json:"rt13_desc,omitempty" sign:"0"`                    // 退款原因/备注,若商户传入,会在下发给用户的退款账单消息中体现退款原因
	RefundOrderAttribute    string                     `json:"rt14_refundOrderAttribute,omitempty" sign:"0"`    // 通道订单属性,UNDIRECT_DEFAULT:间连通道,DIRECT_CHANNEL 直连通道
	AppPayType              string                     `json:"rt15_appPayType,omitempty" sign:"0"`              // 客户端类型,AppPayType
	PayType                 string                     `json:"rt16_payType,omitempty" sign:"0"`                 // 支付类型,PayType
	RefundMarketingRuleJson string                     `json:"refundMarketingRule,omitempty" sign:"0"`          // 退款营销规则串,营销规则,JSON格式字符串,des加密返回,
	RefundPromotionDetail   string                     `json:"refundPromotionDetail,omitempty" sign:"0"`        // 微信/支付宝等卡券优惠详情串
	DebitAmount             string                     `json:"debitAmount,omitempty" sign:"0"`                  // 退款面额,以元为单位,最小金额为0.01
	RefundCashAmount        string                     `json:"refundCashAmount,omitempty" sign:"0"`             // 用户实际退款到账金额,以元为单位,最小金额为0.01
	RetReasonDesc           string                     `json:"retReasonDesc,omitempty" sign:"0"`                // 通道上游业务结果描述
	UpAddData               string                     `json:"upAddData,omitempty" sign:"0"`                    // 银联二维码平台付款方附加数据
	ProductFee              string                     `json:"productFee,omitempty" sign:"0"`                   // 本次退款成功,退还手续费,单位:元
	FeeAccountAmt           string                     `json:"feeAccountAmt,omitempty" sign:"0"`                // 本次退款成功,退还平台商补贴的手续费,单位:元
	Sign                    string                     `json:"sign" sign:"0"`                                   // 签名
	RefundMarketingRule     *AppPayRefundMarketingRule `json:"-" sign:"0"`                                      // 退款营销规则
}

// OrderPayResultNotifyReq 订单支付结果异步通知
type OrderPayResultNotifyReq struct {
	CustomerNumber          string                       `json:"rt1_customerNumber" sign:"1"`                   // 合利宝分配的商户号
	OrderId                 string                       `json:"rt2_orderId" sign:"1"`                          // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一
	SystemSerial            string                       `json:"rt3_systemSerial,omitempty" sign:"1"`           // 合利宝系统唯一支付流水号
	Status                  string                       `json:"rt4_status" sign:"1"`                           // 订单状态,INIT:已接收,DOING:处理中,SUCCESS:成功,FAIL:失败,CLOSE:关闭,CANCEL:撤销
	OrderAmount             float64                      `json:"rt5_orderAmount,string" sign:"1"`               // 订单金额,以元为单位,最小金额为0.01
	Currency                string                       `json:"rt6_currency" sign:"1"`                         // 币种,CNY:人民币
	Timestamp               string                       `json:"rt7_timestamp" sign:"1"`                        // 通知时间,精确到通知时间的毫秒数
	Desc                    string                       `json:"rt8_desc,omitempty" sign:"1"`                   // 备注,微信用户openId(暂不用)
	OpenId                  string                       `json:"rt10_openId,omitempty" sign:"0"`                // 微信用户openId
	ChannelOrderNum         string                       `json:"rt11_channelOrderNum,omitempty" sign:"0"`       // 第三方平台订单号
	OrderCompleteDate       string                       `json:"rt12_orderCompleteDate,omitempty" sign:"0"`     // 订单完成时间,格式:yyyy-MM-dd HH:mm:ss
	OnlineCardType          string                       `json:"rt13_onlineCardType,omitempty" sign:"0"`        // 支付卡类型,DEBIT(借记卡),CREDIT(贷记卡),UNKNOWN(未知),CFT(钱包零钱)
	CashFee                 string                       `json:"rt14_cashFee,omitempty" sign:"0"`               // 上游返回:现金支付金额, (订单总金额-现金券金额=现金支付金额)
	CouponFee               string                       `json:"rt15_couponFee,omitempty" sign:"0"`             // 上游返回:现金券金额,
	FundBillList            string                       `json:"rt16_fundBillList,omitempty" sign:"0"`          // 支付宝使用的资金渠道和优惠信息,支付宝时返回
	OutTransactionOrderId   string                       `json:"rt17_outTransactionOrderId,omitempty" sign:"0"` // 微信支付宝交易订单号,成功时有返回
	BankType                string                       `json:"rt18_bankType,omitempty" sign:"0"`              // 用户付款银行,成功时有返回
	SubOpenId               string                       `json:"rt19_subOpenId,omitempty" sign:"0"`             // 微信子商户subOpenId或支付宝子商户buyer_id
	OrderAttribute          string                       `json:"rt20_orderAttribute,omitempty" sign:"0"`        // 标记通道订单属性,UNDIRECT_DEFAULT:间连通道,DIRECT_CHANNEL:直连通道
	MarketingRuleJson       string                       `json:"rt21_marketingRule,omitempty" sign:"0"`         // 营销规则,JSON格式字符串,des加密传输
	PromotionDetail         string                       `json:"rt22_promotionDetail,omitempty" sign:"0"`       // 上游返回的优惠详情
	PaymentAmount           string                       `json:"rt23_paymentAmount,omitempty" sign:"0"`         // 用户实际支付金额
	CreditAmount            string                       `json:"rt24_creditAmount,omitempty" sign:"0"`          // 入账面额(不扣手续费)
	AppId                   string                       `json:"rt25_appId,omitempty" sign:"0"`                 // 子商户公众号sub_appid
	AppPayType              string                       `json:"rt26_appPayType,omitempty" sign:"0"`            // 客户端类型,AppPayType
	PayType                 string                       `json:"rt27_payType,omitempty" sign:"0"`               // 支付类型,PayType
	SplitRuleJson           string                       `json:"ruleJson,omitempty" sign:"0"`                   // 响应分账结果规则以及对应状态
	ProductFee              string                       `json:"productFee,omitempty" sign:"0"`                 // 该笔交易产生的手续费,成功时返回
	ChannelSettlementAmount string                       `json:"channelSettlementAmount,omitempty" sign:"0"`    // 渠道结算金额,成功时返回
	RealCreditAmount        string                       `json:"realCreditAmount,omitempty" sign:"0"`           // 该笔交易成功后收单商户实际入账发生额,成功时返回
	TradeType               string                       `json:"tradeType,omitempty" sign:"0"`                  // 微信交易类型,微信成功时有返回 https://myshangpu.yuque.com/org-wiki-myshangpu-sfbw9n/qnpgdn/phtfgreqpllk9i2m#p13iF
	ChargeFlag              string                       `json:"chargeFlag,omitempty" sign:"0"`                 // 支付宝费率活动标识,支付宝活动通道成功时有返回 https://myshangpu.yuque.com/org-wiki-myshangpu-sfbw9n/qnpgdn/phtfgreqpllk9i2m#ap0Bq
	UpAddData               string                       `json:"upAddData,omitempty" sign:"0"`                  // 银联二维码返回的付款方附加数据,Base64编码
	ResvData                string                       `json:"resvData,omitempty" sign:"0"`                   // 银联二维码返回的保留数据,Base64编码
	ChannelSubMerchantNo    string                       `json:"subMerchantNo,omitempty" sign:"0"`              // 渠道子商户号(U/A/T)
	FeeRate                 string                       `json:"feeRate,omitempty" sign:"0"`                    // 商户交易的费率,被扫支付成功时返回
	FeeAccountAmt           string                       `json:"feeAccountAmt,omitempty" sign:"0"`              // 平台商补贴的手续费,单位:元,被扫支付成功时返回
	Sign                    string                       `json:"sign" sign:"0"`                                 // MD5 签名结果
	MarketingRule           *AppPayPreOrderMarketingRule `json:"-" sign:"0"`                                    // 营销规则
	SplitRules              []*AppPayPreOrderSplitRule   `json:"-" sign:"0"`                                    // 分账规则
}

// OrderRefundResultNotifyReq 订单退款结果异步通知
type OrderRefundResultNotifyReq struct {
	CustomerNumber          string                     `json:"rt1_customerNumber" sign:"1"`                    // 合利宝分配的商户号
	OrderId                 string                     `json:"rt2_orderId" sign:"1"`                           // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一
	RefundOrderNum          string                     `json:"rt3_refundOrderNum" sign:"1"`                    // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一
	SerialNumber            string                     `json:"rt4_serialNumber,omitempty" sign:"1"`            // 合利宝平台唯一流水号
	Status                  string                     `json:"rt5_status" sign:"1"`                            // 订单状态,INIT:初始化,DOING:处理中,SUCCESS:成功,FAIL:失败,CLOSE:关闭
	Amount                  float64                    `json:"rt6_amount,string" sign:"1"`                     // 退款金额,以元为单位,最小金额为0.01
	Currency                string                     `json:"rt7_currency,omitempty" sign:"1"`                // 币种类型,CNY:人民币
	Timestamp               string                     `json:"rt8_timestamp,omitempty" sign:"1"`               // 通知时间,精确到通知时间的毫秒数
	RefundOrderCompleteDate string                     `json:"rt9_refundOrderCompleteDate,omitempty" sign:"0"` // 退款完成时间,格式:yyyy-MM-dd HH:mm:ss
	RefundChannelOrderNum   string                     `json:"rt10_refundChannelOrderNum,omitempty" sign:"0"`  // 第三方平台退款订单号
	Desc                    string                     `json:"rt11_desc,omitempty" sign:"0"`                   // 退款原因/备注,若商户传入,会在下发给用户的退款账单消息中体现退款原因
	RefundOrderAttribute    string                     `json:"rt12_refundOrderAttribute,omitempty" sign:"0"`   // 通道订单属性,UNDIRECT_DEFAULT:间连通道,DIRECT_CHANNEL 直连通道
	AppPayType              string                     `json:"rt13_appPayType,omitempty" sign:"0"`             // 客户端类型,AppPayType
	PayType                 string                     `json:"rt14_payType,omitempty" sign:"0"`                // 支付类型,PayType
	RefundMarketingRuleJson string                     `json:"refundMarketingRule,omitempty" sign:"0"`         // 退款营销规则串,营销规则,JSON格式字符串,des加密返回,
	RefundPromotionDetail   string                     `json:"refundPromotionDetail,omitempty" sign:"0"`       // 微信/支付宝等卡券优惠详情串
	DebitAmount             string                     `json:"debitAmount,omitempty" sign:"0"`                 // 退款面额,以元为单位,最小金额为0.01
	RefundCashAmount        string                     `json:"refundCashAmount,omitempty" sign:"0"`            // 用户实际退款到账金额,以元为单位,最小金额为0.01
	RetReasonDesc           string                     `json:"retReasonDesc,omitempty" sign:"0"`               // 通道上游业务结果描述
	UpAddData               string                     `json:"upAddData,omitempty" sign:"0"`                   // 银联二维码平台付款方附加数据
	ProductFee              string                     `json:"productFee,omitempty" sign:"0"`                  // 本次退款成功,退还手续费,单位:元
	FeeAccountAmt           string                     `json:"feeAccountAmt,omitempty" sign:"0"`               // 本次退款成功,退还平台商补贴的手续费,单位:元
	Sign                    string                     `json:"sign" sign:"0"`                                  // 签名
	RefundMarketingRule     *AppPayRefundMarketingRule `json:"-" sign:"0"`                                     // 退款营销规则
}

// MerchantRiskNotifyReq 风险商户信息接口推送
type MerchantRiskNotifyReq struct {
	SubMerchantNo      string                   `json:"subMerchantNo,omitempty" sign:"1"`      // 渠道子商户号(U/A/T)
	EventId            string                   `json:"eventId,omitempty" sign:"1"`            // 事件号,风险事件号,唯一值
	AppPayType         string                   `json:"appPayType,omitempty" sign:"1"`         // 客户端,ALIPAY/WXPAY
	ParentMerchantNo   string                   `json:"parentMerchantNo,omitempty" sign:"1"`   // 商户在我方平台系统对应的平台商户号,唯一
	ParentMerchantName string                   `json:"parentMerchantName,omitempty" sign:"0"` // 平台商户名称
	RiskDesc           string                   `json:"riskDesc,omitempty" sign:"0"`           // 风险描述
	Info               []MerchantRiskNotifyInfo `json:"info,omitempty" sign:"0"`               // 详情内容
	Sign               string                   `json:"sign,omitempty" sign:"0"`               // 签名
}

type MerchantRiskNotifyInfo struct {
	Ids        []int64 `json:"ids"`        // 风控的报备id
	MerchantNo string  `json:"merchantNo"` // 涉及的合利宝商户号或子商户号
	SigName    string  `json:"sigName"`    // 涉及的商户或子商户签约名
}
