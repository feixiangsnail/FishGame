package Drive_Disk

import (
	"Drive"
	"reflect"
	"Lib/Tool"
	"io/ioutil"
	"log"
)

type Disk struct{
	Name string `json:"Name" sql:"Disk_Name"`
	Root string `json:"Root" sql:"Disk_Root"`
}

func init(){
	Drive.Register("Disk",reflect.TypeOf(Disk{}))

}
func (this *Disk)InitData(Data map[string]interface{}){
	this.Name = Tool_Lib.Convert_String(Data["Name"])
	this.Root = Tool_Lib.Convert_String(Data["Url"])

}
//获取单挑数据
func (this *Disk) GetObject(Url string,Data interface{}) interface{}{
	return Tool_Lib.Json_Object(this.GetContent(Url),Data)

}
//获取文本数据
func (this *Disk)  GetContent(Url string,)string{
	Url = this.GetUrl(Url)
	dat,err:=ioutil.ReadFile(Url)
	Tool_Lib.Message_Check(err)
	return string(dat)

}
//转换url
func (this *Disk) GetUrl(URL string) string{
	log.Println(this.Root+URL)
	return this.Root+URL
}
//写入数据
func (this *Disk) Set(Url string ,Data interface{}) int{
	return 1
}