package dinpay

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	"github.com/bytedance/sonic"
	"github.com/cyjaysong/dinpay-go/model"
)

// ComplaintNotifyVerify 投诉内容回调验签
func (t *Client) ComplaintNotifyVerify(httpBody []byte) (notifyReq *model.ComplaintNotifyReq, err error) {
	reqReq := http.Request{Method: "POST", Body: io.NopCloser(bytes.NewBuffer(httpBody)), Header: http.Header{}}
	reqReq.Header.Set(`Content-Type`, `application/x-www-form-urlencoded`)
	if err = reqReq.ParseForm(); err != nil {
		return nil, err
	}
	if !t.SM3WithSM2Verify([]byte(reqReq.Form.Get("data")), reqReq.Form.Get("sign")) {
		return nil, errors.New("投诉内容验签失败")
	}
	httpBodyMap := make(map[string]string, len(reqReq.Form))
	for key, strings := range reqReq.Form {
		httpBodyMap[key] = strings[0]
	}
	httpBody, err = sonic.Marshal(httpBodyMap)
	if err != nil {
		return nil, err
	}

	var baseRes *model.NotifyReq[string]
	if err = sonic.Unmarshal(httpBody, &baseRes); err != nil {
		return nil, err
	}
	if notifyReq, err = model.ParseNotifyReq[model.ComplaintNotifyReqBody](baseRes); err != nil {
		return nil, err
	}
	return
}
