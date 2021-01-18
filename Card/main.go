package main

import "fmt"

func main() {

	// initialize variable
	var card string = "Ace of Spades"
	fmt.Println(card)

	// assign value to variable
	card = newCard()
	fmt.Println(card)

	// slice
	cards := []string{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	for i, card := range cards {
		fmt.Println(i, card)
	}

	// type
	cards2 := deck{"Ace of Diamonds", newCard()}
	cards2 = append(cards, "Six of Spades")
	cards2.print()

	// slice and return mutilple value
	cards3 := newDeck()
	hand, remainCards := deal(cards3, 5)
	hand.print()
	remainCards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
