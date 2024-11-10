package main

import (
	"GolangEntryTest/internal/fundamentals"
	"GolangEntryTest/internal/katatask"
	"fmt"
	"log"
)

// Обработчик ошибок
func handleError(err error) {
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}
}

func main() {
	fmt.Println("Запускаем примеры:")

	// Вызов примеров из fundamentals
	err := fundamentals.RunExamples()
	handleError(err)

	// Вызов текстового модификатора из kataTask
	err = katatask.RunTextModifier()
	handleError(err)
}
