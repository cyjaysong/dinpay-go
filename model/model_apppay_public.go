package model

// AppPayPublicPreOrderReq 公众号/JS/服务窗预下单接口
type AppPayPublicPreOrderReq struct {
	InterfaceName  string  `json:"interfaceName"`            // 接口类型,固定为:AppPayPublic
	OrderNo        string  `json:"orderNo"`                  // 商户系统内部订单号，要求50字符以内，同一商户号下订单号唯一
	PaymentType    string  `json:"paymentType"`              // 支付客户端类型,constants.PaymentType
	PaymentMethods string  `json:"paymentMethods"`           // 支付类型,固定为:PUBLIC
	MerchantId     string  `json:"-"`                        // 商户编号
	AppId          string  `json:"appid,omitempty"`          // 交易appId
	OpenId         string  `json:"openid,omitempty"`         // 用户id
	PayAmount      float64 `json:"payAmount"`                // 订单金额,以元为单位，最小金额为0.01
	Currency       string  `json:"currency"`                 // 币种类型,CNY:人民币
	OrderIp        string  `json:"orderIp"`                  // 该笔订单用户付款IP/商户IP
	GoodsName      string  `json:"goodsName"`                // 商品名称,会在用户账单上显示
	GoodsDetail    string  `json:"goodsDetail,omitempty"`    // 商品优惠详情:参加上游单品优惠活动的则该字段必传，且必须按照规范上传，JSON字符串格式
	GoodsTag       string  `json:"goodsTag,omitempty"`       // 商品标记,微信平台配置的商品标记,用于优惠券或者满减使用
	OrderDesc      string  `json:"orderDesc,omitempty"`      // 订单备注信息,商户可自定义上送,原样返回。
	IsNative       string  `json:"isNative,omitempty"`       // 是否原生,是否原生态,1:是;0:否,不传默认是1
	LimitCreditPay string  `json:"limitCreditPay,omitempty"` // 是否限制借贷记,1:禁用贷记卡,0:不限制,2:针对支付宝禁用借记;不传以上游通道风控判断为准
	NotifyUrl      string  `json:"notifyUrl,omitempty"`      // 通知回调地址,接收交易系统支付结果通知的回调地址,通知url必须为外网可访问的url
	SuccessToUrl   string  `json:"successToUrl,omitempty"`   // 页面跳转地址,仅当 IsNative 为0微信公众号非原生态时支付完会跳转到该地址
	TimeExpire     string  `json:"timeExpire,omitempty"`     // 超时时间,微信/银联传参规则:超时时间单位为秒,微信超时时间要大于等于60秒;支付宝传参规则:取值范围:1m～15d。m-分钟,h-小时,d-天,1c-当天(1c-当天的情况下,无论交易何时创建,都在0点关闭)。该参数数值不接受小数点,1.5h,可转换为90m。
	OpenIdType     string  `json:"openIdType,omitempty"`     // userAuthCode(临时授权码)/openId(用户id),此参表示上送的 OpenId 值的类型;不填默认是userAuthCode
	Identity       string  `json:"identity,omitempty"`       // 实名参数,实名支付功能,用于公安和保险类商户使用，该字段为JSON格式数据,微信、支付宝格式不同
	StoreId        string  `json:"storeId,omitempty"`        // 商户自定义的上游门店编码
	AlipayStoreId  string  `json:"alipayStoreId,omitempty"`  // 支付宝店铺编号
	UserAgent      string  `json:"userAgent,omitempty"`      // 银联用户代理,银联JS识别用户的客户端,不填默认是UnionPay/1.0 CloudPay;详见 https://myshangpu.yuque.com/org-wiki-myshangpu-sfbw9n/qnpgdn/phtfgreqpllk9i2m#ENYf3
	TermNo         string  `json:"termNo,omitempty"`         // 终端号
	TermInfo       string  `json:"termInfo,omitempty"`       // 银联终端信息,参加银联官方单品营销必送,json字符串格式,详见 https://myshangpu.yuque.com/org-wiki-myshangpu-sfbw9n/qnpgdn/phtfgreqpllk9i2m#DuXus
	TerminalSysNo  string  `json:"terminalSysNo,omitempty"`  // 终端序列号,在平台已报备过的终端信息的绑定号(采集接口返回);线下场景时必填
	SceneInfo      string  `json:"sceneInfo,omitempty"`      // 场景信息,该字段用于上报场景信息,目前支持上报实际门店信息; 详见 https://myshangpu.yuque.com/org-wiki-myshangpu-sfbw9n/qnpgdn/phtfgreqpllk9i2m#fxUGg
	EduSubject     string  `json:"eduSubject,omitempty"`     // 学校名称、场景名称,该参数传入学校名称、场景名称，格式要求：“学校名称_场景名称_其它描述信息”，其中“其它描述信息部分”不做要求。
	BusinessParams string  `json:"businessParams,omitempty"` // 学校ID和场景ID,商户传入业务信息,具体值要和支付宝约定
	ExtendParams   string  `json:"extendParams,omitempty"`   // 业务扩展参数,目前是扫码支付宝用:https://open.alipay.com/api
	Pid            string  `json:"pid"`                      // 服务商pid,不传以平台配置为准(仅银联二维码生效)
	ReportId       string  `json:"reportId,omitempty"`       // 报备id

	SplitType string `json:"splitType,omitempty"` // 分账类型,FIXED_AMOUNT:固定金额(默认,目前只支持固定金额);RATE:比率
	// Deprecated: 请勿直接赋值,应调用SplitRules添加
	SplitRulesJson string                     `json:"splitRules,omitempty"` // 分账规则,Json格式字符串;分账类型和分账规则串出现时须2个字段都要上送
	SplitRules     []*AppPayPreOrderSplitRule `json:"-"`                    // 分账规则
	// Deprecated: 请勿直接赋值,应调用MarketingRules添加
	MarketingRulesJson string                       `json:"marketingRules,omitempty"` // 营销规则,JSON格式字符串
	MarketingRules     *AppPayPreOrderMarketingRule `json:"-"`                        // 营销规则

}

// AppPayPublicPreOrderRes 公众号/JS/服务窗预下单接口
type AppPayPublicPreOrderRes struct {
	InterfaceName        string                       `json:"interfaceName"`            // 接口类型,固定为:AppPayPublic
	PaymentType          string                       `json:"paymentType,omitempty"`    // 支付客户端类型,constants.PaymentType
	PaymentMethods       string                       `json:"paymentMethods"`           // 支付类型,固定为:PUBLIC
	MerchantId           string                       `json:"merchantId"`               // 商户编号
	OrderNo              string                       `json:"orderNo"`                  // 商户系统内部订单号,要求50字符以内,同一商户号下订单号唯一
	PayAmount            float64                      `json:"payAmount"`                // 交易金额,以元为单位,最小金额为0.01
	Currency             string                       `json:"currency"`                 // 币种类型,CNY:人民币
	AppId                string                       `json:"appid,omitempty"`          // 公众账号ID
	PayInfoJson          string                       `json:"payInfo"`                  // 原生态JS支付信息,is_raw为1时返回,json格式的字符串,作用于原生态JS支付时的参数;为小程序时返回的json串集成小程序JS接口时需要
	PayInfo              *AppPayPublicPreOrderPayInfo `json:"-"`                        // 原生态JS支付信息
	ChannelNumber        string                       `json:"channelNumber,omitempty"`  // 上游请求订单号,智付交易订单号
	ChannelRetCode       string                       `json:"channelRetCode,omitempty"` // 上游返回码,失败时透传上游返回码,仅供参考,请以订单状态为准
	ChannelSubMerchantNo string                       `json:"subMerchantNo,omitempty"`  // 渠道子商户号(U/A/T)
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
