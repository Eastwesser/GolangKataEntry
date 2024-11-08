package fundamentials

import "fmt"

func loopMe(a, b int) int {
	for i := a; i < b; i++ {
		res = enumerate(a, b)
		fmt.Println("The numbers between", {a}, "and", {b} are: {res})
	}
}

func mainLoop() {
	var a, b int
	fmt.Scanln("Please enter a number": &a)
	fmt.Scanln("Please enter the second number": &b)
	result := loopMe(a, b)
	fmt.Println(result)
}

// пишем функцию loopMe(), которая принимает на вход два числа (с новой строки),
// сначала одно, потом энтер, второе, и пихает их в функцию
// main(), которая берет диапазон между ними и считает все числа между ними,
// выводит все числа, их сумму, разность, умножение и деление
