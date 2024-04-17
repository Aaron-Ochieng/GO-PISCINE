package main

import "strings"

func Title(s string) string {
	val := strings.ToUpper(string(s[0]))
	return val + s[1:]
}
