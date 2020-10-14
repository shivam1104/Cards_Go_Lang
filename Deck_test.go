package main
import (
"testing"
"os"
)

func TestNewDeck(t *testing.T) {
	d := newObj()
	if (len(d) != 16){
		t.Errorf("Expected Deck Length of 16 Got %v", len(d))
	}
	if (d[0] != "One Of Diamond"){
		t.Errorf("Not the FIRST Value, got %v" ,d[0])
	}
	if (d[15] != "Four Of Clubs"){
		t.Errorf("Not the Last Value, got %v" ,d[15])
	}
}

func TestSaveToDeckAndNewDeckFromFile (t *testing.T){
	os.Remove("_decktesting")
	deck := newObj()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if(len(loadedDeck) !=16){
		t.Errorf("Cannot re-open FIle preoperly apprently, Read  =  %v" ,len(loadedDeck) )
	}
}

