package interpreter

import "fmt"

func StartRepl() {
	fmt.Println("Bienvenido a nuestro martitrio")
	var firstInput string
	fmt.Scanln(firstInput)
	l := Lexer{}
	t := Token{}
	
	for l.currentChar != "end"{
		fmt.Printf(">>>")
		fmt.Scanln(&l.currentChar)
		next_token(l, t)
		fmt.Println(t.tp)
		fmt.Println(t.Literal)
	}
}

