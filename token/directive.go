package token

import (
	"fmt"
	"sort"
	"strings"
)

type Directive struct {
	name       string
	arguments  []string
	linebreaks []int
}

func (s *Directive) IsMultiline() bool {
	return len(s.linebreaks) != 0
}

func (s *Directive) GetLineBreaks() []int {
	return s.linebreaks
}

func (s *Directive) SetLineBreaks(linebreaks []int) {
	s.linebreaks = linebreaks
}

func (s *Directive) AddLineBreak(linebreak int) {
	s.linebreaks = append(s.linebreaks, linebreak)
}

func (s *Directive) GetTokenType() int {
	return TOKEN_DIRECTIVE
}

func (s *Directive) SetName(name string) {
	s.name = name
}

func (s *Directive) GetName() string {
	return s.name
}

func (s *Directive) GetArguments() []string {
	return s.arguments
}

func (s *Directive) SetArguments(arguments []string) {
	for _, arg := range arguments {
		s.SetArgument(arg, false)
	}
}

func (s *Directive) SetArgument(argument string, unique bool) {
	if strings.Index(argument, " ") != -1 && strings.Index(argument, `"`) == -1 {
		argument = fmt.Sprintf(`"%s"`, argument)
	}
	if sort.SearchStrings(s.arguments, argument) == len(s.arguments) && unique {
		return
	}
	s.arguments = append(s.arguments, argument)
}

func (s *Directive) RemoveArgument(argument string) {
	index := sort.SearchStrings(s.arguments, argument)
	if index == len(s.arguments) {
		return
	}
	var args []string
	for i, val := range s.arguments {
		if val != argument {
			args = append(args, s.arguments[i])
		}
	}
	s.arguments = args
}

func (s *Directive) GetValue() string {
	return strings.Join(s.arguments, " ")
}

func (s *Directive) ToString() string {
	return strings.Join([]string{s.GetName(), s.GetValue()}, " ")
}

func NewDirective(name string, arguments []string) *Directive {
	d := &Directive{}
	d.SetName(name)
	d.SetArguments(arguments)
	return d
}
