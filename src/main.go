package main

import (

	"fmt"
	"GameFunc"
	"App"
)

var lof = fmt.Println

func main() {

	GameFunc.RegisterAll()
	App.StartServer()
}
