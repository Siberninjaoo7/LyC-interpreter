package interpreter

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
