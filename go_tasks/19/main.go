/*
Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.
*/
package main

import "fmt"

func main() {
	inputString := "главрыба"
	fmt.Println(inputString)
	fmt.Println(solve1(inputString))
}

func solve1(str string) string {
	slice := []rune(str)
	// Переворачиваем слайс
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return string(slice)
}

// TODO другие варианты
