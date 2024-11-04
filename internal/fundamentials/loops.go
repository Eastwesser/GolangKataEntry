package fundamentials

import "fmt"

func loopMe(a int, b int) int {
	a := fmt.Scanln("Please enter a number")
	b := fmt.Scanln("Please enter the second number")
	for i := 0; i < 10; i++ {
		fmt.Println("loop me")
	}
}

func main() {
	fmt.Println(loopMe(a, b))
}

// пишем функцию loopMe(), которая принимает на вход два числа, сначала одно, потом энтер, второе, и пихает их в функцию
// main(), которая берет диапазон и считает все числа между ними
