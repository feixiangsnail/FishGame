package WebPageApi

import (
	"net/http"
	"Lib/Tool"
	"Drive/MySQL/Model"
	"Application/User/Model"
)

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	username := r.FormValue("username")
	password := r.FormValue("password")
	password = Tool_Lib.MdFormat(password)
	var mySQLModel =Drive_Mysql.GetSqlDb()
	isExist:=mySQLModel.FindUser(username)
	if isExist{
		Tool_Lib.SendBack(0,"用户已存在",w,nil)

	}else{
		user:=&User_Module.UserModel{
			Username:username,
			Password:password,
		}
		if mySQLModel.CreateUser(user){
			Tool_Lib.SendBack(200,"注册成功",w,nil)
		}else{
			Tool_Lib.SendBack(0,"注册失败",w,nil)
		}
	}

}
