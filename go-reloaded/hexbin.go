package main

import (
	"fmt"
	"strconv"
)

// Converts a given string from binary to int64
func Bindec(s string) int64 {
	value, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return value
}

// Converts a given tring from hexadecimal to int64
func Hexdec(s string) int64 {
	value, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		fmt.Println(err)
	}
	return value
}
