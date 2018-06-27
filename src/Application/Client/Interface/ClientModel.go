package ClientModule

type ClientModel interface {
	Send(data interface{})
	Close()
	OnOpen()
	OnMessage()
	OnClose()
}