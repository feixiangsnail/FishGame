package Tool_Lib

import "fmt"

func Message_Check(err error){
	if err!=nil{
		fmt.Println(err)
	}
}
