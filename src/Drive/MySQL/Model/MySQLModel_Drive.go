package Drive_Http

import (
	"fmt"
	"Drive"
	"reflect"
	"Lib/Tool"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

var lof = fmt.Println

type MySQLModel struct{
	Name string `json:"Name"`
	UserName string `json:"UserName"`
	UserPassword string `json:"UserPassword"`
	DataBase string `json:"DataBase"`
	IP string `json:"IP"`
	Port string `json:"Port"`
	Pre string `json:"Pre"`
}
func init(){
	Drive.Register("MySQL",reflect.TypeOf(MySQLModel{}))
	lof("加载HTTP驱动成功")
}
func (this *MySQLModel) InitData(data map[string]interface{}){
	this.Name = Tool_Lib.Convert_String(data["Name"])
	this.UserName = Tool_Lib.Convert_String(data["UserName"])
	this.UserPassword = Tool_Lib.Convert_String(data["UserPassword"])
	this.DataBase = Tool_Lib.Convert_String(data["DataBase"])
	this.IP = Tool_Lib.Convert_String(data["IP"])
	this.Port = Tool_Lib.Convert_String(data["Port"])
	this.Init()
}

func (this *MySQLModel) Init(){
	this.initMysql()
}
func (this *MySQLModel) initMysql(){
	var databaseInfo = this.UserName+":"+this.UserPassword+"@tcp("+this.IP+":"+this.Port+")/"+this.DataBase+"?charset=utf8"
	db, err := sql.Open("mysql", databaseInfo)
	lof(db,err)
}





//获取单挑数据
func (this *MySQLModel) GetObject(Url string,Data interface{}) interface{}{
	return nil
}
//获取文本数据
func (this *MySQLModel) GetContent(Url string,)string{
	return "mysqlstring"
}
//写入数据
func (this *MySQLModel) Set(Url string ,Data interface{}) int{
	return 0
}