package interpreter

import (
	"regexp"
	"unicode/utf8"
)

type Lexer struct {
	source         string
	currentPos     int
	currentChar    string
	readCurrentPos int
}

// NewLexer crea una nueva instancia de Lexer.
func newLexer(source string) *Lexer {
	mutlexer := &Lexer{source: source}
	mutlexer.ReadCharacter()
	return mutlexer
}
func (l *Lexer) ReadCharacter() {
	if l.readCurrentPos >= utf8.RuneCountInString(l.source) {
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
		l.ReadCharacter()
	}
}

// peekCharacter lee el siguiente carácter sin avanzar el cursor.
func (l *Lexer) peekCharacter() string {
	if l.readCurrentPos >= len(l.source) {
		return ""
	}
	return string(l.source[l.readCurrentPos])
}

func (l *Lexer) IsLetter(character string) bool {
	match, _ := regexp.MatchString(`^[a-zA-ZáéíóúÁÉÍÓÚñÑ_]$`, character)
	return match
}

// isNumber evalúa si el carácter es un número.
func (l *Lexer) IsNumber(character string) bool {
	match, _ := regexp.MatchString(`^\d$`, character)
	return match
}

// readIdentifier lee y devuelve identificadores.
func (l *Lexer) ReadIdentifier() string {
	initialPosition := l.currentPos
	isFirstLetter := true
	for l.IsLetter(l.currentChar) || (!isFirstLetter && l.IsNumber(l.currentChar)) {
		l.ReadCharacter()
		isFirstLetter = false
	}
	return l.source[initialPosition:l.currentPos]
}

func (l *Lexer) ReadNumber() string {
	initialPosition := l.currentPos
	for l.IsNumber(l.currentChar) {
		l.ReadCharacter()
	}
	return l.source[initialPosition:l.currentPos]
}
func next_token(l Lexer, t Token) Token {

	l.skipWhitespace()
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
	l.ReadCharacter()
 return t
}
