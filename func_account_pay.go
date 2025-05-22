package dinpay

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/cyjaysong/dinpay-go/model"
)

// AccountPayOrder 账户支付下单
func (t *Client) AccountPayOrder(reqBody model.AccountPayOrderReq) (res *model.BaseRes[model.AccountPayOrderRes], err error) {
	orderParameterJsonBytes, err := sonic.Marshal(reqBody.OrderParameter)
	if err != nil {
		return nil, err
	}
	reqBody.OrderParameters = string(orderParameterJsonBytes)

	const path = "/trx/api/accountPay/pay"
	reqBody.InterfaceName = "accountPay"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.accountPayPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.AccountPayOrderRes](baseRes)
}

// AccountPayQuery 账户支付订单查询
func (t *Client) AccountPayQuery(reqBody model.AccountPayQueryReq) (res *model.BaseRes[model.AccountPayQueryRes], err error) {
	const path = "/trx/api/accountPay/accountPayQuery"
	reqBody.InterfaceName = "accountPayQuery"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.accountPayPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.AccountPayQueryRes](baseRes)
}

// AccountPayGuaranteeConfirm 账户支付担保确认
func (t *Client) AccountPayGuaranteeConfirm(reqBody model.AccountPayGuaranteeConfirmReq) (res *model.BaseRes[model.AccountPayGuaranteeConfirmRes], err error) {
	const path = "/trx/api/accountPay/guaranteeConfirm"
	reqBody.InterfaceName = "guaranteeConfirm"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.accountPayPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.AccountPayGuaranteeConfirmRes](baseRes)
}

// AccountPayTransferValidateCode 资金划拨类账户支付获取验证码接口
func (t *Client) AccountPayTransferValidateCode(reqBody model.AccountPayTransferValidateCodeReq) (res *model.BaseRes[model.AccountPayTransferValidateCodeRes], err error) {
	const path = "/trx/api/accountPay/accountPayValidateCode"
	reqBody.InterfaceName = "accountPayValidateCode"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.accountPayPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.AccountPayTransferValidateCodeRes](baseRes)
}

// AccountPayTransferConfirm 资金划拨类账户支付确认接口
func (t *Client) AccountPayTransferConfirm(reqBody model.AccountPayTransferConfirmReq) (res *model.BaseRes[model.AccountPayTransferConfirmRes], err error) {
	const path = "/trx/api/accountPay/accountPayConfirmPay"
	reqBody.InterfaceName = "accountPayConfirmPay"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.accountPayPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.AccountPayTransferConfirmRes](baseRes)

}

// AccountPayOrderRefund 账户支付退款接口
func (t *Client) AccountPayOrderRefund(reqBody model.AccountPayOrderRefundReq) (res *model.BaseRes[model.AccountPayOrderRefundRes], err error) {
	refundDetailJsonBytes, err := sonic.Marshal(reqBody.RefundDetails)
	if err != nil {
		return nil, err
	}
	reqBody.RefundDetail = string(refundDetailJsonBytes)

	const path = "/trx/api/accountPay/accountPayRefund"
	reqBody.InterfaceName = "accountPayRefund"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.accountPayPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.AccountPayOrderRefundRes](baseRes)
}

// AccountPayOrderRefundQuery 账户支付退款查询接口
func (t *Client) AccountPayOrderRefundQuery(reqBody model.AccountPayOrderRefundQueryReq) (res *model.BaseRes[model.AccountPayOrderRefundQueryRes], err error) {
	const path = "/trx/api/accountPay/accountPayRefundQuery"
	reqBody.InterfaceName = "accountPayRefundQuery"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.accountPayPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.AccountPayOrderRefundQueryRes](baseRes)
}

// AccountPayOrderResultNotifyVerify 账户支付回调通知验签
func (t *Client) AccountPayOrderResultNotifyVerify(httpBody []byte) (notifyReq *model.AccountPayOrderResultNotifyReq, err error) {
	var baseRes *model.NotifyReq[string]
	if err = sonic.Unmarshal(httpBody, &baseRes); err != nil {
		return nil, err
	}
	if !t.SM3WithSM2Verify([]byte(baseRes.Data), baseRes.Sign) {
		return nil, errors.New("响应内容验签失败")
	}
	if notifyReq, err = model.ParseNotifyReq[model.AccountPayOrderResultNotifyReqBody](baseRes); err != nil {
		return nil, err
	}
	return
}

// AccountPayOrderRefundNotifyVerify 账户支付退款回调通知
func (t *Client) AccountPayOrderRefundNotifyVerify(httpBody []byte) (notifyReq *model.AccountPayOrderRefundNotifyReq, err error) {
	var baseRes *model.NotifyReq[string]
	if err = sonic.Unmarshal(httpBody, &baseRes); err != nil {
		return nil, err
	}
	if !t.SM3WithSM2Verify([]byte(baseRes.Data), baseRes.Sign) {
		return nil, errors.New("响应内容验签失败")
	}
	if notifyReq, err = model.ParseNotifyReq[model.AccountPayOrderRefundNotifyReqBody](baseRes); err != nil {
		return nil, err
	}
	return
}
