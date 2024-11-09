package main

import (
	"fmt"
	"strconv"
)

func betweenNums(a, b int) string {
	var stringOfNumbers string

	if a > b {
		a, b = b, a
	}

	for i := a + 1; i < b; i++ {
		stringOfNumbers += strconv.Itoa(i) + " "
	}

	return stringOfNumbers
}

func main() {
	var a, b int
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	// Вызываем функцию и выводим результат
	res := betweenNums(a, b)
	fmt.Println("Числа между", a, "и", b, ":", res)
}
