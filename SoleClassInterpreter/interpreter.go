package main

import (
	"fmt"
	"regexp"
)

type TokenType int

const (
	ASSIGN TokenType = iota
	COMMA
	DIVISION
	EOF
	EQ
	FALSE
	TRUE
	FUNCTION
	GT
	GTE
	IDENT
	PLUS
	ILLEGAL
	INT
	LET
	LT
	NOE
	NOT
	SEMICOLON
	MINUS
	MULTI
	RETURN
	IF
	ELSE
	CORCHETEI
	CORCEHTED
	NULL
)

var TokenName = []string{
	"ASSIGN",
	"COMMA",
	"DIVISION",
	"EOF",
	"EQ",
	"FALSE",
	"TRUE",
	"FUNCTION",
	"GT",
	"GTE",
	"IDENT",
	"PLUS",
	"ILLEGAL",
	"INT",
	"LET",
	"LT",
	"NOE",
	"NOT",
	"SEMICOLON",
	"MINUS",
	"MULTI",
	"RETURN",
	"IF",
	"ELSE",
	"CORCHETEI",
	"CORCEHTED",
	"NULL",
}

func (t TokenType) String() string {
	return TokenName[t]
}

type Token struct {
	tp      TokenType
	Literal string
}

func lookUpTokenType(literal string) TokenType {
	keywords := map[string]TokenType{
		"false":    FALSE,
		"true":     TRUE,
		"func":     FUNCTION,
		"function": FUNCTION,
		"return":   RETURN,
		"if":       IF,
		"else":     ELSE,
		"let":      LET,
	}

	if tok, ok := keywords[literal]; ok {
		return tok
	}
	return IDENT
}

type Lexer struct {
	source         string
	currentPos     int
	currentChar    string
	readCurrentPos int
}

// NewLexer crea una nueva instancia de Lexer.
func newLexer(source string) Lexer {
	lexer := Lexer{source: source}
	lexer.readCharacter()
	return lexer
}
func (l Lexer) readCharacter() {
	if l.readCurrentPos >= len(l.source) {
		l.currentChar = ""
	} else {
		l.currentChar = string(l.source[l.readCurrentPos])
	}
	l.currentPos = l.readCurrentPos
	l.readCurrentPos++
}

// skipWhitespace evita los espacios en blanco.
func (l Lexer) skipWhitespace() {
	for l.currentChar == " " || l.currentChar == "\t" || l.currentChar == "\n" || l.currentChar == "\r" {
		l.readCharacter()
	}
}

// peekCharacter lee el siguiente carácter sin avanzar el cursor.
func (l Lexer) peekCharacter() string {
	if l.readCurrentPos >= len(l.source) {
		return ""
	}
	return string(l.source[l.readCurrentPos])
}

func (l Lexer) isLetter(character string) bool {
	match, _ := regexp.MatchString(`^[a-zA-ZáéíóúÁÉÍÓÚñÑ_]$`, character)
	return match
}

// isNumber evalúa si el carácter es un número.
func (l Lexer) isNumber(character string) bool {
	match, _ := regexp.MatchString(`^\d$`, character)
	return match
}

// readIdentifier lee y devuelve identificadores.
func (l Lexer) readIdentifier() string {
	initialPosition := l.currentPos
	isFirstLetter := true
	for l.isLetter(l.currentChar) || (!isFirstLetter && l.isNumber(l.currentChar)) {
		l.readCharacter()
		isFirstLetter = false
	}
	return l.source[initialPosition:l.currentPos]
}

func (l Lexer) readNumber() string {
	initialPosition := l.currentPos
	for l.isNumber(l.currentChar) {
		l.readCharacter()
	}
	return l.source[initialPosition:l.currentPos]
}

// evalua el caracter para darle valor al token y su literal
func next_token(l Lexer, t Token) Token {

	if l.currentChar == "=" {
		t.tp = ASSIGN
		t.Literal = "="
	} else if l.currentChar == "+" {
		t.tp = PLUS
		t.Literal = "+"
	} else if l.currentChar == "," {
		t.tp = COMMA
		t.Literal = ","
	} else if l.currentChar == ";" {
		t.tp = SEMICOLON
		t.Literal = ";"
	} else if l.currentChar == "" {
		t.tp = EOF
	} else if l.currentChar == "{" {
		t.tp = CORCHETEI
		t.Literal = "{"
	} else if l.currentChar == "}" {
		t.tp = CORCEHTED
		t.Literal = "}"
	} else if l.currentChar == "-" {
		t.tp = MINUS
		t.Literal = "-"
	} else if l.currentChar == "/" {
		t.tp = DIVISION
		t.Literal = "/"
	} else if l.currentChar == "*" {
		t.tp = MULTI
		t.Literal = "*"
	} else {
		t.tp = ILLEGAL
	}
	l.readCharacter()
	return t
}

func startRepl() {
	fmt.Println("Bienvenido a nuestro lenguaje de programacion")
	var firstInput string
	fmt.Scanln(firstInput)
	l := newLexer(firstInput)
	t := Token{
		tp:      NULL,
		Literal: "",
	}
	for l.source != "end" {
		fmt.Scanln(l.source)
		t := next_token(l, t)
		for t.tp != EOF {
			fmt.Println()
		}

	}
}

func main() {

	startRepl()
}
