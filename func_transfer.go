package dinpay

import (
	"bytes"
	"errors"
	"io"
	"net/http"

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
	reqReq := http.Request{Method: "POST", Body: io.NopCloser(bytes.NewBuffer(httpBody)), Header: http.Header{}}
	reqReq.Header.Set(`Content-Type`, `application/x-www-form-urlencoded`)
	if err = reqReq.ParseForm(); err != nil {
		return nil, err
	}
	if !t.SM3WithSM2Verify([]byte(reqReq.Form.Get("data")), reqReq.Form.Get("sign")) {
		return nil, errors.New("代付结果异步通知验签失败")
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
	if notifyReq, err = model.ParseNotifyReq[model.TransferNotifyReqBody](baseRes); err != nil {
		return nil, err
	}
	return
}
