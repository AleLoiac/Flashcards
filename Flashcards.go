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

var flashcardDeck []flashcard

var numberOfCards int

func createCards() {

	flashcardDeck = make([]flashcard, 0)

	for i := 0; i < numberOfCards; i++ {

		var f flashcard

		fmt.Printf("The term for card #%v: ", i+1)
		reader := bufio.NewReader(os.Stdin)
		ter, _ := reader.ReadString('\n')
		ter = strings.TrimSpace(ter)

		fmt.Printf("The definition for card #%v: ", i+1)
		reader2 := bufio.NewReader(os.Stdin)
		def, _ := reader2.ReadString('\n')
		def = strings.TrimSpace(def)

		f.term = ter
		f.definition = def

		flashcardDeck = append(flashcardDeck, f)
	}
}

func main() {

	fmt.Scan(&numberOfCards)

	createCards()

	fmt.Println(flashcardDeck)

}
