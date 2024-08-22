package dinpay

import (
	"crypto/md5"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/ast"
	"github.com/deatil/go-cryptobin/cryptobin/crypto"
	"github.com/deatil/go-cryptobin/cryptobin/sm2"
	"github.com/deatil/go-cryptobin/hash/sm3"
	"github.com/deatil/go-cryptobin/tool"
	"github.com/imroc/req/v3"
	nanoid "github.com/matoous/go-nanoid/v2"
	"reflect"
	"strings"
	"time"
)

func (t *Client) getApiUrl(path string) string {
	return prodApiUrl + path
}

// merchantPost 商户相关接口请求
func (t *Client) merchantPost(interfaceName string, body any) (response *req.Response, err error) {
	const path = "/trx/merchantEntry/interface.action"
	bodyJsonBytes, _ := sonic.Marshal(body)
	bodyEncryptStr := t.Des3Encrypt(bodyJsonBytes, t.commonEncryptKey)
	merchantBaseReqMap := map[string]string{"interfaceName": interfaceName, "merchantNo": t.platformMerchantId,
		"body": bodyEncryptStr, "sign": t.MD5Sign([]string{bodyEncryptStr, t.platformMerchantId}, t.commonSignKey)}
	if response, err = t.reqClient.R().SetFormData(merchantBaseReqMap).Post(t.getApiUrl(path)); err != nil {
		return nil, err
	}
	return
}

// merchantUploadPost 商户相关接口请求(带上传)
func (t *Client) merchantUploadPost(interfaceName string, body any) (response *req.Response, err error) {
	const path = "/trx/merchantEntry/upload.action"
	bodyJsonBytes, _ := sonic.Marshal(body)
	bodyEncryptStr := t.Des3Encrypt(bodyJsonBytes, t.commonEncryptKey)
	merchantBaseReqMap := map[string]string{"interfaceName": interfaceName, "merchantNo": t.platformMerchantId,
		"body": bodyEncryptStr, "sign": t.MD5Sign([]string{bodyEncryptStr, t.platformMerchantId}, t.commonSignKey)}
	request := t.reqClient.R().SetFormData(merchantBaseReqMap)
	if response, err = request.Post(t.getApiUrl(path)); err != nil {
		return nil, err
	}
	return
}

// merchantAgreementPost 商户签约相关接口请求
func (t *Client) merchantAgreementPost(interfaceName string, body any) (response *req.Response, err error) {
	const path = "/trx/merchantAgreement/ signContract.action"
	bodyJsonBytes, _ := sonic.Marshal(body)
	bodyEncryptStr := t.Des3Encrypt(bodyJsonBytes, t.commonEncryptKey)
	merchantBaseReqMap := map[string]string{"interfaceName": interfaceName, "merchantNo": t.platformMerchantId,
		"body": bodyEncryptStr, "sign": t.MD5Sign([]string{bodyEncryptStr, t.platformMerchantId}, t.commonSignKey)}
	request := t.reqClient.R().SetFormData(merchantBaseReqMap)
	if response, err = request.Post(t.getApiUrl(path)); err != nil {
		return nil, err
	}
	return
}

// appPayPost 支付相关接口请求
func (t *Client) appPayPost(body any) (response *req.Response, err error) {
	const path = "/trx/app/interface.action"
	bodyJsonBytes, _ := sonic.Marshal(body)
	jsonNode, _ := sonic.Get(bodyJsonBytes)
	bodyType := reflect.TypeOf(body)
	if bodyType.Kind() == reflect.Ptr {
		bodyType = bodyType.Elem()
	}
	var signVals = []string{""}
	for i := 0; i < bodyType.NumField(); i++ {
		var field = bodyType.Field(i)
		if field.Tag.Get("sign") != "1" {
			continue
		}
		jsonName := field.Tag.Get("json")
		if len(jsonName) == 0 {
			jsonName = field.Name
		}
		var valStr string
		if val := jsonNode.Get(strings.Split(jsonName, ",")[0]); val.Valid() {
			valStr, _ = val.String()
		}
		signVals = append(signVals, valStr)
	}
	_, _ = jsonNode.Set("sign", ast.NewString(t.MD5Sign(signVals, t.scanSignKey)))
	bodyMap, _ := jsonNode.Map()
	if response, err = t.reqClient.R().SetFormDataAnyType(bodyMap).Post(t.getApiUrl(path)); err != nil {
		return nil, err
	}
	return
}

// 结算相关接口请求
func (t *Client) settlementPost(body any) (response *req.Response, err error) {
	const path = "/trx/transfer/interface.action"
	bodyJsonBytes, _ := sonic.Marshal(body)
	jsonNode, _ := sonic.Get(bodyJsonBytes)
	bodyType := reflect.TypeOf(body)
	if bodyType.Kind() == reflect.Ptr {
		bodyType = bodyType.Elem()
	}
	var signVals = []string{""}
	for i := 0; i < bodyType.NumField(); i++ {
		var field = bodyType.Field(i)
		if field.Tag.Get("sign") != "1" {
			continue
		}
		jsonName := field.Tag.Get("json")
		if len(jsonName) == 0 {
			jsonName = field.Name
		}
		var valStr string
		if val := jsonNode.Get(strings.Split(jsonName, ",")[0]); val.Valid() {
			valStr, _ = val.String()
		}
		signVals = append(signVals, valStr)
	}
	signValsBytes := []byte(strings.Join(append(signVals, "SM3WITHSM2"), "&"))
	_, _ = jsonNode.Set("signType", ast.NewString("SM3WITHSM2"))
	_, _ = jsonNode.Set("sign", ast.NewString(t.SM3WithSM2Sign(signValsBytes, t.settlementSignKey)))
	bodyMap, _ := jsonNode.Map()
	if response, err = t.reqClient.R().SetFormDataAnyType(bodyMap).Post(t.getApiUrl(path)); err != nil {
		return nil, err
	}
	return
}

// 商户余额接口请求
func (t *Client) merchantBalancePost(body any) (response *req.Response, err error) {
	const alphabet = "0123456789ABCDEF"
	var path = "/trx/api/merchant/merchantBalanceQuery"
	timestampStr := time.Now().Format("20060102150405")

	bodyJsonBytes, _ := sonic.Marshal(body)
	bodyEncryptKey := nanoid.MustGenerate(alphabet, 16)
	bodyEncryptStr := t.SM4Encrypt(bodyJsonBytes, bodyEncryptKey)
	merchantBaseReqMap := map[string]string{"body": bodyEncryptStr, "merchantId": t.platformMerchantId,
		"encryptionKey": t.SM2Encrypt(bodyEncryptKey, t.accountPayEncryptKey), "timestamp": timestampStr,
		"signatureMethod": "SM3WITHSM2", "sign": t.SM3WithSM2Sign([]byte(bodyEncryptStr), t.accountPaySignKey)}
	if response, err = t.reqClient.R().SetFormData(merchantBaseReqMap).Post(t.getApiUrl(path)); err != nil {
		return nil, err
	}
	return
}

// 虚拟账户支付接口请求
func (t *Client) accountPayPost(pathSuffix string, body any) (response *req.Response, err error) {
	const alphabet = "0123456789ABCDEF"
	var path = "/trx/api/accountPay/" + pathSuffix
	timestampStr := time.Now().Format("20060102150405")

	bodyJsonBytes, _ := sonic.Marshal(body)
	bodyEncryptKey := nanoid.MustGenerate(alphabet, 16)
	bodyEncryptStr := t.SM4Encrypt(bodyJsonBytes, bodyEncryptKey)
	merchantBaseReqMap := map[string]string{"body": bodyEncryptStr, "merchantId": t.platformMerchantId,
		"encryptionKey": t.SM2Encrypt(bodyEncryptKey, t.accountPayEncryptKey), "timestamp": timestampStr,
		"signatureMethod": "SM3WITHSM2", "sign": t.SM3WithSM2Sign([]byte(bodyEncryptStr), t.accountPaySignKey)}
	if response, err = t.reqClient.R().SetFormData(merchantBaseReqMap).Post(t.getApiUrl(path)); err != nil {
		return nil, err
	}
	return
}

func (t *Client) getSignValBySignTag(body any, bodyJsonBytes []byte) (signVals []string) {
	jsonNode, _ := sonic.Get(bodyJsonBytes)
	bodyType := reflect.TypeOf(body)
	if bodyType.Kind() == reflect.Ptr {
		bodyType = bodyType.Elem()
	}
	signVals = []string{""}
	for i := 0; i < bodyType.NumField(); i++ {
		var field = bodyType.Field(i)
		if field.Tag.Get("sign") != "1" {
			continue
		}
		jsonName := field.Tag.Get("json")
		if len(jsonName) == 0 {
			jsonName = field.Name
		}
		var valStr string
		if val := jsonNode.Get(strings.Split(jsonName, ",")[0]); val.Valid() {
			valStr, _ = val.String()
		}
		signVals = append(signVals, valStr)
	}
	return
}

// MD5Sign MD5 签名
func (t *Client) MD5Sign(signVals []string, key string) string {
	//var c = strings.Join(append(data, key), "&")
	//hash := md5.Sum([]byte(c))
	//fmt.Println(c)
	hash := md5.Sum([]byte(strings.Join(append(signVals, key), "&")))
	return fmt.Sprintf("%x", hash)
}

// SM3WithSM2Sign SM3WithSM2签名
func (t *Client) SM3WithSM2Sign(data []byte, privateKey string) (sign string) {
	sm3Hash := sm3.New()
	sm3Hash.Write(data)
	sm3HashData := sm3Hash.Sum(nil)

	sm2Sign := sm2.New()
	sm2Sign = sm2Sign.FromPrivateKeyString(privateKey).FromBytes(sm3HashData)
	if sign = sm2Sign.Sign().ToString(); sm2Sign.Error() != nil {
		fmt.Println("SM3WithSM2签名出错:", sm2Sign.Error())
	}
	return
}

// Des3Encrypt Des3 加密
func (t *Client) Des3Encrypt(data []byte, key string) (encrypt string) {
	desCrypto := crypto.New().TripleDes().PKCS5Padding()
	desCrypto = desCrypto.FromBytes(data).SetKey(key)
	if desCrypto = desCrypto.Encrypt(); desCrypto.Error() != nil {
		fmt.Println("Des3加密出错:", desCrypto.Error())
	}
	return desCrypto.ToBase64String()
}

// Des3dDecrypt Des3 解密
func (t *Client) Des3dDecrypt(base64StrData string, key string) (decrypt []byte) {
	desCrypto := crypto.New().TripleDes().PKCS5Padding()
	desCrypto = desCrypto.FromBase64String(base64StrData).SetKey(key)
	if desCrypto = desCrypto.Decrypt(); desCrypto.Error() != nil {
		fmt.Println("Des3解密出错:", desCrypto.Error())
	}
	return desCrypto.ToBytes()
}

// SM2Encrypt SM2加密
func (t *Client) SM2Encrypt(data string, key string) (encrypt string) {
	sm2Encrypt := sm2.New().SetMode("C1C3C2")
	sm2Encrypt = sm2Encrypt.FromPublicKeyString(key).FromString(data)
	if sm2Encrypt = sm2Encrypt.Encrypt(); sm2Encrypt.Error() != nil {
		fmt.Println("SM2加密出错:", sm2Encrypt.Error())
	}
	return sm2Encrypt.ToBase64String()
}

// SM4Encrypt SM4加密
func (t *Client) SM4Encrypt(data []byte, key string) (encrypt string) {
	ivBytes, _ := tool.Base64Decode("T172oxqWwkr8wqB9D7aR7g==")
	sm4Crypto := crypto.New().SM4().CBC()
	sm4Crypto = sm4Crypto.PKCS7Padding().SetIv(string(ivBytes))
	sm4Crypto = sm4Crypto.SetKey(key).FromBytes(data)
	if sm4Crypto = sm4Crypto.Encrypt(); sm4Crypto.Error() != nil {
		fmt.Println("SM4加密出错:", sm4Crypto.Error())
	}
	return sm4Crypto.ToBase64String()
}
