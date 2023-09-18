/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	counterByMutex()
	counterByAtomic()
}

// ----------------------------------------------------------------
// Реализация через Mutex

// Структура счетчика
type MutexCounter struct {
	counter int
	mtx     sync.Mutex
}

// Метод инкремента счетчика
func (c *MutexCounter) Incriment() {
	c.mtx.Lock()
	c.counter++
	c.mtx.Unlock()
}

// Метод получения значения счетчика
func (c *MutexCounter) GetCounter() int {
	return c.counter
}

func counterByMutex() {
	// Инициализируем структуру счетчика
	c := MutexCounter{counter: 0, mtx: sync.Mutex{}}
	// Запускаем 10 горутин
	gorutines := 10
	wg := sync.WaitGroup{}
	for x := 0; x < gorutines; x++ {
		wg.Add(1)
		go func() {
			c.Incriment()
			wg.Done()
		}()
	}
	// Ожидаем окончание работы горутин
	wg.Wait()
	// Выводим значение счетчика
	fmt.Println("", c.GetCounter())
}

// ----------------------------------------------------------------
// Реализация через атомики

// Структура счетчика
type AtomicCounter struct {
	counter int64
}

// Метод инкремента счетчика
func (c *AtomicCounter) Incriment() {
	atomic.AddInt64(&c.counter, 1)
}

// Метод получения значения счетчика
func (c *AtomicCounter) GetCounter() int64 {
	return c.counter
}

func counterByAtomic() {
	// Инициализируем структуру счетчика
	c := AtomicCounter{counter: 0}
	// Запускаем 10 горутин
	gorutines := 10
	wg := sync.WaitGroup{}
	for x := 0; x < gorutines; x++ {
		wg.Add(1)
		go func() {
			c.Incriment()
			wg.Done()
		}()
	}
	// Ожидаем окончание работы горутин
	wg.Wait()
	// Выводим значение счетчика
	fmt.Println("", c.GetCounter())
}
