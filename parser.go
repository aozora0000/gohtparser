package gohtparser

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/aozora0000/gohtparser/token"
	"io"
	"strings"
)

type Parser struct {
	file *bufio.Scanner
}

func NewParser(file io.Reader) *Parser {
	parser := &Parser{
		file: bufio.NewScanner(file),
	}
	return parser
}

func (s *Parser) scan() bool {
	return s.file.Scan()
}

func (s *Parser) text() string {
	str := strings.TrimSpace(s.file.Text())
	return str
}

func (s *Parser) multiLine(line string, linebreaks []int) (string, []int) {
	for isMultiLine(line) && s.scan() {
		linebreaks = append(linebreaks, len(line))
		line = strings.TrimRight(line, "\\")
		line = strings.TrimSpace(line + s.text())
	}
	return line, linebreaks
}

func (s *Parser) parse(line string) (token.Token, error) {
	var linebreaks []int

	if isMultiLine(line) {
		line, linebreaks = s.multiLine(line, linebreaks)
	}

	switch {
	case isBlock(line):
		return s.generateBlockLine(line, linebreaks)
	case isWhiteLine(line):
		return s.generateWhiteLine()
	case IsComment(line):
		return s.generateCommentLine(line, linebreaks)
	case isDirective(line):
		return s.generateDirectiveLine(line, linebreaks)

	default:
		return nil, fmt.Errorf("SyntaxError: %s", line)
	}
}

func (s *Parser) generateBlockLine(line string, linebreaks []int) (token.Token, error) {
	args := blockRegex(line)
	name, args := GetName(args)
	if name == "" {
		return nil, errors.New("Could not parse the name of the block")
	}
	t := token.NewBlock(name, args)
	t.SetName(name)
	t.SetArguments(args)
	s.file.Scan()

	for !isBlockEnd(line, name) {
		line = s.text()
		tk, err := s.parse(line)
		if err == nil {
			t.AddChild(tk)
		}
		s.file.Scan()
	}
	return t, nil
}

func (s *Parser) generateDirectiveLine(line string, linebreak []int) (token.Token, error) {
	args := directiveRegex(line)
	name, args := GetName(args)

	t := token.NewDirective(name, args)
	t.SetLineBreaks(linebreak)
	return t, nil
}

func (s *Parser) generateWhiteLine() (token.Token, error) {
	return token.NewWhiteLine(), nil
}

func (s *Parser) generateCommentLine(line string, linebreaks []int) (token.Token, error) {
	comment := token.NewComment(line)
	comment.SetLineBreaks(linebreaks)
	return comment, nil
}

func (s *Parser) GenerateAst() (*Ast, error) {
	ast := NewAst()
	for s.scan() {
		t, err := s.parse(s.text())
		if err != nil {
			return nil, err
		}
		ast.Add(t)
	}
	return ast, nil
}
