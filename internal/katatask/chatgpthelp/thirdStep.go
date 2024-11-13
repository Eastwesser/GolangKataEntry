package main

import (
	"fmt"
	"strings"
)

func replacePlusWithExclamation(text string) string {
	// Используем ReplaceAll для замены всех плюсов на восклицательные знаки
	return strings.ReplaceAll(text, "+", "!")
}

func main() {
	text := "a+b+c+d+e+f"
	fmt.Println("Исходный текст:", text)

	result := replacePlusWithExclamation(text)
	fmt.Println("Текст после замены:", result)
}
