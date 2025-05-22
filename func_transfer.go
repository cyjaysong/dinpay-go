package dinpay

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/cyjaysong/dinpay-go/model"
)

// TransferOrder 商户代付下单
func (t *Client) TransferOrder(reqBody model.TransferOrderReq) (res *model.BaseRes[model.TransferOrderRes], err error) {
	const path = "/trx/api/transfer"
	reqBody.InterfaceName, reqBody.Urgency = "Transfer", true
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.transferPost(reqBody.MerchantId, path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.TransferOrderRes](baseRes)
}

// TransferQuery 商户代付订单查询
func (t *Client) TransferQuery(reqBody model.TransferQueryReq) (res *model.BaseRes[model.TransferQueryRes], err error) {
	const path = "/trx/api/transferQuery"
	reqBody.InterfaceName = "TransferQuery"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.transferPost(reqBody.MerchantId, path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.TransferQueryRes](baseRes)
}

// TransferResultNotifyVerify 商户代付结果异步通知验签
func (t *Client) TransferResultNotifyVerify(httpBody []byte) (notifyReq *model.TransferNotifyReq, err error) {
	notifyReq = new(model.TransferNotifyReq)
	if err = sonic.Unmarshal(httpBody, notifyReq); err != nil {
		return nil, err
	}
	bodyNode, _ := sonic.Get(httpBody, "data")
	bodyBytes, _ := bodyNode.MarshalJSON()
	if !t.SM3WithSM2Verify(bodyBytes, notifyReq.Sign) {
		return nil, errors.New("响应内容验签失败")
	}
	return
}
