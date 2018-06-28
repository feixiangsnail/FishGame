package main

import (
	"fmt"
	"net/http"
	"io"
	"database/sql"
	"crypto/md5"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
)

var lof = fmt.Println

func main() {


	mux:=http.NewServeMux()
	mux.HandleFunc("/api/login",login)
	mux.HandleFunc("/api/register",resister)
	http.ListenAndServe(":8081",mux)
}
type result struct{
	Code int
	Msg string
	Data []string
}

type PostData struct{
	user_name string
	pass_word string
}
type FilterData interface {
	formatData() PostData
}
func (post_data PostData) formatData() PostData{
	post_data.pass_word = mdFormat(post_data.pass_word)
	return post_data
}
func mdFormat(data string) string{
	t:=md5.New()
	io.WriteString(t,data)
	return fmt.Sprintf("%x",t.Sum(nil))
}
func login(w http.ResponseWriter,r *http.Request){

//	w.Header().Set("Access-Control-Allow-Origin","*")
	r.ParseForm()
	username,found1:=r.Form["username"]
	password,found2:=r.Form["password"]
	if !(found1&&found2){
		io.WriteString(w,"参数不正确")
	}
	db,err:=sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err!=nil{
		lof(err,"数据库连接失败")
	}
	defer db.Close()
	var post_data PostData
	post_data.user_name = username[0]
	post_data.pass_word = password[0]
	var filter_data FilterData = post_data
	post_data = filter_data.formatData()
	lof(post_data,"postData")
	var user_name ,pass_word string
	var id int

	err = db.QueryRow("select * from users where id = ?",1).Scan(&id,&user_name,&pass_word)
	lof(err,"err")

	lof(id,user_name,pass_word,"idusernamepassword")
	if err!=nil{
		arr:=&result{
			500,
			"登录失败",
			[]string{},
		}
		b,json_err:=json.Marshal(arr)
		if json_err !=nil{
			lof("encoding faild")
		}else{
			io.WriteString(w,string(b))
		}



	}else{
		arr:=&result{
			200,
			"登录成功",
			[]string{},
		}
		b,json_err:=json.Marshal(arr)
		if json_err!=nil{
			lof("encoding faild")
		}else{
			io.WriteString(w,string(b))
		}
	}







}
func resister(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseForm()
	username, found1 := r.Form["username"]
	password, found2 := r.Form["password"]
	if !(found1&&found2){
		io.WriteString(w,"参数不正确")
	}
	db,err:=sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err!=nil{
		lof(err,"数据库连接失败")
	}
	defer db.Close()
	var post_data PostData
	post_data.user_name = username[0]
	post_data.pass_word = password[0]
	var filter_data FilterData = post_data
	post_data = filter_data.formatData()
	_,err = db.Exec("insert into users(username,password) values(?,?)",post_data.user_name,post_data.pass_word)
	if err != nil {
		arr := &result{
			500,
			"注册失败",
			[]string{},
		}
		b, json_err := json.Marshal(arr) //json化结果集
		if json_err != nil {
			fmt.Println("encoding faild")
		} else {
			io.WriteString(w, string(b)) //返回结果
		}
	} else {
		arr := &result{
			200,
			"注册成功",
			[]string{},
		}
		b, json_err := json.Marshal(arr) //json化结果集
		if json_err != nil {
			fmt.Println("encoding faild")
		} else {
			io.WriteString(w, string(b)) //返回结果
		}
	}
}