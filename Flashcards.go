package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type flashcard struct {
	term       string
	definition string
}

type flashcardDeck struct {
	deck []flashcard
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	card, _ := reader.ReadString('\n')
	card = strings.TrimSpace(card)

	reader2 := bufio.NewReader(os.Stdin)
	definition, _ := reader2.ReadString('\n')
	definition = strings.TrimSpace(definition)

	reader3 := bufio.NewReader(os.Stdin)
	answer, _ := reader3.ReadString('\n')
	answer = strings.TrimSpace(answer)

	if definition == answer {
		fmt.Println("Your answer is right!")
	} else {
		fmt.Println("Your answer is wrong...")
	}

}
