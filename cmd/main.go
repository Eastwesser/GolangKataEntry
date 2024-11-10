package main

import (
	"GolangEntryTest/internal/fundamentals"
	"GolangEntryTest/internal/katatask"
	"fmt"
)

func main() {
	fmt.Println("Запускаем примеры:")

	// Вызов примеров из fundamentals
	fundamentals.Aryth()

	// Вызов текстового модификатора из kataTask
	katatask.RunTextModifier()
}
