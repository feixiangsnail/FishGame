package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"io"
	"log"
	"net/http"
)

var lof = fmt.Println
var store = sessions.NewCookieStore([]byte("itissecret"))

func main() {
	lof("hello go is running")
	router := mux.NewRouter()
	//router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("src/web/"))))

	initApiRouter(router)
	log.Fatal(http.ListenAndServe(":8081", router))
}

func initApiRouter(router *mux.Router) {
	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/login", login).Methods("POST")
	apiRouter.HandleFunc("/register", resister).Methods("POST")
}

type PostData struct {
	user_name string
	pass_word string
}
type FilterData interface {
	formatData() PostData
}

//func mdFormat(data string) string{
//	t:=md5.New()
//
//	io.WriteString(t,data)
//	return fmt.Sprintf("%x",t.Sum(nil))
//}

func mdFormat(data string) string {
	t := md5.New()
	io.WriteString(t, data)
	io.WriteString(t, string(t.Sum(nil))+"newcode1")
	io.WriteString(t, string(t.Sum(nil))+"newcode2")
	return fmt.Sprintf("%x", t.Sum(nil))
}
func login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	r.ParseForm()
	username, found1 := r.Form["username"]
	password, found2 := r.Form["password"]
	if !(found1 && found2) {
		io.WriteString(w, "参数不正确")
	}
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		lof(err, "数据库连接失败")
	}
	defer db.Close()
	var post_data PostData
	post_data.user_name = username[0]
	post_data.pass_word = mdFormat(password[0])

	var user_name, pass_word string
	var id int

	err = db.QueryRow("select * from users where id = ?", 1).Scan(&id, &user_name, &pass_word)
	lof(err, "err")

	lof(id, user_name, pass_word, "idusernamepassword")
	if err != nil {
		sendBack(500, "登录失败", w)
	} else {
		sendBack(200, "登录成功", w)
	}

}

func resister(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	session, err := store.Get(r, "firstSession2")
	if err != nil {
		lof(err, "err")
		return
	}
	lof(session, "session")
	r.ParseForm()
	username, found1 := r.Form["username"]
	password, found2 := r.Form["password"]
	if !(found1 && found2) {
		io.WriteString(w, "参数不正确")
	}

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		lof(err, "数据库连接失败")
	}
	defer db.Close()
	lof(username[0], password[0])
	var post_data PostData
	post_data.user_name = username[0]

	var tempId int
	err = db.QueryRow("select id from users where username = ?", username[0]).Scan(&tempId)
	if err != sql.ErrNoRows {
		sendBack(500, "用户名已存在，注册失败", w)
		return
	}
	post_data.pass_word = mdFormat(password[0])
	session.Values["username"] = post_data.user_name
	session.Values["password"] = post_data.pass_word
	session.Save(r, w)
	_, err = db.Exec("insert into users(username,password) values(?,?)", post_data.user_name, post_data.pass_word)
	if err != nil {
		lof(err, "err")
		sendBack(500, "注册失败", w)

	} else {
		sendBack(200, "注册成功", w)

	}
}

type result struct {
	Code int
	Msg  string
	Data []string
}

func sendBack(c int, s string, w http.ResponseWriter) {
	arr := &result{
		c,
		s,
		[]string{},
	}

	b, json_err := json.Marshal(arr) //json化结果集

	if json_err != nil {
		fmt.Println("encoding faild")
	} else {
		io.WriteString(w, string(b)) //返回结果
	}
}
