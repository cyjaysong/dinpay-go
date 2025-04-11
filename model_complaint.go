package dinpay

import "time"

type ComplaintNotifyReqBody struct {
	MerchantId         string    `json:"merchantId"`                  // 商户编号,支付系统分配的商户号
	ComplaintId        string    `json:"complaintId"`                 // 投诉单号
	OrderNo            string    `json:"orderNo"`                     // 商户订单号
	ComplaintState     string    `json:"complaintState"`              // 投诉单状态
	ProblemType        string    `json:"problemType,omitempty"`       // 问题类型
	ProblemDescription string    `json:"problemDescription"`          //问题描述
	ComplaintDetail    string    `json:"complaintDetail"`             //投诉详情
	UserTagList        string    `json:"userTagList"`                 //用户标签
	ComplaintDate      time.Time `json:"complaintDate"`               //时间格式：yyyy-MM-dd HH:mm:ss
	ImgFileId          string    `json:"imgFileId,omitempty"`         //用户投诉图片id json数组USER_COMPLAINT_IMAGE:用户投诉图片;OPERATION_IMAGE:操作流水图片
	ReplyState         string    `json:"replyState"`                  //回复状态
	ErrorMsg           string    `json:"errorMsg,omitempty"`          //错误信息
	NotifyId           string    `json:"notifyId,omitempty"`          //通知报文id
	Amount             float64   `json:"amount,omitempty"`            //投诉订单金额
	PayerPhone         string    `json:"payerPhone,omitempty"`        //投诉人联系方式
	ComplaintTimes     int64     `json:"complaintTimes,omitempty"`    //用户投诉次数
	AppPayType         string    `json:"appPayType,omitempty"`        //投诉单类型 WXPAY ALIPAY
	ComplainUrl        string    `json:"complainUrl,omitempty"`       //投诉网址,支付宝特有参数
	ProcessDate        string    `json:"processDate,omitempty"`       //处理时间
	ProcessCode        string    `json:"processCode,omitempty"`       //商家处理结果码
	ProcessRemark      string    `json:"processRemark,omitempty"`     //商家处理备注
	ProcessImgUrlList  string    `json:"processImgUrlList,omitempty"` //商家处理备注图片url列表
}

type ComplaintNotifyReq = NotifyReq[ComplaintNotifyReqBody]
