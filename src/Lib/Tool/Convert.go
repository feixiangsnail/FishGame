package Tool_Lib

import (
	"reflect"
	"strings"
	"strconv"
	"fmt"
)

func Convert_Int(Data interface{}) int {
	if v, ok := Data.(int); ok {
		return v;
	}
	return 0;
}
func Convert_String(Data interface{}) string {
	if v, ok := Data.(string); ok {
		return v;
	}
	return "";
}
func Convert_Float32(Data interface{}) float32 {
	if v, ok := Data.(float32); ok {
		return v;
	}
	return 0;
}
func Convert_Float64(Data interface{}) float64 {
	if v, ok := Data.(float64); ok {
		return v;
	}
	return 0;
}
func Convert_Int32(Data interface{}) int32 {
	if v, ok := Data.(int32); ok {
		return v;
	}
	return 0;
}
func Convert_Int64(Data interface{}) int64 {
	if v, ok := Data.(int64); ok {
		return v;
	}
	return 0;
}
func Convert_Int8(Data interface{}) int8 {
	if v, ok := Data.(int8); ok {
		return v;
	}
	return 0;
}
func Convert_FormStruct(m map[string][]string, pointer interface{}) {
	// reflect.Ptr类型 *main.Person
	pointertype := reflect.TypeOf(pointer)
	// reflect.Value类型
	pointervalue := reflect.ValueOf(pointer)
	// struct类型  main.Person
	structType := pointertype.Elem()
	// 遍历结构体字段
	for i := 0; i < structType.NumField(); i++ {
		// 获取指定字段的反射值
		f := pointervalue.Elem().Field(i)
		//获取struct的指定字段
		stf := structType.Field(i)
		// 获取tag
		name := strings.Split(stf.Tag.Get("json"), ",")[0]
		// 判断是否为忽略字段
		if name == "-" {
			continue
		}
		// 判断是否为空，若为空则使用字段本身的名称获取value值
		if name == "" {
			name = stf.Name
		}
		//获取value值
		v, ok := m[name]
		if !ok {
			continue
		}

		//获取指定字段的类型
		kind := pointervalue.Elem().Field(i).Kind()
		// 若字段为指针类型
		if kind == reflect.Ptr {
			// 获取对应字段的kind
			kind = f.Type().Elem().Kind()
		}
		// 设置对应字段的值
		switch kind {
		case reflect.Int:
			res, _ := strconv.ParseInt(fmt.Sprint(v[0]), 10, 64)
			pointervalue.Elem().Field(i).SetInt(res)

		case reflect.String:
			pointervalue.Elem().Field(i).SetString(fmt.Sprint(v[0]))
		}
	}
}
