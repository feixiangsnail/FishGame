package Dao

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"Demo/login2/UserModel"
)

var lof = fmt.Println
type DbInfo struct{
	Db *sql.DB
}

func (this *DbInfo) InitDatabase(){
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	var err error
	this.Db,err = sql.Open("mysql", "root:asaiw48!skdhw1223@tcp(172.104.32.98:3306)/chess?charset=utf8")
	if err!=nil{
		log.Fatal(err.Error())
	}
}

func (this *DbInfo) FindUserPwd(username string,password string) (tempUser *User_Module.UserModel){
	var id ,coins int

	err := this.Db.QueryRow("select id,coins from fish_user where username=? and password=?", username,password).Scan(&id,&coins)

	lof(username,password,err,tempUser,"et")
	if err != sql.ErrNoRows {
		tempUser =&User_Module.UserModel{
			ID:id,
			Username:username,
			Password:password,
			Coins:coins,
		}
	}else{
		tempUser = nil
		log.Println("用户或者密码不正确")

	}
	return
}
func (this *DbInfo) FindUser(username string) (isExsit bool){
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
func (this *DbInfo)  CreateUser(user *User_Module.UserModel) (isSuccess bool){
	_, err := this.Db.Exec("insert into fish_user(username,password,coins) values(?,?,?)", user.Username, user.Password,1000)
	if err!=nil{
		log.Println(err,"创建用户失败")
		isSuccess = false
	}else{
		lof("创建用户成功")
		isSuccess = true
	}
	return
}
func checkError(err error){
	if err!=nil{
		log.Fatal(err)
	}
}