package pepe

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

func inList(l []string, e interface{}) bool {
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
		if inList(p.AllowedLowerLetters, value) {
			llsum++
		} else if inList(p.AllowedCapitalLetters, value) {
			clsum++
		} else if inList(p.AllowedNumbers, value) {
			nsum++
		} else if inList(p.AllowedSpecialCharacters, value) {
			scsum++
		} else {
			return false, fmt.Errorf("Password contains an invalid character: %c", value)
		}

		fmt.Sprintln(llsum, clsum, nsum, scsum)
	}

	if llsum >= p.MinLowerLetters && clsum >= p.MinCapitalLetters && nsum >= p.MinNumbers && scsum >= p.MinSpecialCharacters {
		fmt.Println(true)

		return true, nil
	}
	fmt.Println(false)

	return false, nil
}
