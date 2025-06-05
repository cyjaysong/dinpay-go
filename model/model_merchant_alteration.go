package model

import "github.com/imroc/req/v3"

// MerchantModifyInfoReq 商户信息变更
type MerchantModifyInfoReq struct {
	InterfaceName               string                              `json:"interfaceName"`                    // 固定值:modifyMerchantInfo
	ChangeOrderNo               string                              `json:"changeOrderNo"`                    // 订单号,变更请求订单号
	MerchantId                  string                              `json:"subMerchantId"`                    // 子商户编号
	BusinessLicense             string                              `json:"businessLicense,omitempty"`        // 营业执照号(需人工审核)
	MerchantName                string                              `json:"merchantName,omitempty"`           // 商户名称(需人工审核)
	MerchantShowName            string                              `json:"merchantShowName,omitempty"`       // 商户简称
	BusinessLicenseType         string                              `json:"businessLicenseType,omitempty"`    // 营业证书类型,见常量 constants.CertType,默认为营业执照
	ServicePhone                string                              `json:"servicePhone,omitempty"`           // 客服电话
	MCC                         string                              `json:"mcc,omitempty"`                    // 银联MCC码
	PaymentType                 string                              `json:"paymentType,omitempty"`            // 支付客户端类型,constants.PaymentType,ALL:微信支付宝报备记录都修改
	IsChangeChannel             bool                                `json:"isChangeChannel"`                  // 是否同步通道到修改,需与支付客户端类型同时出现
	BusinessCategory            string                              `json:"businessCategory"`                 // 经营类别,见常量 constants.BusinessCategory
	CategoryId                  string                              `json:"categoryId"`                       // 商户经营子类,具体值见附件
	MicroBizType                string                              `json:"microBizType,omitempty"`           // 小微经营类型,见常量 constants.MicroBizType,小微个人商户必填
	IsChangeChannelIndustryType bool                                `json:"isChangeChannelIndustryType"`      // 是否同步修改微信侧行业类别
	DistrictCode                string                              `json:"districtCode,omitempty"`           // 区县编码
	Address                     string                              `json:"address,omitempty"`                // 地址
	BusinessDateStart           string                              `json:"businessStartDate,omitempty"`      // 经营起始日期,yyyyMMdd,非个人商户必传
	BusinessDateEnd             string                              `json:"businessEndDate,omitempty"`        // 经营结束日期,yyyyMMdd或者长期有效,非个人商户必传
	LegalPerson                 string                              `json:"legalPerson,omitempty"`            // 法人姓名(需人工审核)
	LegalPersonPhone            string                              `json:"legalPersonPhone,omitempty"`       // 法人手机号(需人工审核)
	LegalPersonId               string                              `json:"legalPersonID,omitempty"`          // 法人身份证号(需人工审核)
	LegalIdType                 string                              `json:"idType,omitempty"`                 // 法人证件类型,见常量 constants.IdType
	LegalIdCardStartDate        string                              `json:"idCardStartDate,omitempty"`        // 法人身份证有效期开始,yyyyMMdd,开通微信产品必传
	LegalIdCardEndDate          string                              `json:"idCardEndDate,omitempty"`          // 法人身份证有效期结束,yyyyMMdd或者长期有效,开通微信产品必传
	LegalPersonIdAddress        string                              `json:"legalPersonIdAddress,omitempty"`   // 法人证件居住地址,企业商户必传
	Contact                     string                              `json:"contact,omitempty"`                // 联系人名称
	ContactId                   string                              `json:"contactId,omitempty"`              // 联系人证件号,开通微信产品必填
	ContactIdType               string                              `json:"contactIdType,omitempty"`          // 联系人证件类型,见常量 constants.IdType
	ContactType                 string                              `json:"contactType"`                      // 联系人类型,见常量 constants.ContactType,LEGAL:经营者/法人,SUPER:经办人
	ContactPhone                string                              `json:"contactPhone,omitempty"`           // 联系电话
	ContactIdCardStartDate      string                              `json:"contactIdCardStartDate,omitempty"` // 联系人证件有效期开始时间,yyyyMMdd,联系人类型为经办人时必填
	ContactIdCardEndDate        string                              `json:"contactIdCardEndDate,omitempty"`   // 联系人证件有效期结束时间,yyyyMMdd或长期有效,联系人类型为经办人时必填
	UboType                     bool                                `json:"uboType"`                          // 经营者/法人是否为受益人,企业商户必传
	UboIdType                   string                              `json:"uboIdType,omitempty"`              // 受益人证件类型,受益人非经营者/法人类型时必传
	UboIdName                   string                              `json:"uboIdName,omitempty"`              // 受益人证件姓名,受益人非经营者/法人类型时必传
	UboId                       string                              `json:"uboId,omitempty"`                  // 受益人证件号码,受益人非经营者/法人类型时必传
	UboAddress                  string                              `json:"uboAddress,omitempty"`             // 受益人证件居住地址,受益人非经营者/法人类型时必传
	UboIdStartDate              string                              `json:"uboIdStartDate,omitempty"`         // 受益人证件有效期开始时间,yyyyMMdd,受益人非经营者/法人类型时必传
	UboIdEndDate                string                              `json:"uboIdEndDate,omitempty"`           // 受益人证件有效期结束时间,yyyyMMdd或长期有效,受益人非经营者/法人类型时必传
	ChangeSettleInfo            *MerchantModifyInfoChangeSettleInfo `json:"changeSettleInfo,omitempty"`       // 结算卡信息修改域
	ServiceCodes                string                              `json:"serviceCodes,omitempty"`           // 支付宝商户服务类型,如:"[\"F2F\",\"PRE_F2F\"]",constants.AlipayServiceCode
	SpecialSignName             bool                                `json:"specialSignName,omitempty"`        // 是否需要格式化商户名称,根据微信定义的签约名规则去报备,只有营业执照名字为“*”，或者个人商户才传true
	WebSite                     string                              `json:"webSite,omitempty"`                // 商户交易网站地址,与“其他产品信息”二选一必填
	OtherPdInfo                 string                              `json:"otherPdInfo,omitempty"`            // 其他产品信息,可输入公众号/小程序/app名称,与“商户交易网站地址”二选一必填
	AccessUrl                   string                              `json:"accessUrl,omitempty"`              // 接入地址,空
	IcpFiling                   string                              `json:"icpFiling,omitempty"`              // ICP备案号，平台商类型为服务商时必填，平台商类型为SaaS系统商的子商户默认读取平台商的 ICP 备案号
	FileUrlMap                  map[string]string                   `json:"imageUrlMap,omitempty"`            // 文件URL,constants.MerchantCredentialType
}

// MerchantModifyInfoChangeSettleInfo
// 1、结算账户指定书：变更开户人名称或结算卡号时必须上传
// 2、持卡人身份证正面照：变更开户人名称或结算卡号时并且结算卡类型对私必须上传
// 3、持卡人身份证反面照：变更开户人名称或结算卡号时并且结算卡类型对私必须上传
// 4、持卡人手持身份证照：变更开户人名称或结算卡号时并且结算卡类型对私必须上传
// 5、持卡人手持银行卡照：变更开户人名称或结算卡号时并且结算卡类型对私必须上传
// 6、银行开户证明：变更开户人名称或结算卡号时并且结算卡类型对公必须上传
// 7、转租证明：个体工商户变更开户人名称必须上传
// 8、开户许可证：企业商户非同名变更对私变对公必须上传
// 9、结算卡：企业商户非同名变更对公变更为非法人结算时必须上传需要将文件传到文件列表fileMap域中，并上传对应的文件签名
type MerchantModifyInfoChangeSettleInfo struct {
	CardName             string `json:"cardName"`                       // 开户人名称
	ChangeCardName       string `json:"changeCardName,omitempty"`       // 变更后开户人名称,输入该字段需上传结算账户指定书
	CardNo               string `json:"cardNo"`                         // 原结算卡号
	ChangeCardNo         string `json:"changeCardNo,omitempty"`         // 变更后结算卡号
	CardType             string `json:"cardType"`                       // 结算卡类型,constants.SettleBankType
	ChangeCardType       string `json:"changeCardType,omitempty"`       // 变更后结算卡类型,constants.SettleBankType
	BankBranchCode       string `json:"bankBranchCode"`                 // 结算卡联行号
	ChangeBankBranchCode string `json:"changeBankBranchCode,omitempty"` // 变更后结算卡联行号
	ChangeType           string `json:"changeType"`                     // 变更类型
	CardIdCard           string `json:"cardIdCard,omitempty"`           // 结算人身份证号
	CardIdCardStartDate  string `json:"cardIdCardStartDate,omitempty"`  // 结算人身份证开始日期,yyyyMMdd
	CardIdCardEndDate    string `json:"cardIdCardEndDate,omitempty"`    // 结算人身份证结束日期,yyyyMMdd或者长期有效
	SettlementMode       string `json:"settlementMode,omitempty"`       // 结算方式,constants.SettlementMode
	ArrivalMode          string `json:"arrivalMode,omitempty"`          // 结算到账方式,constants.ArrivalMode
	SettlementPhone      string `json:"settlementPhone,omitempty"`      // 结算人手机号
	SettlementCycle      string `json:"settlementCycle,omitempty"`      // 结算周期类型,constants.SettlementCycle
	SettleMode           string `json:"settleMode,omitempty"`           // 结算模式,constants.SettleMode
	//SettlementCutTim          string `json:"settlementCutTim,omitempty"`          // 待定，是否保留定时结算模式
	//SettleCutTimeEffectiveDat string `json:"settleCutTimeEffectiveDat,omitempty"` // 待定，是否保留定时结算模式
}

// MerchantModifyInfoRes 商户信息变更
type MerchantModifyInfoRes struct {
	ChangeOrderNo string `json:"changeOrderNo"` // 变更请求的订单号
	MerchantId    string `json:"subMerchantId"` // 子商户编号
	ChangeStatus  string `json:"changeStatus"`  // 审核状态,constants.AlterationStatus
	ChangeMsg     string `json:"changeMsg"`     // 审核状态说明
}

// MerchantModifyInfoQueryReq 商户信息变更查询
type MerchantModifyInfoQueryReq struct {
	InterfaceName string `json:"interfaceName"` // 固定值:changeOrderQuery
	ChangeOrderNo string `json:"changeOrderNo"` // 订单号,变更请求订单号
	MerchantId    string `json:"subMerchantId"` // 子商户编号
	ChangeType    string `json:"changeType"`    // 变更类型,固定值:MERCHANT_INFO_SYN
}

// MerchantModifyInfoQueryRes 商户信息变更查询
type MerchantModifyInfoQueryRes struct {
	ChangeOrderNo string `json:"changeOrderNo"` // 变更请求的订单号
	MerchantId    string `json:"subMerchantId"` // 子商户编号
	ChangeStatus  string `json:"changeStatus"`  // 审核状态,constants.AlterationStatus
	ChangeMsg     string `json:"changeMsg"`     // 审核状态说明
	Msg           string `json:"msg,omitempty"` // 备注
}

// MerchantCredentialImageUploadReq 商户资质图片上传
type MerchantCredentialImageUploadReq struct {
	InterfaceName  string             `json:"interfaceName"`           // 固定值:imageUpload
	OrderNo        string             `json:"orderNo,omitempty"`       // 商户入驻请求订单号,为入网待审核商户上传填写此参数
	MerchantId     string             `json:"subMerchantId,omitempty"` // 子商户编号,为进件审核通过商户上传填写此参数
	CredentialType string             `json:"credentialType"`          // 资质类型,constants.MerchantCredentialType
	FileSign       string             `json:"fileSign"`                // 资质文件SM3 HASH
	GetFileContent req.GetContentFunc `json:"-"`                       // 文件内容获取方法
}

// MerchantCredentialImageUrlUploadReq 商户资质图片Url上传
type MerchantCredentialImageUrlUploadReq struct {
	InterfaceName  string `json:"interfaceName"`           // 固定值:imageUrlUpload
	OrderNo        string `json:"orderNo,omitempty"`       // 商户入驻请求订单号,为入网待审核商户上传填写此参数
	MerchantId     string `json:"subMerchantId,omitempty"` // 子商户编号,为进件审核通过商户上传填写此参数
	CredentialType string `json:"credentialType"`          // 资质类型,constants.MerchantCredentialType
	CredentialUrl  string `json:"credentialUrl"`           // 资质文件url,资质文件地址
}

// MerchantCredentialImageUploadQueryReq 商户资质图片上传查询
type MerchantCredentialImageUploadQueryReq struct {
	InterfaceName  string `json:"interfaceName"`           // 固定值:imageUploadQuery或imageUrlUploadQuery
	OrderNo        string `json:"orderNo,omitempty"`       // 商户入驻请求订单号,为入网待审核商户上传填写此参数
	MerchantId     string `json:"subMerchantId,omitempty"` // 子商户编号,为进件审核通过商户上传填写此参数
	CredentialType string `json:"credentialType"`          // 资质类型,constants.MerchantCredentialType
}

// MerchantCredentialImageUploadRes 商户资质图片上传
type MerchantCredentialImageUploadRes struct {
	OrderNo        string `json:"orderNo,omitempty"`       // 订单号,入网待审核商户才有此参数，跟商户入驻请求订单号一致
	MerchantId     string `json:"subMerchantId,omitempty"` // 子商户编号,进件审核通过后才有商户号
	CredentialType string `json:"credentialType"`          // 资质类型,constants.MerchantCredentialType
	UploadStatus   string `json:"uploadStatus"`            // 上传结果,constants.UploadResult
	Msg            string `json:"msg,omitempty"`           // 备注
}

// MerchantCredentialImageChangeUploadReq 商户资质图片变更上传
type MerchantCredentialImageChangeUploadReq struct {
	InterfaceName  string             `json:"interfaceName"`  // 固定值:imageChangeUpload
	ChangeOrderNo  string             `json:"changeOrderNo"`  // 订单号,变更请求订单号
	MerchantId     string             `json:"subMerchantId"`  // 子商户编号
	CredentialType string             `json:"credentialType"` // 资质类型,constants.MerchantCredentialType
	ChangeType     string             `json:"changeType"`     // 变更类型,固定值:MERCHANT_CREDENTIAL
	FileSign       string             `json:"fileSign"`       // 资质文件SM3 HASH
	GetFileContent req.GetContentFunc `json:"-"`              // 文件内容获取方法
}

// MerchantCredentialImageUrlChangeUploadReq 商户资质图片Url变更上传
type MerchantCredentialImageUrlChangeUploadReq struct {
	InterfaceName  string `json:"interfaceName"`  // 固定值:imageUrlChangeUpload
	ChangeOrderNo  string `json:"changeOrderNo"`  // 订单号,变更请求订单号
	MerchantId     string `json:"subMerchantId"`  // 子商户编号
	CredentialType string `json:"credentialType"` // 资质类型,constants.MerchantCredentialType
	ChangeUrl      string `json:"changeUrl"`      // 资质文件url
}

// MerchantCredentialImageChangeUploadQueryReq 商户资质图片变更上传查询
type MerchantCredentialImageChangeUploadQueryReq struct {
	InterfaceName string `json:"interfaceName"` // 固定值:changeOrderQuery
	ChangeOrderNo string `json:"changeOrderNo"` // 订单号,变更请求订单号
	MerchantId    string `json:"subMerchantId"` // 子商户编号
	ChangeType    string `json:"changeType"`    // 变更类型,固定值:MERCHANT_CREDENTIAL
}

// MerchantCredentialImageChangeUploadRes 商户资质图片变更上传查询
type MerchantCredentialImageChangeUploadRes struct {
	ChangeOrderNo string `json:"changeOrderNo"`        // 变更请求的订单号
	MerchantId    string `json:"subMerchantId"`        // 子商户编号
	ChangeStatus  string `json:"changeStatus"`         // 审核状态
	ChangeMsg     string `json:"changeMsg"`            // 审核结果
	ChangeTime    string `json:"changeTime,omitempty"` // 生效时间,审核通过时间
}
