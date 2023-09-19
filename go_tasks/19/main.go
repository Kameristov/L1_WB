/*
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.
*/
package main

import (
	"fmt"

	"slices"
)

func main() {
	inputString := "главрыба1dsf34"
	fmt.Println(inputString)
	fmt.Println(solve1(inputString))
	fmt.Println(solve2(inputString))
}

// Реверсия через for
func solve1(str string) string {
	slice := []rune(str)
	// Переворачиваем слайс
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return string(slice)
}

// Реверсия через функцию пакета slices
func solve2 (str string) string{
	slice := []rune(str)
	slices.Reverse(slice)
	return string(slice)
}

