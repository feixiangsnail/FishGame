package Drive_Mysql

import (
	"fmt"

	"reflect"
	"Lib/Tool"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"log"
	"Application/User/Model"
	"Drive"
	"Lib/Service"
)

var lof = fmt.Println
var sqlDb *MySQLModel
type MySQLModel struct{
	Name string `json:"Name"`
	UserName string `json:"UserName"`
	UserPassword string `json:"UserPassword"`
	DataBase string `json:"DataBase"`
	IP string `json:"IP"`
	Port string `json:"Port"`
	Pre string `json:"Pre"`
	Db *sql.DB
}
func GetSqlDb() *MySQLModel{

	if sqlDb == nil{
		sqlDb = &MySQLModel{}
		sqlDb.Init()
	}
	return sqlDb
}

func init(){
	Drive.Register("MySQL",reflect.TypeOf(MySQLModel{}))

}
var Name,UserName,UserPassword,DataBase,IP,Port string
func (this *MySQLModel) InitData(data map[string]interface{}){

	Name = Tool_Lib.Convert_String(data["Name"])
	UserName = Tool_Lib.Convert_String(data["UserName"])
	UserPassword = Tool_Lib.Convert_String(data["UserPassword"])
	DataBase = Tool_Lib.Convert_String(data["DataBase"])
	IP = Tool_Lib.Convert_String(data["IP"])
	Port = Tool_Lib.Convert_String(data["Port"])

}

func (this *MySQLModel) Init(){
	this.initMysql()
}
func (this *MySQLModel) initMysql(){
	var databaseInfo = UserName+":"+UserPassword+"@tcp("+IP+":"+Port+")/"+DataBase+"?charset=utf8"
	var err error

	this.Db, err = sql.Open("mysql", databaseInfo)

	if err!=nil{
		log.Fatal(err.Error())
	}
}

//获取用户名和密码
func (this *MySQLModel) FindUserPwd(username string,password string) (tempUser *User_Module.UserModel){


	var id ,coins int

	err := this.Db.QueryRow("select id,coins from fish_user where username=? and password=?", username,password).Scan(&id,&coins)

	if err != sql.ErrNoRows {
		var token = Service_Lib.NewClientID()
		tempUser =&User_Module.UserModel{
			ID:id,
			Username:username,
			Coins:coins,
			Token:token,
		}
		this.UpdateToken(token,username,password)
	}else{
		tempUser = nil
		log.Println("用户或者密码不正确")

	}
	return
}

//登录之后更新用户token
func (this *MySQLModel) UpdateToken(token string,username string,password string) (isTrue bool){
	log.Println("更新token")
	_, err := this.Db.Exec(`UPDATE fish_user SET token=? WHERE username=? and password=?`,token,username,password)
	Tool_Lib.Message_Check(err)
	return
}
//查找用户是否存在
func (this *MySQLModel) FindUser(username string) (isExsit bool){
	var tempId int
	err := this.Db.QueryRow("select id from fish_user where username=?", username).Scan(&tempId)

	if err != sql.ErrNoRows {
		log.Println("用户已存在")
		isExsit = true
	}else{
		isExsit = false


	}
	return
}
//创建用户
func (this *MySQLModel)  CreateUser(user *User_Module.UserModel) (isSuccess bool){
	_, err := this.Db.Exec("insert into fish_user(username,password,coins) values(?,?,?)", user.Username, user.Password,10000)
	if err!=nil{
		log.Println(err,"创建用户失败")
		isSuccess = false
	}else{
		lof("创建用户成功")
		isSuccess = true
	}
	return
}

//通过token获取获取用户名和密码
func (this *MySQLModel) FindUserByToken(token string) (tempUser *User_Module.UserModel){

	var username  string
	var id ,coins int

	err := this.Db.QueryRow("select username,id,coins from fish_user where token=? ", token).Scan(&username,&id,&coins)

	if err != sql.ErrNoRows {
		var token = Service_Lib.NewClientID()
		tempUser =&User_Module.UserModel{
			ID:id,
			Username:username,
			Coins:coins,
			Token:token,
		}

	}else{
		tempUser = nil
		log.Println("查不到此token用户")

	}
	return
}

//通过token更新用户金币
func (this *MySQLModel) UpdateCoinsByToken(token string,getCoin int) (isTrue bool){
	log.Println("通过token更新用户金币")
	_, err := this.Db.Exec(`UPDATE fish_user SET coins = coins+? WHERE token=?`,getCoin,token)
	Tool_Lib.Message_Check(err)
	return
}


func checkError(err error){
	if err!=nil{
		log.Fatal(err)
	}
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
