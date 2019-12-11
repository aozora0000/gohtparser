package token

const (
	TOKEN_DIRECTIVE int = iota
	TOKEN_BLOCK     int = iota
	TOKEN_COMMENT   int = iota
	TOKEN_WHITELINE int = iota
)

type Token interface {
	IsMultiline() bool
	GetLineBreaks() []int
	SetLineBreaks(linebreaks []int)
	AddLineBreak(linebreak int)
	GetTokenType() int
	SetName(name string)
	GetName() string
	GetArguments() []string
	SetArguments(arguments []string)
	GetValue() string
	ToString() string
}
