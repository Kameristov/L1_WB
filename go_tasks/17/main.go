/*
Реализовать бинарный поиск встроенными методами языка.
*/
package main

import "fmt"

func main() {
	// Бинарный поиск происходит над сортированным массивом(слайсом), поэтому используется сортированный слайс
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 8
	fmt.Println(binarySearch(list, target))
}

func binarySearch(slice []int, target int) (int, error) {
	// Находим крайние индексы слайса
	low := 0
	high := len(slice) - 1

	// Пока индексы не станут однаковыми
	for low <= high {

		// Находим серидину участка слайса
		mid := (low + high) / 2

		if slice[mid] == target {
			// Если нали элементы выходим
			return mid, nil
		} else if slice[mid] < target {
			// Если средний элемент меньше искомого присваеваем правое от середины значение нижнему индексу
			low = mid + 1
		} else if slice[mid] > target {
			// Если средний элемент больше искомого присваеваем левое от середины значение верхнему индексу
			high = mid - 1
		}
	}
	return 0, fmt.Errorf("искомое значение не найдено")
}
