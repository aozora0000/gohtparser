package token

type WhiteLine struct {
	name       string
	linebreaks []int
}

func (s *WhiteLine) IsMultiline() bool {
	return len(s.linebreaks) != 0
}

func (s *WhiteLine) GetLineBreaks() []int {
	return s.linebreaks
}

func (s *WhiteLine) SetLineBreaks(linebreaks []int) {
	s.linebreaks = linebreaks
}

func (s *WhiteLine) AddLineBreak(linebreak int) {
	s.linebreaks = append(s.linebreaks, linebreak)
}

func (s *WhiteLine) GetTokenType() int {
	return TOKEN_WHITELINE
}

func (s *WhiteLine) SetName(name string) {
	s.name = name
}

func (s *WhiteLine) GetName() string {
	return s.name
}

func (s *WhiteLine) GetArguments() []string {
	return []string{}
}

func (s *WhiteLine) SetArguments(arguments []string) {

}

func (s *WhiteLine) GetValue() string {
	return ""
}

func (s *WhiteLine) ToString() string {
	return ""
}

func NewWhiteLine() *WhiteLine {
	return &WhiteLine{
		name: "WhiteLine",
	}
}
