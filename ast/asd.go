package ast

import "interpreter/token"

type Node interface {
	TokenLiteranl() string
}

type Statment interface {
	Node
	statmentNode() string
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statments []Statment
}

func (p *Program) TokenLiteranl() string {
	if len(p.Statments) > 0 {
		return p.Statments[0].TokenLiteranl()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statmentNode()         {}
func (ls *LetStatement) TokenLiteranl() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value Expression
}

func (ls *Identifier) expressionNode()       {}
func (ls *Identifier) TokenLiteranl() string { return ls.Token.Literal }
