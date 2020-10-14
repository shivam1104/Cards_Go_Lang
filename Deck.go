package main

import (
"fmt"
"strings" 
"os"
"io/ioutil"
"math/rand"
"time"
)

type deck []string

func newObj() deck {
	cards := deck{}
	cardsSuite := deck{"Diamond", "Spade", "Hearts", "Clubs"} 
	cardsNumber := deck{"One", "Two", "Three", "Four"} 
	for _,suite := range cardsSuite {
		for _,number := range cardsNumber {
			cards = append(cards, number  + " Of " + suite)

		}
	}
	return cards
}
 
func (d deck) Print() {
	for i,p := range d {
		fmt.Println(i,p)
	}
}
func (d deck) toString() string {
	p:=strings.Join([]string(d),",")
	return(p)

}

func (d deck) saveToFile (filename string) error{
 return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
 

}

func newDeckFromFile (filename string) deck{
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Print("Error: ", err)
		os.Exit(1)
	}
	return (deck(strings.Split(string(bs),",")))
	
}

func (d deck) shuffle()  {
	source := rand.NewSource(time.Now().UnixNano())
	r:= rand.New(source)
	for i := range d{
		newPosition := r.Intn(len(d)-1)
		d[i], d[newPosition] = d[newPosition], d[i]

	}

}