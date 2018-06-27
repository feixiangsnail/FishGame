package Client_Control

import "Application/User/Model"

func Send(model User_Module.UserModel,Send interface{}){
	model.Client.Send(Send)
}
