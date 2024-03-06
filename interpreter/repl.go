package main

import "fmt"

func start_repl(){
	fmt.Println("Buenas mi estimado")
	l := Lexer{

	}
	t := Token{

	}
	for l.source != "end"{
     for next_token(l , t).tokenType != EOF{
       fmt.Println(next_token(l,t).Literal)
	 }
	}

}