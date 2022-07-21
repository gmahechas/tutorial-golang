package main

import "fmt"

func main() {
	/* 	cards := deck{"Ace of Diamonds", newCard()}
	   	cards = append(cards, "Six of Spades")
	   	cards.print("this is arg1") */

	/* 	cards := newDesk()
	   	hand, remainingDesk := deal(cards, 5)
	   	hand.print()
	   	remainingDesk.print() */

	/* 	cards := newDesk()
	   	fmt.Println(cards.toString())
	   	cards.saveToFile("my_cards") */

	cards := newDeckFromFile("my_cards")
	cards.print()
	cards.shuffle()
	fmt.Print("#######\n")
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
