package model

// GetMerchantSignContractUrlReq 获取商户签章合同链接
type GetMerchantSignContractUrlReq struct {
	InterfaceName      string `json:"interfaceName"`                // 接口名称,固定值:getSignContractUrl
	MerchantId         string `json:"merchantId"`                   // 子商户商编
	CallBackUrl        string `json:"callBackUrl,omitempty"`        // 子商户签约成功后的通知地址
	SuccessCallBackUrl string `json:"successCallBackUrl,omitempty"` // 子商户签约成功后的页面跳转地址
}

// GetMerchantSignContractUrlRes 获取商户签章合同链接
type GetMerchantSignContractUrlRes struct {
	MerchantId string `json:"merchantId"`        // 子商户商编
	SignUrl    string `json:"signUrl,omitempty"` // 链接
}

// MerchantSignContractQueryReq 获取商户签章查询
type MerchantSignContractQueryReq struct {
	InterfaceName string `json:"interfaceName"` // 接口名称,固定值:contractQuery
	MerchantId    string `json:"merchantId"`    // 子商户商编
}

// MerchantSignContractQueryRes 获取商户签章查询
type MerchantSignContractQueryRes struct {
	MerchantId string `json:"merchantId"`           // 子商户商编
	SignStatus string `json:"signStatus,omitempty"` // 签章状态,UNSIGNED:未签约,SIGNED:已签约,NO_PRODUCE:未生成协议
}
