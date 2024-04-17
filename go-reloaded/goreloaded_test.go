package main

import "testing"

// arg means argument and the expected stands for the 'result we expect'
type sentencesTest struct {
	arg, expected string
}

var addTests = []sentencesTest{
	// Add your argument and expected result
	{"it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair.", "It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, it was the winter of despair."},
	{"Simply add 42 (hex) and 10 (bin) and you will see the result is 68.", "Simply add 66 and 2 and you will see the result is 68."},
	{"There is no greater agony than bearing a untold story inside you.", "There is no greater agony than bearing an untold story inside you."},
	{"Punctuation tests are ... kinda boring ,don't you think !?", "Punctuation tests are... kinda boring, don't you think!?"},
}

func TestSentence(t *testing.T) {
	for _, test := range addTests {
		if output := CreateSentence(test.arg); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
