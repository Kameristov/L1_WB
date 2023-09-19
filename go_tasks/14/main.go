/*
Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/
package main

import (
	"fmt"
	"reflect"
)

func main() {
	d := 42
	a := "Hello"
	b := false
	c := make(chan int)
	e := 3.14


	fmt.Println(CheckType(d)) // Выведет:  int
	fmt.Println(CheckType(a)) // Выведет:  string
	fmt.Println(CheckType(b)) // Выведет:  bool
	fmt.Println(CheckType(c)) // Выведет: channel int
	fmt.Println(CheckType(e)) // Выведет: Неизвестный тип


	fmt.Println(CheckType2(d))
	fmt.Println(CheckType2(a))
	fmt.Println(CheckType2(b))
	fmt.Println(CheckType2(c))
	fmt.Println(CheckType2(e))
}

// ----------------------------------------------------------------
// Функция определения типа входящей переменой с помощью Type switch
func CheckType(variable interface{}) string {
	switch variable.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	default:
		return reflect.TypeOf(variable).String()
	}
}

// ----------------------------------------------------------------
// Функция определения типа входящей переменой с помощью пакета reflect
func CheckType2(variable interface{}) string {
	return reflect.TypeOf(variable).String()
}
