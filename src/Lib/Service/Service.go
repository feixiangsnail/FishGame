package Service_Lib

import (
	"fmt"
	"reflect"
	"Lib/Tool"
	"github.com/satori/go.uuid"
)

var lof = fmt.Println

var AllServiceActin = make(map[string]*ServiceActin)
type ServiceActin struct{
	Name string
	Type reflect.Type
	Action interface{}
}
func (this *ServiceActin) Run(Content string ){
	obj:=reflect.New(this.Type)
	Tool_Lib.Json_Object(Content,obj.Interface())

	fv:=reflect.ValueOf(this.Action)
	FunParams:=make([]reflect.Value,1)
	FunParams[0] = reflect.ValueOf(obj.Interface())
	fv.Call(FunParams)

}
func NewClientID() string{
	ClientID,_:=uuid.NewV4()
	return ClientID.String()
}
func Register(Name string,Fun interface{}) interface{}{

	_Fun := reflect.TypeOf(Fun)


	AllServiceActin[Name] = &ServiceActin{
		Name:Name,
		Type:_Fun.In(0).Elem(),
		Action:Fun,
	}

	return Fun
}












