package Drive

import (
	"reflect"
	"Drive/Base/Interface"
	"fmt"
	"Lib/Tool"
)
var lof = fmt.Println
//驱动配置文件
var data []map[string]interface{}
//所有驱动
var allDriveType = map[string]reflect.Type{}
//所有驱动实体
var allDrive = map[string]Drive_Base.Drive_Interface{}

func Init(URL string){

	Base:=reflect.New(allDriveType["Disk"]).Interface().(Drive_Base.Drive_Interface)
	Base.InitData(map[string]interface{}{"Name":"Run","Url":URL,})
	data = make([]map[string]interface{},0)
	Base.GetObject("config.json",&data)
	for _,v:=range data{
		InitDrive(v)
	}
}
func InitDrive(Data map[string]interface{}){
	DriveItem:=allDriveType[Tool_Lib.Convert_String(Data["Drive"])]
	if DriveItem !=nil{
		DriveItemBase:= reflect.New(DriveItem).Interface().(Drive_Base.Drive_Interface)
		DriveItemBase.InitData(Data)
		allDrive[Tool_Lib.Convert_String(Data["Drive"])] = DriveItemBase
	}else{
		//lof(Data,"Data")
	}
}
func Register(Name string,Drive reflect.Type){
	allDriveType[Name] = Drive
}

