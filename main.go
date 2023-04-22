package main

import (
	"fmt"
	"os"
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

type lexer struct {
	in      chan rune
	out     chan token
	seqStr  string
	seqType tokenType
}

func (l *lexer) sendSeq() {
	if len(l.seqStr) == 0 {
		return
	}

	l.out <- token{l.seqType, l.seqStr}
	l.seqStr = ""
	l.seqType = ttIllegal
}

func (l *lexer) tokenize() {
	r, ok := <-l.in

	if !ok {
		l.sendSeq()
		l.out <- token{ttEOF, ""}
		close(l.out)
		return
	}

	if unicode.IsSpace(r) {
		l.sendSeq()
		l.tokenize()
		return
	}

	if unicode.IsDigit(r) {
		if l.seqType != ttNum {
			l.sendSeq()
		}
		l.seqStr += string(r)
		l.seqType = ttNum
		l.tokenize()
		return
	}

	l.sendSeq()
	var tt tokenType
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

	l.out <- token{tt, string(r)}

	if tt == ttIllegal {
		close(l.out)
		return
	}

	l.tokenize()
}

func main() {
	if len(os.Args) < 0 {
		panic("missing math expression")
	}

	l := lexer{
		in:  make(chan rune),
		out: make(chan token),
	}

	code := os.Args[1]

	go func() {
		for _, r := range code {
			l.in <- r
		}
		close(l.in)
	}()

	go func() {
		l.tokenize()
	}()

	for t := range l.out {
		fmt.Println(t)
		if t.typ == ttIllegal {
			panic(fmt.Sprintf("illegal character %q", t.val))
		}
	}
}
