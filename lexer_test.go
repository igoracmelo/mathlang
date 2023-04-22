package main

import (
	"strings"
	"testing"
)

func TestTokensMustBeEquivToInput(t *testing.T) {
	wants := []string{
		"1 + 2 * 5 / ( 5 - 2 ) - 0 ",
		"1 + ",
	}

	for _, want := range wants {
		l := lexer{
			in:  make(chan rune),
			out: make(chan token),
		}

		go l.tokenizeStr(want)
		tokensVals := []string{}

		for t := range l.out {
			tokensVals = append(tokensVals, t.val)
		}

		got := strings.Join(tokensVals, " ")

		if got != want {
			t.Errorf("want: '%s', got: '%s'", want, got)
		}
	}
}
