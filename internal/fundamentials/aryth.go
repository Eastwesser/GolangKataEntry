package fundamentials

import "fmt"

// Пример базовых математических функций
func Add(bigNum int, lowNum int) int {
	return bigNum + lowNum
}

func Del(bigNum int, lowNum int) int {
	return bigNum - lowNum
}

func Mul(bigNum int, lowNum int) int {
	return bigNum * lowNum
}

func Div(bigNum int, lowNum int) int {
	if lowNum == 0 {
		fmt.Println("Деление на ноль!")
		return 0
	}
	return bigNum / lowNum
}

func Aryth() {
	resultSum := Add(10000, 2)
	resultDel := Del(1000, 2)
	resultMul := Mul(2000, 2)
	resultDiv := Div(1000, 5)

	fmt.Println("Сумма:", resultSum, "Разность:", resultDel, "Умножение:", resultMul, "Деление:", resultDiv)
}
