package RpStatusEnum

import (
	"encoding/json"
	"errors"

	"github.com/liubai047/generate/enum"
)

var _ json.Marshaler = &Enum{}
var _ json.Unmarshaler = &Enum{}

type Enum struct {
	enum.ValString
	enum.AntString
}

var Success = &Enum{"sc01", "成功咯"}
var Fail = &Enum{"fl01", "失败咯"}
var Wait = &Enum{"wt01", "待拉咯"}
var Deal = &Enum{"el01", "待处理咯"}

// GetVal 获取枚举值
func (r *Enum) GetVal() string {
	return r.ValString.String()
}

// GetAnnotation 获取枚举注释
func (r *Enum) GetAnnotation() string {
	return r.AntString.String()
}

// MarshalJSON 实现MarshalJSON接口
func (r *Enum) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.GetVal())
}

// UnmarshalJSON 实现UnmarshalJSON接口
func (r *Enum) UnmarshalJSON(data []byte) error {
	var val string
	err := json.Unmarshal(data, &val)
	if err != nil {
		return errors.New("RpStatusEnum Unmarshal error: " + err.Error())
	}
	res := Convert(val)
	if res == nil {
		return errors.New("RpStatusEnum enum value error")
	}
	*r = *res
	return nil
}

// Convert 将数据转换成类型，如果该数据值未定义，则直接转换为对应空值
//
// 判定空值可以使用Convert(xx) == nil
func Convert(s string) *Enum {
	switch s {
	case "sc01":
		return Success
	case "fl01":
		return Fail
	case "wt01":
		return Wait
	case "el01":
		return Deal
	}
	return nil
}

// Maps 获取枚举值和枚举注释组成的map
func Maps() map[string]string {
	return map[string]string{
		"sc01": "成功咯",
		"fl01": "失败咯",
		"wt01": "待拉咯",
		"el01": "待处理咯",
	}
}

// AntMaps 获取枚举注释和枚举值组成的map
func AntMaps() map[string]string {
	return map[string]string{
		"成功咯":  "sc01",
		"失败咯":  "fl01",
		"待拉咯":  "wt01",
		"待处理咯": "el01",
	}
}

// Options 获取枚举注释和枚举值组成的下拉列表
func Options() []map[string]interface{} {
	return []map[string]interface{}{
		{"label": "成功咯", "value": "sc01"},
		{"label": "失败咯", "value": "fl01"},
		{"label": "待拉咯", "value": "wt01"},
		{"label": "待处理咯", "value": "el01"},
	}
}
