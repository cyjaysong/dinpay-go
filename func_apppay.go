package dinpay

import (
	"bytes"
	"errors"
	"github.com/bytedance/sonic"
	"github.com/cyjaysong/dinpay-go/model"
	"io"
	"net/http"
)

// AppPayScanOrder 主扫:用户扫商户/被扫:商户扫用户 下单
func (t *Client) AppPayScanOrder(reqBody model.AppPayScanOrderReq) (res *model.BaseRes[model.AppPayScanOrderRes], err error) {
	const path = "/trx/api/appPay/pay"
	reqBody.InterfaceName = "AppPay"
	if reqBody.MarketingRules != nil {
		reqBody.MarketingRulesJson, _ = sonic.MarshalString(reqBody.MarketingRules)
	}
	reqBody.SplitType = ""
	if len(reqBody.SplitRules) >= 1 {
		reqBody.SplitType = "FIXED_AMOUNT"
		reqBody.SplitRulesJson, _ = sonic.MarshalString(reqBody.SplitRules)
	}
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.appPayJsonPost(reqBody.MerchantId, path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.AppPayScanOrderRes](baseRes)
}

// AppPayPublicPreOrder 公众号/JS/服务窗预下单
func (t *Client) AppPayPublicPreOrder(reqBody model.AppPayPublicPreOrderReq) (res *model.BaseRes[model.AppPayPublicPreOrderRes], err error) {
	const path = "/trx/api/appPay/payPublic"
	reqBody.InterfaceName, reqBody.PaymentMethods = "AppPayPublic", "PUBLIC"
	if reqBody.MarketingRules != nil {
		reqBody.MarketingRulesJson, _ = sonic.MarshalString(reqBody.MarketingRules)
	}
	reqBody.SplitType = ""
	if len(reqBody.SplitRules) >= 1 {
		reqBody.SplitType = "FIXED_AMOUNT"
		reqBody.SplitRulesJson, _ = sonic.MarshalString(reqBody.SplitRules)
	}
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.appPayJsonPost(reqBody.MerchantId, path, reqBody); err != nil {
		return nil, err
	}
	if res, err = model.ParseRes[model.AppPayPublicPreOrderRes](baseRes); err != nil {
		return
	} else if res.Code == "0000" && res.Data.PayInfoJson != "" {
		res.Data.PayInfo = &model.AppPayPublicPreOrderPayInfo{}
		err = sonic.UnmarshalString(res.Data.PayInfoJson, res.Data.PayInfo)
	}
	return
}

// AppPayWapPreOrder WAP(H5)预下单
func (t *Client) AppPayWapPreOrder(reqBody model.AppPayWapPreOrderReq) (res *model.BaseRes[model.AppPayWapPreOrderRes], err error) {
	const path = "/trx/api/appPay/payH5"
	reqBody.InterfaceName, reqBody.PaymentMethods = "AppPayH5WFT", "WAP"
	if reqBody.MarketingRules != nil {
		reqBody.MarketingRulesJson, _ = sonic.MarshalString(reqBody.MarketingRules)
	}
	reqBody.SplitType = ""
	if len(reqBody.SplitRules) >= 1 {
		reqBody.SplitType = "FIXED_AMOUNT"
		reqBody.SplitRulesJson, _ = sonic.MarshalString(reqBody.SplitRules)
	}
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.appPayJsonPost(reqBody.MerchantId, path, reqBody); err != nil {
		return nil, err
	}
	return model.ParseRes[model.AppPayWapPreOrderRes](baseRes)
}

// AppPaySdkPreOrder SDK(APP)预下单接口
func (t *Client) AppPaySdkPreOrder(reqBody model.AppPaySdkPreOrderReq) (res *model.BaseRes[model.AppPaySdkPreOrderRes], err error) {
	const path = "/trx/api/appPay/paySdk"
	reqBody.InterfaceName, reqBody.PaymentMethods = "AppPaySdk", "SDK"
	if reqBody.MarketingRules != nil {
		reqBody.MarketingRulesJson, _ = sonic.MarshalString(reqBody.MarketingRules)
	}
	reqBody.SplitType = ""
	if len(reqBody.SplitRules) >= 1 {
		reqBody.SplitType = "FIXED_AMOUNT"
		reqBody.SplitRulesJson, _ = sonic.MarshalString(reqBody.SplitRules)
	}
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.appPayJsonPost(reqBody.MerchantId, path, reqBody); err != nil {
		return nil, err
	}
	return model.ParseRes[model.AppPaySdkPreOrderRes](baseRes)
}

// AppPayAppletPreOrder 小程序预下单
func (t *Client) AppPayAppletPreOrder(reqBody model.AppPayAppletPreOrderReq) (res *model.BaseRes[model.AppPayAppletPreOrderRes], err error) {
	const path = "/trx/api/appPay/payApplet"
	reqBody.InterfaceName, reqBody.PaymentMethods = "AppPayApplet", "APPLET"
	if reqBody.MarketingRules != nil {
		reqBody.MarketingRulesJson, _ = sonic.MarshalString(reqBody.MarketingRules)
	}
	reqBody.SplitType = ""
	if len(reqBody.SplitRules) >= 1 {
		reqBody.SplitType = "FIXED_AMOUNT"
		reqBody.SplitRulesJson, _ = sonic.MarshalString(reqBody.SplitRules)
	}
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.appPayJsonPost(reqBody.MerchantId, path, reqBody); err != nil {
		return nil, err
	}
	if res, err = model.ParseRes[model.AppPayAppletPreOrderRes](baseRes); err != nil {
		return
	} else if res.Code == "0000" && res.Data.PayInfoJson != "" {
		res.Data.PayInfo = &model.AppPayAppletPreOrderPayInfo{}
		err = sonic.UnmarshalString(res.Data.PayInfoJson, res.Data.PayInfo)
	}
	return
}

// AppPayOrderQuery 交易订单查询
func (t *Client) AppPayOrderQuery(reqBody model.AppPayOrderQueryReq) (res *model.BaseRes[model.AppPayOrderQueryRes], err error) {
	const path = "/trx/api/appPay/payQuery"
	reqBody.InterfaceName = "AppPayQuery"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.appPayJsonPost(reqBody.MerchantId, path, reqBody); err != nil {
		return nil, err
	}
	res, err = model.ParseRes[model.AppPayOrderQueryRes](baseRes)
	if err != nil {
		return nil, err
	}
	if len(res.Data.MarketingRulesJson) > 0 {
		if err = sonic.Unmarshal([]byte(res.Data.MarketingRulesJson), res.Data.MarketingRules); err != nil {
			return nil, err
		}
	}
	if len(res.Data.SplitRulesJson) > 0 {
		if err = sonic.Unmarshal([]byte(res.Data.SplitRulesJson), &res.Data.SplitRules); err != nil {
			return nil, err
		}
	}
	return
}

// OrderPayResultNotifyVerify 订单支付结果异步通知验签
func (t *Client) OrderPayResultNotifyVerify(httpBody []byte) (notifyReq *model.OrderPayResultNotifyReq, err error) {
	reqReq := http.Request{Method: "POST", Body: io.NopCloser(bytes.NewBuffer(httpBody)), Header: http.Header{}}
	reqReq.Header.Set(`Content-Type`, `application/x-www-form-urlencoded`)
	if err = reqReq.ParseForm(); err != nil {
		return nil, err
	}
	if !t.SM3WithSM2Verify([]byte(reqReq.Form.Get("data")), reqReq.Form.Get("sign")) {
		return nil, errors.New("响应内容验签失败")
	}
	httpBodyMap := make(map[string]string, len(reqReq.Form))
	for key, strings := range reqReq.Form {
		httpBodyMap[key] = strings[0]
	}
	httpBody, err = sonic.Marshal(httpBodyMap)
	if err != nil {
		return nil, err
	}

	var baseRes *model.NotifyReq[string]
	if err = sonic.Unmarshal(httpBody, &baseRes); err != nil {
		return nil, err
	}
	if notifyReq, err = model.ParseNotifyReq[model.OrderPayResultNotifyReqBody](baseRes); err != nil {
		return nil, err
	}
	if len(notifyReq.Data.MarketingRulesJson) > 0 {
		_ = sonic.Unmarshal([]byte(notifyReq.Data.MarketingRulesJson), &notifyReq.Data.MarketingRule)
	}
	if len(notifyReq.Data.SplitRulesJson) > 0 {
		_ = sonic.Unmarshal([]byte(notifyReq.Data.SplitRulesJson), &notifyReq.Data.SplitRules)
	}
	return
}
