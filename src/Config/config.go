package Config

const BanefitRate = 0.2
var FishIndexStr = []string{"null", "fish1", "fish2", "fish3", "fish4", "fish5", "fish6", "fish8", "fish9", "fish10", "fish7", "shark1", "shark2"}
var FishCoins = map[string]int{
	"null":0,"fish1":1,"fish2":3,"fish3":5,"fish4":8,"fish5":10,"fish6":20,"fish8":40,"fish9":50,"fish10":60,"fish7":30,"shark1":100,"shark2":200,
}
var FishRates = SetFishRates(FishCoins)

func SetFishRates(fishCoins map[string]int) map[string]float64{
	fishRates:=map[string]float64{}
	for k,v:= range fishCoins{
		fishRates[k] =(1-BanefitRate)/float64(v)
	}
	return fishRates
}

//var BaseRate = 0.2