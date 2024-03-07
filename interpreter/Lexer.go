package main

import (
	"regexp"
)

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
func next_token(l Lexer, t Token) Token {

	if l.currentChar == "=" {
		t.tokenType = ASSIGN
		t.Literal = "="
	} else if l.currentChar == "+" {
		t.tokenType = PLUS
		t.Literal = "+"
	} else if l.currentChar == "," {
		t.tokenType = COMMA
		t.Literal = ","
	} else if l.currentChar == ";" {
		t.tokenType = SEMICOLON
		t.Literal = ";"
	} else if l.currentChar == "" {
		t.tokenType = EOF
	} else if l.currentChar == "{" {
		t.tokenType = CORCHETEI
		t.Literal = "{"
	} else if l.currentChar == "}" {
		t.tokenType = CORCEHTED
		t.Literal = "}"
	} else if l.currentChar == "-" {
		t.tokenType = MINUS
		t.Literal = "-"
	} else if l.currentChar == "/" {
		t.tokenType = DIVISION
		t.Literal = "/"
	} else if l.currentChar == "*" {
		t.tokenType = MULTI
		t.Literal = "*"
	} else {
		t.tokenType = ILLEGAL
	}
	l.readCharacter()
	return t
}
