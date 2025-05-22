package dinpay

import "github.com/cyjaysong/dinpay-go/model"

// MerchantBankcardQuery 商户银行卡查询
func (t *Client) MerchantBankcardQuery(cardNo, orderNo string) (res *model.BaseRes[model.MerchantBankcardQueryRes], err error) {
	const path = "/trx/api/merchant/queryCardNo"
	reqBody := map[string]string{"interfaceName": "MerchantQueryCardNo", "merchantId": t.platformMerchantId,
		"cardNo": cardNo, "orderNum": orderNo}
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return nil, err
	}
	return model.ParseRes[model.MerchantBankcardQueryRes](baseRes)
}
