package main

import "fmt"
func aCard() string {
	return ("Ace of Clubs")
}

func main() {

	x := newObj()
	//fmt.Println(x.toString())
	x.saveToFile("Cards")
	p:= newDeckFromFile("Cards")
	p.Print()
	p.shuffle()
	fmt.Println("Is It Random Below?")
	p.Print()
}
