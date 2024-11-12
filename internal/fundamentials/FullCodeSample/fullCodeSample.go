package FullCodeSample

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

// Функция для обработки ввода с консоли
func getInput() (int, error) {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return strconv.Atoi(input)
}

// Функция деления с обработкой ошибок
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Пример функции, использующей defer
func greet() {
	defer fmt.Println("Thank you for using the program.")
	fmt.Println("Hello!")
}

func main() {
	greet()

	fmt.Println("Enter a number:")
	a, err := getInput()
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	fmt.Println("Enter another number:")
	b, err := getInput()
	if err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	result, err := divide(a, b)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result of division:", result)
	}

	person := Person{Name: "Alice", Age: 25}
	fmt.Printf("Person: %+v\n", person)

	switch a {
	case 0:
		fmt.Println("Zero entered")
	case 1:
		fmt.Println("One entered")
	default:
		fmt.Println("A different number entered")
	}

	var numbers []int
	for i := 0; i < a; i++ {
		numbers = append(numbers, i)
	}

	fmt.Println("Numbers up to a:", numbers)
	fmt.Println("Done with the program.")
}
