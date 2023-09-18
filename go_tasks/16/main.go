/*
Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
*/
package main

import "fmt"

var iterations int

func main() {
	list := []int{9, 7, 5, 5, 11, 5, 12, 2, 14, 1, 1}
	fmt.Println(list)
	iterations = 0
	fmt.Println(quickSort(list), iterations)
	iterations = 0
	fmt.Println(quickSort2(list), iterations)
}

func quickSort(slice []int) []int {
	iterations++
	// Если слайс меньше или равен 1, считаем что он уже отсортирован
	if len(slice) <= 1 {
		return slice
	}

	// выбираем последний элемент слайса как опорный элемент
	pivot := slice[len(slice)-1]

	// Подгатавливаем слайсы в которые будем записывать значения больше или меньше опорного элемента
	less := []int{}
	greater := []int{}

	// Разделляем слайс на слайсы относительно опорного элемента
	for _, val := range slice[:len(slice)-1] {
		if val <= pivot {
			less = append(less, val)
		} else {
			greater = append(greater, val)
		}
	}

	// Сортируем получившиеся слайсы
	less = quickSort(less)
	greater = quickSort(greater)

	// Объединение слайсов в один
	return append(append(less, pivot), greater...)
}

// Вариант быстрой сортировки чуть эфиктивнее если в слайсе есть много повторяющихся элементов
func quickSort2(slice []int) []int {
	iterations++
	// Если слайс меньше или равен 1, считаем что он уже отсортирован
	if len(slice) <= 1 {
		return slice
	}

	// Выбираем последний элемент слайса как опорный элемент
	pivot := slice[len(slice)-1]

	// Подгатавливаем слайсы в которые будем записывать значения больше или меньше опорного элемента
	less := []int{}
	greater := []int{}
	same := []int{pivot}

	// Разделляем слайс на слайсы относительно опорного элемента
	for _, val := range slice[:len(slice)-1] {
		switch {
		case val < pivot:
			less = append(less, val)
		case val > pivot:
			greater = append(greater, val)
		case val == pivot:
			same = append(same, val)
		}
	}

	// Сортируем получившиеся слайсы
	less = quickSort2(less)
	greater = quickSort2(greater)

	// Объединение слайсов в один
	return append(append(less, same...), greater...)
}
