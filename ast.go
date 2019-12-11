package gohtparser

import "github.com/aozora0000/gohtparser/token"
import "github.com/k0kubun/pp"

type Ast struct {
	tree []token.Token
}

func (s *Ast) Add(token token.Token) {
	s.tree = append(s.tree, token)
}

func (s *Ast) Dump() {
	pp.Println(s.tree)
}

func (s *Ast) ToHtAccess() chan string {
	c := make(chan string)
	go func() {
		for _, t := range s.tree {
			c <- t.ToString()
		}
		close(c)
	}()

	return c
}

func NewAst() *Ast {
	return &Ast{}
}
