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

func newToken(t TokenType, lit string) Token {
	tok := Token{
		tp:      t,
		Literal: lit,
	}
	return tok
}
func strToken(t *Token) string {
	return t.Literal
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
func newLexer(source string) *Lexer {
	lexer := Lexer{source: source}
	lexer.readCharacter()
	return &lexer
}
func (l *Lexer) readCharacter() {
	if l.readCurrentPos >= len(l.source) {
		l.currentChar = ""
	} else {
		l.currentChar = string(l.source[l.readCurrentPos])
	}
	l.currentPos = l.readCurrentPos
	l.readCurrentPos++
}

// skipWhitespace evita los espacios en blanco.
func (l *Lexer) skipWhitespace() {
	for l.currentChar == " " || l.currentChar == "\t" || l.currentChar == "\n" || l.currentChar == "\r" {
		l.readCharacter()
	}
}

// peekCharacter lee el siguiente carácter sin avanzar el cursor.
func (l *Lexer) peekCharacter() string {
	if l.readCurrentPos >= len(l.source) {
		return ""
	}
	return string(l.source[l.readCurrentPos])
}

func (l *Lexer) isLetter(character string) bool {
	match, _ := regexp.MatchString(`^[a-zA-ZáéíóúÁÉÍÓÚñÑ_]$`, character)
	return match
}

// isNumber evalúa si el carácter es un número.
func (l *Lexer) isNumber(character string) bool {
	match, _ := regexp.MatchString(`^\d$`, character)
	return match
}

// readIdentifier lee y devuelve identificadores.
func (l *Lexer) readIdentifier() string {
	initialPosition := l.currentPos
	isFirstLetter := true
	for l.isLetter(l.currentChar) || (!isFirstLetter && l.isNumber(l.currentChar)) {
		l.readCharacter()
		isFirstLetter = false
	}
	return l.source[initialPosition:l.currentPos]
}

func (l *Lexer) readNumber() string {
	initialPosition := l.currentPos
	for l.isNumber(l.currentChar) {
		l.readCharacter()
	}
	return l.source[initialPosition:l.currentPos]
}

// evalua el caracter para darle valor al token y su literal
func next_token(l Lexer, t Token) {

	if l.currentChar == "=" {
		fmt.Println("ASSIGN"+"=")
	} else if l.currentChar == "+" {
		t = newToken(PLUS, "+")
	} else if l.currentChar == "," {
		t = newToken(COMMA,",")
	} else if l.currentChar == ";" {
		t = newToken(SEMICOLON, ";")
	} else if l.currentChar == "" {
		t = newToken(EOF,"EOF")
	} else if l.currentChar == "{" {
		t = newToken(CORCHETEI, "{")
	} else if l.currentChar == "}" {
		t = newToken(CORCEHTED, "}")
	} else if l.currentChar == "-" {
		t = newToken(MINUS,"-")
	} else if l.currentChar == "/" {
		t = newToken(DIVISION, "/")
	} else if l.currentChar == "*" {
		t = newToken(MULTI, "*")
	} else {
		t = newToken(ILLEGAL,"ILLEGAL")
	}
	l.readCharacter()
	fmt.Println(t.tp)
	fmt.Println(t.Literal)
	
}

func startRepl() {
	fmt.Println("Bienvenido a nuestro martitrio")
	var firstInput string
	fmt.Scanln(firstInput)
	l := newLexer(firstInput)
	t := newToken(NULL,"NULL") 
	for l.source != "end" {
		fmt.Printf(">>>")
		fmt.Scanln(&l.source)
		next_token(*l, t)
		
	}
}

func main() {
	startRepl()
}
