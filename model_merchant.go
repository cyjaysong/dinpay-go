package dinpay

// MerchantInfoQueryReq 商户信息查询
type MerchantInfoQueryReq struct {
	FirstClassMerchantId string `json:"firstClassMerchantNo"` // 平台商商编,智付下发的
	MerchantId           string `json:"merchantNo"`           // 子商户号
}

// MerchantInfoQueryRes 商户信息查询
type MerchantInfoQueryRes struct {
	MerchantId               string `json:"merchantNo"`                 // 子商户商编
	MerchantName             string `json:"signName"`                   // 商户名称
	MerchantShowName         string `json:"showName"`                   // 商户简称,用于支付收银台的展示名
	CreateAt                 string `json:"createDate"`                 // 入网时间
	MerchantType             string `json:"merchantType"`               // 商户类型,见常量 constants.MerchantType
	OrgNum                   string `json:"orgNum,omitempty"`           // 组织机构代码
	BizCategory              string `json:"merchantCategory"`           // 经营类别,见常量 constants.BusinessCategory
	LegalName                string `json:"legalPerson"`                // 法人名字
	LegalIdNo                string `json:"legalPersonID"`              // 法人身份证号
	LegalIdCardStartDate     string `json:"idCardStartDate,omitempty"`  // 法人身份证有效起始日期
	LegalIdCardEndDate       string `json:"idCardEndDate,omitempty"`    // 法人身份证有效终止日期,1111-11-11代表长期有效
	Province                 string `json:"province"`                   // 子商户所在省份
	City                     string `json:"city"`                       // 子商户所在城市
	Region                   string `json:"region,omitempty"`           // 子商户所在区县
	Address                  string `json:"address"`                    // 通讯地址
	ContactName              string `json:"linkman"`                    // 联系人
	ContactPhone             string `json:"linkPhone"`                  // 联系电话
	ContactEmail             string `json:"email"`                      // 联系邮箱
	BindMobile               string `json:"bindMobile,omitempty"`       // 绑定手机
	ServicePhone             string `json:"servicePhone,omitempty"`     // 客服联系电话
	MerchantStatus           string `json:"merchantStatus"`             // 商户状态,见常量 constants.MerchantStatus
	AccountStatus            string `json:"accountStatus"`              // 账户状态,见常量 constants.AccountStatus
	BusinessLicenseNo        string `json:"businessLicense"`            // 营业执照号
	BusinessLicenseStartDate string `json:"busLiceStartDate,omitempty"` // 营业执照有效起始日期
	BusinessLicenseEndDate   string `json:"busLiceEndDate,omitempty"`   // 营业执照有效终止日期,1111-11-11代表长期有效
	PostalAddress            string `json:"postalAddress,omitempty"`    // 注册地址
}

// MerchantRegisterReq 商户入驻
type MerchantRegisterReq struct {
	InterfaceName            string `json:"interfaceName"`                      // 接口名称,固定值:register
	OrderNo                  string `json:"orderNo"`                            // 商户入驻请求订单号,服务商下唯一
	MerchantName             string `json:"merchantName"`                       // 商户名称,企业、个体工商户、事业单位、其他组织:营业执照上的工商注册名称;个人小微商户:商户_XXX(法人名称)
	MerchantShowName         string `json:"merchantShowName"`                   // 商户简称,用于支付收银台的展示名,进件测试的可在简称中添加“测试”字样;注意:简称长度不能小于2大于20
	MerchantType             string `json:"merchantType"`                       // 商户类型,见常量 constants.MerchantType,个人、企业、个体商户交易限额不同,切勿传错
	OrgNum                   string `json:"orgNum"`                             // 组织机构代码,1.若是个人商户,统一填写身份证号;2.三证合一,则统一填营业证件号
	BizCategory              string `json:"businessCategory"`                   // 经营类别,见常量 constants.BusinessCategory
	BizCategoryId            string `json:"categoryId"`                         // 商户经营子类id,具体值见附件
	MicroBizType             string `json:"microBizType,omitempty"`             // 小微经营类型,见常量 constants.MicroBizType,小微个人商户必填
	MccCode                  string `json:"mcc,omitempty"`                      // 银联MCC码
	BizAddress               string `json:"address"`                            // 经营地址
	BizDistrictCode          string `json:"districtCode"`                       // 经营区县编码
	ServicePhone             string `json:"servicePhone,omitempty"`             // 用户支付后,可通过此号码进行咨询,如不填会送联系人电话。
	BindMobile               string `json:"bindMobile,omitempty"`               // 绑定手机号
	BusinessLicenseType      string `json:"businessLicenseType,omitempty"`      // 营业证书类型,见常量 constants.CertType,默认为营业执照
	BusinessLicenseNo        string `json:"businessLicense,omitempty"`          // 营业执照号
	BusinessLicenseStartDate string `json:"businessStartDate,omitempty"`        // 营业执照起始日期,yyyyMMdd,非个人商户必传
	BusinessLicenseEndDate   string `json:"businessEndDate,omitempty"`          // 营业执照结束日期,yyyyMMdd或者长期有效,非个人商户必传
	LegalPhone               string `json:"legalPersonPhone"`                   // 法人手机号已被使用时法人名称必须相同，法人手机号传过的情况下，后续重复使用进件时法人证件、名称要与之前入网的一致
	LegalName                string `json:"legalPerson"`                        // 法人名字
	LegalIdType              string `json:"idType"`                             // 法人证件类型,见常量 constants.IdType
	LegalIdNo                string `json:"legalPersonID"`                      // 法人证件号
	LegalIdAddress           string `json:"legalPersonIdAddress,omitempty"`     // 法人证件居住地址,企业商户必传
	LegalIdStartDate         string `json:"idCardStartDate,omitempty"`          // 法人身份证开始日期,yyyyMMdd,开通微信产品必传
	LegalIdEndDate           string `json:"idCardEndDate,omitempty"`            // 法人身份证结束日期,yyyyMMdd或者长期有效,开通微信产品必传
	ContactEmail             string `json:"contactEmail"`                       // 各子商户邮箱地址唯一（可不用于接收邮件）
	ContactPhone             string `json:"contactPhone"`                       // 联系电话
	ContactName              string `json:"contact"`                            // 联系人名称
	ContactType              string `json:"contactType"`                        // 联系人类型,见常量 constants.ContactType,LEGAL:经营者/法人,SUPER:经办人
	ContactIdType            string `json:"contactIdType,omitempty"`            // 联系人证件类型,见常量 constants.IdType
	ContactIdNo              string `json:"contactId,omitempty"`                // 联系人证件号,开通微信产品必填
	ContactIdStartDate       string `json:"contactIdCardStartDate,omitempty"`   // 联系人证件有效期开始时间,yyyyMMdd,联系人类型为经办人时必填
	ContactIdEndDate         string `json:"contactIdCardEndDate,omitempty"`     // 联系人证件有效期结束时间,yyyyMMdd或长期有效,联系人类型为经办人时必填
	SettlementCycle          string `json:"settlementCycle"`                    // 结算周期,constants.SettlementCycle;T1:工作日隔天结算,D1:自然日隔天结算,D0:当日结算(目前不支持)
	SettlementMode           string `json:"settlementMode"`                     // 结算方式,constants.SettlementMode;NOTOPEN:不开通;AUTO:自动结算;SELF:自主结算;
	SettlementPostscript     string `json:"settlementPostscript,omitempty"`     // 结算附言,自动结算/自主结算所使用的备注信息
	SettleMode               string `json:"settleMode"`                         // 结算模式,constants.SettleMode,MERCHANT:按商户结算;MERGE:按结算人结算,同平台商下,结算卡号一致的结算单合并成一笔出款
	SettlementPhone          string `json:"settlementPhone,omitempty"`          // 结算人手机号
	ArrivalMode              string `json:"arrivalMode,omitempty"`              // 结算到账方式,constants.ArrivalMode;NORMAL:银行卡,ELECTRONIC:银行电子账户
	CardIdCard               string `json:"cardIdCard,omitempty"`               // 开户人身份证(非法人结算时传递)
	CardNo                   string `json:"cardNo"`                             // 结算银行卡账号
	CardName                 string `json:"cardName"`                           // 结算银行卡开户名
	CardType                 string `json:"cardType"`                           // 结算卡类型,constants.SettleBankType;TOPRIVATE:对私,TOPUBLIC:对公
	BankBranchCode           string `json:"bankBranchCode"`                     // 结算账户开户行联行号,注:需商户电话联系银行核实具体联行号
	ElectronicAccountNo      string `json:"electronicAccountNo,omitempty"`      // 电子账户
	ElectronicAccountName    string `json:"electronicAccountName,omitempty"`    // 电子账户户名
	ElectronicSettleBankType string `json:"electronicSettleBankType,omitempty"` // 电子账户类型,constants.SettleBankType;TOPRIVATE:对私,TOPUBLIC:对公
	UboType                  bool   `json:"uboType"`                            // 经营者/法人是否为受益人,企业商户必传
	UboName                  string `json:"uboIdName,omitempty"`                // 受益人证件姓名,受益人非经营者/法人类型时必传
	UboIdType                string `json:"uboIdType,omitempty"`                // 受益人证件类型,受益人非经营者/法人类型时必传
	UboIdNo                  string `json:"uboId,omitempty"`                    // 受益人证件号码,受益人非经营者/法人类型时必传
	UboIdAddress             string `json:"uboAddress,omitempty"`               // 受益人证件居住地址,受益人非经营者/法人类型时必传
	UboIdStartDate           string `json:"uboIdStartDate,omitempty"`           // 受益人证件有效期开始时间,yyyyMMdd,受益人非经营者/法人类型时必传
	UboIdEndDate             string `json:"uboIdEndDate,omitempty"`             // 受益人证件有效期结束时间,yyMMdd或长期有效,受益人非经营者/法人类型时必传
	Longitude                string `json:"longitude,omitempty"`                // 经度,如:113.292482
	Latitude                 string `json:"latitude,omitempty"`                 // 纬度,如:23.201091
	NeedAuth                 bool   `json:"needAuth,omitempty"`                 // 是否需要接口提交商家注册申请单,true:系统自动去微信、支付宝提交商家注册申请单
	SpecialSignName          bool   `json:"specialSignName,omitempty"`          // 是否需要格式化商户名称,根据微信定义的签约名规则去报备,只有营业执照名字为“*”，或者个人商户才传true
	NotifyUrl                string `json:"notifyUrl,omitempty"`                // 异步通知地址,人工审核后异步回调地址(返回参数为JSON数据)注:只有人工审核完成的会异步通知;自动审核通过,不会通知
	AuthorizationKey         bool   `json:"authorizationKey"`                   // 是否统一使用平台商秘钥,true代表授权,false代表不授权,注:平台密钥,通常默认为true
	WebSite                  string `json:"webSite,omitempty"`                  // 商户交易网站地址,与“其他产品信息”二选一必填
	OtherPdInfo              string `json:"otherPdInfo,omitempty"`              // 其他产品信息,可输入公众号/小程序/app名称,与“商户交易网站地址”二选一必填
	AccessUrl                string `json:"accessUrl,omitempty"`                // 接入地址,空
	IcpFiling                string `json:"icpFiling"`                          // ICP备案号，平台商类型为服务商时必填，平台商类型为SaaS系统商的子商户默认读取平台商的 ICP 备案号

	// Deprecated: 请勿直接赋值,应调用ServiceCodes添加
	ServiceCodesJson string   `json:"serviceCodes,omitempty"` // 支付宝商户申请服务类型,如:"[\"F2F\",\"PRE_F2F\"]",constants.AlipayServiceCode
	ServiceCodes     []string `json:"-"`                      // 支付宝商户申请服务类型,如:["F2F","PRE_F2F"],constants.AlipayServiceCode
}

// MerchantRegisterRes 商户入驻
type MerchantRegisterRes struct {
	OrderNo        string `json:"orderNo"`              // 订单号 进件时的订单号
	RegisterStatus string `json:"registeStatus"`        // 进件状态 INIT:待审核,OVERRULE:申请驳回;AUDITED:审核通过
	Msg            string `json:"msg,omitempty"`        // 进件状态说明,RegisterStatus为INIT时返回
	MerchantId     string `json:"merchantId,omitempty"` // 子商户商编,只有进件状态为审核通过才返回
}

// MerchantRegisterQueryReq 商户入驻查询
type MerchantRegisterQueryReq struct {
	InterfaceName string `json:"interfaceName"` // 接口名称,固定值:registeredQuery
	OrderNo       string `json:"orderNo"`       // 订单号,要求50字符以内,同一商户号下订单号唯一
}

// MerchantRegisterQueryRes 商户入驻查询
type MerchantRegisterQueryRes struct {
	InterfaceName  string `json:"interfaceName"`        // 接口名称 registeredQuery
	OrderNo        string `json:"orderNo"`              // 订单号,商户入驻时的请求订单号
	RegisterStatus string `json:"registeStatus"`        // 入网状态(INIT:待审核,OVERRULE:申请驳回,AUDITED:审核通过)
	Msg            string `json:"msg,omitempty"`        // 状态说明, RegisterStatus 为INIT、AUDITED时返回
	MerchantId     string `json:"merchantId,omitempty"` // 子商户商编 只有进件状态为审核通过才返回
	MerchantName   string `json:"merchantName"`         // 商户名称
}

// MerchantRegisterNotifyReqBody 商户入驻结果异步通知Body
type MerchantRegisterNotifyReqBody struct {
	OrderNo        string `json:"orderNo"`              // 订单号 进件时的订单号
	RegisterStatus string `json:"registeStatus"`        // 进件状态 INIT:待审核;OVERRULE:申请驳回;AUDITED:审核通过
	Msg            string `json:"msg,omitempty"`        // 进件状态说明 当RegisterStatus为 INIT,AUDITED时返回
	MerchantId     string `json:"merchantId,omitempty"` // 子商户商编 只有进件状态为审核通过才返回
}
type MerchantRegisterNotifyReq = NotifyReq[MerchantRegisterNotifyReqBody]

// MerchantSignContractNotifyReqBody 商户签章异步通知Body
type MerchantSignContractNotifyReqBody struct {
	MerchantId string `json:"merchantId"`           // 子商户商编 只有进件状态为审核通过才返回
	SignStatus string `json:"signStatus,omitempty"` // 签章状态,SIGNED:已签约
}
type MerchantSignContractNotifyReq = NotifyReq[MerchantSignContractNotifyReqBody]

// MerchantStatusAlterNotifyReqBody 商户状态变更异步通知Body
type MerchantStatusAlterNotifyReqBody struct {
	NotifyId    string `json:"notifyId"`             // 订单号 进件时的订单号
	MerchantId  string `json:"merchantId,omitempty"` // 子商户商编 只有进件状态为审核通过才返回
	StatusType  string `json:"statusType"`           // 状态类型,商户状态:MERCHANT_STATUS;账户状态:ACCOUNT_STATUS
	Status      string `json:"status"`               // 状态变更类型,商户状态:constants.MerchantStatus,账户状态:constants.AccountStatus
	Reason      string `json:"reason,omitempty"`     // 变更原因
	OperateTime string `json:"operateTime"`          // 操作时间,如:2022-09-08 17:36:24
}
type MerchantStatusAlterNotifyReq = NotifyReq[MerchantStatusAlterNotifyReqBody]
