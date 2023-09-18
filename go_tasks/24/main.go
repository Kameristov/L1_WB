/*
Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	a := NewPoint(1, -1)
	b := NewPoint(1, 5)

	res := distance(a, b)
	fmt.Println(res)
}

// Структура точки, хранящая координаты по осям X и Y
type Point struct {
	x float64
	y float64
}

// Создание новой точки
func NewPoint(x float64, y float64) *Point {
	return &Point{x: x, y: y}
}

// Вычисление дистанции между точками
func distance(a, b *Point) float64 {
	return math.Sqrt(math.Pow(math.Abs(a.x-b.x), 2) + math.Pow(math.Abs(a.y-b.y), 2)) 

}
