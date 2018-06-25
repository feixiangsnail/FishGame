package Sockets

import (
	// "../Lib/Service"
	"Lib/Tool"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"Lib/Service"
)

var lof = fmt.Println
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type MessageInfo struct {

	Action string `action`
}

type WebSocketClientModel struct {
	Connect      *websocket.Conn
	AllOnMessage map[string]func(Connect *websocket.Conn)
	//打开触发事件
	AllOnOpen map[string]func(Connect *websocket.Conn)
	//关闭触发事件
	AllOnClose map[string]func(Connect *websocket.Conn)
	//发送触发事件
	AllOnSend map[string]func(Connect *websocket.Conn)
}

func (this *WebSocketClientModel) Init(w http.ResponseWriter, r *http.Request) bool {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		lof(err)
		return false
	}
	this.Connect = conn
	lof("已经成功连接200")
	return true

}
func (this *WebSocketClientModel) Send(data interface{}) {
	err := this.Connect.WriteJSON(data)
	if err != nil {
		lof(err)
		return
	}
	this.OnSend()
}

//执行关闭事件
func (this *WebSocketClientModel) Close() {
	this.OnClose()
}

//执行发送数据事件
func (this *WebSocketClientModel) OnSend() {
	for _, v := range this.AllOnSend {
		v(this.Connect)
	}
}

//执行打开事件
func (this *WebSocketClientModel) OnOpen() {
	for _, v := range this.AllOnOpen {
		v(this.Connect)
	}
}

//执行关闭事件
func (this *WebSocketClientModel) OnClose() {
	for _, v := range this.AllOnClose {
		v(this.Connect)
	}
}

func (this *WebSocketClientModel) OnMessage() {

	for {
		_, p, err := this.Connect.ReadMessage()
		if err != nil {
			lof(err)
			return
		}

		Content := string(p)
		Mess := MessageInfo{}
		Tool_Lib.Json_Object(Content, &Mess)

		if Service_Lib.AllServiceActin[Mess.Action]!=nil{
			Service_Lib.AllServiceActin[Mess.Action].Run(Content)
		}
	}
}


