package enum

// 该文件用来展示最终生成的
import (
	"encoding/json"
)

var _ json.Marshaler = ReportStatusEnum{}
var _ json.Unmarshaler = ReportStatusEnum{}

type ReportStatusEnum struct {
	ValString
	AntString
}

var ReportStatusSuccess = ReportStatusEnum{"success", "成功"}
var ReportStatusFail = ReportStatusEnum{"fail", "失败"}
var ReportStatusWaitPull = ReportStatusEnum{"start", "开始"}
var ReportStatusWaitDeal = ReportStatusEnum{"stop", "停止"}

// GetVal 获取枚举值
func (r ReportStatusEnum) GetVal() string {
	return r.ValString.String()
}

// GetAnnotation 获取枚举注释
func (r ReportStatusEnum) GetAnnotation() string {
	return r.AntString.String()
}

// MarshalJSON 实现MarshalJSON接口
func (r ReportStatusEnum) MarshalJSON() ([]byte, error) {
	return []byte(r.GetVal()), nil
}

// UnmarshalJSON 实现UnmarshalJSON接口
func (r ReportStatusEnum) UnmarshalJSON(data []byte) error {
	type Temp ReportStatusEnum
	tmp := &struct {
		Temp
	}{
		Temp: Temp(r),
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	return nil
}

// Convert 将数据转换成类型，如果该数据值未定义，则直接转换为对应空值
//
// 判定空值可以使用Convert(xx) == xx.ReportStatusEnum{}
func Convert(s string) ReportStatusEnum {
	switch s {
	case "success":
		return ReportStatusSuccess
	case "fail":
		return ReportStatusFail
	case "start":
		return ReportStatusWaitPull
	case "stop":
		return ReportStatusWaitDeal
	}
	return ReportStatusEnum{}
}

// Maps 获取枚举值和枚举注释组成的map
func Maps() map[string]string {
	return map[string]string{
		"success": "成功",
		"fail":    "失败",
		"start":   "开始",
		"stop":    "停止",
	}
}

// AntMaps 获取枚举注释和枚举值组成的map
func AntMaps() map[string]string {
	return map[string]string{
		"成功": "success",
		"失败": "fail",
		"开始": "start",
		"停止": "stop",
	}
}

// Options 获取枚举注释和枚举值组成的下拉列表
func Options() []map[string]string {
	return []map[string]string{
		{"label": "成功", "value": "success"},
		{"label": "失败", "value": "fail"},
		{"label": "开始", "value": "start"},
		{"label": "停止", "value": "stop"},
	}
}
