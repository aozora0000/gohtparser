package token

import "strings"

type Comment struct {
	text       string
	linebreaks []int
}

func (s *Comment) IsMultiline() bool {
	return len(s.linebreaks) != 0
}

func (s *Comment) GetLineBreaks() []int {
	return s.linebreaks
}

func (s *Comment) SetLineBreaks(linebreaks []int) {
	s.linebreaks = linebreaks
}

func (s *Comment) AddLineBreak(linebreak int) {
	s.linebreaks = append(s.linebreaks, linebreak)
}

func (s *Comment) GetTokenType() int {
	return TOKEN_COMMENT
}

func (s *Comment) SetName(name string) {

}

func (s *Comment) GetName() string {
	return "#comment"
}

func (s *Comment) GetArguments() []string {
	return []string{s.text}
}

func (s *Comment) SetArguments(arguments []string) {
	s.SetText(arguments[0])
}

func (s *Comment) GetValue() string {
	return s.text
}

func (s *Comment) ToString() string {
	return s.text
}

func (s *Comment) SetText(text string) {
	s.text = "# " + strings.TrimPrefix(text, "#")
}

func NewComment(text string) *Comment {
	return &Comment{text: text}
}
