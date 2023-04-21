package main

import (
	"os"
	"fmt"
	"unicode"
)

type tokenType uint

const (
	ttIllegal tokenType = iota
	ttNum
	ttAdd
	ttSub
	ttMul
	ttDiv
	ttLParen
	ttRParen
	ttEOF
)

var ttStr = []string{
	ttIllegal: "ttIllegal",
	ttNum:     "ttNum",
	ttAdd:     "ttAdd",
	ttSub:     "ttSub",
	ttMul:     "ttMul",
	ttDiv:     "ttDiv",
	ttLParen:  "ttParenL",
	ttRParen:  "ttParenR",
	ttEOF:     "ttEOF",
}

type token struct {
	typ tokenType
	val string
}

func (t token) String() string {
	return fmt.Sprintf("{ %s `%s` }", ttStr[t.typ], t.val)
}

func main() {
	if len(os.Args) < 2 {
		panic("missing math expression")
	}
	code := []rune(os.Args[1])

	i := 0
	for i <= len(code) {
		t, d := nextToken(code[i:])
		i += d + 1
		fmt.Println(t)
		if t.typ == ttEOF {
			break
		}
	}
}

func nextToken(rs []rune) (token, int) {
	var r rune
	var i int

	for i, r = range rs {
		if unicode.IsSpace(r) {
			continue
		}

		var tt tokenType

		if unicode.IsDigit(r) {
			tt = ttNum
			for j := i + 1; j < len(rs); j++ {
				r := rs[j]
				if !unicode.IsDigit(r) {
					break
				}
				i++
			}
			return token{ttNum, string(rs[:i+1])}, i
		}

		switch r {
		case '+':
			tt = ttAdd
		case '-':
			tt = ttSub
		case '*':
			tt = ttMul
		case '/':
			tt = ttDiv
		case '(':
			tt = ttLParen
		case ')':
			tt = ttRParen
		default:
			tt = ttIllegal
		}

		return token{tt, string(r)}, i
	}

	return token{ttEOF, ""}, i
}
