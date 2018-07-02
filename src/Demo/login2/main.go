package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"io"
	"Demo/login2/Dao"
	"encoding/json"
	"crypto/md5"
	"Demo/login2/UserModel"
	"Demo/login2/session"
)

var lof = fmt.Println
var m = map[string]string{"a":"a1"}
var dbInfo =Dao.DbInfo{}
type result struct {
	Code int
	Msg  string
	Data interface{}
}
func init(){
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	dbInfo.InitDatabase()

}
func initApiRouter(router *mux.Router){
	router.HandleFunc("/",home).Methods("GET")
	apiRouter:=router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/userinfo", userinfo).Methods("POST")
	apiRouter.HandleFunc("/logout", logout).Methods("POST")
	apiRouter.HandleFunc("/login",login).Methods("POST")
	apiRouter.HandleFunc("/register",register).Methods("POST")
}
func main(){
	var router = mux.NewRouter()
	initApiRouter(router)
	log.Println("Server is running at http://localhost:8083")
	log.Fatal(http.ListenAndServe(":8083",router))

}

func home(w http.ResponseWriter,r *http.Request){
	sess := session.GetSession(w, r)
	user, exist := sess.GetAttr("user")


	if !exist{
		sendBack(0,"请登录",w,user)

	}else{

		sendBack(200,"用户数据",w,user)
	}

}
func userinfo(w http.ResponseWriter,r *http.Request){

}
func logout(w http.ResponseWriter,r *http.Request){
	lof(w,r,"logout")
}
func login(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//http.Redirect(w, r, "/", 302)
	username := r.FormValue("username")
	password := r.FormValue("password")
	password = mdFormat(password)
	user:=dbInfo.FindUserPwd(username,password)
	if user !=nil{
		sendBack(200,"登录成功2",w,user)
		sess := session.GetSession(w, r)
		sess.SetAttr("user", user)
	}else{
		sendBack(0,"用户名或密码错误",w,user)
	}


}
func register(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//http.Redirect(w, r, "/", 302)
	username := r.FormValue("username")
	password := r.FormValue("password")
	password = mdFormat(password)
	isExist:=dbInfo.FindUser(username)
	if isExist{
		sendBack(0,"用户已存在",w,nil)

	}else{
		user:=&User_Module.UserModel{
				Username:username,
				Password:password,
		}
		if dbInfo.CreateUser(user){
			sendBack(200,"注册成功",w,nil)
		}else{
			sendBack(0,"注册失败",w,nil)
		}
	}
}

func checkError(err error){
	if err!=nil{
		log.Fatal(err)
	}
}

func sendBack(c int, s string, w http.ResponseWriter,d interface{}) {
	arr := &result{
		c,
		s,
		d,
	}

	b, json_err := json.Marshal(arr) //json化结果集

	if json_err != nil {
		fmt.Println("encoding faild")
	} else {
		io.WriteString(w, string(b)) //返回结果
	}
}
func mdFormat(data string) string {
	t := md5.New()
	io.WriteString(t, data)
	io.WriteString(t, string(t.Sum(nil))+"newcode1")
	io.WriteString(t, string(t.Sum(nil))+"newcode2")
	return fmt.Sprintf("%x", t.Sum(nil))
}