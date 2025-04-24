package constants

var (
	// Code 返回码
	Code = map[string]string{
		"0000": "成功",
		"0001": "文件IO异常",
		"0002": "接入IP不匹配",
		"0003": "系统繁忙，请重试一次",
		"0004": "请求太频繁，请稍后再试",
		"0005": "请勿重新提交",
		"0006": "商户信息不完整，请先完善",
		"1000": "子商户未开通此产品",
		"1001": "参数错误",
		"1002": "非法的身份证号码",
		"1003": "非法的区域编码",
		"1004": "非法的行业编码",
		"1005": "邮箱已被占用",
		"1006": "签约名已经存在",
		"1007": "营业执照号已存在",
		"1008": "非法的结算卡号",
		"1009": "找不着联行号对应的记录",
		"1010": "非法的商户号",
		"1011": "商户状态异常",
		"1012": "未开通商户进件接口",
		"1013": "验签失败了",
		"1014": "非法的 JSON 串",
		"1015": "body 解密失败了",
		"1016": "非法的费率值",
		"1017": "费率不能小于自身产品的费率",
		"1018": "非法最低费率金额",
		"1019": "浮动值不能小于自身的浮动值",
		"1020": "最低费率金额不能小于自身的金额",
		"1021": "平台商未开通此产品",
		"1022": "不允许对该产品进行费率设置",
		"1023": "非法的浮动值",
		"1024": "未找着对应的子商户",
		"1025": "子商户尚未通过审核",
		"1026": "进件信息待审核",
		"1027": "商户已进件，不允许重复提交",
		"1028": "重新进件失败，订单不匹配",
		"1029": "订单号重复了",
		"1030": "产品开通失败",
		"1031": "产品费率设置失败了",
		"1032": "未找着对应商品费率记录",
		"1033": "尚未支持此产品",
		"1034": "邮箱重复了",
		"1035": "营业执照号重复了",
		"1036": "非法最高费率金额",
		"1037": "最高费率金额不能大于自身的金额",
		"1038": "进件信息重新提交申请，营业执照号已存在",
		"1039": "进件信息重新提交申请，邮箱已存在",
		"1040": "子商户状态异常",
		"1049": "非法的结算账户名", // 结算卡户名与法人不一致
		"1103": "未上传身份证正面照",
		"1104": "未上传身份证反面照",
		"1105": "未上传营业执照",
		"1107": "未上传税务登记证",
		"1109": "未上传开户许可证",
		"1111": "未上传结算账户指定书",
		"1113": "未上传持卡人手持身份证照",
		"1114": "未上传持卡人手持银行卡照",
		"1117": "未上传结算卡照",
		"1119": "未上传转租证明",
		"1120": "未上传门头照",
		"1121": "未上传门头照",
		"1122": "未上传结算人身份证正面照",
		"1123": "未上传结算人身份证反面照",
		"1200": "提交失败，请稍后重试",
		"1201": "正在审核中",
		"1202": "支付服务协议签署提交成功",
		"1203": "支付服务协议签署提交成功",
		"1204": "未在数据库中找到对应的记录",
		"1205": "记录已存在",
		"1206": "结算变更超过设置最大次数",
		"1207": "上传合同图片失败",
		"1208": "微信公众号报备次数不能大于10",
		"1209": "该报备记录不存在",
		"1210": "子商户号或邮箱二选一",
		"1211": "businessCode或applymentId二选一",
		"1212": "报备状态非审核驳回或已作废",
		"1213": "商户鉴权失败",
		"1214": "资质文件url已上传",
		"1215": "资质文件url正在处理中",
		"1216": "结算卡鉴权失败+原因",
		"1217": "缺少结算人手机号",
		"1218": "未开通电子账户",
		"1219": "活动已过期",
		"1220": "只支持对审核驳回的进行修改",
	}

	// AuthenticationType 鉴权类型
	AuthenticationType = map[string]string{
		"TRADE":        "交易鉴权",
		"INDEPENDENCE": "独立鉴权",
	}

	// CalcType 计算类型
	CalcType = map[string]string{
		"SINGLE": "单笔收费",
		"RATIO":  "比率收费",
	}

	// BusinessCategory 经营类别
	BusinessCategory = map[string]string{
		"OFFLINE_RETAIL":                 "线下零售",
		"FOOD_BEVERAGE":                  "餐饮/食品",
		"TICKETING_TRAVEL":               "票务/旅游",
		"EDUCATION_TRAINING":             "教育/培训",
		"LIFE_ADVISORY_SERVICE":          "生活/咨询服务",
		"ENTERTAINMENT_FITNESS_SERVICES": "娱乐/健身服务",
		"MEDICAL_CARE":                   "医疗",
		"COLLECTION_AUCTION":             "收藏/拍卖",
		"LOGISTICS_EXPRESS":              "物流/快递",
		"PUBLIC_WELFARE":                 "公益",
		"COMMUNICATION":                  "通讯",
		"FINANCE_INSURANCE":              "金融/保险",
		"NETWORK_VIRTUAL_SERVICE":        "网络虚拟服务",
		"LIVING_PAYMENT":                 "生活缴费",
		"HOTEL":                          "酒店",
		"HOME_FURNISHING":                "家居",
		"GROUP_PURCHASE":                 "电商团购",
		"LOTTERY":                        "彩票",
		"FASHION":                        "时尚",
		"UTILITIES_EXPENSE":              "公共事业缴费",
		"REAL_ESTATE":                    "房地产",
		"TRANSPORTATION_SERVICE":         "交通运输服务类",
		"MACHINE_ELECTRON":               "机械/电子",
		"SOFTWARE":                       "软件",
		"DECORATION":                     "装饰",
		"GREEN_SEEDLING":                 "苗木/绿化",
		"MATERNAL_CHILD_PRODUCT":         "母婴/玩具",
		"COLLECTION_PET":                 "收藏/宠物",
		"BOOK_STATIONERY_AUDIO_VIDEO":    "书籍/音像/文具",
		"BUSINESS_PLATFORM":              "平台商",
		"DIGITAL":                        "数码",
		"DIGITAL_ENTERTAINMENT":          "数字娱乐",
		"ABROAD":                         "境外",
		"PREPAID_CARD":                   "预付卡",
		"DIRECT_SELLING":                 "直销",
		"CROWD_FUNDING":                  "众筹",
		"OTHER":                          "其他",
		"CAMPUS_GROUPMEAL":               "校园团餐",
		"UNIVERSITIES_SECONDARY_CAMPUS":  "高校和中职校园",
	}

	// MerchantEntryStatus 审核状态
	MerchantEntryStatus = map[string]string{
		"INIT":     "待审核",
		"OVERRULE": "申请驳回",
		"AUDITED":  "审核通过",
	}

	// MerchantType 商户类型
	MerchantType = map[string]string{
		"ENTERPRISE":          "企业",
		"INSTITUTION":         "事业单位",
		"INDIVIDUALBISS":      "个体工商户",
		"PERSON":              "小微个人商户",
		"SUBJECT_TYPE_OTHERS": "其他组织",
	}

	// SettleBankType 结算卡类型
	SettleBankType = map[string]string{
		"TOPRIVATE": "对私",
		"TOPUBLIC":  "对公",
	}

	// SettlementCycle 结算周期
	SettlementCycle = map[string]string{
		"T1": "工作日隔天结算",
		"D0": "当日结算",
		"D1": "自然日隔天结算",
	}

	// ProductType 产品类型
	ProductType = map[string]string{
		"APPPAY":      "扫码",
		"QUICKPAY":    "快捷",
		"PAYMENT":     "虚拟账户支付",
		"ACCOUNT_PAY": "转账充值",
		"SETTLEMENT":  "结算",
		"TRANSFER":    "代付",
	}

	// OnlineBankType 账户属性
	OnlineBankType = map[string]string{
		"B2C": "对私",
		"B2B": "对公",
	}

	// CardType 卡类型
	CardType = map[string]string{
		"DEBIT":  "借记卡",
		"CREDIT": "贷记卡",
	}

	// IntegralType 积分计费类型
	IntegralType = map[string]string{
		"DISCOUNTS": "优惠",
		"STANDARD":  "标准",
		"PUBLIC":    "公共",
	}

	// SettlementMode 结算方式
	SettlementMode = map[string]string{
		"NOTOPEN": "不开通",
		"AUTO":    "自动结算",
		"SELF":    "自主结算",
	}

	// ReportResult 报备结果
	ReportResult = map[string]string{
		"NOT_APPLY_YET": "尚未报备",
		"DOING":         "正在处理中",
		"SUCCESS":       "报备成功",
		"FAIL":          "报备失败",
	}

	// MerchantEntryAlterationType 变更类型
	MerchantEntryAlterationType = map[string]string{
		"SETTLE_BANKCARD":     "结算卡",
		"MERCHANT_INFO":       "商户信息",
		"MERCHANT_CREDENTIAL": "资质",
		"MERCHANT_INFO_SYN":   "商户信息变更查询(新接口)",
	}

	// AlterationStatus 变更状态
	AlterationStatus = map[string]string{
		"INIT":    "初始化",
		"WAIT":    "待审核",
		"REFUSE":  "审核拒绝",
		"AUDITED": "审核通过",
		"DOING":   "处理中",
	}

	// MerchantAgreementStatus 合同审核状态
	MerchantAgreementStatus = map[string]string{
		"INIT":    "初始化",
		"WAIT":    "待审核",
		"REFUSE":  "审核拒绝",
		"AUDITED": "审核通过",
	}

	// MerchantCredentialType 资质类型
	MerchantCredentialType = map[string]string{
		"UNIFIED_CODE_CERTIFICATE":                "三证合一营业执照",
		"BUSINESS_LICENSE":                        "营业执照",
		"ORG_CERTIFICATE":                         "组织机构证",
		"TAX_REG_CERTIFICATE":                     "税务登记证",
		"PERMIT_FOR_BANK_ACCOUNT":                 "开户许可证",
		"SIGN_BOARD":                              "门头照",      // 商户门头照
		"SIGN_BOARD_ACTIVITY":                     "门头照_活动报名", // 活动报名用
		"FRONT_OF_ID_CARD":                        "法人身份证正面",
		"BACK_OF_ID_CARD":                         "法人身份证反面",
		"AUTHORIZATION_FOR_SETTLEMENT":            "结算账户指定书", // 非法人结算进件，必传
		"HANDHELD_OF_ID_CARD":                     "手持身份证照",
		"HANDHELD_OF_BANK_CARD":                   "手持银行卡照正面",
		"CHECKOUT_COUNTER":                        "收银台照",     // 蓝海使用
		"INTERIOR_PHOTO":                          "室内照",      // 商户室内照
		"INTERIOR_PHOTO_ACTIVITY":                 "室内照_活动报名", // 活动报名用
		"BANK_CARD":                               "结算卡",
		"SETTLE_FRONT_OF_ID_CARD":                 "结算人身份证正面",
		"SETTLE_BACK_OF_ID_CARD":                  "结算人身份证反面",
		"ATTACHMENTINFO_1":                        "附件1", // 报名微信官方的教育食堂行业活动时,合作资质证明有多个图片请分别上传至附件1-4
		"ATTACHMENTINFO_2":                        "附件2",
		"ATTACHMENTINFO_3":                        "附件3",
		"ATTACHMENTINFO_4":                        "附件4",
		"ATTACHMENTINFO_5":                        "附件5",
		"ACTIVITY_CHECKOUT_COUNTER":               "活动收银台照", // 此字段仅供绿洲报名使用
		"SUBLEASE_CERTIFICATE":                    "转租证明",
		"FINANCE_ROOM":                            "财务室",
		"RUN_SCHOOL_LICENSE_PIC":                  "办学资质图片",
		"CHARGE_SAMPLE":                           "收费样本",
		"PRIVATE_NONENTERPRISE_UNITS":             "民办非企业单位登记证书图片",
		"DIPLOMATIC_NOTE":                         "照会",
		"CERTIFICATE_FILE":                        "证明文件图片",
		"MEDICAL_INSTRUMENT_PRACTICE_LICENSE_PIC": "医疗执业许可证图片",
		"UNIONPAY_CHECKOUT_COUNTER":               "银联标识收银台照",
		"UNIONPAY_SIGNBOARD":                      "银联标识门头照",
		"HUABEI_MATTER":                           "花呗物料",
		"ACTIVITY_MATTER":                         "支付宝天天活动物料",
		"UNIONPAY_AGREEMENT":                      "加盟协议",
		"UNIONPAY_TERMINAL":                       "银联轻型终端照",
		"UNIONPAY_STUFF":                          "银联物料回传照片",
		"UNIONPAY_DEVICE":                         "银联语音报备设备照片",
		"HANDHELD_OF_BANK_CARD_BACK":              "手持银行卡照反面",
		"ORGANIZATION_FOUND":                      "党组织成立文件",
		"INSTITUTIONAL_ORGANIZATION_PIC":          "法人登记证书",
		"ACTIVITY_RATE_COMMITMENT":                "优惠费率承诺函照", // 官方活动报名上传此资质
		"BUSINESS_CERTIFICATE":                    "经营许可证照",
		"FOOD_SERVICE_LIC":                        "餐饮服务许可证照",
		"FOOD_HEALTH_LIC":                         "食品卫生许可证照",
		"FOOD_BUSINESS_LIC":                       "食品经营许可证照",
		"FOOD_CIRCULATE_LIC":                      "食品流通许可证照",
		"FOOD_PRODUCTION_LIC":                     "食品生产许可证照",
		"TOBACCO_LIC":                             "烟草专卖零售许可证照",
		"MERCHANT_ENTERPRISE_MEALS_COOPERATION":   "商户与企业团餐合作协议(支付宝)", // 报名支付宝未来校园团餐活动时必传
	}

	// CollectMethod 产品手续费收取方式
	CollectMethod = map[string]string{
		"FEEACCOUNT":      "从平台商手续费账户收取",
		"OWN_CASHACCOUNT": "子商户资金账户",
	}

	// FeeRateEffectAt 费率生效时间
	FeeRateEffectAt = map[string]string{
		"NOW":      "即时生效",
		"TOMORROW": "次日凌晨生效",
	}

	// CredentialFileStatus 资质文件url上传状态类型
	CredentialFileStatus = map[string]string{
		"FAIL":    "失败",
		"SUCCESS": "成功",
		"DOING":   "处理中",
	}

	// SplittingStatus 分账状态
	SplittingStatus = map[string]string{
		"FAIL":    "失败",
		"SUCCESS": "成功",
		"DOING":   "处理中",
	}

	// ReportType 报备类型
	ReportType = map[string]string{
		"BlueOcean":      "蓝海绿洲",
		"PAYMENT":        "缴费类",
		"PUBLIC_WELFARE": "公益类",
		"OTHER":          "其他类",
		"ACADEMY":        "私立院校类",
		"INSURANCE":      "保险教育类",
		"ONLINE":         "线上类",
	}

	// BlueOceanEntryStatus 蓝海绿洲报名状态
	BlueOceanEntryStatus = map[string]string{
		"INIT":     "初始化", // 处理中
		"SUBMIT":   "待审核", // 待通道方审核
		"REJECT":   "拒绝",
		"PASS":     "通过",
		"ABNORMAL": "编辑中", // 处理中
	}

	// FeeMode 费率模式
	FeeMode = map[string]string{
		"DEFAULT": "默认模式", // 按百分比
		"RANGE":   "区间模式", // 区间固定金额， 注：1.商户计费模式若采用分段模式，那么此商户不允许切换到默认模式；2.开通分段费率，必须开通手续费账户，否则无法进行平台商补贴；
	}

	// SettleMode 结算模式
	SettleMode = map[string]string{
		"MERCHANT": "商户结算，按商户号维度结算",
		"MERGE":    "合并结算，按结算人结算，支持同服务商下，结算卡号一致的结算单合并成一笔出款",
	}

	// PosActiveStatus POS产品开通状态
	PosActiveStatus = map[string]string{
		"INIT":          "处理中", // 处理中
		"FAIL":          "失败",  // 各种原因开通失败，可再次申请开通
		"SUCCESS":       "成功",
		"CHECK_PENDING": "待审核",
		"DISMISSAL":     "审核驳回", // 审核驳回，可补充资料后再次申请开
	}

	// UploadResult 上传结果
	UploadResult = map[string]string{
		"UPLOADED":     "已上传",
		"NOT_UPLOADED": "未上传",
	}

	// FeeType 费率类型
	FeeType = map[string]string{
		"DEFAULT": "默认类型",
		// 以下类型费率仅支持分段计费类型
		"PUBLIC_WELFARE":    "公益类费率",
		"PAYMENT":           "缴费类费率",
		"ACADEMY":           "私立院校类",
		"INSURANCE":         "保险教育类",
		"ONLINE":            "线上类",
		"SCHOOLCANTEEN_001": "高校食堂行业活动",
		"EDUCATION_001":     "线下教培机构行业活动",
		"BLUESEA_1":         "支付宝新蓝海活动",
		"ALISCHOOL_1":       "支付宝未来校园活动",
		"MEDICAL_0":         "医疗卫生活动",
		"GOVERNMENT_0":      "政务组织活动",
		"SCHOOLCANTEEN_002": "教育行业学校主体食堂活动",
		"CAMPUSDINING_001":  "教育行业非学校主体商户餐饮活动",
		"PARKING_001":       "停车缴费活动",
		"HEALTHCARE_001":    "商业医疗行业活动",
		"VIRTUAL":           "虚拟类",
	}

	// MerchantStatus 商户状态
	MerchantStatus = map[string]string{
		"INIT":                 "入网中",
		"AUDITING":             "运营审核中",
		"AUDITERROR":           "运营审核未通过",
		"AVAILABLE":            "正常",
		"FROZEN":               "关闭",
		"CANCELLED":            "已注销",
		"UNDO_REG_SUCCESS":     "商户入网已撤销",
		"PASS_RISK_CONTROL":    "风控审核通过",
		"NOPASS_RISK_CONTROL":  "风控审核未通过",
		"REFUSED":              "运营拒绝",
		"REFUSED_RISK_CONTROL": "风控拒绝",
	}

	// AccountStatus 账号状态
	AccountStatus = map[string]string{
		"FREEZE_DEBIT":  "冻结出帐",
		"FREEZE_CREDIT": "冻结入账",
		"CANCELLED":     "已注销",
		"AVAILABLE":     "正常",
		"FROZEN":        "冻结",
		"NONE":          "未开通商户资金账户",
	}

	// MicroBizType 小微经营类型
	MicroBizType = map[string]string{
		"MICRO_TYPE_STORE":  "门店场所",
		"MICRO_TYPE_MOBILE": "流动经营/便民服务",
		"MICRO_TYPE_ONLINE": "线上商品/服务交易",
	}

	// CertType 证书类型
	CertType = map[string]string{
		"CERTIFICATE_TYPE_2388": "事业单位法人证书",
		"CERTIFICATE_TYPE_2389": "统一社会信用代码证书",
		"CERTIFICATE_TYPE_2390": "有偿服务许可证（军队医院适用）",
		"CERTIFICATE_TYPE_2391": "医疗机构执业许可证（军队医院适用）",
		"CERTIFICATE_TYPE_2392": "企业营业执照（挂靠企业的党组织适用）",
		"CERTIFICATE_TYPE_2393": "组织机构代码证（政府机关适用）",
		"CERTIFICATE_TYPE_2394": "社会团体法人登记证书",
		"CERTIFICATE_TYPE_2395": "民办非企业单位登记证书",
		"CERTIFICATE_TYPE_2396": "基金会法人登记证书",
		"CERTIFICATE_TYPE_2397": "慈善组织公开募捐资格证书",
		"CERTIFICATE_TYPE_2398": "农民专业合作社法人营业执照",
		"CERTIFICATE_TYPE_2399": "宗教活动场所登记证",
		"CERTIFICATE_TYPE_2400": "其他证书/批文/证明",
	}

	// AuthorizeStatus 授权状态
	AuthorizeStatus = map[string]string{
		"PRE_AUTHORIZED":  "待授权",
		"AUTHORIZED":      "已授权",
		"NOT_AUTHORIZED":  "未授权",
		"AUTHORIZED_FAIL": "授权失败",
	}

	// ConfirmStatus 认证/绑定状态
	ConfirmStatus = map[string]string{
		"PRE_CONFIRM":         "待认证",
		"CONFIRMING":          "审核中",
		"CONFIRM_CONTACT":     "待确认联系信息",
		"CONFIRM_LEGALPERSON": "待账户验证",
		"CONFIRM_PASSED":      "审核通过",
		"CONFIRM_REJECTED":    "审核驳回",
		"CONFIRM_CANCELED":    "已作废",
		"CONFIRM_FREEZED":     "已冻结",
	}

	// IdType 身份证件类型
	IdType = map[string]string{
		"IDCARD":                    "身份证",
		"PASSPORT":                  "护照",
		"SOLDIERSCERTIFICATE":       "士兵证",
		"OFFICERSCERTIFICATE":       "军官证",
		"GATXCERTIFICATE":           "香港居民来往内地通行证",
		"TWNDCERTIFICATE":           "台湾同胞来往内地通行证",
		"MACAOCERTIFICATE":          "澳门来往内地通行证",
		"FOREIGNERCERTIFICATE":      "外国人居留证",
		"HKANDMACAORESIDENCEPERMIT": "港澳居民证",
		"TAIWANRESIDENCEPERMIT":     "台湾居民证",
	}

	// ArrivalMode 结算到账方式
	ArrivalMode = map[string]string{
		"NORMAL":     "银行卡",
		"ELECTRONIC": "银行电子账户",
	}

	// ElectronicAccountStatus 电子账户鉴权状态
	ElectronicAccountStatus = map[string]string{
		"INIT":   "未鉴权",
		"DONE":   "鉴权成功",
		"FAILED": "鉴权失败",
	}

	// OnlinePayType 在线支付类型
	OnlinePayType = map[string]string{
		"ONLINE":         "网银",
		"ONLINETRANSFER": "网银代付",
	}

	// ActivityRegisterStatus 活动报名状态
	ActivityRegisterStatus = map[string]string{
		"INIT":     "初始化",
		"AUDITING": "审核中",
		"PASS":     "报名通过",
		"REJECT":   "审核驳回",
		"MODIFY":   "报名修改", // 微信只支持对审核驳回的进行修改
	}

	// ActivityId 活动标识id
	ActivityId = map[string]string{
		"SCHOOLCANTEEN_001": "高校食堂行业活动",   // 微信
		"EDUCATION_001":     "线下教培机构行业活动", // 微信
		"BLUESEA_1":         "新蓝海活动",      // 支付宝
		"ALISCHOOL_1":       "支付宝未来校园活动",  // 支付宝
	}

	// ActivityScene 活动场景
	ActivityScene = map[string]string{
		"BLUE_SEA_FOOD_INDIRECT_APPLY": "支付宝新蓝海活动,间连餐饮",
		"BLUE_SEA_FMCG_INDIRECT_APPLY": "支付宝新蓝海活动,间连快消",
	}

	// AppIdConfigurationStatus AppId配置状态
	AppIdConfigurationStatus = map[string]string{
		"APPID_PATH_SUBSCRIBEAPPID_INIT":    "支付appId配置成功，授权目录配置成功，关注appId未配置",
		"APPID_PATH_SUBSCRIBEAPPID_FAIL":    "支付appId配置成功，授权目录配置成功，关注appId配置失败",
		"APPID_PATH_SUBSCRIBEAPPID_SUCCESS": "完全配置成功",
		"APPID_INIT":                        "支付appId未配置",
		"APPID_FAIL":                        "支付appId配置失败",
		"APPID_PATH_INIT":                   "支付appId配置成功，授权目录未配置",
		"APPID_PATH_FAIL":                   "支付appId配置成功，授权目录配置失败",
	}

	// ContactType 联系人类型
	ContactType = map[string]string{
		"LEGAL": "经营者/法人",
		"SUPER": "经办人",
	}

	// AlipayServiceCode 支付宝商户申请服务类型
	AlipayServiceCode = map[string]string{
		"F2F":     "当面付", // 对应线下
		"PRE_F2F": "线下预授权",
	}

	ContractSignStatus = map[string]string{
		"UNSIGNED":   "未签约",
		"SIGNED":     "已签约",
		"NO_PRODUCE": "未生成协议",
	}
)
