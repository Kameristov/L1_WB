/*
Реализовать все возможные способы остановки выполнения горутины.
*/
package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	solve1()
	solve2()
	solve3()
	solve4()
	solve5()
	solve6()
}
//----------------------------------------------------------------
// 1. Использование сигнализирующего канала
func solve1() {
	ch := make(chan struct{})

	// Запуск горутины
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(id int) {
		for {
			select {
			case <-ch:
				fmt.Printf("Prog %d END work\n\n", id)
				wg.Done()
				return
			default:
				fmt.Printf("Prog %d work\n", id)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}(1)

	time.Sleep(time.Second) // Пауза
	ch <- struct{}{}        // Отправка сиглана горутине через канал
	close(ch)               // Закрытие канала
	wg.Wait()
}

//----------------------------------------------------------------
// 2. Использование функции отмены контекста
func solve2() {
	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(id int) {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("Prog %d END work\n\n", id)
				wg.Done()
				return
			default:
				fmt.Printf("Prog %d work\n", id)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}(2)
	time.Sleep(time.Second)
	cancel() // завершаем горутину через вызов функции завершения контекста
	wg.Wait()
}
//----------------------------------------------------------------
// 3. Использование таймера в контексте
func solve3() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	go func(id int) {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("Prog %d END work\n\n", id)
				wg.Done()
				return
			default:
				fmt.Printf("Prog %d work\n", id)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}(3)
	wg.Wait()
	cancel()
}
//----------------------------------------------------------------
// 4. Использование переменной-флага с мъютексом
func solve4() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	mut := sync.Mutex{}
	stat := false

	go func(id int) {
		for {
			mut.Lock()
			s := stat
			mut.Unlock()
			if s {
				fmt.Printf("Prog %d END work\n\n", id)
				wg.Done()
				return
			}
			fmt.Printf("Prog %d work\n", id)
			time.Sleep(100 * time.Millisecond)
		}
	}(4)

	time.Sleep(time.Second)
	mut.Lock()
	stat = true
	mut.Unlock()
	wg.Wait()
}
//----------------------------------------------------------------
// 5. Использование функции runtime.Goexit() которая немедленно останавливает выполнение текущей горутины.
func solve5() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func(id int) {
		for {
			fmt.Printf("Prog %d work\n", id)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Prog %d work\n", id)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Prog %d END work\n\n", id)
			wg.Done()
			runtime.Goexit()
			fmt.Printf("I'm Alive\n")
		}
	}(5)
	wg.Wait()
	time.Sleep(time.Second)
}
//----------------------------------------------------------------
// 6. Простое завершение через return
func solve6() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func(id int) {
		for i := 0; i < 15; i++ {
			if i >= 10 {
				fmt.Printf("Prog %d END work\n\n", id)
				wg.Done()
				return
			}
			fmt.Printf("Prog %d work\n", id)
		}
	}(6)
	wg.Wait()
}
