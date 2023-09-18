/*
Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
*/

package main

import "fmt"

func main() {
	i := 6
	var number int64 = 0b11111100000011111

	fmt.Printf("Number 1: %b\n", number)

	// установка бита
	if setBit(&number, i) != nil {
		fmt.Printf("Incorrect number")
	} else {
		fmt.Printf("Number 2: %b\n", number)
	}

	// обнуление бита
	if resetBit(&number, i) != nil {
		fmt.Printf("Incorrect number")
	} else {
		fmt.Printf("Number 3: %b\n", number)
	}
}

// Установка бита
func setBit(number *int64, bitNum int) error {
	// Проверка изменяемого бита на выход за пределы int64
	if bitNum >= 64 && bitNum < 0 {
		return fmt.Errorf("invalid bit number: %d", bitNum)
	}
	// Установка бита в 1 через побитовое OR со сдвигом на нужный бит
	*number = *number | (1 << bitNum)
	return nil
}

// Обнуление бита
func resetBit(number *int64, bitNum int) error {
	// Проверка изменяемого бита на выход за пределы int64
	if bitNum >= 64 && bitNum < 0 {
		return fmt.Errorf("invalid bit number: %d", bitNum)
	}
	// Установка бита в 0 через AND с инвертированым числом
	*number = *number & ^(1 << bitNum)
	return nil
}
