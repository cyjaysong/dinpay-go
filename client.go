package dinpay

import (
	"github.com/deatil/go-cryptobin/gm/sm2"
	"github.com/deatil/go-cryptobin/pkcs12"
	"github.com/deatil/go-cryptobin/tool/pem"
	"github.com/deatil/go-cryptobin/x509"
	reqclient "github.com/imroc/req/v3"
	"os"
	"time"
)

const alphabet = "0123456789ABCDEF"
const prodApiUrl = "https://payment.dinpay.com"

type Client struct {
	platformMerchantId         string // 平台商户Id
	merchantSM2PrivatePassword string // 商户SM2私钥证书密码
	merchantSM2PrivateKeyPath  string // 商户SM2私钥证书路径
	platformSM2PublicKeyPath   string // 平台SM2公钥证书路径
	// 内部变量
	merchantPrivateKey *sm2.PrivateKey
	platformPublicKey  *sm2.PublicKey
	reqClient          *reqclient.Client
	devMode            bool
}

func NewClient(platformMerchantId, merchantSM2PrivatePassword, merchantSM2PrivateKeyPath, platformSM2PublicKeyPath string,
	devMode bool) (client *Client, err error) {
	privateBytes, err := os.ReadFile(merchantSM2PrivateKeyPath)
	if err != nil {
		return nil, err
	}
	privateKey, _, err := pkcs12.Decode(privateBytes, merchantSM2PrivatePassword)
	if err != nil {
		return nil, err
	}
	merchantPrivateKey := privateKey.(*sm2.PrivateKey)

	publicFileBytes, err := os.ReadFile(platformSM2PublicKeyPath)
	if err != nil {
		return nil, err
	}
	publicFileBytes = []byte("-----BEGIN CERTIFICATE-----\n" + string(publicFileBytes) + "\n-----END CERTIFICATE-----")
	block, _ := pem.Decode(publicFileBytes)
	certBody, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, err
	}
	platformPublicKey := certBody.PublicKey.(*sm2.PublicKey)

	client = &Client{platformMerchantId: platformMerchantId, merchantSM2PrivatePassword: merchantSM2PrivatePassword,
		merchantSM2PrivateKeyPath: merchantSM2PrivateKeyPath, platformSM2PublicKeyPath: platformSM2PublicKeyPath,
		devMode: devMode,
	}

	client.merchantPrivateKey, client.platformPublicKey = merchantPrivateKey, platformPublicKey

	client.reqClient = reqclient.C().SetTimeout(time.Second * 10).SetCommonRetryCount(1)
	client.reqClient.SetUserAgent("")
	if devMode {
		client.reqClient.DevMode()
	}
	return
}
