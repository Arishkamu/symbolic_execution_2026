package symbolic

import (
	"strings"
	"testing"
)

func TestStringAndType(t *testing.T) {
	x := NewSymbolicVariable("x", IntType)
	y := NewSymbolicVariable("y", IntType)
	a := NewSymbolicVariable("a", BoolType)
	b := NewSymbolicVariable("b", BoolType)

	const10 := NewIntConstant(10)
	constFalse := NewBoolConstant(false)

	sum := NewBinaryOperation(x, y, ADD)
	mod := NewBinaryOperation(x, const10, MOD)
	eq := NewBinaryOperation(x, y, EQ)
	eqBool := NewBinaryOperation(a, b, EQ)
	le := NewBinaryOperation(const10, y, LE)

	log := NewLogicalOperation([]SymbolicExpression{a, b}, AND)

	complicated1 := NewBinaryOperation(sum, const10, GT)
	complicated2 := NewLogicalOperation([]SymbolicExpression{complicated1, NewLogicalOperation([]SymbolicExpression{constFalse}, NOT), a}, AND)
	complicated3 := NewBinaryOperation(sum, mod, SUB)

	testPrimitives := []struct {
		expr SymbolicExpression
		str  string
		tpe  ExpressionType
	}{
		{x, "x", IntType},
		{a, "a", BoolType},
		{const10, "10", IntType},
		{constFalse, "false", BoolType},

	}
	testOps := []struct {
		expr SymbolicExpression
		str  string
		tpe  ExpressionType
	}{
		{sum, "(x + y)", IntType},
		{mod, "(x % 10)", IntType},
		{eq, "(x == y)", BoolType},
		{eqBool, "(a == b)", BoolType},
		{le, "(10 <= y)", BoolType},
		{log, "(a && b)", BoolType},
	}
	testComplicated := []struct {
		expr SymbolicExpression
		str  string
		tpe  ExpressionType
	}{
		{complicated1, "((x + y) > 10)", BoolType},
		{complicated2, "(((x + y) > 10) && !false && a)", BoolType},
		{complicated3, "((x + y) - (x % 10))", IntType},
	}
	for _, res := range testPrimitives {
		if res.expr.String() != res.str || res.expr.Type() != res.tpe {
			t.Errorf("Error testing primitive expression.\n  Expected: type=%s; string=%s\n       Got: type=%s; string=%s", res.tpe, res.str, res.expr.Type(), res.expr.String())
		}
	}
	for _, res := range testOps {
		if res.expr.String() != res.str || res.expr.Type() != res.tpe {
			t.Errorf("Error testing operation expression.\n  Expected: type=%s; string=%s\n       Got: type=%s; string=%s", res.tpe, res.str, res.expr.Type(), res.expr.String())
		}
	}
	for _, res := range testComplicated {
		if res.expr.String() != res.str || res.expr.Type() != res.tpe {
			t.Errorf("Error testing complicated expression.\n  Expected: type=%s; string=%s\n       Got: type=%s; string=%s", res.tpe, res.str, res.expr.Type(), res.expr.String())
		}
	}
}

func TestPanics(t *testing.T) {
	x := NewSymbolicVariable("x", IntType)
	a := NewSymbolicVariable("a", BoolType)

	panicCheck := func(panicTag string, fn func()) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic %s, but no panic occurred", panicTag)
			}
		}()
		fn()
	}

	panicCheck("int + nil", func() { NewBinaryOperation(x, nil, ADD) })
	panicCheck("int + bool", func() { NewBinaryOperation(x, a, ADD) })
	panicCheck("bool + bool", func() { NewBinaryOperation(a, a, ADD) })
	panicCheck("bool >= bool", func() { NewBinaryOperation(a, a, GE) })
	
	panicCheck("a &&", func() { NewLogicalOperation([]SymbolicExpression{a}, AND) })
	panicCheck("!a a", func() { NewLogicalOperation([]SymbolicExpression{a, a}, NOT) })
	panicCheck("a => a a", func() { NewLogicalOperation([]SymbolicExpression{a, a, a}, IMPLIES) })
}

func TestVisitorStringPrinter(t *testing.T) {
	x := NewSymbolicVariable("x", IntType)
	y := NewSymbolicVariable("y", IntType)
	a := NewSymbolicVariable("a", BoolType)

	// ((x % 10 == 0) || (x + y) > 10) || !true || a
	expr := NewLogicalOperation([]SymbolicExpression{
		NewBinaryOperation(NewBinaryOperation(x, NewIntConstant(10), MOD), NewIntConstant(0), EQ),
	    NewBinaryOperation(NewBinaryOperation(x, y, ADD), NewIntConstant(10), GT),
		NewLogicalOperation([]SymbolicExpression{NewBoolConstant(true)}, NOT),
		a,
		}, OR)
	strPr := &StringPrinter{result: strings.Builder{}}
	expr.Accept(strPr)
	result := strPr.result.String()
	if result != "(((x % 10) == 0) || ((x + y) > 10) || !true || a)"{
		t.Errorf("incorrect visitor StringPrinter: %s", result)
	}
}