package main

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

func (t *Text) textModifier() {
	// 1. Убираем лишние пробелы
	t.Content = strings.Join(strings.Fields(t.Content), " ")

	// 2. Обрабатываем знаки минус и плюс
	var result strings.Builder
	var sumOfDigits int

	for i := 0; i < len(t.Content); i++ {
		if t.Content[i] == '-' {
			if i > 0 && i < len(t.Content)-1 {
				// Меняем местами символы слева и справа от '-'
				result.WriteByte(t.Content[i-1])
				result.WriteByte(t.Content[i+1])
				i++ // Пропускаем следующий символ, так как мы его уже обработали
			}
		} else if t.Content[i] == '+' {
			result.WriteByte('!')
		} else if unicode.IsDigit(rune(t.Content[i])) {
			sumOfDigits += int(t.Content[i] - '0') // Суммируем цифры
		} else {
			result.WriteByte(t.Content[i]) // Добавляем обычные символы
		}
	}

	t.Content = result.String()

	// 3. Удаляем цифры и добавляем их сумму в конец строки
	if sumOfDigits > 0 {
		t.Content += fmt.Sprintf(" %d", sumOfDigits)
	}

	fmt.Println(t.Content)
}

func main() {
	text := &Text{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter text: ")

	for scanner.Scan() {
		text.Content = scanner.Text()
		text.textModifier()
	}
}
