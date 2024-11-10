package fundamentials

import (
	"fmt"
	"strconv"
)

// Функция для вывода чисел между `a` и `b`
func BetweenNums(a, b int) string {
	var stringOfNumbers string
	if a > b {
		a, b = b, a
	}
	for i := a + 1; i < b; i++ {
		stringOfNumbers += strconv.Itoa(i) + " "
	}
	return stringOfNumbers
}

func PrintBetweenNums() {
	var a, b int
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	res := BetweenNums(a, b)
	fmt.Println("Числа между", a, "и", b, ":", res)
}
