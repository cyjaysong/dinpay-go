package dinpay

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/cyjaysong/dinpay-go/model"
)

// MerchantSettlement 商户结算接口
func (t *Client) MerchantSettlement(reqBody model.MerchantSettlementReq) (res *model.BaseRes[model.MerchantSettlementRes], err error) {
	const path = "/trx/api/settle"
	reqBody.InterfaceName = "MerchantSettlement"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.settlementPost(reqBody.MerchantId, path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantSettlementRes](baseRes)
}

// MerchantSettlementQuery 商户结算查询接口
func (t *Client) MerchantSettlementQuery(reqBody model.MerchantSettlementQueryReq) (res *model.BaseRes[model.MerchantSettlementQueryRes], err error) {
	const path = "/trx/api/settleQuery"
	reqBody.InterfaceName = "MerchantSettlementQuery"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.settlementPost(reqBody.MerchantId, path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantSettlementQueryRes](baseRes)
}

// MerchantSettlementResultNotifyVerify 商户结算结果异步通知验签
func (t *Client) MerchantSettlementResultNotifyVerify(httpBody []byte) (notifyReq *model.MerchantSettlementNotifyReq, err error) {
	notifyReq = new(model.MerchantSettlementNotifyReq)
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
