/*
Реализовать конкурентную запись данных в map.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	BlockChan()
	NonBlockChan()
	MutexMap()
	RWMutexMap()
	SyncMap()
}

//----------------------------------------------------------------
// Передача данных через буфферезированный канал в горутину, где происходит запись

// Структура данных передаваемых через канал
type Data struct {
	id   int
	data string
}

func BlockChan() {
	storage := make(map[int]string)
	writeChan := make(chan Data, 3) // буфферезированный канал
	gorutines := 5

	// Горутина записи в мапу
	go func(map[int]string) {
		for val := range writeChan {
			storage[val.id] = val.data
		}
	}(storage)

	// Создание группы горутин для записи данных
	wg := sync.WaitGroup{}
	for count := 0; count < gorutines; count++ {
		wg.Add(1)
		go func(count int) {
			writeChan <- Data{id: count, data: "123"}
			wg.Done()
		}(count)
	}
	// Ожидание завершения
	wg.Wait()
	close(writeChan)
	fmt.Println(storage)
}

// ----------------------------------------------------------------
// Передача данных через не буфферезированный канал в горутину, где происходит запись
func NonBlockChan() {
	storage := make(map[int]string)
	closeChan := make(chan struct{})
	writeChan := make(chan Data) // не буфферизированный канал
	gorutines := 5

	go func(map[int]string) {
		for {
			select {
			case val := <-writeChan:
				storage[val.id] = val.data
			case <-closeChan:
				return
			}
		}
	}(storage)

	wg := sync.WaitGroup{}
	for count := 0; count < gorutines; count++ {
		wg.Add(1)
		go func(count int) {
			writeChan <- Data{id: count, data: "123"}
			wg.Done()
		}(count)
	}
	wg.Wait()
	closeChan <- struct{}{}
	fmt.Println(storage)
}

// ----------------------------------------------------------------
// Запись в мап через мъютекс
func MutexMap() {
	storage := make(map[int]string)
	gorutines := 5

	mux := sync.Mutex{}
	wg := sync.WaitGroup{}
	for count := 0; count < gorutines; count++ {
		wg.Add(1)
		go func(count int) {
			mux.Lock()
			storage[count] = "123"
			mux.Unlock()
			wg.Done()
		}(count)
	}
	wg.Wait()
	fmt.Println(storage)
}

// ----------------------------------------------------------------
// Запись в мап через RW мъютекс, при записи происходит блокировка, при чтение нет блокировки если не происходит запись
func RWMutexMap() {
	storage := make(map[int]string)
	gorutines := 5

	mux := sync.RWMutex{}
	wg := sync.WaitGroup{}
	for count := 0; count < gorutines; count++ {
		wg.Add(1)
		go func(count int) {
			mux.Lock()
			storage[count] = "123"
			mux.Unlock()
			wg.Done()
		}(count)
	}
	wg.Wait()
	fmt.Println(storage)
}

//----------------------------------------------------------------
// Запись в sync мап
/*
Тип Map оптимизирован для двух распространенных случаев использования:
(1) когда запись для данного ключа записывается только один раз, но читается много раз, как в кэшах, которые только растут.
(2) когда несколько горутин читают, записывают и перезаписывать записи для непересекающихся наборов ключей.
В этих двух случаях использование карты может значительно уменьшить конфликты блокировок по сравнению с картой Go в сочетании с отдельным мьютексом или RWMutex.
*/
func SyncMap() {
	storage := sync.Map{}
	gorutines := 5

	wg := sync.WaitGroup{}
	for count := 0; count < gorutines; count++ {
		wg.Add(1)
		go func(count int) {
			storage.Store(count, "123")
			wg.Done()
		}(count)
	}
	wg.Wait()
	storage.Range(func(key, value any) bool {
		fmt.Println(key, value)
		return true
	})
}
