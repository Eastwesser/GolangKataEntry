package main

import (
	"fmt"
	"strings"
)

func removeSpaces(textLine string) string {
	// Убираем все двойные пробелы, заменяя их на одинарные
	for strings.Contains(textLine, "  ") { // Проверяем, есть ли двойные пробелы
		textLine = strings.ReplaceAll(textLine, "  ", " ")
	}
	return textLine
}

func main() {
	text := "Это    пример   текста     с     лишними пробелами."
	fmt.Println("Исходный текст:", text)

	result := removeSpaces(text)
	fmt.Println("Текст после удаления лишних пробелов:", result)
}
