/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/
package main

import (
	"fmt"
	"strings"
)

const (
	phrase = "snow dog sun"
)

func main() {
	fmt.Println(phrase)
	fmt.Println(splitRevertJoin(phrase))
}

// Функция перестановки элементов в строке местами
func splitRevertJoin(input string) string {
	slice := strings.Split(input, " ") // разделяем строку по пробелу
	// Переворачиваем слайс
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return strings.Join(slice, " ") // Собраем строку с пробелами
}

