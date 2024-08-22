package dinpay

import (
	reqclient "github.com/imroc/req/v3"
	"time"
)

const testApiUrl = "https://payment.dinpay.com"
const prodApiUrl = "https://payment.dinpay.com"

type Client struct {
	platformMerchantId   string // 接入商ID
	commonSignKey        string // 公共产品签名秘钥
	commonEncryptKey     string // 公共产品加密秘钥
	scanSignKey          string // 扫码产品签名秘钥
	scanEncryptKey       string // 扫码产品加密秘钥
	settlementSignKey    string // 结算产品签名秘钥
	settlementEncryptKey string // 结算产品加密秘钥
	accountPaySignKey    string // 虚拟账号支付产品签名秘钥
	accountPayEncryptKey string // 虚拟账号支付产品加密秘钥
	reqClient            *reqclient.Client
}

func NewClient(platformMerchantId, commonSignKey, commonEncryptKey, scanSignKey, scanEncryptKey,
	settlementSignKey, settlementEncryptKey, accountPaySignKey, accountPayEncryptKey string, isProd bool) (client *Client) {
	client = &Client{platformMerchantId: platformMerchantId,
		commonSignKey: commonSignKey, commonEncryptKey: commonEncryptKey,
		scanSignKey: scanSignKey, scanEncryptKey: scanEncryptKey,
		settlementSignKey: settlementSignKey, settlementEncryptKey: settlementEncryptKey,
		accountPaySignKey: accountPaySignKey, accountPayEncryptKey: accountPayEncryptKey,
	}
	client.reqClient = reqclient.C().SetTimeout(time.Second * 10).SetCommonRetryCount(1)
	if !isProd {
		client.reqClient.DevMode()
	}
	return
}
