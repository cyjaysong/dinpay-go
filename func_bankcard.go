package dinpay

// MerchantBankcardQuery 商户银行卡查询
func (t *Client) MerchantBankcardQuery(cardNo, orderNo string) (res *BaseRes[MerchantBankcardQueryRes], err error) {
	const path = "/trx/api/merchant/queryCardNo"
	reqBody := map[string]string{"interfaceName": "MerchantQueryCardNo", "merchantId": t.platformMerchantId,
		"cardNo": cardNo, "orderNum": orderNo}
	var baseRes *BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return nil, err
	}
	return ParseRes[MerchantBankcardQueryRes](baseRes)
}
