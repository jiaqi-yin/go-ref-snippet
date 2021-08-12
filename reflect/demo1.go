package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)

	rValue := reflect.ValueOf(b)
	n1 := 200 + rValue.Int()
	fmt.Printf("rValue=%v type=%T\n", rValue, rValue)
	fmt.Printf("n1=%v type=%T\n", n1, n1)

	rKind := rValue.Kind()    // (1)
	rTypeKind := rType.Kind() // (2)
	fmt.Println("rKind=", rKind)
	fmt.Println("rTypeKind=", rTypeKind)

	iV := rValue.Interface()
	n2 := iV.(int)
	fmt.Printf("n2=%v type=%T\n", n2, n2)
}

type Student struct {
	Name string
	Age  int
}

func reflectTest02(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)

	rValue := reflect.ValueOf(b)

	rKind := rValue.Kind()    // (1)
	rTypeKind := rType.Kind() // (2)
	fmt.Println("rKind=", rKind)
	fmt.Println("rTypeKind=", rTypeKind)

	iV := rValue.Interface()
	fmt.Printf("iV=%v type=%T\n", iV, iV)
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("name=%v age=%v\n", stu.Name, stu.Age)
	}
}

func reflectTest03(b interface{}) {
	rValue := reflect.ValueOf(b)
	rValue.Elem().FieldByName("Name").SetString("Jerry")
}

func main() {
	var num int = 100
	reflectTest01(num)

	stu := Student{
		Name: "Tom",
		Age:  20,
	}
	reflectTest02(stu)

	fmt.Println("Before change", stu)
	reflectTest03(&stu)
	fmt.Println("After change", stu)
}
