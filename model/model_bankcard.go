package model

// MerchantBankcardQueryRes 商户银行卡查询
type MerchantBankcardQueryRes struct {
	OrderNo        string `json:"orderId"`        // 订单号,商户入驻时的请求订单号
	BankCode       string `json:"bankCode"`       // 银行编码
	BankName       string `json:"bankName"`       // 银行名称
	OnlineCardType string `json:"onlineCardType"` // 卡类型
	CardNo         string `json:"cardNo"`         // 银行卡号
}
