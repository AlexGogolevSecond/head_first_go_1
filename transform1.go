package main

import (
	"fmt"
	"reflect"
)

func main3() {
	var a int = 1
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))
	// f := 1.22
	var b float64
	b = float64(a) // вот так осуществляется приведение типов
	fmt.Println(b)
	fmt.Println(reflect.TypeOf(b))

}
