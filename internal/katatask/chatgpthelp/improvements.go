//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//	"unicode"
//)
//
//type Text struct {
//	Content string
//}
//
//// A. Функция для перестановки символов вокруг "-" и удаления "-"
//func swapChars(textLine string) string {
//	runes := []rune(textLine)
//	var res []rune
//
//	for i := 0; i < len(runes); i++ {
//		if runes[i] == '-' && i > 0 && i < len(runes)-1 {
//			// Если слева и справа от "-" есть символы, меняем их местами
//			res = append(res, runes[i+1], runes[i-1])
//			i++ // Пропускаем следующий символ, так как он уже добавлен
//		} else if runes[i] != '-' {
//			res = append(res, runes[i])
//		}
//	}
//
//	return string(res)
//}
//
//// B. Функция для замены "+" на "!"
//func swapPlus(textLine string) string {
//	return strings.ReplaceAll(textLine, "+", "!")
//}
//
//// C. Функция для удаления цифр и подсчёта их суммы
//func countSum(textLine string) string {
//	var sum int
//	var builder strings.Builder
//
//	for _, char := range textLine {
//		if unicode.IsDigit(char) {
//			digit, _ := strconv.Atoi(string(char))
//			sum += digit
//		} else {
//			builder.WriteRune(char)
//		}
//	}
//
//	if sum > 0 {
//		builder.WriteString(" ")
//		builder.WriteString(strconv.Itoa(sum))
//	}
//
//	return builder.String()
//}
//
//// D. Функция для удаления лишних пробелов
//func removeSpaces(textLine string) string {
//	for strings.Contains(textLine, "  ") {
//		textLine = strings.ReplaceAll(textLine, "  ", " ")
//	}
//	return textLine
//}
//
//func (t *Text) TextModifier() {
//	// A
//	t.Content = swapChars(t.Content)
//	// B
//	t.Content = swapPlus(t.Content)
//	// C
//	t.Content = countSum(t.Content)
//	// D
//	t.Content = removeSpaces(t.Content)
//
//	fmt.Println(t.Content)
//}
//
//func main() {
//	text := &Text{}
//	scanner := bufio.NewScanner(os.Stdin)
//	fmt.Print("Введите строку: ")
//	if scanner.Scan() {
//		text.Content = scanner.Text()
//		text.TextModifier()
//	} else {
//		fmt.Println("Ошибка ввода:", scanner.Err())
//	}
//}
