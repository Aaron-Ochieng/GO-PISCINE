package main

import (
	"strconv"
	"strings"
)

func StringToInt(s string) int {
	res := ""
	for _, char := range s {
		if char >= '0' && char <= '9' {
			res += string(char)
		}
	}
	num, _ := strconv.Atoi(res)
	return num
}

func Convert(words []string) []string {
	for i, word := range words {
		if strings.Contains(word, "(up)") {
			words[i-1] = strings.ToUpper(words[i-1])
			words[i] = ""

		} else if strings.Contains(word, "(up,") {
			n := StringToInt(words[i+1])
			for x := i - n; x <= i-1; x++ {
				words[x] = strings.ToUpper(words[x])
			}
			punctuation := Punctuate(words[i+1])
			words[i], words[i+1] = "", ""+punctuation

		} else if strings.Contains(word, "(low)") {
			words[i-1] = strings.ToLower(words[i-1])
			words[i] = ""

		} else if strings.Contains(word, "(low,") {
			n := StringToInt(words[i] + words[i+1])
			for x := i - n; x <= i-1; x++ {
				words[x] = strings.ToLower(words[x])
			}
			punctuation := Punctuate(words[i+1])
			words[i], words[i+1] = "", ""+punctuation

		} else if strings.Contains(word, "(cap)") {
			words[i-1] = Title(words[i-1])
			words[i] = ""

		} else if strings.Contains(word, "(cap,") {
			n := StringToInt(words[i] + words[i+1])
			for x := i - n; x <= i-1; x++ {
				words[x] = Title(words[x])
			}
			punctuation := Punctuate(words[i+1])
			words[i], words[i+1] = "", ""+punctuation

		} else if strings.Contains(word, "hex") {
			val := Hexdec(words[i-1])
			// convert the value back to string using Itoa
			xy := strconv.Itoa(int(val)) // cast int64 back to int
			words[i-1] = xy
			words[i] = ""

		} else if strings.Contains(word, "bin") {

			xb := Bindec(words[i-1])
			// Convert the int to string
			xc := strconv.Itoa(int(xb))
			words[i-1] = xc
			words[i] = ""

		} else if words[i] == "..." {
			words[i-1] = words[i-1] + words[i]
			words[i] = ""

		}
	}
	return RemoveEmptyValues(words)
}
