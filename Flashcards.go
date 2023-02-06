package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type flashcard struct {
	term       string
	definition string
}

var flashcardDeck []flashcard

var numberOfCards int

func createCards(reader *bufio.Reader) {

	flashcardDeck = make([]flashcard, 0)

	for i := 0; i < numberOfCards; i++ {

		var f flashcard

		fmt.Printf("The term for card #%v: ", i+1)
		ter, _ := reader.ReadString('\n')
		ter = strings.TrimSpace(ter)

		fmt.Printf("The definition for card #%v: ", i+1)
		def, _ := reader.ReadString('\n')
		def = strings.TrimSpace(def)

		f.term = ter
		f.definition = def

		flashcardDeck = append(flashcardDeck, f)
	}
}

func quiz(reader *bufio.Reader) {
	for i := range flashcardDeck {
		fmt.Printf("Print the definition of %v:\n", flashcardDeck[i].term)
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)
		if answer == flashcardDeck[i].definition {
			fmt.Println("Correct!")
		} else {
			fmt.Printf("Wrong. The right answer is %v.\n", flashcardDeck[i].definition)
		}
	}
}

func main() {

	fmt.Println("Input the number of cards:")

	reader := bufio.NewReader(os.Stdin)
	num, _ := reader.ReadString('\n')
	num = strings.TrimSpace(num)
	numberOfCards, _ = strconv.Atoi(num)

	createCards(reader)

	quiz(reader)

}
