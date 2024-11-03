package internal

import (
	"bufio"
	"fmt"
	"os"
)

type Text struct {
	Content string
}

func (t *Text) textModifier() {
	fmt.Println(t.Content) // We have to work out this method
}

func main() {
	text := &Text{}
	// Create new scanner for reading from enter
	scanner := bufio.NewScanner(os.Stdin)

	// Ask user to enter the string
	fmt.Print("Enter text: ")

	for scanner.Scan() {
		text.Content = scanner.Text()
		text.textModifier()
	}
}
