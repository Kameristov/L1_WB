/*
Реализовать собственную функцию sleep.
*/

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	t := time.Second * 3
	fmt.Println("sleep1")
	sleep1(t)
	fmt.Println("sleep2")
	sleep2(t)
	fmt.Println("sleep3")
	sleep3(t)
	fmt.Println("end")
}

// Проверка в бесконечном цикле
func sleep1(duration time.Duration) {
	end := time.Now().Add(duration) // находим время завершения сна 
	for {
		if time.Now().After(end) {
			return // выходим если текущее время больше время завершения сна
		}
	}
}

// Сон через блокировку канала таймером
func sleep2(duration time.Duration) {
	<-time.After(duration)
}

// Сон через установка таймера в контексте и канал
func sleep3(duration time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	<-ctx.Done()
	cancel()
}
