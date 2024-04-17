package main

// For checking punctioation after every cap,up,low that contains a number
func Punctuate(s string) string {
	punctuation := ""
	for _, p := range s {
		if p == ',' {
			punctuation = ","
		} else if p == ':' {
			punctuation = ":"
		} else if p == '!' {
			punctuation = "!"
		} else if p == '.' {
			punctuation = "."
		}
	}
	return punctuation
}

func Punctuate1(s []string) []string {
	for i, char := range s {
		if char == "..." {
			s[i-1] = s[i-1] + string(char[0])
		} else if char == "!?" || char == "!!" {
			s[i-1] = s[i-1] + s[i]
			s[i] = ""
		} else if char[0] == ',' || char[0] == '.' || char[0] == '!' || char[0] == '?' || char[0] == ':' || char[0] == ';' {
			s[i-1] = s[i-1] + string(char[0])
			s[i] = char[1:]
		} else if char == "," || char == "." || char == "!" || char == "?" || char == ":" || char == ";" {
			s[i-1] = s[i-1] + s[i]
			s[i] = char[1:]
		}
	}
	return RemoveEmptyValues(s)
}

// For handling the first part of the single quote'
func SingleQuote1(s []string) []string {
	count := false
	for i, char := range s {
		if char == "'" && !count {
			s[i] = s[i] + s[i+1]
			s[i+1] = " "
			count = true
		} else if char == "'" && count {
			count = false
			continue
		}
	}
	return s
}

// For handling the second part of the single quote (')
func SingleQuote2(s []string) []string {
	for i, char := range s {
		if char == "'" && (s[i-1] != " ") {
			// if char
			s[i-1] = s[i-1] + s[i]
			s[i] = " "
		}
	}
	return s
}
