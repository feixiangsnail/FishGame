package Tool_Lib

import (
	"crypto/md5"
	"io"
	"fmt"
	"net/http"
	"encoding/json"
)
type result struct {
	Code int
	Msg  string
	Data interface{}
}
func MdFormat(data string) string {
	t := md5.New()
	io.WriteString(t, data)
	io.WriteString(t, string(t.Sum(nil))+"newcode1")
	io.WriteString(t, string(t.Sum(nil))+"newcode2")
	return fmt.Sprintf("%x", t.Sum(nil))
}

func SendBack(c int, s string, w http.ResponseWriter,d interface{}) {
	arr := &result{
		c,
		s,
		d,
	}

	b, json_err := json.Marshal(arr) //json化结果集

	if json_err != nil {
		fmt.Println("encoding faild")
	} else {
		io.WriteString(w, string(b)) //返回结果
	}
}