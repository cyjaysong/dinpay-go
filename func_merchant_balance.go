package dinpay

import "github.com/cyjaysong/dinpay-go/model"

// MerchantBalanceQuery 商户余额查询接口
func (t *Client) MerchantBalanceQuery(reqBody model.MerchantBalanceQueryReq) (res *model.BaseRes[model.MerchantBalanceQueryRes], err error) {
	const path = "/trx/api/merchant/merchantBalanceQuery"
	reqBody.InterfaceName = "merchantBalanceQuery"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantBalanceQueryRes](baseRes)
}
