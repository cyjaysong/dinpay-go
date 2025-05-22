package model

type WxPayGoods struct {
	CostPrice   int               `json:"cost_price,omitempty"` // 订单原价 订单总金额，与total_fee一致
	ReceiptId   string            `json:"receipt_id,omitempty"` // 商品小票ID 商家小票ID
	GoodsDetail []GoodsDetailItem `json:"goods_detail"`         // 单品列表 单品信息，使用Json数组格式提交
}

type GoodsDetailItem struct {
	GoodsId      string `json:"goods_id"`                 // 商品编码：由半角的大小写字母、数字、中划线、下划线中的一种或几种组成
	WxPayGoodsId string `json:"wxpay_goods_id,omitempty"` // 微信侧商品编码：微信支付定义的统一商品编号（没有可不传）
	GoodsName    string `json:"goods_name,omitempty"`     // 商品名称：商品的实际名称
	Quantity     int    `json:"quantity"`                 // 商品数量：用户购买的数量
	Price        int    `json:"price"`                    // 商品单价：单位为分。如果商户有优惠，需传输商户优惠后的单价
}

type AliPayGoods struct {
	GoodsId       string  `json:"goods_id"`                  // 商品的编号
	AlipayGoodsId string  `json:"alipay_goods_id,omitempty"` // 支付宝定义的统一商品编号
	GoodsName     string  `json:"goods_name"`                // 商品名称
	Quantity      int     `json:"quantity"`                  // 商品数量
	Price         float64 `json:"price"`                     // 商品单价,单位为元
	GoodsCategory string  `json:"goods_category,omitempty"`  // 商品类目
	Body          string  `json:"body,omitempty"`            // 商品描述信息
	ShowURL       string  `json:"show_url,omitempty"`        // 商品的展示地址
}
