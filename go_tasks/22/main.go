/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
*/
package main

import (
 "fmt"
 "math"
)

func main() {
// Возводим в степень
 a := math.Pow(2, 21)   // a = 2^21
 b := math.Pow(2, 22)   // b = 2^22

 // Умножение
 multiply := a * b
 fmt.Printf("Умножение a * b = %.2f\n", multiply)

  // Деление
 divide := b / a
 fmt.Printf("Деление b / a = %.2f\n", divide)

 // Сложение
 sum := a + b
 fmt.Printf("Сложение a + b = %.2f\n", sum)

 // Вычитание
 subtract := b - a
 fmt.Printf("Вычитание b - a = %.2f\n", subtract)
}