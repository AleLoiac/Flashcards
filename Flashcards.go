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

// need add cards
func createCards(reader *bufio.Reader) {

	flashcardDeck = make([]flashcard, 0)

	for i := 0; i < numberOfCards; i++ {

		var f flashcard

		fmt.Printf("The term for card #%v:\n", i+1)
	Loop1:
		ter, _ := reader.ReadString('\n')
		ter = strings.TrimSpace(ter)

		for j := range flashcardDeck {
			if ter == flashcardDeck[j].term {
				fmt.Printf("The term \"%v\" already exists. Try again:\n", flashcardDeck[j].term)
				goto Loop1
			}
		}

		fmt.Printf("The definition for card #%v:\n", i+1)
	Loop2:
		def, _ := reader.ReadString('\n')
		def = strings.TrimSpace(def)

		for z := range flashcardDeck {
			if def == flashcardDeck[z].definition {
				fmt.Printf("The definition \"%v\" already exists. Try again:\n", flashcardDeck[z].definition)
				goto Loop2
			}
		}

		f.term = ter
		f.definition = def

		flashcardDeck = append(flashcardDeck, f)
	}
}

func quiz(reader *bufio.Reader) {
	for i := range flashcardDeck {
		fmt.Printf("Print the definition of \"%v\":\n", flashcardDeck[i].term)
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		if answer == flashcardDeck[i].definition {
			fmt.Println("Correct!")
		} else {
			control := false
			for j := range flashcardDeck {
				if answer == flashcardDeck[j].definition {
					fmt.Printf("Wrong. The right answer is \"%v\", but your definition is correct for \"%v\".\n", flashcardDeck[i].definition, flashcardDeck[j].term)
					control = true
					break
				}
			}
			if control == false {
				fmt.Printf("Wrong. The right answer is \"%v\".\n", flashcardDeck[i].definition)
			}
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
