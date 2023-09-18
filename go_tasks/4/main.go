/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.

 Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
*/

package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {
	worckersNum := 5
	// Настраеваем перехват приерывания Ctrl+C
	interupt := make(chan os.Signal, 1)
	signal.Notify(interupt, os.Interrupt)

	// Создаем канал завершающий горутину
	done := make(chan struct{})
	mainChan := make(chan rune, 5)

	// Запуск генератора символов
	go func() {
		for {
			select {
			case <-done:
				return
			default:
				mainChan <- rune(charset[rand.Intn(len(charset))]) // отправляем случайные символы из константной строки
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()

	// Запуск воркеров
	wg := sync.WaitGroup{}
	for i := 0; i < worckersNum; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for val := range mainChan {
				fmt.Println(val)
			}
		}()
	}

	// Блокируемся до получения прерывания
	<-interupt
	done <- struct{}{} // отправляем сигнал на закрытие горутины генератора символов
	close(mainChan)    // Закрываем канал
	wg.Wait()
}
