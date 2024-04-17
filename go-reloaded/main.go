package main

import (
	"fmt"
	"os"
	"strings"
)

func CreateSentence(x string) string {
	words := Vowels(x)
	converted := Convert(words)

	newSentence := ArrayToString(converted)

	// Convert again to Array
	a := ToArray(newSentence)

	// handle punctuation of (,.:;?!) within the code
	c := Punctuate1(a)
	d := SingleQuote1(c)
	e := SingleQuote2(RemoveEmptyValues(d))
	xy := RemoveEmptyValues(e)
	return strings.Join(xy, " ")
}

func main() {
	filepath := os.Args[1:]
	if len(filepath) < 2 {
		fmt.Println("Less arguments provided")
	} else if len(filepath) > 2 {
		fmt.Println("Too many arguments passed")
	}
	b, err := os.ReadFile(filepath[0])
	if err != nil {
		fmt.Print(err)
	}
	filecontent := string(b)
	sentence := CreateSentence(filecontent)
	newFile, _ := os.Create(filepath[1])

	// Write to a file
	newFile.WriteString(sentence)
	defer newFile.Close()
}
