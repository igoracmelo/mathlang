package main

import (
	"fmt"
	"os"
)

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
