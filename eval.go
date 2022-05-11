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

func (p Password) IsCompliant(s string) (bool, error) {
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
			return false, fmt.Errorf("Password contains an invalid character: %c", value)
		}
	}

	if llsum >= p.MinLowerLetters && clsum >= p.MinCapitalLetters && nsum >= p.MinNumbers && scsum >= p.MinSpecialCharacters {
		return true, nil
	}
	return false, nil
}
