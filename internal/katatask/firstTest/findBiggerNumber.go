package main

import "fmt"

func whichIsBigger(a, b int) string {
	var biggerNums string

	// a = 10, b = 100
	if a < b {
		biggerNums = "Число Б больше"
	} else if a > b {
		biggerNums = "Число А больше"
	} else {
		biggerNums = "Ваши числа равны!"
	}
	//

	return biggerNums
}

func main() {
	var a int
	var b int

	fmt.Println("Введите первое число:")
	fmt.Scanln(&a)
	fmt.Println("Введите первое число:")
	fmt.Scanln(&b)
	res := whichIsBigger(a, b)
	fmt.Println(res)
}
