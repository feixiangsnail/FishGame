package User_Module

import "Application/Client/Interface"

type UserModel struct {
	ID        int                      `sql:"User_ID" json:"ID"`
	Name      string                   `sql:"User_Name" json:"Name"`
	Sex       int                      `sql:"User_sex" json:"sex"`
	Birthday  string                   `sql:"User_Birthday" json:"birthday"`
	WeiXinID  string                   `sql:"User_WeChatID" json:"_"`
	Src       string                   `sql:"User_Src" json:"src"`
	RoomCard  int                      `sql:"User_RoomCard" json:"src"`
	Client    ClientModule.ClientModel  `json:"_"`
}
