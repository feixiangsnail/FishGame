package Tool_Lib

import (
	"log"
)

func Message_Check(err error){
	if err!=nil{
		log.Println(err)
	}
}
