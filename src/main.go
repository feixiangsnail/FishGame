package main

import (

	"fmt"
	"Drive"
	_ "Drive/Disk/Model"
	_ "Drive/Http/Model"
	_ "Application/Game"
)

var lof = fmt.Println

func main() {

	Drive.Init("D:/workspace/games/FishGame/")
	select {}

	//App.StartServer()
}
