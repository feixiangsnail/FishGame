package User_Module

import "Application/Client/Interface"

//type UserModel struct {
//	ID        int                      `sql:"User_ID" json:"ID"`
//	Name      string                   `sql:"User_Name" json:"Name"`
//	Sex       int                      `sql:"User_sex" json:"sex"`
//	Birthday  string                   `sql:"User_Birthday" json:"birthday"`
//	WeiXinID  string                   `sql:"User_WeChatID" json:"_"`
//	Src       string                   `sql:"User_Src" json:"src"`
//	RoomCard  int                      `sql:"User_RoomCard" json:"src"`
//	Client    ClientModule.ClientModel  `json:"_"`
//}

type UserModel struct {
	ID        int                      `json:"id"`
	Username     string                  `sql:"User_Name" json:"username"`
	Password 	string 					`json:"password"`
	WeiXinID  string                   `sql:"User_WeChatID" json:"weixinid"`
	Coins  int                      `sql:"User_RoomCard" json:"coins"`
	Client    ClientModule.ClientModel  `json:"_"`
	Name      string                   `sql:"User_Name" json:"Name"`
	Src       string                   `sql:"User_Src" json:"src"`
	RoomCard  int                      `sql:"User_RoomCard" json:"src"`
	Sex       int                      `sql:"User_sex" json:"sex"`
	Birthday  string                   `sql:"User_Birthday" json:"birthday"`
}

type UserModel2 struct {
	ID        int                      `sql:"User_ID" json:"id"`
	Username     string                  `sql:"User_Name" json:"username"`
	Password 	string 					`json:"password"`
	WeiXinID  string                   `sql:"User_WeChatID" json:"weixinid"`
	Coins  int                      `sql:"User_RoomCard" json:"coins"`
}