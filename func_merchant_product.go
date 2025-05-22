package dinpay

import "github.com/cyjaysong/dinpay-go/model"

// MerchantAppPaySetting 商户扫码产品开通
func (t *Client) MerchantAppPaySetting(reqBody model.MerchantAppPaySettingReq) (res *model.BaseRes[model.MerchantAppPaySettingRes], err error) {
	const path = "/trx/api/product/appPaySetting"
	reqBody.InterfaceName, reqBody.ProductType = "productSettings", "APPPAY"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantAppPaySettingRes](baseRes)
}

// MerchantAppPayQuery 商户扫码产品查询
func (t *Client) MerchantAppPayQuery(reqBody model.MerchantAppPayQueryReq) (res *model.BaseRes[model.MerchantAppPayQueryRes], err error) {
	const path = "/trx/api/product/appPaySetting"
	reqBody.InterfaceName, reqBody.ProductType = "productQuery", "APPPAY"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantAppPayQueryRes](baseRes)
}

// MerchantSettlementSetting 商户开通结算产品
func (t *Client) MerchantSettlementSetting(reqBody model.MerchantSettlementSettingReq) (res *model.BaseRes[model.MerchantSettlementSettingRes], err error) {
	const path = "/trx/api/product/settlementSetting"
	reqBody.InterfaceName, reqBody.ProductType = "productSettings", "SETTLEMENT"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantSettlementSettingRes](baseRes)
}

// MerchantSettlementProductQuery 商户结算产品查询
func (t *Client) MerchantSettlementProductQuery(reqBody model.MerchantSettlementProductQueryReq) (res *model.BaseRes[model.MerchantSettlementProductQueryRes], err error) {
	const path = "/trx/api/product/settlementQuery"
	reqBody.InterfaceName, reqBody.ProductType = "productQuery", "SETTLEMENT"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantSettlementProductQueryRes](baseRes)
}

// MerchantTransferSetting 商户开通代付产品
func (t *Client) MerchantTransferSetting(reqBody model.MerchantTransferSettingReq) (res *model.BaseRes[model.MerchantTransferSettingRes], err error) {
	const path = "/trx/api/product/transferSetting"
	reqBody.InterfaceName, reqBody.ProductType = "productSettings", "TRANSFER"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantTransferSettingRes](baseRes)
}

// MerchantTransferQuery 商户代付产品查询
func (t *Client) MerchantTransferQuery(reqBody model.MerchantTransferQueryReq) (res *model.BaseRes[model.MerchantTransferQueryRes], err error) {
	const path = "/trx/api/product/transferQuery"
	reqBody.InterfaceName, reqBody.ProductType = "productQuery", "TRANSFER"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantTransferQueryRes](baseRes)

}

// MerchantQuickPaySetting 商户开通快捷产品
func (t *Client) MerchantQuickPaySetting(reqBody model.MerchantQuickPaySettingReq) (res *model.BaseRes[model.MerchantQuickPaySettingRes], err error) {
	const path = "/trx/api/product/quickPaySetting"
	reqBody.InterfaceName, reqBody.ProductType = "productSettings", "QUICKPAY"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantQuickPaySettingRes](baseRes)

}

// MerchantQuickPayQuery 商户快捷产品查询
func (t *Client) MerchantQuickPayQuery(reqBody model.MerchantQuickPayQueryReq) (res *model.BaseRes[model.MerchantQuickPayQueryRes], err error) {
	const path = "/trx/api/product/quickPayQuery"
	reqBody.InterfaceName, reqBody.ProductType = "productQuery", "QUICKPAY"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantQuickPayQueryRes](baseRes)

}

// MerchantAccountPaySetting 商户开通虚拟账户支付产品
func (t *Client) MerchantAccountPaySetting(reqBody model.MerchantAccountPaySettingReq) (res *model.BaseRes[model.MerchantAccountPaySettingRes], err error) {
	const path = "/trx/api/product/accountPaySetting"
	reqBody.InterfaceName, reqBody.ProductType = "productSettings", "ACCOUNT_PAY"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantAccountPaySettingRes](baseRes)

}

// MerchantAccountPayQuery 商户虚拟账户支付产品查询
func (t *Client) MerchantAccountPayQuery(reqBody model.MerchantAccountPayQueryReq) (res *model.BaseRes[model.MerchantAccountPayQueryRes], err error) {
	const path = "/trx/api/product/accountPayQuery"
	reqBody.InterfaceName, reqBody.ProductType = "productQuery", "ACCOUNT_PAY"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantAccountPayQueryRes](baseRes)
}

// MerchantTransferDepositSetting 商户开通转账充值产品
func (t *Client) MerchantTransferDepositSetting(reqBody model.MerchantTransferDepositSettingReq) (res *model.BaseRes[model.MerchantTransferDepositSettingRes], err error) {
	const path = "/trx/api/product/transferDepositOpen"
	reqBody.InterfaceName, reqBody.ProductType = "productSettings", "TRANSFERDEPOSIT"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantTransferDepositSettingRes](baseRes)
}

// MerchantTransferDepositQuery 商户转账充值产品查询
func (t *Client) MerchantTransferDepositQuery(reqBody model.MerchantTransferDepositQueryReq) (res *model.BaseRes[model.MerchantTransferDepositQueryRes], err error) {
	const path = "/trx/api/product/transferDepositQuery"
	reqBody.InterfaceName, reqBody.ProductType = "productQuery", "TRANSFERDEPOSIT"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantTransferDepositQueryRes](baseRes)
}

// MerchantModifyFeeConfig 商户产品手续费配置修改
func (t *Client) MerchantModifyFeeConfig(reqBody model.MerchantModifyFeeConfigReq) (res *model.BaseRes[model.MerchantModifyFeeConfigRes], err error) {
	const path = "/trx/api/product/modifyFeeConfig"
	//configModifi 没写错,智付的历史遗留问题
	reqBody.InterfaceName, reqBody.ModifyType = "configModifi", "FeeCollection"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantModifyFeeConfigRes](baseRes)

}

// MerchantQueryFeeConfig 商户产品手续费配置查询
func (t *Client) MerchantQueryFeeConfig(reqBody model.MerchantQueryFeeConfigReq) (res *model.BaseRes[model.MerchantQueryFeeConfigRes], err error) {
	const path = "/trx/api/product/queryFeeConfig"
	reqBody.InterfaceName, reqBody.ModifyType = "configQuery", "FeeCollection"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantQueryFeeConfigRes](baseRes)
}
