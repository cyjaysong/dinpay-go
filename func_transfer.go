package dinpay

import (
	"errors"
	"github.com/bytedance/sonic"
)

// TransferOrder 商户代付下单
func (t *Client) TransferOrder(reqBody TransferOrderReq) (res *BaseRes[TransferOrderRes], err error) {
	const path = "/trx/api/transfer"
	reqBody.InterfaceName, reqBody.Urgency = "Transfer", true
	var baseRes *BaseRes[string]
	if baseRes, err = t.transferPost(reqBody.MerchantId, path, reqBody); err != nil {
		return
	}
	return ParseRes[TransferOrderRes](baseRes)
}

// TransferQuery 商户代付订单查询
func (t *Client) TransferQuery(reqBody TransferQueryReq) (res *BaseRes[TransferQueryRes], err error) {
	const path = "/trx/api/transferQuery"
	reqBody.InterfaceName = "TransferQuery"
	var baseRes *BaseRes[string]
	if baseRes, err = t.transferPost(reqBody.MerchantId, path, reqBody); err != nil {
		return
	}
	return ParseRes[TransferQueryRes](baseRes)
}

// TransferResultNotifyVerify 商户代付结果异步通知验签
func (t *Client) TransferResultNotifyVerify(httpBody []byte) (notifyReq *TransferNotifyReq, err error) {
	notifyReq = new(TransferNotifyReq)
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
