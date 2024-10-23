# 通过配置文件生成对应go结构体

    1. go install github.com/ChimeraCoder/gojson/gojson
    2. gojson -fmt=yaml -input=./config/config.yaml -name=Conf -o=./config/conf.go -pkg=config -subStruct=false -tags=json

    参数说明：    
    fmt         表示类型，默认json，只支持yaml和json两种
    input       表示配置文件路径，推荐使用相对路径。文件不存在会报错
    name        表示生成的结构体的名称
    o           表示生成的文件路径。文件目录必须存在，文件会自动创建
    pkg         表示生成文件的包名
    subStruct   表示是否生成子结构体，默认false，推荐false
    tags        表示生成生成的结构体tag标签名，默认为yaml，推荐改为json