package enum

type Item struct {
	Name       string      // 变量名
	Val        interface{} // 映射值
	Annotation interface{} // 注释
}

type GenEnum struct {
	PkgName  string
	EnumPath string
	Data     []Item
}
