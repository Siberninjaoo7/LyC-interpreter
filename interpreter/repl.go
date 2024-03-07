package main

import "fmt"

func start_repl() {
	var inpt string
	fmt.Println("Buenas mi estimado")
	l := newLexer(inpt)
	t := Token{}
	for l.source != "end" {
		for next_token(l, t).tokenType != EOF {
			fmt.Println(next_token(l, t).Literal)
		}
	}

}
