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

	testP1 := "abcd#"
	testP2 := "abcdfAE1234,."
	testP3 := "123455.,"
	testP4 := "carAbamOs.,"
	testP5 := "carAbamOs1234"
	testP6 := "abamos1234.,"

	if com, err := p.IsCompliant(testP3); com || err != nil {
		t.Log(com, err)
		t.Errorf("Password 3 should not be compliant and error should be nil")
	}

	if com, err := p.IsCompliant(testP4); com || err != nil {
		t.Log(com, err)
		t.Errorf("Password 4 should not be compliant and error should be nil")
	}

	if com, err := p.IsCompliant(testP5); com || err != nil {
		t.Log(com, err)
		t.Errorf("Password 5 should not be compliant and error should be nil")
	}

	if com, err := p.IsCompliant(testP6); com || err != nil {
		t.Log(com, err)
		t.Errorf("Password 6 should not be compliant and error should be nil")
	}

	if _, err := p.IsCompliant(testP1); err == nil {
		t.Errorf("Password 1 contains an invalid character but the function didn't returned an error")
	}

	if com, err := p.IsCompliant(testP2); !com || err != nil {
		t.Log(com, err)
		t.Errorf("Password 2 should be compliant and error should be nil")
	}

}
