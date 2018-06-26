package Drive_Base
type Drive_Interface interface{
InitData(map[string]interface{})
//获取单挑数据
GetObject(Url string,Data interface{}) interface{}
//获取文本数据
GetContent(Url string,)string
//写入数据
Set(Url string ,Data interface{}) int
}