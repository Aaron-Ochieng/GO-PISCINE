package main

import "strings"

func ArrayToString(s []string) string {
	res := ""
	for i, char := range s {
		if char != " " && (i != len(s)-1) {
			res += char + " "
		} else if char != " " && (i == len(s)-1) {
			res += char
		}
	}

	return strings.TrimSpace(res)
}
