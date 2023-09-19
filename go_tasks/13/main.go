/*
13. Поменять местами два числа без создания временной переменной.
*/
package main

import "fmt"

func main() {
	firstNumber := 11
	secondNumber := 22

	fmt.Printf("firstNumber = %d\nsecondNumber = %d\n\n", firstNumber, secondNumber)
	changeNumbers(&firstNumber, &secondNumber)
	fmt.Printf("firstNumber = %d\nsecondNumber = %d\n\n", firstNumber, secondNumber)

	changeNumbers2(&firstNumber, &secondNumber)
	fmt.Printf("firstNumber = %d\nsecondNumber = %d\n\n", firstNumber, secondNumber)

	firstNumber, secondNumber =changeNumbers3(&firstNumber, &secondNumber)
	fmt.Printf("firstNumber = %d\nsecondNumber = %d\n\n", firstNumber, secondNumber)
}

//----------------------------------------------------------------
// Функция обмена значениями переменных математически
func changeNumbers(first, second *int) {
	*first = *first + *second // суммируем два числа, записываем в первое (11 + 22 => 33)
	*second = *first - *second // из суммы вычитаем 2 число, записываем в первое (33 - 22 => 11)
	*first = *first - *second // из суммы вычитаем новое 2 число записываем в 1 число (33 - 11 => 22)
}

//----------------------------------------------------------------
// Функция обмена значениями переменных по значению
func changeNumbers2(first, second *int) {
	*first, *second = *second, *first 
}

//----------------------------------------------------------------
// Функция обмена значениями переменных через возврат по значению
func changeNumbers3(first, second *int) (int,int){
	return  *second, *first
}