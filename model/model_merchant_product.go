package model

// MerchantAppPaySettingReq 商户扫码产品开通
type MerchantAppPaySettingReq struct {
	InterfaceName  string                   `json:"interfaceName"`            // 接口名称,固定值:productSettings
	MerchantId     string                   `json:"subMerchantId"`            // 子商户编号,进件审核通过后才有的商户号
	ProductType    string                   `json:"productType"`              // 产品类型:固定值:APPPAY:扫码产品
	PaymentType    string                   `json:"paymentType"`              // 支付客户端类型,constants.PaymentType
	PaymentMethods string                   `json:"paymentMethods,omitempty"` // 支付方式,constants.PaymentMethods,传ALL根据服务商开通的支付类型配置所有类型的费率,不传是修改当前商户已开通的所有类型费率
	FeeMode        string                   `json:"feeMode,omitempty"`        // 费率模式,constants.FeeMode,DEFAULT:默认模式,按百分比;RANGE:区间模式,区间固定金额
	FeeType        string                   `json:"feeType,omitempty"`        // 费率类型,constants.FeeType
	CardType       string                   `json:"cardType,omitempty"`       // 卡类型,constants.CardType
	Fee            float64                  `json:"fee,omitempty"`            // 费率,单位(%或元),小数后两位
	MinFee         float64                  `json:"minFee,omitempty"`         // 最低费率金额,单位(元),小数后两位
	MaxFee         float64                  `json:"maxFee"`                   // 最高费率金额,单位(元),小数后两位
	FeeRanges      []MerchantAppPayFeeRange `json:"feeRanges,omitempty"`      // 分段费率,需要与服务商区间数目保持一致,区间上下限都包含为双闭区间[0,1000]
}

// MerchantAppPaySettingRes 商户扫码产品开通
type MerchantAppPaySettingRes struct {
	MerchantId     string                   `json:"subMerchantId"`            // 子商户编号
	PaymentType    string                   `json:"paymentType"`              // 支付客户端类型,constants.PaymentType
	PaymentMethods string                   `json:"paymentMethods,omitempty"` // 支付方式,constants.PaymentMethods,传ALL根据服务商开通的支付类型配置所有类型的费率,不传是修改当前商户已开通的所有类型费率
	Fee            float64                  `json:"fee,omitempty"`            // 费率,单位(%或元),小数后两位,注:分段费率返回值默认以%为单位
	MinFee         float64                  `json:"minFee,omitempty"`         // 最低费率金额,单位(元),小数后两位
	MaxFee         float64                  `json:"maxFee"`                   // 最高费率金额,单位(元),小数后两位
	FeeRanges      []MerchantAppPayFeeRange `json:"feeRanges,omitempty"`      // 银联扫码/wap分段费率
}

type MerchantAppPayFeeRange struct {
	Start    float64 `json:"startAmount"` // 开始费率区间值,闭区间
	End      float64 `json:"endAmount"`   // 结束费率区间值,闭区间
	CalcType string  `json:"calcType"`    // 计算类型;constants.CalcType,SINGLE:单笔收费;RATIO:比率收费
	CardType string  `json:"cardType"`    // 费率类型,constants.CardType
	Fee      float64 `json:"fee"`         // 费率值,0.50代表0.5元或0.5%(单位,见计算类型 CalcType)
	MinFee   float64 `json:"minFee"`      // 最低费率值,单位(元)
	MaxFee   float64 `json:"maxFee"`      // 最高费率金额,单位(元),小数后两位
}

// MerchantAppPayQueryReq 商户扫码产品查询
type MerchantAppPayQueryReq struct {
	InterfaceName  string `json:"interfaceName"`            // 接口名称,固定值:productQuery
	MerchantId     string `json:"subMerchantId"`            // 子商户编号,进件审核通过后才有的商户号
	ProductType    string `json:"productType"`              // 产品类型:固定值:APPPAY:扫码产品
	PaymentType    string `json:"paymentType"`              // 支付客户端类型,constants.PaymentType
	PaymentMethods string `json:"paymentMethods,omitempty"` // 支付方式,constants.PaymentMethods,传ALL根据服务商开通的支付类型配置所有类型的费率,不传是修改当前商户已开通的所有类型费率
	FeeMode        string `json:"feeMode,omitempty"`        // 费率模式,constants.FeeMode,DEFAULT:默认模式,按百分比;RANGE:区间模式,区间固定金额
	FeeType        string `json:"feeType,omitempty"`        // 费率类型,constants.FeeType
	CardType       string `json:"cardType,omitempty"`       // 卡类型,constants.CardType
}

// MerchantAppPayQueryRes 商户扫码产品查询
type MerchantAppPayQueryRes struct {
	MerchantId     string                   `json:"subMerchantId"`            // 子商户编号
	PaymentType    string                   `json:"paymentType"`              // 支付客户端类型,constants.PaymentType
	PaymentMethods string                   `json:"paymentMethods,omitempty"` // 支付方式,constants.PaymentMethods,传ALL根据服务商开通的支付类型配置所有类型的费率,不传是修改当前商户已开通的所有类型费率
	Fee            float64                  `json:"fee,omitempty"`            // 费率,单位(%或元),小数后两位
	MinFee         float64                  `json:"minFee,omitempty"`         // 最低费率金额,单位(元),小数后两位
	MaxFee         float64                  `json:"maxFee"`                   // 最高费率金额,单位(元),小数后两位
	FeeRanges      []MerchantAppPayFeeRange `json:"feeRanges,omitempty"`      // 分段费率,需要与服务商区间数目保持一致,区间上下限都包含为双闭区间[0,1000]
}

// MerchantSettlementSettingReq 商户开通结算产品
type MerchantSettlementSettingReq struct {
	InterfaceName   string  `json:"interfaceName"`      // 接口名称,固定值:productSettings
	MerchantId      string  `json:"subMerchantId"`      // 子商户编号,进件审核通过后才有的商户号
	ProductType     string  `json:"productType"`        // 产品类型:固定值:SETTLEMENT:结算产品
	SettlementCycle string  `json:"settlementCycle"`    // 资金周期:T1,D0,D1,constants.SettlementCycle
	CardType        string  `json:"cardType"`           // 账户属性:对公结算(TOPUBLIC)或对私结算(TOPRIVATE),constants.SettleBankType
	CalcType        string  `json:"calcType"`           // 计算类型;constants.CalcType,SINGLE:单笔收费;RATIO:比率收费
	Fee             float64 `json:"fee,omitempty"`      // 费率,0.50代表0.5元或0.5%(单位,见计算类型 CalcType)
	Floating        float64 `json:"floating,omitempty"` // 加收费率:百分比(%),小数后两位(非必填)
	MinFee          float64 `json:"minFee,omitempty"`   // 最低费率金额:单位(元),小数后两位(非必填)
}

// MerchantSettlementSettingRes 商户开通结算产品
type MerchantSettlementSettingRes struct {
	MerchantId      string  `json:"subMerchantId"`      // 子商户编号,进件审核通过后才有的商户号
	SettlementCycle string  `json:"settlementCycle"`    // 资金周期:T1,D0,D1,constants.SettlementCycle
	CardType        string  `json:"cardType"`           // 账户属性:对公结算(TOPUBLIC)或对私结算(TOPRIVATE),constants.SettleBankType
	CalcType        string  `json:"calcType"`           // 计算类型;constants.CalcType,SINGLE:单笔收费;RATIO:比率收费
	Fee             float64 `json:"fee,omitempty"`      // 费率,0.50代表0.5元或0.5%(单位,见计算类型 CalcType)
	Floating        float64 `json:"floating,omitempty"` // 加收费率:百分比(%),小数后两位(非必填)
	MinFee          float64 `json:"minFee,omitempty"`   // 最低费率金额:单位(元),小数后两位(非必填)
}

// MerchantSettlementProductQueryReq 商户结算产品查询
type MerchantSettlementProductQueryReq struct {
	InterfaceName   string `json:"interfaceName"`   // 接口名称,固定值:productQuery
	MerchantId      string `json:"subMerchantId"`   // 子商户编号:进件审核通过后才有的商户号
	ProductType     string `json:"productType"`     // 产品类型:固定值:SETTLEMENT:结算产品
	SettlementCycle string `json:"settlementCycle"` // 资金周期:T1,D0,D1,constants.SettlementCycle
	CardType        string `json:"cardType"`        // 账户属性:对公结算(TOPUBLIC)或对私结算(TOPRIVATE),constants.SettleBankType
	CalcType        string `json:"calcType"`        // 计算类型;constants.CalcType,SINGLE:单笔收费;RATIO:比率收费
}

// MerchantSettlementProductQueryRes 商户结算产品查询
type MerchantSettlementProductQueryRes struct {
	MerchantId      string  `json:"subMerchantId"`      // 子商户编号
	SettlementCycle string  `json:"settlementCycle"`    // 资金周期:T1,D0,D1,constants.SettlementCycle
	CardType        string  `json:"cardType"`           // 账户属性:对公结算(TOPUBLIC)或对私结算(TOPRIVATE),constants.SettleBankType
	CalcType        string  `json:"calcType"`           // 计算类型;constants.CalcType,SINGLE:单笔收费;RATIO:比率收费
	Fee             float64 `json:"fee,omitempty"`      // 费率,0.50代表0.5元或0.5%(单位,见计算类型 CalcType)
	Floating        float64 `json:"floating,omitempty"` // 加收费率:百分比(%),小数后两位(非必填)
	MinFee          float64 `json:"minFee,omitempty"`   // 最低费率金额:单位(元),小数后两位(非必填)
}

// MerchantTransferSettingReq 商户开通代付产品
type MerchantTransferSettingReq struct {
	InterfaceName   string  `json:"interfaceName"`      // 接口名称,固定值:productSettings
	MerchantId      string  `json:"subMerchantId"`      // 子商户编号,进件审核通过后才有的商户号
	ProductType     string  `json:"productType"`        // 产品类型:固定值:TRANSFER:代付产品
	SettlementCycle string  `json:"settlementCycle"`    // 资金周期:T1,D0,D1,constants.SettlementCycle
	CardType        string  `json:"cardType"`           // 账户属性:constants.OnlineBankType,ALL,对公(B2B)或对私结算(B2C)
	CalcType        string  `json:"calcType"`           // 计算类型;constants.CalcType,SINGLE:单笔收费;RATIO:比率收费
	Fee             float64 `json:"fee,omitempty"`      // 费率,0.50代表0.5元或0.5%(单位,见计算类型 CalcType)
	Floating        float64 `json:"floating,omitempty"` // 加收费率:百分比(%),小数后两位(非必填)
	MinFee          float64 `json:"minFee,omitempty"`   // 最低费率金额:单位(元),小数后两位(非必填)
}

// MerchantTransferSettingRes 商户开通代付产品
type MerchantTransferSettingRes struct {
	MerchantId      string  `json:"subMerchantId"`      // 子商户编号,进件审核通过后才有的商户号
	SettlementCycle string  `json:"settlementCycle"`    // 资金周期:T1,D0,D1,constants.SettlementCycle
	CardType        string  `json:"cardType"`           // 账户属性:constants.OnlineBankType,ALL,对公(B2B)或对私结算(B2C)
	CalcType        string  `json:"calcType"`           // 计算类型;constants.CalcType,SINGLE:单笔收费;RATIO:比率收费
	Fee             float64 `json:"fee,omitempty"`      // 费率,0.50代表0.5元或0.5%(单位,见计算类型 CalcType)
	Floating        float64 `json:"floating,omitempty"` // 加收费率:百分比(%),小数后两位(非必填)
	MinFee          float64 `json:"minFee,omitempty"`   // 最低费率金额:单位(元),小数后两位(非必填)
}

// MerchantTransferQueryReq 商户代付产品查询
type MerchantTransferQueryReq struct {
	InterfaceName   string `json:"interfaceName"`   // 接口名称,固定值:productQuery
	MerchantId      string `json:"subMerchantId"`   // 子商户编号:进件审核通过后才有的商户号
	ProductType     string `json:"productType"`     // 产品类型:固定值:TRANSFER:代付产品
	SettlementCycle string `json:"settlementCycle"` // 资金周期:T1,D0,D1,constants.SettlementCycle
	CardType        string `json:"cardType"`        // 账户属性:对公结算(TOPUBLIC)或对私结算(TOPRIVATE),constants.SettleBankType
	CalcType        string `json:"calcType"`        // 计算类型;constants.CalcType,SINGLE:单笔收费;RATIO:比率收费
}

// MerchantTransferQueryRes 商户代付产品查询
type MerchantTransferQueryRes struct {
	MerchantId      string  `json:"subMerchantId"`      // 子商户编号
	SettlementCycle string  `json:"settlementCycle"`    // 资金周期:T1,D0,D1,constants.SettlementCycle
	CardType        string  `json:"cardType"`           // 账户属性:对公结算(TOPUBLIC)或对私结算(TOPRIVATE),constants.SettleBankType
	CalcType        string  `json:"calcType"`           // 计算类型;constants.CalcType,SINGLE:单笔收费;RATIO:比率收费
	Fee             float64 `json:"fee,omitempty"`      // 费率,0.50代表0.5元或0.5%(单位,见计算类型 CalcType)
	Floating        float64 `json:"floating,omitempty"` // 加收费率:百分比(%),小数后两位(非必填)
	MinFee          float64 `json:"minFee,omitempty"`   // 最低费率金额:单位(元),小数后两位(非必填)
}

// MerchantQuickPaySettingReq 商户开通快捷产品
type MerchantQuickPaySettingReq struct {
	InterfaceName string  `json:"interfaceName"`          // 接口名称,固定值:productSettings
	MerchantId    string  `json:"subMerchantId"`          // 子商户编号,进件审核通过后才有的商户号
	ProductType   string  `json:"productType"`            // 产品类型:固定值:QUICKPAY:快捷产品
	CardType      string  `json:"cardType"`               // 账户属性:DEBIT:借记卡,CREDIT:贷记卡,constants.CardType
	IntegralType  string  `json:"integralType,omitempty"` // 积分计费: CardType 为 CREDIT 必填,constants.IntegralType
	StartAmount   float64 `json:"startAmount"`            // 费率区间开始金额
	EndAmount     float64 `json:"endAmount"`              // 费率区间结束金额
	CalcType      string  `json:"calcType"`               // 计算类型;constants.CalcType,SINGLE:单笔收费;RATIO:比率收费
	Fee           float64 `json:"fee,omitempty"`          // 费率,0.50代表0.5元或0.5%(单位,见计算类型 CalcType)
	MinFee        float64 `json:"minFee,omitempty"`       // 最低费率金额:单位(元),小数后两位(非必填)
	MaxFee        float64 `json:"maxFee,omitempty"`       // 最高费率金额,单位(元),小数后两位
}

// MerchantQuickPaySettingRes 商户开通快捷产品
type MerchantQuickPaySettingRes struct {
	MerchantId   string  `json:"subMerchantId"`          // 子商户编号,进件审核通过后才有的商户号
	CardType     string  `json:"cardType"`               // 账户属性:DEBIT:借记卡,CREDIT:贷记卡,constants.CardType
	IntegralType string  `json:"integralType,omitempty"` // 积分计费: CardType 为 CREDIT 必填,constants.IntegralType
	StartAmount  float64 `json:"startAmount"`            // 费率区间开始金额
	EndAmount    float64 `json:"endAmount"`              // 费率区间结束金额
	CalcType     string  `json:"calcType"`               // 计算类型;constants.CalcType,SINGLE:单笔收费;RATIO:比率收费
	Fee          float64 `json:"fee,omitempty"`          // 费率,0.50代表0.5元或0.5%(单位,见计算类型 CalcType)
	MinFee       float64 `json:"minFee,omitempty"`       // 最低费率金额:单位(元),小数后两位(非必填)
	MaxFee       float64 `json:"maxFee,omitempty"`       // 最高费率金额,单位(元),小数后两位
}

// MerchantQuickPayQueryReq 商户快捷产品查询
type MerchantQuickPayQueryReq struct {
	InterfaceName string  `json:"interfaceName"`          // 接口名称,固定值:productQuery
	MerchantId    string  `json:"subMerchantId"`          // 子商户编号:进件审核通过后才有的商户号
	ProductType   string  `json:"productType"`            // 产品类型:固定值:QUICKPAY:快捷产品
	CardType      string  `json:"cardType"`               // 账户属性:对公结算(TOPUBLIC)或对私结算(TOPRIVATE),constants.SettleBankType
	IntegralType  string  `json:"integralType,omitempty"` // 积分计费: CardType 为 CREDIT 必填,constants.IntegralType
	StartAmount   float64 `json:"startAmount"`            // 费率区间开始金额
	EndAmount     float64 `json:"endAmount"`              // 费率区间结束金额
	CalcType      string  `json:"calcType"`               // 计算类型;constants.CalcType,SINGLE:单笔收费;RATIO:比率收费
}

// MerchantQuickPayQueryRes 商户快捷产品查询
type MerchantQuickPayQueryRes struct {
	MerchantId   string  `json:"subMerchantId"`          // 子商户编号
	CardType     string  `json:"cardType"`               // 账户属性:DEBIT:借记卡,CREDIT:贷记卡,constants.CardType
	IntegralType string  `json:"integralType,omitempty"` // 积分计费: CardType 为 CREDIT 必填,constants.IntegralType
	StartAmount  float64 `json:"startAmount"`            // 费率区间开始金额
	EndAmount    float64 `json:"endAmount"`              // 费率区间结束金额
	CalcType     string  `json:"calcType"`               // 计算类型;constants.CalcType,SINGLE:单笔收费;RATIO:比率收费
	Fee          float64 `json:"fee,omitempty"`          // 费率,0.50代表0.5元或0.5%(单位,见计算类型 CalcType)
	MinFee       float64 `json:"minFee,omitempty"`       // 最低费率金额:单位(元),小数后两位(非必填)
	MaxFee       float64 `json:"maxFee,omitempty"`       // 最高费率金额,单位(元),小数后两位
}

// MerchantAccountPaySettingReq 商户开通虚拟账户支付产品
type MerchantAccountPaySettingReq struct {
	InterfaceName string  `json:"interfaceName"`    // 接口名称,固定值:productSettings
	MerchantId    string  `json:"subMerchantId"`    // 子商户编号,进件审核通过后才有的商户号
	ProductType   string  `json:"productType"`      // 产品类型:固定值:ACCOUNT_PAY:虚拟账户支付产品
	CalcType      string  `json:"calcType"`         // 计算类型;SINGLE:单笔收费;RATIO:比率收费
	Fee           float64 `json:"fee,omitempty"`    // 费率:单位(%或元),小数后两位(非必填)
	MinFee        float64 `json:"minFee,omitempty"` // 最低费率金额:单位(元),小数后两位(非必填)
}

// MerchantAccountPaySettingRes 商户开通虚拟账户支付产品
type MerchantAccountPaySettingRes struct {
	MerchantId string  `json:"subMerchantId"`    // 子商户编号,进件审核通过后才有的商户号
	CalcType   string  `json:"calcType"`         // 计算类型;SINGLE:单笔收费;RATIO:比率收费
	Fee        float64 `json:"fee,omitempty"`    // 费率:单位(%或元),小数后两位(非必填)
	MinFee     float64 `json:"minFee,omitempty"` // 最低费率金额:单位(元),小数后两位(非必填)
}

// MerchantAccountPayQueryReq 商户虚拟账户支付产品查询
type MerchantAccountPayQueryReq struct {
	InterfaceName string `json:"interfaceName"` // 接口名称,固定值:productSettings
	MerchantId    string `json:"subMerchantId"` // 子商户编号,进件审核通过后才有的商户号
	ProductType   string `json:"productType"`   // 产品类型:固定值:ACCOUNT_PAY:虚拟账户支付产品
	CalcType      string `json:"calcType"`      // 计算类型;SINGLE:单笔收费;RATIO:比率收费
}

// MerchantAccountPayQueryRes 商户虚拟账户支付产品查询
type MerchantAccountPayQueryRes struct {
	MerchantId string  `json:"subMerchantId"`    // 子商户编号,进件审核通过后才有的商户号
	CalcType   string  `json:"calcType"`         // 计算类型;SINGLE:单笔收费;RATIO:比率收费
	Fee        float64 `json:"fee,omitempty"`    // 费率:单位(%或元),小数后两位(非必填)
	MinFee     float64 `json:"minFee,omitempty"` // 最低费率金额:单位(元),小数后两位(非必填)
}

// MerchantTransferDepositSettingReq 商户开通转账充值产品
type MerchantTransferDepositSettingReq struct {
	InterfaceName string  `json:"interfaceName"`    // 接口名称,固定值:productSettings
	MerchantId    string  `json:"subMerchantId"`    // 子商户编号,进件审核通过后才有的商户号
	ProductType   string  `json:"productType"`      // 产品类型:固定值:TRANSFERDEPOSIT:转账充值产品
	CalcType      string  `json:"calcType"`         // 计算类型;SINGLE:单笔收费;RATIO:比率收费
	Fee           float64 `json:"fee,omitempty"`    // 费率:单位(%或元),小数后两位(非必填)
	MinFee        float64 `json:"minFee,omitempty"` // 最低费率金额:单位(元),小数后两位(非必填)
}

// MerchantTransferDepositSettingRes 商户开通转账充值产品
type MerchantTransferDepositSettingRes struct {
	MerchantId string  `json:"subMerchantId"`    // 子商户编号,进件审核通过后才有的商户号
	CalcType   string  `json:"calcType"`         // 计算类型;SINGLE:单笔收费;RATIO:比率收费
	Fee        float64 `json:"fee,omitempty"`    // 费率:单位(%或元),小数后两位(非必填)
	MinFee     float64 `json:"minFee,omitempty"` // 最低费率金额:单位(元),小数后两位(非必填)
}

// MerchantTransferDepositQueryReq 商户转账充值产品查询
type MerchantTransferDepositQueryReq struct {
	InterfaceName string `json:"interfaceName"` // 接口名称,固定值:productSettings
	MerchantId    string `json:"subMerchantId"` // 子商户编号,进件审核通过后才有的商户号
	ProductType   string `json:"productType"`   // 产品类型:固定值:TRANSFERDEPOSIT:转账充值产品
	CalcType      string `json:"calcType"`      // 计算类型;SINGLE:单笔收费;RATIO:比率收费
}

// MerchantTransferDepositQueryRes 商户转账充值产品查询
type MerchantTransferDepositQueryRes struct {
	MerchantId string  `json:"subMerchantId"`    // 子商户编号,进件审核通过后才有的商户号
	CalcType   string  `json:"calcType"`         // 计算类型;SINGLE:单笔收费;RATIO:比率收费
	Fee        float64 `json:"fee,omitempty"`    // 费率:单位(%或元),小数后两位(非必填)
	MinFee     float64 `json:"minFee,omitempty"` // 最低费率金额:单位(元),小数后两位(非必填)
}

// MerchantModifyFeeConfigReq 商户产品手续费配置修改
type MerchantModifyFeeConfigReq struct {
	InterfaceName string `json:"interfaceName"` // 接口名称,固定值:configModifi
	MerchantId    string `json:"subMerchantId"` // 子商户编号,进件审核通过后才有的商户号
	ProductType   string `json:"productType"`   // 产品类型:固定值:APPPAY:扫码,SETTLEMENT:结算
	ModifyType    string `json:"modifyType"`    // 变更类型,固定值:FeeCollection
	CollectMethod string `json:"collectMethod"` // 收取方式,constants.CollectMethod
}

// MerchantModifyFeeConfigRes 商户产品手续费配置修改
type MerchantModifyFeeConfigRes struct {
	MerchantId   string `json:"subMerchantId"` // 子商户编号,进件审核通过后才有的商户号
	ModifyStatus string `json:"modifyStatus"`  // 状态:SUCCESS:成功,FAIL:失败
}

// MerchantQueryFeeConfigReq 商户产品手续费配置查询
type MerchantQueryFeeConfigReq struct {
	InterfaceName string `json:"interfaceName"` // 接口名称,固定值:configModifi
	MerchantId    string `json:"subMerchantId"` // 子商户编号,进件审核通过后才有的商户号
	ModifyType    string `json:"modifyType"`    // 查询类型,固定值:固定值:FeeCollection
	ProductType   string `json:"productType"`   // 产品类型:APPPAY:扫码,SETTLEMENT:结算
}

// MerchantQueryFeeConfigRes 商户产品手续费配置查询
type MerchantQueryFeeConfigRes struct {
	MerchantId    string `json:"subMerchantId"` // 子商户编号,进件审核通过后才有的商户号
	CollectMethod string `json:"collectMethod"` // 收取方式
}
