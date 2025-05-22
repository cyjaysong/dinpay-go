package model

import (
	"github.com/bytedance/sonic"
)

type BaseMerchantOldRes[T any] struct {
	Success  bool   `json:"success"`        // 响应码
	Code     string `json:"code"`           // 响应码
	Message  string `json:"message"`        // 返回信息
	Data     T      `json:"data,omitempty"` // 响应实体
	Sign     string `json:"sign,omitempty"` // 签名
	Hostname string `json:"hostname"`       // 处理机别名
}

// BaseRes 响应公共结构体
type BaseRes[T any] struct {
	Code            string `json:"code"`                      // 响应码,详情见附录返回码,不参与签名
	Msg             string `json:"msg"`                       // 返回信息,响应结果文本信息,不参与签名
	SignatureMethod string `json:"signatureMethod,omitempty"` // 签名方式,本次接口请求使用的签名算法,固定:SM3WITHSM2,不参与签名
	Sign            string `json:"sign,omitempty"`            // 签名,响应参数data明文转成json格式后使用SM3WITHSM2的签名值,不参与签名
	Data            T      `json:"data,omitempty"`            // 响应实体
	MerchantId      string `json:"merchantId"`                // 商户编号,由系统生成的唯一商户标识,不参与签名
}

func ParseRes[T any](baseRes *BaseRes[string]) (res *BaseRes[T], err error) {
	res = &BaseRes[T]{Code: baseRes.Code, Msg: baseRes.Msg, MerchantId: baseRes.MerchantId,
		SignatureMethod: baseRes.SignatureMethod, Sign: baseRes.Sign}
	if baseRes.Data == "" {
		return
	}
	err = sonic.Unmarshal([]byte(baseRes.Data), &res.Data)
	return
}

// NotifyReq 回调公共结构体
type NotifyReq[T any] struct {
	Code            string `json:"code"`                      // 响应码,详情见附录返回码,不参与签名
	Msg             string `json:"msg"`                       // 返回信息,响应结果文本信息,不参与签名
	SignatureMethod string `json:"signatureMethod,omitempty"` // 签名方式,本次接口请求使用的签名算法,固定:SM3WITHSM2,不参与签名
	Sign            string `json:"sign,omitempty"`            // 签名,响应参数data明文转成json格式后使用SM3WITHSM2的签名值,不参与签名
	Data            T      `json:"data,omitempty"`            // 响应实体
	MerchantId      string `json:"merchantId"`                // 商户编号,由系统生成的唯一商户标识,不参与签名
}

func ParseNotifyReq[T any](baseRes *NotifyReq[string]) (notifyReq *NotifyReq[T], err error) {
	notifyReq = &NotifyReq[T]{Code: baseRes.Code, Msg: baseRes.Msg, MerchantId: baseRes.MerchantId,
		SignatureMethod: baseRes.SignatureMethod, Sign: baseRes.Sign}
	if baseRes.Data == "" {
		return
	}
	err = sonic.Unmarshal([]byte(baseRes.Data), &notifyReq.Data)
	return
}
