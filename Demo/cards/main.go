package main

import "fmt"

func main() {
	//var card string = "Aces of Spades"
	//card := "Aces of Spades" //other way declare variable
	card := newCard()
	fmt.Println(card)
}

//define function
func newCard() string {

	return "Five of Diamonds"
}
