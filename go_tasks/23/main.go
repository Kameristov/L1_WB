/*
Удалить i-ый элемент из слайса.
*/
package main

import (
	"errors"
	"fmt"
)

func main() {
	if res, err := delete([]int{1, 2, 3}, 2); err == nil {
		fmt.Println(res)
	}
}

// Функция удаления элемента из слайса
func delete(s []int, index int) ([]int, error) {
	// Проверка что индекс больше нуля и меньше длины слайса
	if index > len(s) && index < 0 {
		return s, errors.New("index out of incoming slice length")
	}
	// Создаем новый слайс
	r := make([]int, len(s)-1)
	// Копируем в новый слайс все элементы кроме ненужного
	copy(r, s[:index])
	copy(r[:index], s[index+1:])
	return r, nil
}
