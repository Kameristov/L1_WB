/*
Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
*/
package main

import "fmt"

// Структура Human
type Human struct {
	Weight float64
}

// Метод структуры Human
func (h Human) GetWeight() float64 {
	return h.Weight
}

// Структура Action в коротрую встроена структура Human
type Action struct {
	Human
	Man  Human
	Name string
}

// Метод структуры Action
func (a Action) GetName() string {
	return a.Name
}

func main() {
	// Инициализация структуры Action
	action := Action{Human: Human{Weight: 12.0}, Man: Human{Weight: 17.0}, Name: "action"}
	fmt.Printf("%v\n", action)
	fmt.Printf("GetWeight() = %v, Man.GetWeight()() = %v, GetName() = %v\n", action.GetWeight(), action.Man.GetWeight(), action.GetName())
}
