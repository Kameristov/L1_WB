/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {

	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}

	solve1(array)
	solve2(array)
	solve3(array)
}

// ----------------------------------------------------------------
// Реализация через не буфферезированный канал
func solve1(array []int) {
	xchan := make(chan int)
	x2chan := make(chan int)

	// горутина записи в первый канал
	go func() {
		for _, val := range array {
			xchan <- val
		}
		close(xchan)
	}()

	// горутина чтения из первого канал, умножение на 2 и отправка во второй канал
	go func() {
		for val := range xchan {
			x2chan <- val * 2
		}
		close(x2chan)
	}()

	// чтение данных из второго канала
	for val := range x2chan {
		fmt.Printf("%d\n", val)
	}
}

// ----------------------------------------------------------------
// Реализация через буфферезированный канал
func solve2(array []int) {
	xchan := make(chan int, 10)
	x2chan := make(chan int, 10)

	// горутина записи в первый канал
	go func() {
		for _, val := range array {
			xchan <- val
		}
		close(xchan)
	}()

	// горутина чтения из первого канал, умножение на 2 и отправка во второй канал

	go func() {
		for val := range xchan {
			x2chan <- val * 2
		}
		close(x2chan)
	}()

	// чтение данных из второго канала
	for val := range x2chan {
		fmt.Printf("%d\n", val)
	}
}

// ----------------------------------------------------------------
// Реализация через буфферезированный канал и несколько горутин
func solve3(array []int) {
	wg := sync.WaitGroup{}
	xchan := make(chan int, 10)
	x2chan := make(chan int, 10)

	// горутина записи в первый канал
	go func() {
		for _, val := range array {
			xchan <- val
		}
		close(xchan)
	}()

	// горутины чтения из первого канал, умножение на 2 и отправка во второй канал
	gorutin := 5
	for x := 0; x < gorutin; x++ {
		wg.Add(1)
		go func() {
			for {
				val, ok := <-xchan
				// если канал закрылся, тоже заканчиваем работу
				if !ok {
					wg.Done()
					return
				}
				x2chan <- val * 2
			}
		}()
	}

	// горутина ожидания завершения работы группы горутин
	go func() {
		wg.Wait()
		close(x2chan)
	}()

	// чтение данных из второго канала
	for val := range x2chan {
		fmt.Printf("%d\n", val)
	}
}
