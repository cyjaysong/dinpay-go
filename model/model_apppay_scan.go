package model

// AppPayScanOrderReq 主扫/被扫下单接口
type AppPayScanOrderReq struct {
	InterfaceName  string  `json:"interfaceName"`  // 接口类型,固定为:AppPay
	PaymentType    string  `json:"paymentType"`    // 支付客户端类型,constants.PaymentType
	PaymentMethods string  `json:"paymentMethods"` // 支付类型,constants.PaymentMethods SCAN:主扫;SWIPE:被扫
	MerchantId     string  `json:"-"`              // 商户编号
	PayAmount      float64 `json:"payAmount"`      // 订单金额,以元为单位，最小金额为0.01
	Currency       string  `json:"currency"`       // 币种类型,CNY:人民币
	PaymentCode    string  `json:"paymentCode"`    // 付款授权码, PaymentMethods 为SCAN时传入1即可;SWIPE时传入一组字符串(用户付款码).
	// 微信:(注:用户条形码规则:18 位纯数字，以 10、11、12、13、 14、15 开头).
	// 支付宝:支付授权码，25~30 开头的长度为 16~24 位的数字
	// 银联:62开头的银联支付授权码数字
	// 京东:62开头的银联支付授权码数字
	// 注：渠道侧规则有可能发送变动，以实际获取值为主
	OrderNo        string `json:"orderNo"`                  // 商户系统内部订单号，要求50字符以内，同一商户号下订单号唯一
	OrderIp        string `json:"orderIp"`                  // 该笔订单用户付款IP/商户IP
	GoodsName      string `json:"goodsName"`                // 商品名称,会在用户账单上显示
	GoodsDetail    string `json:"goodsDetail,omitempty"`    // 商品优惠详情:参加上游单品优惠活动的则该字段必传，且必须按照规范上传，JSON字符串格式
	GoodsTag       string `json:"goodsTag,omitempty"`       // 商品标记,微信平台配置的商品标记,用于优惠券或者满减使用
	OrderDesc      string `json:"orderDesc,omitempty"`      // 订单备注信息,商户可自定义上送,原样返回。
	NotifyUrl      string `json:"notifyUrl,omitempty"`      // 通知回调地址,接收交易系统支付结果通知的回调地址,通知url必须为外网可访问的url
	AppId          string `json:"appId,omitempty"`          // 交易appId
	LimitCreditPay string `json:"limitCreditPay,omitempty"` // 是否限制借贷记,1:禁用贷记卡,0:不限制,2:针对支付宝禁用借记;不传以上游通道风控判断为准
	TimeExpire     string `json:"timeExpire,omitempty"`     // 超时时间,微信/银联传参规则:超时时间单位为秒,微信超时时间要大于等于60秒;支付宝传参规则:取值范围:1m～15d。m-分钟,h-小时,d-天,1c-当天(1c-当天的情况下,无论交易何时创建,都在0点关闭)。该参数数值不接受小数点,1.5h,可转换为90m。
	Identity       string `json:"identity,omitempty"`       // 实名参数,实名支付功能,用于公安和保险类商户使用，该字段为JSON格式数据,微信、支付宝格式不同
	StoreId        string `json:"storeId,omitempty"`        // 商户自定义的上游门店编码
	AlipayStoreId  string `json:"alipayStoreId,omitempty"`  // 支付宝店铺编号
	TermNo         string `json:"termNo,omitempty"`         // 终端号
	TermInfo       string `json:"termInfo,omitempty"`       // 银联终端信息,参加银联官方单品营销必送,json字符串格式,详见 https://myshangpu.yuque.com/org-wiki-myshangpu-sfbw9n/qnpgdn/phtfgreqpllk9i2m#DuXus
	TerminalSysNo  string `json:"terminalSysNo,omitempty"`  // 终端序列号,在平台已报备过的终端信息的绑定号(采集接口返回);线下场景时必填
	EncryptRandNum string `json:"encryptRandNum,omitempty"` // 加密随机因子,仅在被扫支付类交易报文中出现:若付款码为19位数字,则取后6位;若付款码码为EMV二维码,则取其tag57的卡号/token号的后6位
	SecretText     string `json:"secretText,omitempty"`     // 密文数据,仅在被扫支付类交易报文中出现：64bit 的密文数据，对终端硬件序列号和加密随机因子加密后的结果。本子域取值为：64bit 密文数据进行base64 编码后的结果
	SceneInfo      string `json:"sceneInfo,omitempty"`      // 场景信息,该字段用于上报场景信息,目前支持上报实际门店信息; 详见 https://myshangpu.yuque.com/org-wiki-myshangpu-sfbw9n/qnpgdn/phtfgreqpllk9i2m#fxUGg
	EduSubject     string `json:"eduSubject,omitempty"`     // 学校名称、场景名称,该参数传入学校名称、场景名称，格式要求：“学校名称_场景名称_其它描述信息”，其中“其它描述信息部分”不做要求。
	BusinessParams string `json:"businessParams,omitempty"` // 学校ID和场景ID,商户传入业务信息,具体值要和支付宝约定
	ExtendParams   string `json:"extendParams,omitempty"`   // 业务扩展参数,目前是扫码支付宝用:https://open.alipay.com/api
	Pid            string `json:"pid"`                      // 服务商pid,不传以平台配置为准(仅银联二维码生效)
	OpenId         string `json:"openId,omitempty"`         // 用户id
	ReportId       string `json:"reportId,omitempty"`       // 报备id

	SplitType string `json:"splitType,omitempty"` // 分账类型,FIXED_AMOUNT:固定金额(默认,目前只支持固定金额);RATE:比率
	// Deprecated: 请勿直接赋值,应调用SplitRules添加
	SplitRulesJson string                     `json:"splitRules,omitempty"` // 分账规则,Json格式字符串;分账类型和分账规则串出现时须2个字段都要上送
	SplitRules     []*AppPayPreOrderSplitRule `json:"-"`                    // 分账规则
	// Deprecated: 请勿直接赋值,应调用MarketingRules添加
	MarketingRulesJson string                       `json:"marketingRules,omitempty"` // 营销规则,JSON格式字符串
	MarketingRules     *AppPayPreOrderMarketingRule `json:"-"`                        // 营销规则
}

// AppPayScanOrderRes 主扫/被扫下单接口
type AppPayScanOrderRes struct {
	InterfaceName           string                       `json:"interfaceName"`                     // 接口类型,固定为:AppPay
	PaymentType             string                       `json:"paymentType,omitempty"`             // 支付客户端类型,constants.PaymentType
	PaymentMethods          string                       `json:"paymentMethods"`                    // 支付类型,constants.PaymentMethods SCAN:主扫;SWIPE:被扫
	MerchantId              string                       `json:"merchantId"`                        // 商户编号
	OrderNo                 string                       `json:"orderNo"`                           // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一
	PayAmount               float64                      `json:"payAmount"`                         // 交易金额,以元为单位,最小金额为0.01
	Currency                string                       `json:"currency"`                          // 币种类型,CNY:人民币
	Qrcode                  string                       `json:"qrcode,omitempty"`                  // 二维码信息,用于生成二维码图片，展示给用户进行扫码支付
	OrderStatus             string                       `json:"orderStatus,omitempty"`             // 订单状态,被扫支付时响应状态
	ChannelNumber           string                       `json:"channelNumber,omitempty"`           // 上游请求订单号,智付交易订单号
	TransactionOrderNo      string                       `json:"transactionorderNo,omitempty"`      // 上游交易单号
	ChannelRetCode          string                       `json:"channelRetCode,omitempty"`          // 上游返回码
	ChannelSubMerchantNo    string                       `json:"subMerchantNo,omitempty"`           // 上游子商户号(U/A/T)
	AppId                   string                       `json:"appid,omitempty"`                   // 子商户公众号sub_appid
	OpenId                  string                       `json:"openid,omitempty"`                  // 用户openId
	SubOpenId               string                       `json:"subopenid,omitempty"`               // subOpenId;微信子商户subOpenId.或支付宝子商户用户buyer_id
	BankType                string                       `json:"bankType,omitempty"`                // 用户付款银行
	PayCardType             string                       `json:"payCardType,omitempty"`             // 支付卡类型,DEBIT(借记卡);CREDIT(贷记卡);UNKNOWN(未知);CFT(钱包零钱)
	OrderCreditAmount       float64                      `json:"orderCreditAmount,omitempty"`       // 订单入账面额(不扣手续费)
	MerchantCreditAmount    float64                      `json:"merchantCreditAmount,omitempty"`    // 商户实际入账发生额,该笔交易成功后收单商户实际入账发生额,成功时返回
	PaymentAmount           float64                      `json:"paymentAmount,omitempty"`           // 用户实际支付金额
	ChannelSettlementAmount float64                      `json:"channelSettlementAmount,omitempty"` // 渠道结算金额,成功时返回
	OrderFee                float64                      `json:"orderFee,omitempty"`                // 该笔交易产生的手续费,成功时返回
	ChargeFlag              string                       `json:"chargeFlag,omitempty"`              // 渠道支付宝费率活动标识
	MerchantFee             float64                      `json:"merchantFee,omitempty"`             // 商户费率,商户交易的费率,被扫支付成功时返回
	FeeAccountAmt           float64                      `json:"feeAccountAmt,omitempty"`           // 平台商补贴的手续费,单位:元,被扫支付成功时返回
	ReceiverFee             float64                      `json:"receiverFee,omitempty"`             // 实收手续费,单位:元,被扫支付成功时返回
	OfflineFee              float64                      `json:"offlineFee,omitempty"`              // 后收手续费,单位:元,被扫支付成功时返回
	OrderPayDate            string                       `json:"orderPayDate,omitempty"`            // 订单完成时间,格式:yyyy-MM-dd HH:mm:ss
	WxTradeType             string                       `json:"wxTradeType,omitempty"`             // 微信交易类型,被扫支付成功时有返回
	UpAddData               string                       `json:"upAddData,omitempty"`               // 银联二维码返回的付款方附加数据,Base64编码
	ResvData                string                       `json:"resvData,omitempty"`                // 银联二维码返回的保留数据,Base64编码的
	FundBillList            string                       `json:"fundBillList,omitempty"`            // 交易支付使用的资金渠道和优惠信息-支付宝
	PromotionDetail         string                       `json:"promotionDetail,omitempty"`         // 上游返回的优惠详情,JSON字符串
	VoucherDetailList       string                       `json:"voucherDetailList,omitempty"`       // 支付宝优惠信息详情,JSON字符串
	MarketingRulesJson      string                       `json:"marketingRules,omitempty"`          // 营销规则,JSON格式字符串，des加密传输
	MarketingRules          *AppPayPreOrderMarketingRule `json:"-"`                                 // 营销规则
	SplitRulesJson          string                       `json:"splitRules,omitempty"`              // 分账规则及状态,响应分账结果规则以及对应状态
	SplitRules              []*AppPayPreOrderSplitRule   `json:"-"`                                 // 分账规则
}
