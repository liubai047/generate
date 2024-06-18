package test

import "generate/enum"

//go:generate generate --struct_file=gen.go --struct_name=tsRpStsEnum
var tsRpStsEnum = enum.GenEnum{
	PkgName:  "RpStatusEnum",
	EnumPath: "generate/enum",
	Data: []enum.Item{
		{Name: "Success4", Val: "assc", Annotation: "成功咯"},
		{Name: "Fail3", Val: "fffc", Annotation: "失败咯"},
		{Name: "Wait2", Val: "wwwc", Annotation: "待拉咯"},
		{Name: "Deal1", Val: "dddc", Annotation: "待处理咯"},
	},
}
