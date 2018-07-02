package User_Module

type UserModel struct {
	ID        int                      `sql:"User_ID" json:"id"`
	Username     string                  `sql:"User_Name" json:"username"`
	Password 	string 					`json:"password"`
	WeiXinID  string                   `sql:"User_WeChatID" json:"weixinid"`
	Coins  int                      `sql:"User_RoomCard" json:"coins"`

}