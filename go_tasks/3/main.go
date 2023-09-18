/*
Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	inputArray := []int{2, 4, 6, 8, 10}
	solve1(inputArray)
}

func solve1(inputData []int) {
	wg := sync.WaitGroup{}

	resChan := make(chan int, len(inputData))
	// Запускаем горутину на каждую переменную в слайсе
	for _, val := range inputData {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			val = val * val
			resChan <- val
		}(val)
	}
	// Запуск горутины которая ожидает завершения работы горутин и закрывает канал
	go func() {
		wg.Wait()
		close(resChan)
	}()

	// Выводит сумирует значения, пока канал работает
	sum := 0
	for v := range resChan {
		sum = sum + v
	}
	fmt.Println(sum)
}
