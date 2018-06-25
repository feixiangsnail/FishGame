package Tool_Lib

import (
	"fmt"
	"github.com/pquerna/ffjson/ffjson"
)

func Json_Object(Content string, Arg interface{}) interface{} {
	buf := []byte(Content)
	err := ffjson.Unmarshal(buf, Arg)
	if (err != nil) {
		fmt.Println(err);
	}
	return Arg;
}
func Json_String(Content interface{}) string {
	data, err := ffjson.Marshal(&Content);
	if (err != nil) {
		fmt.Println(err);
	}
	return string(data);
}

