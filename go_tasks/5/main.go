/*
Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
*/

package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var workTime time.Duration = 5 // Задаем время работы

	channel := make(chan int, 2)

	wg := sync.WaitGroup{}
	wg.Add(1)

	// Горутина получения даннных из канала. Завершается при закрытии канала.
	go func() {
		defer wg.Done()
		for val := range channel {
			fmt.Printf("Recieve: %d\n", val)
		}
	}()

	// Создание контекста с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * workTime)
	defer cancel()

	// Горутина отправки данных в канал. При завершении таймаута контекста, закрывает канал и завершается.
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				close(channel)
				return
			default:
				channel <- rand.Intn(10)
			}
		}
	}(ctx)

	// Ождиание зарвешения работы горутины
	wg.Wait()
}
