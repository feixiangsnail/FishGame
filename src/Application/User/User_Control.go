package User_Control

import (
	"Application/User/Model"
	"sync"
	"fmt"
)

var UserCount = 0
var loginUser map[string]*User_Module.UserModel

func init(){
	loginUser = make(map[string]*User_Module.UserModel)

}
func LoginUser(ID string) *User_Module.UserModel{
	var mutex sync.Mutex
	mutex.Lock()
	if loginUser[ID] == nil{
		UserID:=UserCount
		User:=User_Module.UserModel{
			ID:UserID,
			Name:fmt.Sprintf("%d用户",UserID),
		}
		loginUser[ID] = &User

	}
	UserCount = UserCount+1
	mutex.Unlock()
	return loginUser[ID]
}
