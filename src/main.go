package main

import (

	"fmt"
	"Drive"
	_ "Drive/Disk/Model"
	_ "Drive/Http/Model"
	"GameFunc"
)

var lof = fmt.Println

func main() {
	GameFunc.RegisterAll()
	Drive.Init("D:/workspace/games/FishGame/")
	select {}

	//App.StartServer()
}
