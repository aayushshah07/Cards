package main

func main() {

	//Colon is required when we declare varible first time only
	cards := newDeck()
	// hand, reamainingCards := deal(cards, 5)

	// hand.print()
	// fmt.Println("Remaining Hands")
	// reamainingCards.print()
	// fmt.Println(cards.toString())

	cards.saveToFile("my_cards.txt")
}
