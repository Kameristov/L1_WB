/*
Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
*/
package main

import "fmt"

func main() {
	d := 42
	fmt.Println(CheckType(d)) // Выведет:  int

	a := interface{}("Hello")
	fmt.Println(CheckType(a)) // Выведет:  string

	b := false
	fmt.Println(CheckType(b)) // Выведет:  bool

	c := make(chan int)
	fmt.Println(CheckType(c)) // Выведет: channel int

	e := 3.14
	fmt.Println(CheckType(e)) // Выведет: Неизвестный тип
}

//----------------------------------------------------------------
// Функция определения типа входящей переменой
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
		return "Неизвестный тип"
	}
}
