package symbolic

import (
	"strings"
)

// Visitor интерфейс для обхода символьных выражений (Visitor Pattern)
type Visitor interface {
	VisitVariable(expr *SymbolicVariable) interface{}
	VisitIntConstant(expr *IntConstant) interface{}
	VisitBoolConstant(expr *BoolConstant) interface{}
	VisitBinaryOperation(expr *BinaryOperation) interface{}
	VisitLogicalOperation(expr *LogicalOperation) interface{}
	// TODO: Добавьте методы для других типов выражений по мере необходимости
}

// Пример реализации - вывод выражения в строку
type StringPrinter struct {
	result strings.Builder
}

func (sp *StringPrinter) VisitVariable(expr *SymbolicVariable) interface{} {
	sp.result.WriteString(expr.String())
	return nil
}

func (sp *StringPrinter) VisitIntConstant(expr *IntConstant) interface{} {
	sp.result.WriteString(expr.String())
	return nil
}

func (sp *StringPrinter) VisitBoolConstant(expr *BoolConstant) interface{} {
	sp.result.WriteString(expr.String())
	return nil
}

func (sp *StringPrinter) VisitBinaryOperation(expr *BinaryOperation) interface{} {
	sp.result.WriteString("(")
	expr.Left.Accept(sp)
	sp.result.WriteString(" ")
	sp.result.WriteString(expr.Operator.String())
	sp.result.WriteString(" ")
	expr.Right.Accept(sp)
	sp.result.WriteString(")")

	return nil
}

func (sp *StringPrinter) VisitLogicalOperation(expr *LogicalOperation) interface{} {
	switch expr.Operator {
	case NOT:
		sp.result.WriteString("!")
		expr.Operands[0].Accept(sp)
		return nil
	case AND, OR:
		sp.result.WriteString("(")
		for i, operand := range expr.Operands {
			if i > 0 {
				sp.result.WriteString(" ")
				sp.result.WriteString(expr.Operator.String())
				sp.result.WriteString(" ")
			}
			operand.Accept(sp)
		}
		sp.result.WriteString(")")
		return nil
	case IMPLIES:
		sp.result.WriteString("(")
		sp.result.WriteString(expr.Operands[0].String())
		sp.result.WriteString(" =>")
		expr.Operands[1].Accept(sp)
		sp.result.WriteString(")")
		return nil
	default:
		return nil
	}
}
