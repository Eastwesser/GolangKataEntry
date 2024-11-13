package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func removeDigitsAndSum(text string) string {
	var sum int                 // Переменная для хранения суммы цифр
	var builder strings.Builder // Builder для построения результирующей строки

	// Проходим по каждому символу текста
	for _, char := range text {
		if unicode.IsDigit(char) {
			// Конвертируем символ в целое число
			digit, _ := strconv.Atoi(string(char))
			sum += digit // Добавляем к сумме
		} else {
			// Если символ не цифра, добавляем его в результат
			builder.WriteRune(char)
		}
	}

	// Если сумма больше нуля, добавляем её в конец строки
	if sum > 0 {
		builder.WriteString(" ") // Добавляем пробел перед суммой
		builder.WriteString(strconv.Itoa(sum))
	}

	return builder.String()
}

func main() {
	text := "abc123def456"
	fmt.Println("Исходный текст:", text)

	result := removeDigitsAndSum(text)
	fmt.Println("Текст после удаления цифр и добавления суммы:", result)
}
