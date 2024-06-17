package ReportStatus

import "generate/enum"

type Enum struct {
	enum.ValString
	enum.AntString
}

var Success = Enum{"sss", "成功咯"}
var Fail = Enum{"fff", "失败咯"}
var Wait = Enum{"www", "待拉咯"}
var Deal = Enum{"ddd", "待处理咯"}

func (r Enum) GetVal() string {
	return r.ValString.String()
}

func (r Enum) GetAnnotation() string {
	return r.AntString.String()
}

// Convert 将数据转换成类型，如果该数据值未定义，则直接转换为对应空值
func Convert(s string) Enum {
	switch s {
	case "sss":
		return Success
	case "fff":
		return Fail
	case "www":
		return Wait
	case "ddd":
		return Deal
	}
	return Enum{}
}

func Maps() map[string]string {
	return map[string]string{
		"sss": "成功咯",
		"fff": "失败咯",
		"www": "待拉咯",
		"ddd": "待处理咯",
	}
}

func AntMaps() map[string]string {
	return map[string]string{
		"成功咯":  "sss",
		"失败咯":  "fff",
		"待拉咯":  "www",
		"待处理咯": "ddd",
	}
}

func Options() []map[string]interface{} {
	return []map[string]interface{}{
		{"label": "成功咯", "value": "sss"},
		{"label": "失败咯", "value": "fff"},
		{"label": "待拉咯", "value": "www"},
		{"label": "待处理咯", "value": "ddd"},
	}
}
