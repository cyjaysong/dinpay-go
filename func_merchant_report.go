package dinpay

import (
	"github.com/bytedance/sonic"
	"github.com/cyjaysong/dinpay-go/model"
)

// MerchantApplyAppReport 商户上游报备
func (t *Client) MerchantApplyAppReport(reqBody model.MerchantApplyAppReportReq) (res *model.BaseRes[model.MerchantApplyAppReportRes], err error) {
	const path = "/trx/api/product/applyAppReport"
	reqBody.InterfaceName = "reporting"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantApplyAppReportRes](baseRes)

}

// MerchantAppApplyQuery 商户上游报备查询
func (t *Client) MerchantAppApplyQuery(reqBody model.MerchantAppApplyQueryReq) (res *model.BaseRes[model.MerchantAppApplyQueryRes], err error) {
	const path = "/trx/api/product/appApplyReportQuery"
	reqBody.InterfaceName = "reportingQuery"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantAppApplyQueryRes](baseRes)

}

// MerchantReIdentify 商户重新发起认证
func (t *Client) MerchantReIdentify(reqBody model.MerchantReIdentifyReq) (res *model.BaseRes[model.MerchantReIdentifyRes], err error) {
	const path = "/trx/api/product/reIdentify"
	reqBody.InterfaceName = "reCertification"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantReIdentifyRes](baseRes)

}

// MerchantWxPublicApply 商户微信交易配置
func (t *Client) MerchantWxPublicApply(reqBody model.MerchantWxPublicApplyReq) (res *model.BaseRes[model.MerchantWxPublicApplyRes], err error) {
	const path = "/trx/api/product/wxPublicApply"
	reqBody.InterfaceName = "wxAppidConfig"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantWxPublicApplyRes](baseRes)

}

// MerchantWxPayDirsApply 商户微信配置支付目录
func (t *Client) MerchantWxPayDirsApply(reqBody model.MerchantWxPayDirsApplyReq) (res *model.BaseRes[model.MerchantWxPayDirsApplyRes], err error) {
	const path = "/trx/api/product/weiXinAddAuthPayDirs"
	reqBody.InterfaceName = "wxAddConfig"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantWxPayDirsApplyRes](baseRes)

}

// MerchantWxPublicApplyQuery 商户微信交易配置查询
func (t *Client) MerchantWxPublicApplyQuery(reqBody model.MerchantWxPublicApplyQueryReq) (res *model.BaseRes[model.MerchantWxPublicApplyQueryRes], err error) {
	const path = "/trx/api/product/wxPublicApplyQuery"
	reqBody.InterfaceName = "wxAppidConfigQuery"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	if dataJson, err := sonic.GetFromString(baseRes.Data); err == nil {
		upstreamData := dataJson.GetByPath("upstreamData")
		upstreamDataArray, _ := dataJson.GetByPath("upstreamData").ArrayUseNode()
		for i, dataItem := range upstreamDataArray {
			appIdsStr, _ := dataItem.Get("appIds").String()
			appIdsNode, _ := sonic.GetFromString(appIdsStr)
			_, _ = dataItem.Set("appIds", appIdsNode)

			payCatalogsStr, _ := dataItem.Get("payCatalogs").String()
			payCatalogsNode, _ := sonic.GetFromString(payCatalogsStr)
			_, _ = dataItem.Set("payCatalogs", payCatalogsNode)

			receiptAppIdsStr, _ := dataItem.Get("receiptAppIds").String()
			receiptAppIdsNode, _ := sonic.GetFromString(receiptAppIdsStr)
			_, _ = dataItem.Set("receiptAppIds", receiptAppIdsNode)

			_, _ = upstreamData.SetByIndex(i, dataItem)
		}
		newDataBytes, _ := dataJson.MarshalJSON()
		baseRes.Data = string(newDataBytes)
	}
	return model.ParseRes[model.MerchantWxPublicApplyQueryRes](baseRes)

}
