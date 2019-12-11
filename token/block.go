package token

import (
	"fmt"
	"strings"
)

type Block struct {
	name        string
	arguments   []string
	children    []Token
	indentation int
	linebreaks  []int
}

func (s *Block) IsMultiline() bool {
	return len(s.linebreaks) != 0
}

func (s *Block) GetLineBreaks() []int {
	return s.linebreaks
}

func (s *Block) SetLineBreaks(linebreaks []int) {
	s.linebreaks = linebreaks
}

func (s *Block) AddLineBreak(linebreak int) {
	s.linebreaks = append(s.linebreaks, linebreak)
}

func (s *Block) SetName(name string) {
	s.name = name
}

func (s *Block) SetArguments(arguments []string) {
	s.arguments = arguments
}

func NewBlock(name string, argument []string) *Block {
	return &Block{name: name, arguments: argument, indentation: 4}
}

func (s *Block) GetTokenType() int {
	return TOKEN_BLOCK
}

func (s *Block) GetName() string {
	return s.name
}

func (s *Block) GetArguments() []string {
	return s.arguments
}

func (s *Block) GetValue() string {
	return strings.Join(s.GetArguments(), " ")
}

func (s *Block) AddArgument(argument string) {
	s.arguments = append(s.arguments, argument)
}

func (s *Block) RemoveArgument(key string) {
	var args []string
	for i, val := range s.arguments {
		if val != key {
			args = append(args, s.arguments[i])
		}
	}
	s.arguments = args
}

func (s *Block) AddChild(token Token) {
	s.children = append(s.children, token)
}

func (s *Block) RemoveChild(token Token) {
	children := []Token{}
	for i, t := range s.children {
		if t != token {
			children = append(children, s.children[i])
		}
	}
	s.children = children
}

func (s *Block) CountChild() int {
	return len(s.children)
}

func (s *Block) HasChild() bool {
	return s.CountChild() != 0
}

func (s *Block) SetIndentation(space int) {
	s.indentation = space
}

func (s *Block) ToString() string {
	block := ""
	block = block + fmt.Sprintf("<%s %s>\n", s.GetName(), strings.Join(s.arguments, " "))
	for _, child := range s.children {
		block = block + strings.Repeat(" ", s.indentation) + child.ToString() + "\n"
	}
	block = block + fmt.Sprintf("</%s>\n", s.GetName())
	return block
}
