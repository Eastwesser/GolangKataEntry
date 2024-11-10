package katatask

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

func (t *Text) TextModifier() error {
	fmt.Println("Исходное содержание:", t.Content)

	textMod := strings.ReplaceAll(t.Content, "\n", " ")
	var result strings.Builder

	for i := 0; i < len(textMod); i++ {
		switch textMod[i] {
		case '-':
			result.WriteByte('~')
		case '+':
			result.WriteByte('!')
		default:
			if unicode.IsDigit(rune(textMod[i])) {
				doubledValue := string(textMod[i]) + string(textMod[i])
				result.WriteString(doubledValue)
			} else {
				result.WriteByte(textMod[i])
			}
		}
	}

	t.Content = strings.Join(strings.Fields(result.String()), " ")
	fmt.Println("Модифицированное содержание:", t.Content)
	return nil
}

func RunTextModifier() error {
	text := &Text{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Введите текст: ")
	for scanner.Scan() {
		text.Content = scanner.Text()
		err := text.TextModifier()
		if err != nil {
			return fmt.Errorf("ошибка при модификации текста: %w", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка ввода текста: %w", err)
	}
	return nil
}
