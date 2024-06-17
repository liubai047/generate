package enum

type ReportStatusEnum struct {
	ValString
	AntString
}

var ReportStatusSuccess = ReportStatusEnum{"success", "成功"}
var ReportStatusFail = ReportStatusEnum{"fail", "失败"}
var ReportStatusWaitPull = ReportStatusEnum{"wait_pull", "等待拉取"}
var ReportStatusWaitDeal = ReportStatusEnum{"wait_deal", "等待处理"}

func (r ReportStatusEnum) GetVal() string {
	return r.ValString.String()
}

func (r ReportStatusEnum) GetAnnotation() string {
	return r.AntString.String()
}

// Convert 将数据转换成类型，如果该数据值未定义，则直接转换为对应空值
func Convert(s string) ReportStatusEnum {
	switch s {
	case "success":
		return ReportStatusSuccess
	case "fail":
		return ReportStatusFail
	case "wait_pull":
		return ReportStatusWaitPull
	case "wait_deal":
		return ReportStatusWaitDeal
	}
	return ReportStatusEnum{}
}

func Maps() map[string]string {
	return map[string]string{
		"success":   "成功",
		"fail":      "失败",
		"wait_pull": "等待拉取",
		"wait_deal": "等待处理",
	}
}

func AntMaps() map[string]string {
	return map[string]string{
		"成功":   "success",
		"失败":   "fail",
		"等待拉取": "wait_pull",
		"等待处理": "wait_deal",
	}
}

func Options() []map[string]string {
	return []map[string]string{
		{"label": "成功", "value": "success"},
		{"label": "失败", "value": "fail"},
		{"label": "等待拉取", "value": "wait_pull"},
		{"label": "等待处理", "value": "wait_deal"},
	}
}
