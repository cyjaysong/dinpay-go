package dinpay

import (
	"crypto/rand"
	"errors"
	"github.com/bytedance/sonic"
	"github.com/cyjaysong/dinpay-go/model"
	"github.com/deatil/go-cryptobin/cryptobin/crypto"
	"github.com/deatil/go-cryptobin/gm/sm2"
	"github.com/deatil/go-cryptobin/tool/encoding"
	"github.com/imroc/req/v3"
	nanoid "github.com/matoous/go-nanoid/v2"
	"log"
	"net/url"
	"time"
)

func (t *Client) getApiUrl(path string) string {
	return prodApiUrl + path
}

// 公共相关接口请求
func (t *Client) commonJsonPost(path string, body any) (baseRes *model.BaseRes[string], err error) {
	timestampStr := time.Now().Format("20060102150405")
	bodyJsonBytes, _ := sonic.Marshal(body)
	if t.devMode {
		log.Println("data:", string(bodyJsonBytes))
	}
	bodyEncryptKey := nanoid.MustGenerate(alphabet, 16)
	bodyEncryptStr := t.SM4Encrypt(bodyJsonBytes, bodyEncryptKey)
	merchantBaseReqMap := map[string]string{"data": bodyEncryptStr, "merchantId": t.platformMerchantId,
		"encryptionKey": t.SM2Encrypt(bodyEncryptKey), "timestamp": timestampStr,
		"signatureMethod": "SM3WITHSM2", "sign": t.SM3WithSM2Sign([]byte(bodyEncryptStr))}
	response, err := t.reqClient.R().SetBodyJsonMarshal(merchantBaseReqMap).Post(t.getApiUrl(path))
	if err != nil {
		return nil, err
	}
	return t.returnStringBaseRes(response)
}

// 公共相关接口请求
func (t *Client) commonMultipartFormPost(path string, body any) (baseRes *model.BaseRes[string], err error) {
	timestampStr := time.Now().Format("20060102150405")
	bodyJsonBytes, _ := sonic.Marshal(body)
	if t.devMode {
		log.Println("data:", string(bodyJsonBytes))
	}
	bodyEncryptKey := nanoid.MustGenerate(alphabet, 16)
	bodyEncryptStr := t.SM4Encrypt(bodyJsonBytes, bodyEncryptKey)
	var formData = url.Values{"data": []string{bodyEncryptStr}, "merchantId": []string{t.platformMerchantId},
		"encryptionKey": []string{t.SM2Encrypt(bodyEncryptKey)}, "timestamp": []string{timestampStr},
		"signatureMethod": []string{"SM3WITHSM2"}, "sign": []string{t.SM3WithSM2Sign([]byte(bodyEncryptStr))}}
	response, err := t.reqClient.R().SetFormDataFromValues(formData).EnableForceMultipart().Post(t.getApiUrl(path))
	if err != nil {
		return nil, err
	}
	return t.returnStringBaseRes(response)
}

// 商户图片上传
func (t *Client) merchantImageFormUpload(path string, body any, credentialType string, getContentFunc req.GetContentFunc) (baseRes *model.BaseRes[string], err error) {
	timestampStr := time.Now().Format("20060102150405")
	bodyJsonBytes, _ := sonic.Marshal(body)
	if t.devMode {
		log.Println("data:", string(bodyJsonBytes))
	}
	bodyEncryptKey := nanoid.MustGenerate(alphabet, 16)
	bodyEncryptStr := t.SM4Encrypt(bodyJsonBytes, bodyEncryptKey)
	var formData = url.Values{"data": []string{bodyEncryptStr}, "merchantId": []string{t.platformMerchantId},
		"encryptionKey": []string{t.SM2Encrypt(bodyEncryptKey)}, "timestamp": []string{timestampStr},
		"signatureMethod": []string{"SM3WITHSM2"}, "sign": []string{t.SM3WithSM2Sign([]byte(bodyEncryptStr))}}
	request := t.reqClient.R().SetFileUpload(req.FileUpload{ParamName: "file",
		FileName: credentialType + "-" + timestampStr, GetFileContent: getContentFunc})
	response, err := request.SetFormDataFromValues(formData).Post(t.getApiUrl(path))
	if err != nil {
		return nil, err
	}
	return t.returnStringBaseRes(response)
}

// 虚拟账户支付接口请求
func (t *Client) accountPayPost(path string, body any) (baseRes *model.BaseRes[string], err error) {
	timestampStr := time.Now().Format("20060102150405")
	bodyJsonBytes, _ := sonic.Marshal(body)
	if t.devMode {
		log.Println("data:", string(bodyJsonBytes))
	}
	bodyEncryptKey := nanoid.MustGenerate(alphabet, 16)
	bodyEncryptStr := t.SM4Encrypt(bodyJsonBytes, bodyEncryptKey)
	merchantBaseReqMap := map[string]string{"data": bodyEncryptStr, "merchantId": t.platformMerchantId,
		"encryptionKey": t.SM2Encrypt(bodyEncryptKey), "timestamp": timestampStr,
		"signatureMethod": "SM3WITHSM2", "sign": t.SM3WithSM2Sign([]byte(bodyEncryptStr))}
	response, err := t.reqClient.R().SetBodyJsonMarshal(merchantBaseReqMap).Post(t.getApiUrl(path))
	if err != nil {
		return nil, err
	}
	return t.returnStringBaseRes(response)
}

// 支付相关接口请求
func (t *Client) appPayJsonPost(merchantId, path string, body any) (baseRes *model.BaseRes[string], err error) {
	timestampStr := time.Now().Format("20060102150405")
	bodyJsonBytes, _ := sonic.Marshal(body)
	if t.devMode {
		log.Println("data:", string(bodyJsonBytes))
	}
	bodyEncryptKey := nanoid.MustGenerate(alphabet, 16)
	bodyEncryptStr := t.SM4Encrypt(bodyJsonBytes, bodyEncryptKey)
	merchantBaseReqMap := map[string]string{"data": bodyEncryptStr, "merchantId": merchantId,
		"encryptionKey": t.SM2Encrypt(bodyEncryptKey), "timestamp": timestampStr,
		"signatureMethod": "SM3WITHSM2", "sign": t.SM3WithSM2Sign([]byte(bodyEncryptStr))}
	response, err := t.reqClient.R().SetBodyJsonMarshal(merchantBaseReqMap).Post(t.getApiUrl(path))
	if err != nil {
		return nil, err
	}
	return t.returnStringBaseRes(response)
}

// 结算相关接口请求
func (t *Client) settlementPost(merchantId, path string, body any) (baseRes *model.BaseRes[string], err error) {
	timestampStr := time.Now().Format("20060102150405")
	bodyJsonBytes, _ := sonic.Marshal(body)
	if t.devMode {
		log.Println("data:", string(bodyJsonBytes))
	}
	bodyEncryptKey := nanoid.MustGenerate(alphabet, 16)
	bodyEncryptStr := t.SM4Encrypt(bodyJsonBytes, bodyEncryptKey)
	merchantBaseReqMap := map[string]string{"data": bodyEncryptStr, "merchantId": merchantId,
		"encryptionKey": t.SM2Encrypt(bodyEncryptKey), "timestamp": timestampStr,
		"signatureMethod": "SM3WITHSM2", "sign": t.SM3WithSM2Sign([]byte(bodyEncryptStr))}
	response, err := t.reqClient.R().SetBodyJsonMarshal(merchantBaseReqMap).Post(t.getApiUrl(path))
	if err != nil {
		return nil, err
	}
	return t.returnStringBaseRes(response)
}

// 代付相关接口请求
func (t *Client) transferPost(merchantId, path string, body any) (baseRes *model.BaseRes[string], err error) {
	timestampStr := time.Now().Format("20060102150405")
	bodyJsonBytes, _ := sonic.Marshal(body)
	if t.devMode {
		log.Println("data:", string(bodyJsonBytes))
	}
	bodyEncryptKey := nanoid.MustGenerate(alphabet, 16)
	bodyEncryptStr := t.SM4Encrypt(bodyJsonBytes, bodyEncryptKey)
	merchantBaseReqMap := map[string]string{"data": bodyEncryptStr, "merchantId": merchantId,
		"encryptionKey": t.SM2Encrypt(bodyEncryptKey), "timestamp": timestampStr,
		"signatureMethod": "SM3WITHSM2", "sign": t.SM3WithSM2Sign([]byte(bodyEncryptStr))}
	response, err := t.reqClient.R().SetBodyJsonMarshal(merchantBaseReqMap).Post(t.getApiUrl(path))
	if err != nil {
		return nil, err
	}
	return t.returnStringBaseRes(response)
}

func (t *Client) returnStringBaseRes(response *req.Response) (baseRes *model.BaseRes[string], err error) {
	baseRes = new(model.BaseRes[string])
	if err = response.UnmarshalJson(baseRes); err != nil {
		return nil, err
	} else if baseRes.Code != "0000" {
		return baseRes, nil
	}
	if !t.SM3WithSM2Verify([]byte(baseRes.Data), baseRes.Sign) {
		return baseRes, errors.New("响应内容验签失败")
	}
	return
}

// SM3WithSM2Sign 签名
func (t *Client) SM3WithSM2Sign(data []byte) (sign string) {
	signBytes, _ := t.merchantPrivateKey.Sign(rand.Reader, data, sm2.DefaultSignerOpts)
	return encoding.Base64Encode(signBytes)
}

// SM3WithSM2Verify 验签
func (t *Client) SM3WithSM2Verify(data []byte, signBase64 string) bool {
	sign, _ := encoding.Base64Decode(signBase64)
	return t.platformPublicKey.Verify(data, sign, sm2.DefaultSignerOpts)
}

// SM2Encrypt SM2加密
func (t *Client) SM2Encrypt(data string) (encryptStr string) {
	encryptBytes, _ := t.platformPublicKey.Encrypt(rand.Reader, []byte(data), sm2.EncrypterOpts{Mode: sm2.C1C3C2})
	return encoding.Base64Encode(encryptBytes)
}

var ivBytes, _ = encoding.Base64Decode("T172oxqWwkr8wqB9D7aR7g==")

// SM4Encrypt SM4加密
func (t *Client) SM4Encrypt(data []byte, key string) (encryptStr string) {
	sm4Crypto := crypto.New().SM4().CBC()
	sm4Crypto = sm4Crypto.PKCS7Padding().SetIv(string(ivBytes))
	sm4Crypto = sm4Crypto.SetKey(key).FromBytes(data)
	if sm4Crypto = sm4Crypto.Encrypt(); sm4Crypto.Error() != nil && t.devMode {
		log.Print("Dinpay SM4加密出错:", sm4Crypto.Error())
	}
	return sm4Crypto.ToBase64String()
}
