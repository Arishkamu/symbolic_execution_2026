package translator

import (
	. "symbolic-execution-course/internal/symbolic"
	"testing"

	"github.com/ebukreev/go-z3/z3"
)

func translateValue(t *testing.T, translator *Z3Translator, expr SymbolicExpression) z3.Value {
	res, err := translator.TranslateExpression(expr)
	if err != nil {
		t.Errorf("Error translating expression: %v", err)
	}
	resValue, ok := res.(z3.Value)
	if !ok {
		t.Errorf("Error asserting result to z3.Value: %v", err)
	}
	return resValue
}

func TestTranslatedStringAndType(t *testing.T) {
	translator := NewZ3Translator()
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

	tests := []struct {
		expr SymbolicExpression
		str  string
		tpe  z3.Kind
	}{
		{x, "x", z3.KindInt},
		{a, "a", z3.KindBool},
		{const10, "10", z3.KindInt},
		{constFalse, "false", z3.KindBool},

		{sum, "(+ x y)", z3.KindInt},
		{mod, "(mod x 10)", z3.KindInt},
		{eq, "(= x y)", z3.KindBool},
		{eqBool, "(= a b)", z3.KindBool},
		{le, "(<= 10 y)", z3.KindBool},
		{log, "(and a b)", z3.KindBool},
	}

	for _, tc := range tests {
		translator.Reset()
		resValue := translateValue(t, translator, tc.expr)
		if resValue.Sort().Kind() != tc.tpe {
			t.Errorf("Expected type %v, got %v", tc.tpe, resValue.Sort().Kind())
		}
		if resValue.String() != tc.str {
			t.Errorf("Expected string %v, got %v", tc.str, resValue.String())
		}
	}
}

func checkAssert(t *testing.T, translator *Z3Translator, expr SymbolicExpression) bool {
    res, err := translator.TranslateExpression(expr)
    if err != nil {
        t.Errorf("Error translating expression: %v", err)
        return false
    }

	solver := z3.NewSolver(translator.GetContext().(*z3.Context))
	solver.Assert(res.(z3.Bool))
	ch, err := solver.Check()
	if err != nil {
		t.Errorf("Checker on %s, error: %v", expr.String(), err)
	}
	return ch
}

func expectSat(t *testing.T, translator *Z3Translator, expr SymbolicExpression, expected bool) {
    if checkAssert(t, translator, expr) != expected {
        t.Errorf("Expected expression %s to be %v", expr.String(), expected)
    }
}

func TestTranslateOperations(t *testing.T) {
	translator := NewZ3Translator()

	ci := func(i int64) *IntConstant { return NewIntConstant(i) }
	cFalse := NewBoolConstant(false)
	cTrue := NewBoolConstant(true)

	binOp := func(left, right SymbolicExpression, op BinaryOperator) *BinaryOperation {
		return NewBinaryOperation(left, right, op)
	}
	equal := func(left, right SymbolicExpression) *BinaryOperation { return binOp(left, right, EQ) }

	for _, tExpr := range []SymbolicExpression{
		equal(binOp(ci(40), ci(2), ADD), ci(42)),
		equal(binOp(ci(6), ci(7), MUL), ci(42)),
		equal(binOp(ci(-6), ci(2), DIV), ci(-6/2)),
		equal(binOp(ci(239), ci(566), NE), cTrue),
		equal(binOp(ci(239), ci(566), LE), cTrue),
		equal(binOp(ci(123), ci(456), GT), cFalse),
		equal(binOp(cTrue, cFalse, NE), cTrue),
	} {
		translator.Reset()
		expectSat(t, translator, tExpr, true)
	}
	
}

func TestTranslateFormula(t *testing.T) {
	translator := NewZ3Translator()
	x := NewSymbolicVariable("x", IntType)
	y := NewSymbolicVariable("y", IntType)
	z := NewSymbolicVariable("z", IntType)

	ci := func(i int64) *IntConstant { return NewIntConstant(i) }
	cTrue := NewBoolConstant(true)

	binOp := func(left, right SymbolicExpression, op BinaryOperator) *BinaryOperation {
		return NewBinaryOperation(left, right, op)
	}
	logOp := func(left, right SymbolicExpression, op LogicalOperator) *LogicalOperation {
		return NewLogicalOperation([]SymbolicExpression{left, right}, op)
	}
	equal := func(left, right SymbolicExpression) *BinaryOperation { return binOp(left, right, EQ) }

	// x > y && x < 10 && x == 7 && y > 100
	chainExprs := []SymbolicExpression{
		binOp(x, y, GT),
		binOp(x, ci(10), LT),
		binOp(x, ci(7), EQ),
	}

	for _, tExpr := range []SymbolicExpression{
		equal(logOp(logOp(binOp(x, y, LT), binOp(y, z, LT), AND), binOp(x, z, LT), IMPLIES), cTrue),
		binOp(binOp(x, ci(10), ADD), x, GT),
		NewLogicalOperation(chainExprs, AND),
	} {
		translator.Reset()
		expectSat(t, translator, tExpr, true)
	}

	chainExprs = append(chainExprs, binOp(y, ci(100), GT))
	expectSat(t, translator, NewLogicalOperation(chainExprs, AND), false)
	
}

func TestIdentety(t *testing.T) {
	translator := NewZ3Translator()
	x1 := translateValue(t, translator, NewSymbolicVariable("x", IntType))
	x2 := translateValue(t, translator, NewSymbolicVariable("x", IntType))
	
	if !x1.AsAST().Equal(x2.AsAST()) {
		t.Errorf("Error. Expected two variables with same name have same identety")
	}

	y := NewSymbolicVariable("y", IntType)
	plusZero := NewBinaryOperation(y, NewIntConstant(0), ADD)

	yVal := translateValue(t, translator, y)
	plusZeroVal := translateValue(t, translator, plusZero)
	if yVal.AsAST().Equal(plusZeroVal.AsAST()) {
		t.Errorf("Error. Expected two expressions with same value but different history to be different")
	}

	expectSat(t, translator, NewBinaryOperation(y, plusZero, EQ), true)
}
