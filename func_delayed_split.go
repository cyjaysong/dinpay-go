package dinpay

import "github.com/cyjaysong/dinpay-go/model"

func (t *Client) DelayedSplit(reqBody model.DelayedSplitReq) (res *model.BaseRes[model.DelayedSplitRes], err error) {
	const path = "/trx/api/delayed/applySplit"
	reqBody.InterfaceName = "delaySplitting"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.appPayJsonPost(reqBody.MerchantId, path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.DelayedSplitRes](baseRes)
}

func (t *Client) DelayedSplitQuery(reqBody model.DelayedSplitQueryReq) (res *model.BaseRes[model.DelayedSplitQueryRes], err error) {
	const path = "/trx/api/delayed/querySplit"
	reqBody.InterfaceName = "delaySplittingQuery"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.appPayJsonPost(reqBody.MerchantId, path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.DelayedSplitQueryRes](baseRes)
}

func (t *Client) DelayedSplitBack(reqBody model.DelayedSplitBackReq) (res *model.BaseRes[model.DelayedSplitBackRes], err error) {
	const path = "/trx/api/delayed/backSplit"
	//delaySplittingRfund 没写错,智付的历史遗留问题
	reqBody.InterfaceName = "delaySplittingRfund"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.appPayJsonPost(reqBody.MerchantId, path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.DelayedSplitBackRes](baseRes)
}

func (t *Client) DelayedSplitBackQuery(reqBody model.DelayedSplitBackQueryReq) (res *model.BaseRes[model.DelayedSplitBackQueryRes], err error) {
	const path = "/trx/api/delayed/queryBackSplit"
	//delaySplittingRfundQuery 没写错,智付的历史遗留问题
	reqBody.InterfaceName = "delaySplittingRfundQuery"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.appPayJsonPost(reqBody.MerchantId, path, reqBody); err != nil {
		return
	}
	res = new(model.BaseRes[model.DelayedSplitBackQueryRes])
	return model.ParseRes[model.DelayedSplitBackQueryRes](baseRes)
}
