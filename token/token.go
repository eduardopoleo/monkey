package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

// what is const
const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	ASSING    = "="
	PLUS      = "+"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "}"
	RBRACE    = "{"
	FUNCTION  = "FUNCTION"
	LET       = "LET"
)
