package fundamentials

import "fmt"

func LoopMe(a, b int) string {
	if a > b {
		a, b = b, a
	}
	var res string
	for i := a; i <= b; i++ {
		res += fmt.Sprintf("%d ", i)
	}
	return res
}

func MainLoop() {
	var a, b int
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)
	result := LoopMe(a, b)
	fmt.Println("Числа в диапазоне:", result)
}

// пишем функцию loopMe(), которая принимает на вход два числа (с новой строки),
// сначала одно, потом энтер, второе, и пихает их в функцию
// main(), которая берет диапазон между ними и считает все числа между ними,
// выводит все числа, их сумму, разность, умножение и деление
