package ast

import (
	"donkey/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&DonkStatement{
				Token: token.Token{Type: token.DONK, Literal: "donk"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "fooVar"},
					Value: "fooVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "barVar"},
					Value: "barVar",
				},
			},
		},
	}

	if program.String() != "donk fooVar = barVar;" {
		t.Errorf("program.String() is wrong. got=%q", program.String())
	}
}
