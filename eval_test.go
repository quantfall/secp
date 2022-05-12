package secp

import (
	"testing"
)

func TestIsCompliant(t *testing.T) {

	p := New()

	p.AllowedLowerLetters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s"}
	p.AllowedCapitalLetters = []string{"A", "E", "I", "O", "U"}
	p.AllowedNumbers = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	p.AllowedSpecialCharacters = []string{".", ","}
	p.MinCapitalLetters = 2
	p.MinLowerLetters = 5
	p.MinNumbers = 4
	p.MinSpecialCharacters = 2

	testP1 := "abcdfAE1234,."
	testP2 := "abcdfAE1234,."
	testP3 := "123455.,"
	testP4 := "carAbamOs.,"
	testP5 := "carAbamOs1234"
	testP6 := "abamos1234.,"

	if err := p.IsCompliant(testP3); err != nil {
		t.Log(err)
	} else {
		t.Error("Password 3 should not be compliant")
	}

	if err := p.IsCompliant(testP4); err != nil {
		t.Log(err)
	} else {
		t.Error("Password 4 should not be compliant")
	}

	if err := p.IsCompliant(testP5); err != nil {
		t.Log(err)
	} else {
		t.Error("Password 5 should not be compliant")
	}

	if err := p.IsCompliant(testP6); err != nil {
		t.Log(err)
	} else {
		t.Error("Password 6 should not be compliant")
	}

	if err := p.IsCompliant(testP1); err != nil {
		t.Log(err)
	} else {
		t.Error("Password 1 contains an invalid character but the function didn't returned an error")
	}

	if err := p.IsCompliant(testP2); err == nil {
		t.Log("Password 2 is compliant")
	} else {
		t.Error(err)
	}

}
