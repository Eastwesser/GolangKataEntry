package main

import (
	"fmt"
	"strconv"
)

// Функция, возвращающая строку с числами между a и b
func betweenNums(a, b int) string {
	// Если a больше b, меняем их местами
	if a > b {
		a, b = b, a
	}

	var stringOfNumbers string
	for i := a + 1; i < b; i++ {
		// Преобразуем число в строку и добавляем к общей строке
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
