package dinpay

import "github.com/cyjaysong/dinpay-go/model"

// MerchantModifyInfo 商户信息变更
func (t *Client) MerchantModifyInfo(reqBody model.MerchantModifyInfoReq) (res *model.BaseRes[model.MerchantModifyInfoRes], err error) {
	const path = "/trx/api/merchantEntryAlteration/modifyMerchantInfo"
	reqBody.InterfaceName = "modifyMerchantInfo"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonMultipartFormPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantModifyInfoRes](baseRes)
}

// MerchantModifyInfoQuery 商户信息变更查询
func (t *Client) MerchantModifyInfoQuery(reqBody model.MerchantModifyInfoQueryReq) (res *model.BaseRes[model.MerchantModifyInfoQueryRes], err error) {
	const path = "/trx/api/merchantCredential/changeOrderQuery"
	reqBody.InterfaceName, reqBody.ChangeType = "changeOrderQuery", "MERCHANT_INFO_SYN"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantModifyInfoQueryRes](baseRes)
}

// MerchantCredentialImageUpload 商户资质图片上传
func (t *Client) MerchantCredentialImageUpload(reqBody model.MerchantCredentialImageUploadReq) (res *model.BaseRes[model.MerchantCredentialImageUploadRes], err error) {
	const path = "/trx/api/merchantCredential/imageUpload"
	reqBody.InterfaceName = "imageUpload"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.merchantImageFormUpload(path, reqBody, reqBody.CredentialType, reqBody.GetFileContent); err != nil {
		return
	}
	return model.ParseRes[model.MerchantCredentialImageUploadRes](baseRes)
}

// MerchantCredentialImageUrlUpload 商户资质图片Url上传
func (t *Client) MerchantCredentialImageUrlUpload(reqBody model.MerchantCredentialImageUrlUploadReq) (res *model.BaseRes[model.MerchantCredentialImageUploadRes], err error) {
	const path = "/trx/api/merchantCredential/imageUrlUpload"
	reqBody.InterfaceName = "imageUrlUpload"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantCredentialImageUploadRes](baseRes)
}

// MerchantCredentialImageUploadQuery 商户资质图片上传查询
func (t *Client) MerchantCredentialImageUploadQuery(reqBody model.MerchantCredentialImageUploadQueryReq) (res *model.BaseRes[model.MerchantCredentialImageUploadRes], err error) {
	const path = "/trx/api/merchantCredential/imageUploadQuery"
	reqBody.InterfaceName = "imageUploadQuery"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantCredentialImageUploadRes](baseRes)
}

// MerchantCredentialImageUrlUploadQuery 商户资质图片Url上传查询
func (t *Client) MerchantCredentialImageUrlUploadQuery(reqBody model.MerchantCredentialImageUploadQueryReq) (res *model.BaseRes[model.MerchantCredentialImageUploadRes], err error) {
	const path = "/trx/api/merchantCredential/imageUrlUploadQuery"
	reqBody.InterfaceName = "imageUrlUploadQuery"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantCredentialImageUploadRes](baseRes)
}

// MerchantCredentialImageChangeUpload 商户资质图片变更上传
func (t *Client) MerchantCredentialImageChangeUpload(reqBody model.MerchantCredentialImageChangeUploadReq) (res *model.BaseRes[model.MerchantCredentialImageChangeUploadRes], err error) {
	const path = "/trx/api/merchantCredential/imageChangeUpload"
	reqBody.InterfaceName, reqBody.ChangeType = "imageChangeUpload", "MERCHANT_CREDENTIAL"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.merchantImageFormUpload(path, reqBody, reqBody.CredentialType, reqBody.GetFileContent); err != nil {
		return
	}
	return model.ParseRes[model.MerchantCredentialImageChangeUploadRes](baseRes)
}

// MerchantCredentialImageUrlChangeUpload 商户资质图片Url变更上传
func (t *Client) MerchantCredentialImageUrlChangeUpload(reqBody model.MerchantCredentialImageUrlChangeUploadReq) (res *model.BaseRes[model.MerchantCredentialImageChangeUploadRes], err error) {
	const path = "/trx/api/merchantCredential/imageUrlChangeUpload"
	reqBody.InterfaceName = "imageUrlChangeUpload"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantCredentialImageChangeUploadRes](baseRes)
}

// MerchantCredentialImageChangeUploadQuery 商户资质图片变更上传查询
func (t *Client) MerchantCredentialImageChangeUploadQuery(reqBody model.MerchantCredentialImageChangeUploadQueryReq) (res *model.BaseRes[model.MerchantCredentialImageChangeUploadRes], err error) {
	const path = "/trx/api/merchantCredential/changeOrderQuery"
	reqBody.InterfaceName, reqBody.ChangeType = "changeOrderQuery", "MERCHANT_CREDENTIAL"
	var baseRes *model.BaseRes[string]
	if baseRes, err = t.commonJsonPost(path, reqBody); err != nil {
		return
	}
	return model.ParseRes[model.MerchantCredentialImageChangeUploadRes](baseRes)
}
