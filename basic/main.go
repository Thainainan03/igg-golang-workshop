package main

import "fmt"

func main() {
	cards := newDeckFromFile("deck.txt")
	fmt.Println(cards.shuffle())
	// cards.print()
	// fmt.Println(time.Now().UnixNano())
}
