package main

import (
	"fmt"
	"reflect"
)

type Cal struct {
	Num1 int `json:"num1"`
	Num2 int `json:"num2"`
}

func (c Cal) GetSub(name string) {
	fmt.Printf("%s completed %d - %d = %d\n", name, c.Num1, c.Num2, c.Num1-c.Num2)
}

func TestStruct(num1 int, num2 int, a interface{}) {
	rType := reflect.TypeOf(a)
	rVal := reflect.ValueOf(a)
	rElement := rVal.Elem()
	rKind := rVal.Kind()
	if rKind != reflect.Ptr || rElement.Kind() != reflect.Struct {
		fmt.Printf("Want %v, got %v\n", reflect.Ptr, rKind)
		fmt.Printf("Want %v, got %v\n", reflect.Struct, rElement.Kind())
	}

	numField := rElement.NumField()
	numMethod := rElement.NumMethod()
	fmt.Printf("Has %d fields\n", numField)
	fmt.Printf("Has %d methods\n", numMethod)

	for i := 0; i < numField; i++ {
		fmt.Printf("Field %d\tvalue=%v\ttag=%v\n",
			i, rElement.Field(i), rType.Elem().Field(i).Tag.Get("json"))
	}

	rElement.Field(0).SetInt(int64(num1))
	rElement.Field(1).SetInt(int64(num2))

	var params []reflect.Value
	params = append(params, reflect.ValueOf("Tom"))
	rElement.Method(0).Call(params)
}

func main() {
	var c Cal = Cal{}
	TestStruct(8, 3, &c)
}
