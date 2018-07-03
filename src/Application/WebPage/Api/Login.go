package WebPageApi

import (
	"net/http"

	"Lib/Tool"
	"fmt"

	"Drive/MySQL/Model"


)

var lof = fmt.Println



func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	username := r.FormValue("username")
	password := r.FormValue("password")
	password = Tool_Lib.MdFormat(password)

	var mySQLModel =Drive_Mysql.GetSqlDb()
	user := mySQLModel.FindUserPwd(username, password)



	if user != nil {
		Tool_Lib.SendBack(200, "登录成功2", w, user)
	} else {
		Tool_Lib.SendBack(0, "用户名或密码错误", w, user)
	}

}
