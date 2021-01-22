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

	// convert to bytes and print
	greeting := "Hello word"
	fmt.Println([]byte(greeting))

	// join string slide
	cards4 := newDeck()
	fmt.Println(cards4.toString())

	// save file
	cards5 := newDeck()
	cards5.saveToFile("cc.txt")

	// create random function
	cards6 := newDeck()
	cards6.shuffle()
	cards6.print()
}

func newCard() string {
	return "Five of Diamonds"
}
