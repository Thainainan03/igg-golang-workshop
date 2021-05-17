package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Heart", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0777)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error!", err)
	}
	// 1. []byte -> string
	// 2. string -> []string
	// 3. []string -> deck
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() deck {
	for index, v := range d {
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		tmp := v
		d[index] = d[r.Intn(len(d)-1)]
		d[r.Intn(len(d)-1)] = tmp
	}
	return d
}
