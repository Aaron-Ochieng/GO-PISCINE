package main

func ToArray(s string) []string {
	words := []string{}
	word := ""
	for _, char := range s {
		if char != ' ' {
			word += string(char)
		} else if char == ' ' && word != "" {
			words = append(words, word)
			word = ""
		}
	}
	if word != "" {
		words = append(words, word)

	}
	return RemoveEmptyValues(words)
}
func Vowels(s string) []string {
	words := ToArray(s)
	// Rules for vowels a,e,i,o,u and h
	for i, word := range words {
		if (word == "a" || word == "A") && (words[i+1][0] == 'a' || words[i+1][0] == 'e' || words[i+1][0] == 'i' || words[i+1][0] == 'o' || words[i+1][0] == 'u' || words[i+1][0] == 'h') {
			if word == "a" {
				words[i] = "an"
			} else {
				words[i] = "An"
			}
		}
	}
	return words
}
