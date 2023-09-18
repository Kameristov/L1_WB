/*
Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	inputArray := []int{2, 4, 6, 8, 10}
	solve1(inputArray)
	solve2(inputArray)
}

// Расчет квадратов в отдельных горутинах и сразу вывод в консоль
func solve1(inputData []int) {
	wg := sync.WaitGroup{}

	// Запуск горутины на каждый элемент слайса
	for _, val := range inputData {
		wg.Add(1) // Инкрементация счетчика waitGroup
		go func(val int) {
			defer wg.Done()        // Отложенное декримент счетчика WaitGroup
			fmt.Println(val * val) // возведение в квадрат
		}(val)
	}
	wg.Wait() // Ожидание завершения работы всех горутин
}

func solve2(inputData []int) {
	wg := sync.WaitGroup{}

	resChan := make(chan int, len(inputData)) // Создаем канал для получения результата работы

	// Запускаем горутину на каждую переменную в слайсе
	for _, val := range inputData {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()      // По завершению уменьшаем счетчик в sync.WaitGroup
			resChan <- val * val // отправляем решение в канал
		}(val)
	}

	// Запуск горутины которая ожидает завершения работы горутин и закрывает канал
	go func() {
		wg.Wait()
		close(resChan)
	}()

	// Выводит значения пока канал работает
	for v := range resChan {
		fmt.Println(v)
	}
}
