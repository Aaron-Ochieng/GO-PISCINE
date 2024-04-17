package main

import "strings"

func RemoveExtraSpaces(sentence string) string {
	var result string
	prevSpace := false

	for _, char := range sentence {
		if char != ' ' {
			result += string(char)
			prevSpace = false
		} else {
			if !prevSpace {
				result += " "
			}
			prevSpace = true
		}
	}

	return strings.TrimSpace(result)
}

func RemoveEmptyValues(s []string) []string {
	res := []string{}
	for _, char := range s {
		if strings.TrimSpace(char) != "" {
			res = append(res, char)
		}
	}

	return res
}
