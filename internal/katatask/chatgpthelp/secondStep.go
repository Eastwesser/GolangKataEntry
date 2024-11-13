package main

import (
	"fmt"
)

func swapAroundMinus(text string) string {
	runes := []rune(text) // Преобразуем строку в срез рун для работы с символами

	// Создаем новый срез для сборки результата без минусов
	var result []rune

	for i := 0; i < len(runes); i++ {
		if runes[i] == '-' && i > 0 && i < len(runes)-1 { // Проверка на символ "-" с символами слева и справа
			// Меняем местами символы слева и справа от "-"
			result = append(result, runes[i+1], runes[i-1])
			i++ // Пропускаем следующий символ, т.к. он уже добавлен
		} else {
			result = append(result, runes[i]) // Добавляем символ, если это не "-"
		}
	}

	return string(result)
}

func main() {
	text := "a-b c-d e- f g-h 5-4"
	fmt.Println("Исходный текст:", text)

	result := swapAroundMinus(text)
	fmt.Println("Текст после замены:", result)
}
