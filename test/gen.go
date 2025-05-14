package test

import "github.com/liubai047/generate/enum"

//go:generate go run github.com/liubai047/generate --struct_file=gen.go --struct_name=tsRpStsEnum
var tsRpStsEnum = enum.GenEnum{
	PkgName:  "RpStatusEnum",
	EnumPath: "github.com/liubai047/generate/enum",
	Data: []enum.Item{
		{Name: "Success", Val: "sc01", Annotation: "成功咯"},
		{Name: "Fail", Val: "fl01", Annotation: "失败咯"},
		{Name: "Wait", Val: "wt01", Annotation: "待拉咯"},
		{Name: "Deal", Val: "el01", Annotation: "待处理咯"},
	},
}
