package GameFunc

import (
	"fmt"
	"math/rand"
	"Config"
	"Lib/Service"
	"Drive/Http/Model"
)

var lof = fmt.Println
type FishInfo struct{
	Id int `json:"id"` //鱼的id
	TypeIndex int `json:"typeIndex"` //鱼的类型
	IsHit bool `json:"isHit"`	//是否击落
	Action string `json:"action"` //回调的方法
	Power float64 `json:"power"` //炮弹威力
}



//基础倍率：是0.3
//乘以随机数:如果还是
//血量
func RegisterAll(){
	Service_Lib.Register("HitMethod",HitMethod)
}

func HitMethod(f *FishInfo){
	//鱼的倍率 *  随机数 >= 当前倍率，如果当前倍率大于死的倍率表明击落
	var captureRate = Config.FishRates[Config.FishIndexStr[f.TypeIndex]] //鱼的倍率
	//var cononRate float64 = f.Power
	var rad = rand.Float64()
	var isHit = false
	if captureRate >rad{            //鱼的捕获率与随机数的对比,大于随机数就是击落
		isHit = true
	}
	lof(f,"f")

	f.IsHit = isHit
	Drive_Http.Client.Send(f)
}


