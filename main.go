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

	go func() {
		l.tokenizeStr(os.Args[1])
	}()

	for t := range l.out {
		fmt.Println(t)
		if t.typ == ttIllegal {
			panic(fmt.Sprintf("illegal character %q", t.val))
		}
	}
}
