package RpStatusEnum

import (
	"encoding/json"

	"github.com/liubai047/generate/enum"
)

var _ json.Marshaler = Enum{}
var _ json.Unmarshaler = Enum{}

type Enum struct {
	enum.ValString
	enum.AntString
}

var Success4 = Enum{"assc", "成功咯"}
var Fail3 = Enum{"fffc", "失败咯"}
var Wait2 = Enum{"wwwc", "待拉咯"}
var Deal1 = Enum{"dddc", "待处理咯"}

// GetVal 获取枚举值
func (r Enum) GetVal() string {
	return r.ValString.String()
}

// GetAnnotation 获取枚举注释
func (r Enum) GetAnnotation() string {
	return r.AntString.String()
}

// MarshalJSON 实现MarshalJSON接口
func (r Enum) MarshalJSON() ([]byte, error) {
	return []byte(r.GetVal()), nil
}

// UnmarshalJSON 实现UnmarshalJSON接口
func (r Enum) UnmarshalJSON(data []byte) error {
	type Temp Enum
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
func Convert(s string) Enum {
	switch s {
	case "assc":
		return Success4
	case "fffc":
		return Fail3
	case "wwwc":
		return Wait2
	case "dddc":
		return Deal1
	}
	return Enum{}
}

// Maps 获取枚举值和枚举注释组成的map
func Maps() map[string]string {
	return map[string]string{
		"assc": "成功咯",
		"fffc": "失败咯",
		"wwwc": "待拉咯",
		"dddc": "待处理咯",
	}
}

// AntMaps 获取枚举注释和枚举值组成的map
func AntMaps() map[string]string {
	return map[string]string{
		"成功咯":  "assc",
		"失败咯":  "fffc",
		"待拉咯":  "wwwc",
		"待处理咯": "dddc",
	}
}

// Options 获取枚举注释和枚举值组成的下拉列表
func Options() []map[string]interface{} {
	return []map[string]interface{}{
		{"label": "成功咯", "value": "assc"},
		{"label": "失败咯", "value": "fffc"},
		{"label": "待拉咯", "value": "wwwc"},
		{"label": "待处理咯", "value": "dddc"},
	}
}
