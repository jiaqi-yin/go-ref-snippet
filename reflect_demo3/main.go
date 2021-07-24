package main

import (
	"fmt"
	"reflect"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster_age"`
	Score float32
	Sex   string
}

func (s Monster) Print() {
	fmt.Println("----Start----")
	fmt.Println(s)
	fmt.Println("----End----")
}

func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

func (s Monster) Set(name string, age int, score float32, sex string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Sex = sex
}

func TestStruct(a interface{}) {
	typ := reflect.TypeOf(a)
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		fmt.Println("expect struct")
	}

	num := val.Elem().NumField()

	val.Elem().Field(0).SetString("Elf")

	fmt.Printf("struct has %d fields\n", num)
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d: value=%v kind=%v\n", i, val.Elem().Field(i), val.Elem().Field(i).Kind())
		tagVal := typ.Elem().Field(i).Tag.Get("json")
		if tagVal != "" {
			fmt.Printf("Field %d: tag=%v\n", i, tagVal)
		}
	}

	numOfMethod := val.Elem().NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)
	val.Elem().Method(1).Call(nil)

	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := val.Elem().Method(0).Call(params)
	fmt.Println("res=", res[0].Int())
}

func main() {
	var a Monster = Monster{
		Name:  "Troll",
		Age:   500,
		Score: 99.9,
	}
	fmt.Println("Before: ", a)
	TestStruct(&a)
	fmt.Println("After: ", a)
}
