package main

import (
	_ "Application/Game"
	"Drive"
	_ "Drive/Disk/Model"
	_ "Drive/Http/Model"
	"fmt"
)

var lof = fmt.Println

func main() {

	Drive.Init("D:/workspace/games/FishGame/")
	select {}

	//App.StartServer()
}
