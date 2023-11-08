package lexer

import (
	"testing"

	"github.com/eduardopoleo/monkey/token"
)

func TestNextToken(t *testing.T) {
	// Do not scape new lines etc
	input := `=+(){},;`

	// array of structs
	// tokenType | literal
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	// Initializing the lexer. This is a bit magical
	// because you'll need to know that is the lexer
	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test[%d] - toketype wong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tested[%d] - literal wrong. expected =%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
