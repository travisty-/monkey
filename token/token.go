package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	Illegal = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	Ident = "IDENT" // add, foo, x, y, ...
	Int   = "INT"   // 123456

	// Operators
	Assign   = "="
	Plus     = "+"
	Minus    = "-"
	Bang     = "!"
	Asterisk = "*"
	Slash    = "/"

	LT = "<"
	GT = ">"

	Eq    = "=="
	NotEq = "!="

	// Delimiters
	Comma     = ","
	Semicolon = ";"

	LParen = "("
	RParen = ")"
	LBrace = "{"
	RBrace = "}"

	// Keywords
	Function = "FUNCTION"
	Let      = "LET"
	True     = "TRUE"
	False    = "FALSE"
	If       = "IF"
	Else     = "ELSE"
	Return   = "RETURN"
)

var keywords = map[string]Type{
	"fn":     Function,
	"let":    Let,
	"true":   True,
	"false":  False,
	"if":     If,
	"else":   Else,
	"return": Return,
}

func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return Ident
}
