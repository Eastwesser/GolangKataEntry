// package main
//
// import (
//
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//	"unicode"
//
// )
//
//	type Text struct {
//		Content string
//	}
//
// // A. Функция для удаления лишних пробелов
// func removeSpaces(textLine string) string {
//
//		for strings.Contains(textLine, "  ") {
//			textLine = strings.ReplaceAll(textLine, "  ", " ")
//		}
//
//		return textLine
//	}
//
// // B. Функция для смены символов вокруг "-" (и удаления "-")
//
//	func swapChars(textLine string) string {
//		runes := []rune(textLine)
//		var res []rune
//
//		for i := 0; i < len(runes); i++ {
//			if runes[i] == '-' && i > 0 && i < len(runes)-1 {
//				res[len(res)-1], runes[i+1] = runes[i+1], res[len(res)-1]
//			} else {
//				res = append(res, runes[i])
//			}
//		}
//
//		return string(res)
//	}
//
// // C. Функция для замены "+" на "!"
// func swapPlus(textLine string) string {
//
//		for strings.Contains(textLine, "+") {
//			textLine = strings.ReplaceAll(textLine, "+", "!")
//		}
//
//		return textLine
//	}
//
// // D. Функция для удаления цифр и подсчета их суммы
//
//	func countSum(textLine string) string {
//		var sum int
//		var builder strings.Builder
//
//		for _, char := range textLine {
//			if unicode.IsDigit(char) {
//				digit, _ := strconv.Atoi(string(char))
//				sum += digit
//			} else {
//				builder.WriteRune(char)
//			}
//		}
//
//		if sum > 0 {
//			builder.WriteString(" ")
//			builder.WriteString(strconv.Itoa(sum))
//
//		}
//
//		return builder.String()
//	}
//
// func (t *Text) TextModifier() {
//
//		// A
//		t.Content = removeSpaces(t.Content)
//		// B
//		t.Content = swapChars(t.Content)
//		// C
//		t.Content = swapPlus(t.Content)
//		// D
//		t.Content = countSum(t.Content)
//
//		fmt.Println(t.Content)
//	}
//
//	func main() {
//		text := &Text{}
//		scanner := bufio.NewScanner(os.Stdin)
//		fmt.Print("Введите строку: ")
//		if scanner.Scan() {
//			text.Content = scanner.Text()
//			text.TextModifier()
//		} else {
//			fmt.Println("Ошибка ввода:", scanner.Err())
//		}
//	}
//
// package main
//
// import "fmt"
//
//	func sumWord(a, b int) string {
//		res := a + b
//
//		switch res {
//		case 0:
//			return "Zero"
//		case 1:
//			return "One"
//		case 2:
//			return "Two"
//		case 3:
//			return "Three"
//		case 4:
//			return "Four"
//		case 5:
//			return "Five"
//		case 6:
//			return "Six"
//		case 7:
//			return "Seven"
//		case 8:
//			return "Eight"
//		case 9:
//			return "Nine"
//		case 10:
//			return "Ten"
//		default:
//			return "Out of range."
//		}
//	}
//
//	func main() {
//		var a, b int
//		fmt.Println("Введите два числа (сумма не более 10):")
//		fmt.Scanln(&a, &b)
//
//		// Проверяем, что сумма не превышает 10
//		if a+b > 10 {
//			fmt.Println("Сумма превышает 10.")
//			return
//		}
//
//		fmt.Println("Ваш результат: ", sumWord(a, b))
//	}
//package main
//
//import (
//	"fmt"
//	"strings"
//)
//
//func main() {
//	names := [3]string{"Den", "Ben", "Zen"}
//	fmt.Println(names, len(names))
//
//	for i := 0; i < len(names); i++ {
//		fmt.Println(names[i]) // в столбик
//	}
//	sprinty := strings.Join(names[:], " ")
//	fmt.Println(sprinty) // в строчку
//
//	sprinty2 := strings.Join(names[:], ", ")
//	fmt.Println(sprinty2) // в строчку с запятой
//
//	arrayNum := [3][4]int{
//		{1, 2, 3, 4},
//		{5, 6, 7, 8},
//		{9, 10, 11, 12},
//	}
//
//	arrayStr := [3][4]string{
//		{"Quas", "Wex", "Exort", "Ult"},
//		{"Quas", "Wex", "Quas", "Ult"},
//		{"Exort", "Wex", "Exort", "Ult"},
//	}
//
//	for i := 0; i < len(arrayNum); i++ {
//		fmt.Println(arrayNum[i], arrayStr[i])
//	}
//}

//package main
//
//import (
//	"fmt"
//)
//
//func main() {
//	var a, b int
//	fmt.Scan(&a, &b)
//
//	for i := a; i <= b; i++ {
//		if i%2 == 0 {
//			fmt.Println(i)
//		}
//	}
//}

// Напиши программу, которая принимает от пользователя строку и число N. Программа должна:
// Если число N чётное: вывести все символы строки в обратном порядке.
// Если число N нечётное: продублировать каждый символ строки N раз и вывести результат.
//Пример ввода и вывода
//Пример 1:
//Ввод:
//Введите строку: hello
//Введите число: 4
//Вывод:
//Обратный порядок: olleh
//Пример 2:
//Ввод:
//Введите строку:
//Введите число: 3
//Вывод:
//Дублированная строка: gggooo

package main

import (
	"fmt"
	"strings"
)

func stringOnum(textLine string, numOne int) string {
	var res string

	if numOne%2 == 0 {
		// REVERSE POLARITY
		for i := len(textLine) - 1; i >= 0; i-- {
			res += string(textLine[i])
		}

	} else if numOne%2 != 0 {
		for _, char := range textLine {
			res += strings.Repeat(string(char), numOne)
		}
	}
	return res
}

func main() {
	var a string
	var b int

	fmt.Println("Введите строку: ")
	fmt.Scan(&a)
	fmt.Println("Введите число: ")
	fmt.Scan(&b)

	res := stringOnum(a, b)
	fmt.Println(res)
	s := "привет"
	r := []rune(s)
	fmt.Println(string(r)) // [1087 1088 1080 1074 1077 1090] если без string
}

//Задача на инкремент
//Напишем программу, которая:
//
//Читает строку и число.
//Если число чётное, добавляет в результат символы строки с чётными индексами.
//Если число нечётное, добавляет символы строки с нечётными индексами.
//
