package parser

import (
	"donkey/ast"
	"donkey/lexer"
	"testing"
)

func TestDonkStatement(t *testing.T) {
	input := `
	donk x = 5;
	donk y = 10;
	donk foobar = 838383;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	total_staements := len(program.Statements)
	if total_staements != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", total_staements)
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testDonkStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testDonkStatement(t *testing.T, s ast.Statement, name string) bool {
	literal := s.TokenLiteral()
	if literal != "donk" {
		t.Errorf("s.TokenLiteral not 'donk'. got=%q", literal)
	}

	donkStmt, ok := s.(*ast.DonkStatement)
	if !ok {
		t.Errorf("s not *asf.DonkStatement. got=%T", s)
		return false
	}

	val := donkStmt.Name.Value
	if val != name {
		t.Errorf("donkStmt.Name.Value not '%s'. got=%s", name, val)
		return false
	}

	tokenLiteral := donkStmt.Name.TokenLiteral()
	if tokenLiteral != name {
		t.Errorf("donkStmt.Name.TokenLiteral() not '%s'. got=%s", name, tokenLiteral)
		return false
	}

	return true
}

func TestReturnStatement(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 929292;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	total_statements := len(program.Statements)
	if total_statements != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", total_statements)
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement. got=%T", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return'. got=%q", returnStmt.TokenLiteral())
		}
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	total_erros := len(errors)
	if total_erros == 0 {
		return
	}

	t.Errorf("parser has %d errors", total_erros)

	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}

	t.FailNow()
}
