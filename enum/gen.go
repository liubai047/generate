package enum

type item struct {
	Name       string      // 变量名
	Val        interface{} // 映射值
	Annotation interface{} // 注释
}

type GenEnum struct {
	PkgName  string
	EnumPath string
	Data     []item
}

var genReportStatus = GenEnum{
	PkgName:  "ReportStatus",
	EnumPath: "generate/enum",
	Data: []item{
		{Name: "Success", Val: "sss", Annotation: "成功咯"},
		{Name: "Fail", Val: "fff", Annotation: "失败咯"},
		{Name: "Wait", Val: "www", Annotation: "待拉咯"},
		{Name: "Deal", Val: "ddd", Annotation: "待处理咯"},
	},
}
