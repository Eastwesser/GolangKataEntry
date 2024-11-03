package fundamentials

import "fmt"

// ЦЕЛОЕ ЧИСЛО
var fullNumber int = 1
fullNumberOne := 1

// ДРОБНОЕ ЧИСЛО
var floatNumberOne float64 = 3.14
floatNumberTwo := 3.15

// СТРОКА
var word string = "Hello World"
words := "Hello, man!"

// ЛОГИЧЕСКИЙ ТИП
var booType bool = true
booTyped := false

// МАССИВ & КАРТЫ
//[]int

// СЛАЙС
//map[string]int

fmt.Println(
fullNumber, fullNumberOne,
floatNumberOne, floatNumberTwo,
word, words,
booType, booTyped,
)


func add(bigNum int, lowNum int) int {
	return bigNum + lowNum
}

func del(bigNum int, lowNum int) int {
	return bigNum - lowNum
}

func mul(bigNum int, lowNum int) int {
	return bigNum * lowNum
}

func div(bigNum int, lowNum int) int {
	return bigNum / lowNum
}

func ifElseLogic(x int, y int) int {
	resIfElse = x + y

	if x > y {
		fmt.Println("x больше y")
	} else {
		fmt.Println("x меньше или равно y")
	}
	return resIfElse
}

func cycleChecker(i int) bool {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

func switchChecker(i int) bool {

	day := "Monday"
	switch day {
	case "Monday":
		fmt.Println("TODAY IS MONDAY")
	case "Friday":
		fmt.Println("TODAY IS FRIDAY")
	default:
		fmt.Println("Today is some day :c")

	}
}

func aryth() {

	resultSum := add(10000, 2)
	resultDel := del(1000, 2)
	resultMul := mul(2000, 2)
	resultDiv := div(1000, 5)

	fmt.Println(
		"Ваша сумма:", resultSum,
		"Ваша разность:", resultDel,
		"Ваше умножение:", resultMul,
		"Ваше деление:", resultDiv,
	)

}
