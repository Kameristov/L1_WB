/*
12. Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
*/

package main

import (
	"fmt"
	"slices"
)

func main() {
	animalsString := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(createPlenty(animalsString))
	fmt.Println(createPlentyAsString(animalsString))
	fmt.Println(sortCompact(animalsString))
}

// ----------------------------------------------------------------
// Функция возвращает собственное множество в виде мапы элемент множества - ключ мапы.
// Клчючи в мапе не повторяются, что обеспечивает индивидуальность элементов.
func createPlenty(slice []string) map[string]struct{} {
	// создаем пустую мапу под результат
	plenty := map[string]struct{}{}
	// Проходим по слайсу
	for _, val := range slice {
		// каждое значение заносив в мапу как ключ,
		plenty[val] = struct{}{}
	}
	return plenty
}

// ----------------------------------------------------------------
// Функция возвращает собственное множество в виде слайса.
// Принцип действия как и у функции createPlenty() за исключением
// вывода финального результата в виде строки.
func createPlentyAsString(slice []string) []string {
	// создаем пустую мапу под результат
	plenty := map[string]struct{}{}
	// Проходим по слайсу
	for _, val := range slice {
		plenty[val] = struct{}{}
	}
	// создаем слайс размером с мапу plenty
	stringPlenty := make([]string, len(plenty))
	for key := range plenty {
		// аписываем каждый элемент мапы в слайс
		stringPlenty = append(stringPlenty, key)
	}
	return stringPlenty
}

// ----------------------------------------------------------------
// Функция возвращает собственное множество в виде слайса.
// через библиотеку slices
func sortCompact(slice []string) []string {
	slices.Sort(slice)            // сортировка слайса
	slice = slices.Compact(slice) // объединение одинаковых элемементов в слайсе
	return slice
}
