package language

var StringMap map[string]string

func init() {
	StringMap = map[string]string{
		"str_ca_ba":   "保安",
		"str_ca_xcsf": "修车师傅",
		"str_ca_gjsj": "公交司机",
		"str_ca_qyms": "企业秘书",
		"str_ca_jc":   "警察",
		"str_ca_yyhs": "医院护士",
		"str_ca_zxjs": "中学教师",
		"str_ca_jl":   "企业经理",
		"str_ca_gcs":  "工程师",
		"str_ca_ls":   "律师",
		"str_ca_jz":   "飞机机长",
		"str_ca_ys":   "主任医生",
		"str_ca_dkzz": "住宅抵押贷款",
		"str_ca_dkjy": "教育贷款",
		"str_ca_dkgc": "购车贷款",
		"str_ca_dkxy": "信用卡",
		"str_ca_dkew": "额外负债",
		"str_ca_zcsj": "税金",
		"str_ca_zczz": "住房抵押贷款/房租",
		"str_ca_zcjy": "教育贷款",
		"str_ca_zcgc": "购车贷款",
		"str_ca_zcxy": "信用卡支出",
		"str_ca_zcew": "额外支出",
		"str_ca_zcot": "其他支出",
		"str_ca_zcxh": "小孩支出",
		"str_ca_gxdk": "高息贷款支出",
	}
}

func Get(key string) string {
	if _, ok := StringMap[key]; ok {
		return StringMap[key]
	}

	return ""
}
