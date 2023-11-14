package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

// This should be named lookup keyword
// Basically if it's not a keyword it's an identifier so
// return that instead
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}

// what is const
const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	ASTERISK  = "*"
	SLASH     = "/"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "}"
	RBRACE    = "{"
	FUNCTION  = "FUNCTION"
	LET       = "LET"
	LT        = "<"
	GT        = ">"
	TRUE      = "true"
	FALSE     = "false"
	IF        = "if"
	ELSE      = "else"
	RETURN    = "return"
	EQ        = "=="
	NOT_EQ    = "!="
)
