package main

import (
	"fmt"
	"math"

	//"strings"
	"reflect"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

/*
go mod init types.go
go get golang.org/x/text/cases
go get golang.org/x/text/language
*/
func main3() {
	// Создайте тег для русского языка
	tag := language.Russian
	fmt.Println(math.Floor(3.26))
	//fmt.Println(strings.Title("hello world"))
	fmt.Println(cases.Title(tag).String("hello мир"))
	// fmt.Println(strings.ToUpper(123))
	var quantity = 4
	var lenght, width = 1.2, 5.3
	fmt.Println(reflect.TypeOf(quantity)) // reflect для работы с типами
	fmt.Println(reflect.TypeOf(lenght))
	fmt.Println(reflect.TypeOf(width))
}
