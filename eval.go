package secp

import "fmt"

type Password struct {
	MinLowerLetters          int
	MinCapitalLetters        int
	MinNumbers               int
	MinSpecialCharacters     int
	AllowedLowerLetters      []string
	AllowedCapitalLetters    []string
	AllowedNumbers           []string
	AllowedSpecialCharacters []string
}

func New() *Password {
	p := new(Password)
	return p
}

func inList(l []string, e string) bool {
	for _, value := range l {
		if e == value {
			return true
		}
	}
	return false
}

func (p Password) IsCompliant(s string) error {
	var llsum int
	var clsum int
	var nsum int
	var scsum int

	for _, value := range s {
		if inList(p.AllowedLowerLetters, string(value)) {
			llsum++
		} else if inList(p.AllowedCapitalLetters, string(value)) {
			clsum++
		} else if inList(p.AllowedNumbers, string(value)) {
			nsum++
		} else if inList(p.AllowedSpecialCharacters, string(value)) {
			scsum++
		} else {
			return fmt.Errorf("the password contains an invalid character: %c", value)
		}
	}

	if llsum < p.MinLowerLetters {
		return fmt.Errorf("the password does not contains at least %d lower case letters ", p.MinLowerLetters)
	}

	if clsum < p.MinCapitalLetters {
		return fmt.Errorf("the password does not contains at least %d capital case letters ", p.MinCapitalLetters)
	}

	if nsum < p.MinNumbers {
		return fmt.Errorf("the password does not contains at least %d numbers", p.MinNumbers)
	}

	if scsum < p.MinSpecialCharacters {
		return fmt.Errorf("the password does not contains at least %d special characters", p.MinSpecialCharacters)
	}
	
	return nil
}
