package symbolic

type Simplifier struct {
    result SymbolicExpression
}

func SimplifyExpression(expr SymbolicExpression) SymbolicExpression {
    return (&Simplifier{}).Simplify(expr)
}

func (s *Simplifier) Simplify(expr SymbolicExpression) SymbolicExpression {
    return expr.Accept(s).(SymbolicExpression)
}

func (s *Simplifier) VisitVariable(expr *SymbolicVariable) interface{} {
	return expr
}

func (s *Simplifier) VisitIntConstant(expr *IntConstant) interface{} {
	return expr
}
func (s *Simplifier) VisitFloatConstant(expr *FloatConstant) interface{} {
	return expr
}

func (s *Simplifier) VisitBoolConstant(expr *BoolConstant) interface{} {
	return expr
}

func (s *Simplifier) VisitBinaryOperation(expr *BinaryOperation) interface{} {
	leftSim := s.Simplify(expr.Left)
	rightSim := s.Simplify(expr.Right)
	return NewBinaryOperation(leftSim, rightSim, expr.Operator)
}

// VisitUnaryOperation implements [Visitor].
func (s *Simplifier) VisitUnaryOperation(expr *UnaryOperation) interface{} {
	operandSim := s.Simplify(expr.Operand)
	newUnary := NewUnaryOperation(operandSim, expr.Operator)
	switch expr.Operator {
	case UN_NEG:
		return newUnary
	case UN_XOR:
		return newUnary
	case UN_NOT:
		switch e := operandSim.(type) {
		case *BoolConstant:
			return NewBoolConstant(!e.Value)
		case *LogicalOperation:
			if e.Operator == NOT {
				return e.Operands[0]
			}
		}
		return NewLogicalOperation([]SymbolicExpression{operandSim}, NOT)
	}
	return newUnary
}

func (s *Simplifier) VisitLogicalOperation(expr *LogicalOperation) interface{} {
	for i, operand := range expr.Operands {
		expr.Operands[i] = s.Simplify(operand)
	}
	return expr
}
