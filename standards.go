package secp

func AllEnglishLowerLetters() []string {
	return []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
}

func AllEnglishCapitalLetters() []string {
	return []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
}

func AllNumbers() []string {
	return []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
}

func BinaryNumbers() []string {
	return []string{"1", "0"}
}

func HexadecimalCapitalLetters() []string {
	return []string{"A", "B", "C", "D", "E", "F"}
}

func HexadecimalLowerLetters() []string {
	return []string{"a", "b", "c", "d", "e", "f"}
}

func CommonKeyboardSymbols() []string {
	return []string{"!", `"`, "#", "€", "%", "&", "/", "(", ")", "=", "?", "+", "´", "`", "°", "§", "^", "¨", ",", ".", "-", "_", ";", ":", ">", "<", "@", "$", "*", "'", "|", `\`, "[", "]", "{", "}"}
}

func SymbolsNonReservedByYAML() []string {
	return []string{"€", "/", "(", ")", "=", "+", "´", "°", "§", "^", "¨", ",", ".", "_", ";", "<", "$"}
}