package kataTask

import (
	"bufio"
	"fmt"
	"os"
)

type Text struct {
	Content string
}

func (t *Text) TextModifier() {
	fmt.Println(t.Content)
	// CONTINUE THE CODE HERE

}

func main() {
	text := &Text{}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите строку:")
	for scanner.Scan() {
		text.Content = scanner.Text()
		text.TextModifier()
	}
}
