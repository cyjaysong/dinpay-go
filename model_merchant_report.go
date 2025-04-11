package dinpay

// MerchantApplyAppReportReq 商户上游报备
type MerchantApplyAppReportReq struct {
	InterfaceName string `json:"interfaceName"`           // 接口名称,固定值:reporting
	OrderNo       string `json:"orderNo"`                 // 请求单号
	MerchantId    string `json:"subMerchantId"`           // 子商户编号,进件审核通过后才有的商户号
	PaymentType   string `json:"paymentType"`             // 支付客户端类型,constants.PaymentType
	ReportingType string `json:"reportingType,omitempty"` // 报备类型 constants.ReportType
}

// MerchantApplyAppReportRes 商户上游报备
type MerchantApplyAppReportRes struct {
	OrderNo         string `json:"orderNo"`         // 请求单号
	MerchantId      string `json:"subMerchantId"`   // 子商户编号,进件审核通过后才有的商户号
	PaymentType     string `json:"paymentType"`     // 支付客户端类型,constants.PaymentType
	ReportingStatus string `json:"reportingStatus"` // 报备结果,constants.ReportResult,只有DOING/SUCCESS两种状态
}

// MerchantAppApplyQueryReq 商户上游报备查询
type MerchantAppApplyQueryReq struct {
	InterfaceName string `json:"interfaceName"`           // 接口名称,固定值:reporting
	OrderNo       string `json:"orderNo"`                 // 请求单号
	MerchantId    string `json:"subMerchantId"`           // 子商户编号,进件审核通过后才有的商户号
	PaymentType   string `json:"paymentType"`             // 支付客户端类型,constants.PaymentType
	ReportingType string `json:"reportingType,omitempty"` // 报备类型 constants.ReportType
}

// MerchantAppApplyQueryRes 商户上游报备查询
type MerchantAppApplyQueryRes struct {
	OrderNo         string             `json:"orderNo"`                // 请求单号
	MerchantId      string             `json:"subMerchantId"`          // 子商户编号,进件审核通过后才有的商户号
	PaymentType     string             `json:"paymentType"`            // 支付客户端类型,constants.PaymentType
	ReportingStatus string             `json:"reportingStatus"`        // 报备结果,constants.ReportResult,只有DOING/SUCCESS两种状态
	ReportingMsg    string             `json:"reportingMsg,omitempty"` // 报备失败原因,当报备结果为DOING、FAIL 时可返回
	CashierUrl      string             `json:"cashierUrl,omitempty"`   // 商户收银台url,预留字段
	UpstreamData    []UpstreamDataItem `json:"upstreamData,omitempty"` // 上游平台信息
}

// UpstreamDataItem 上游平台信息
type UpstreamDataItem struct {
	ChannelMerchantNo string `json:"channelMerchantNo,omitempty"` // 上游平台渠道号
	WxRuleId          string `json:"wxRuleId"`                    // 微信结算规则ID
	UpstreamNo        string `json:"upstreamNo,omitempty"`        // 上游商户号
	AuthorizeStatus   string `json:"authorizeStatus,omitempty"`   // 授权状态
	ConfirmStatus     string `json:"confirmStatus,omitempty"`     // 认证状态
	QrCode            string `json:"qrCode,omitempty"`            // 联系人信息确认二维码,base64
	FailedMsg         string `json:"failedMsg"`                   // 失败原因
}

// MerchantReIdentifyReq 商户重新发起认证
type MerchantReIdentifyReq struct {
	InterfaceName string `json:"interfaceName"` // 接口名称,固定值:reCertification
	MerchantId    string `json:"subMerchantId"` // 子商户号
	PaymentType   string `json:"paymentType"`   // 支付客户端类型,constants.PaymentType
	UpstreamNo    string `json:"upstreamNo"`    // 通过 Client.MerchantAppApplyQuery 接口获取
}

// MerchantReIdentifyRes 商户重新发起认证
type MerchantReIdentifyRes struct {
	ReStatus string `json:"reStatus"`           // 重新认证状态,SUCCESS/FAIL
	ErrorMsg string `json:"errorMsg,omitempty"` // 认证失败返回错误原因
}

// MerchantWxPublicApplyReq 商户微信交易配置
type MerchantWxPublicApplyReq struct {
	InterfaceName string   `json:"interfaceName"`           // 接口名称,固定值:wxAppidConfig
	OrderNo       string   `json:"orderNo"`                 // 请求单号
	MerchantId    string   `json:"subMerchantId"`           // 子商户号
	ChannelName   string   `json:"channelName"`             // 渠道名
	AppIds        []string `json:"appIds"`                  // 支付公众号/小程序appid列表
	PayCatalogs   []string `json:"payCatalogs"`             // 支付授权目录列表,url格式参考微信公众号文档https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=7_3
	ReceiptAppIds []string `json:"receiptAppIds,omitempty"` // 关注appid列表,不传，默认与 AppIds 一致
	ReportingType string   `json:"reportingType,omitempty"` // 报备类型 constants.ReportType
}

// MerchantWxPublicApplyRes 商户微信交易配置
type MerchantWxPublicApplyRes struct {
	OrderNo      string `json:"orderNo"`       // 请求单号
	MerchantId   string `json:"subMerchantId"` // 子商户号
	ChannelName  string `json:"channelName"`   // 渠道名
	ConfigStatus string `json:"configStatus"`  // 结果,constants.ReportResult
}

// MerchantWxPayDirsApplyReq 商户微信配置支付目录
type MerchantWxPayDirsApplyReq struct {
	InterfaceName string `json:"interfaceName"`           // 接口名称,固定值:wxAddConfig
	OrderNo       string `json:"orderNo"`                 // 请求单号
	MerchantId    string `json:"subMerchantId"`           // 子商户号
	UpstreamNo    string `json:"upstreamNo"`              // 上游商户号
	PayCatalog    string `json:"payCatalog"`              // 支付授权目录,url格式参考微信公众号文档https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=7_3
	ReportingType string `json:"reportingType,omitempty"` // 报备类型 constants.ReportType
}

// MerchantWxPayDirsApplyRes 商户微信配置支付目录
type MerchantWxPayDirsApplyRes struct {
	OrderNo      string `json:"orderNo"`        // 请求单号
	MerchantId   string `json:"subMerchantId"`  // 子商户号
	ConfigStatus string `json:"status"`         // 结果,constants.ReportResult
	ConfigMsg    string `json:"wxAddConfigMsg"` // 配置结果描述,失败时返回
}

// MerchantWxPublicApplyQueryReq 商户微信交易配置查询
type MerchantWxPublicApplyQueryReq struct {
	InterfaceName string   `json:"interfaceName"`           // 接口名称,固定值:wxAppidConfigQuery
	OrderNo       string   `json:"orderNo"`                 // 请求单号
	MerchantId    string   `json:"subMerchantId"`           // 子商户号
	ChannelName   string   `json:"channelName"`             // 渠道名
	ReceiptAppIds []string `json:"receiptAppIds,omitempty"` // 关注appid列表,不传，默认与 AppIds 一致
	ReportingType string   `json:"reportingType,omitempty"` // 报备类型 constants.ReportType
}

// MerchantWxPublicApplyQueryRes 商户微信交易配置查询
type MerchantWxPublicApplyQueryRes struct {
	OrderNo      string `json:"orderNo"`             // 请求单号
	MerchantId   string `json:"subMerchantId"`       // 子商户号
	ChannelName  string `json:"channelName"`         // 渠道名
	ConfigStatus string `json:"configStatus"`        // 结果,constants.ReportResult
	FailedMsg    string `json:"failedMsg,omitempty"` // 失败原因
	// 以下在configStatus为’SUCCESS’时才有返回;  receiptAppIds, channelName都不传则返回JSON数组upstreamData
	UpstreamData []MerchantWxPublicApplyUpstreamDataItem `json:"upstreamData,omitempty"` // 上游报备信息
}

type MerchantWxPublicApplyUpstreamDataItem struct {
	ReportId          int64    `json:"reportId,omitempty"`          // 报备ID
	ChannelId         string   `json:"channelId,omitempty"`         // 渠道
	UpstreamNo        string   `json:"upstreamNo,omitempty"`        // 上游商户号
	AppIds            []string `json:"appIds,omitempty"`            // 支付公众号
	PayCatalogs       []string `json:"payCatalogs,omitempty"`       // 授权目录
	ReceiptAppIds     []string `json:"receiptAppIds,omitempty"`     // 关注小程序
	AppidConfigStatus string   `json:"appidConfigStatus,omitempty"` // 支付目录,appId配置状态
	ConfigChannelMsg  string   `json:"configChannelMsg,omitempty"`  // 配置信息
}
