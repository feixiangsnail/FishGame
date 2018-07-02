package Drive_Http

import (
	"Application/User"
	"Application/User/Model"
	"Drive"
	"Lib/Service"
	"Lib/Tool"
	"Drive/webSocket/model"
	"log"
	"net/http"
	"reflect"
)

var lof = log.Println
var Client WebSocketModel.WebSocketClientModel

type HttpMode_Drive struct {
	Name          string `json:"Name"`
	Port          string `json:"Port"`
	Host          string `json:"host"`
	WebSocketPath string `json:"WebSocket"`
}

func init() {
	Drive.Register("HTTP", reflect.TypeOf(HttpMode_Drive{}))
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("加载http驱动成功")

}

func (this *HttpMode_Drive) InitData(data map[string]interface{}) {
	this.Name = Tool_Lib.Convert_String(data["Name"])
	this.Port = Tool_Lib.Convert_String(data["Port"])
	this.Host = Tool_Lib.Convert_String(data["Host"])
	this.WebSocketPath = Tool_Lib.Convert_String(data["WebSocket"])
	this.Init()
}
func (this *HttpMode_Drive) Init() {

	this.HttpService()
	this.WebSocketRun()
	go log.Fatal(http.ListenAndServe(this.Host+":"+this.Port, nil))

}

func (this *HttpMode_Drive) GetObject(Url string, Data interface{}) interface{} {
	return nil
}

//获取文本数据
func (this *HttpMode_Drive) GetContent(Url string) string {
	return "http"
}

//写入数据
func (this *HttpMode_Drive) Set(Url string, Data interface{}) int {
	return 666
}
func (this *HttpMode_Drive) HttpService() {
	http.Handle("/", http.FileServer(http.Dir("./views/")))
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	io.WriteString(w, "首页")
	//})
}

func (this *HttpMode_Drive) WebSocketRun() {
	if this.WebSocketPath == "" {
		return
	}

	http.HandleFunc("/"+this.WebSocketPath, func(w http.ResponseWriter, r *http.Request) {
		User := this.LoginUser(w, r)
		lof(User, "User")
		Client = WebSocketModel.WebSocketClientModel{}

		if !Client.Init(User,w, r) {
			return
		}
		go User.Client.OnMessage()
		User.Client.Send(User)
	})

}

func (this *HttpMode_Drive) LoginUser(w http.ResponseWriter, r *http.Request) *User_Module.UserModel {
	ID := InitID(w, r)

	user := User_Control.LoginUser(ID)

	return user

}
func InitID(w http.ResponseWriter, r *http.Request) string {
	MyCookie, err := r.Cookie("ClientID")
	if err != nil && MyCookie != nil {
		return MyCookie.Value
	}
	ID := Service_Lib.NewClientID()
	rc := http.Cookie{
		Value: ID,
		Name:  "ClientID",
	}
	http.SetCookie(w.(http.ResponseWriter), &rc)
	return ID

}
