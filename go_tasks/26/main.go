/*
Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	have := []string{"abcd", "abCdefAaf", "aabcd"}
	want := []bool{true, false, false}

	for i := 0; i < len(want); i++ {
		res := repeatCheck(have[i])
		if res != want[i] {
			fmt.Printf("Err: have: %v, want: %v", res, want[i])
		}
	}
	fmt.Println("repeatCheck() OK")

	for i := 0; i < len(want); i++ {
		res := repeatCheck2(have[i])
		if res != want[i] {
			fmt.Printf("Err: have: %v, want: %v", res, want[i])
		}
	}
	fmt.Println("repeatCheck2() OK")
}

// ------------------------------------------------------------------------------
// Функция проверки на повторение символов в строке
func repeatCheck(str string) bool {
	str = strings.ToLower(str) // перевод букв из верхнего регистра в нижний
	m := make(map[rune]uint)
	// Заполнение мапы символами из строки, каждый символ - ключ, значение количество повторений символа в строке
	for _, val := range str {
		m[val]++
	}
	// Если символ повторился больше одного раза, возврат false
	for _, val := range m {
		if val > 1 {
			return false
		}
	}
	return true // если символы не повторялись
}

// ------------------------------------------------------------------------------
// Функция проверки на повторение символов в строке
// оптимизация repeatCheck()
func repeatCheck2(str string) bool {
	str = strings.ToLower(str) // перевод букв из верхнего регистра в нижний
	m := make(map[rune]struct{})
	// Заполнение мапы символами из строки, каждый символ - ключ, значение количество повторений символа в строке
	for _, val := range str {
		// если мы нашли в мапе ключ
		if _, ok := m[val]; ok {
			return false // выход
		}
		m[val] = struct{}{} // создание элемента мапы
	}
	return true // если символы не повторялись
}
