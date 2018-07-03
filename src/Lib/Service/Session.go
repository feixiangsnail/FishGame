package Service_Lib

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Session struct {
	attrMap map[string]interface{}
}

func (session *Session) SetAttr(name string, value interface{}) {
	session.attrMap[name] = value
}

func (session *Session) GetAttr(name string) (value interface{}, exist bool) {
	value = session.attrMap[name]
	if value != nil {
		exist = true
	} else {
		exist = false

	}
	return
}
func (session *Session) DelAttr(name string) {
	delete(session.attrMap, name)
}

var sessionMap = make(map[string]*Session)

func GetSession(w http.ResponseWriter, r *http.Request) (session *Session) {
	cookie, _ := r.Cookie("SESSIONID")
	log.Println(cookie,"cookie")
	session = sessionMap[cookie.String()]
	if session == nil {
		cookie = &http.Cookie{Name: "SESSIONID", Value: generateSessionID()}
		http.SetCookie(w, cookie)
		fmt.Println("设置cookie")
		session = &Session{attrMap: make(map[string]interface{})}
		sessionMap[cookie.String()] = session
	}

	return
}
func generateSessionID() (sessionID string) {
	buf := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, buf)
	checkError(err)
	sessionID = base64.URLEncoding.EncodeToString(buf)
	return

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
