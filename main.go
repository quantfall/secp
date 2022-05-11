package main

import (
	"log"

	"github.com/quantfall/secp/pepe"
)

func main() {
	p := pepe.New()

	p.AllowedLowerLetters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r"}
	p.AllowedCapitalLetters = []string{"A", "E", "I", "O", "U"}
	p.AllowedNumbers = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	p.AllowedSpecialCharacters = []string{".", ","}
	p.MinCapitalLetters = 2
	p.MinLowerLetters = 5
	p.MinNumbers = 4
	p.MinSpecialCharacters = 2

	testP1 := "abcd"
	testP2 := "abcdefAE1234,."
	testP3 := "carbamOs1234.,"
	testP4 := "carAbamOs123.,"
	testP5 := "carAbamOs1234."
	testP6 := "AbamOs1234.,"

	if _, err := p.IsCompliant(testP1); err == nil {
		log.Println("Password 1 contains an invalid character but the function didn't returned an error")
	}

	if com, _ := p.IsCompliant(testP2); !com {
		log.Println("Password 2 should be compliant")
	}

	if com, _ := p.IsCompliant(testP3); com {
		log.Println("Password 3 should not be compliant")
	}

	if com, _ := p.IsCompliant(testP4); com {
		log.Println("Password 4 should not be compliant")
	}

	if com, _ := p.IsCompliant(testP5); com {
		log.Println("Password 5 should not be compliant")
	}

	if com, _ := p.IsCompliant(testP6); com {
		log.Println("Password 6 should not be compliant")
	}

}
