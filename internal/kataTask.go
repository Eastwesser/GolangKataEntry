package kataTask

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

type Text struct {
	Content string
}

func (t *Text) TextModifier() {
	fmt.Println("Исходное содержание:", t.Content)

	textMod := strings.ReplaceAll(t.Content, "\n", " ")
	for i := 0; i < len(textMod); i++ {
		switch textMod[i] {
		case '-':
			// Обработка символа '-'
		case '+':
			textMod = strings.ReplaceAll(textMod, "+", "!")
		default:
			if unicode.IsDigit(rune(textMod[i])) {
				// Действия для цифр
			}
		}
	}

	t.Content = strings.Join(strings.Fields(textMod), " ")
	fmt.Println("Модифицированное содержание:", t.Content)
}

func InputAndModifyText() {
	text := &Text{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите текст: ")
	for scanner.Scan() {
		text.Content = scanner.Text()
		text.TextModifier()
	}
}

//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strings"
//	"unicode"
//)
//
//type Text struct {
//	Content string
//}
//
////func (t *Text) textModifier() {
////	fmt.Println(t.Content) // We have to work out this method
////	textMod := strings.ReplaceAll(t.Content, "\n", " ")
////	t.Content = strings.Join(strings.Fields(textMod), " ")
////	for i := len(textMod) - 1; i >= 0; i-- {
////		// if '-' in i {
////		if textMod[i] == '-' {
////			fmt.swap([1] with [2])
////		} else if '+' in i {
////			// fmt.ReplaceAll("+", "!")
////			textMod := strings.ReplaceAll(textMod, "+", "!")
////		} else {
////			for num in len(textMod ) {
////				fmt.unicode.IsDigit(Checksum)
////
////				//for num := 0; num < len(textMod); num++ {
////				//	// Действия с каждым символом
////				//}
////			}
////		}
////	}
////	return strings.Fields(strings.Join(textMod))
////	if unicode.IsDigit(rune(textMod[num])) {
////		// Действия для цифр
////	}
////}
//
//func (t *Text) textModifier() {
//	fmt.Println("Original Content:", t.Content)
//
//	// Заменяем все переносы строк пробелами
//	textMod := strings.ReplaceAll(t.Content, "\n", " ")
//
//	// Проходим по каждому символу строки
//	for i := 0; i < len(textMod); i++ {
//		switch textMod[i] {
//		case '-':
//			// Обработка символа '-'
//		case '+':
//			textMod = strings.ReplaceAll(textMod, "+", "!")
//		default:
//			if unicode.IsDigit(rune(textMod[i])) {
//				// Действия, если символ является цифрой
//			}
//		}
//	}
//
//	// Преобразуем строку в массив слов и соединяем обратно с пробелами
//	t.Content = strings.Join(strings.Fields(textMod), " ")
//	fmt.Println("Modified Content:", t.Content)
//}
//
//func main() {
//	text := &Text{}
//	// Create new scanner for reading from enter
//	scanner := bufio.NewScanner(os.Stdin)
//
//	// Ask user to enter the string
//	fmt.Print("Enter text: ")
//
//	for scanner.Scan() {
//		text.Content = scanner.Text()
//		text.textModifier()
//	}
//}
