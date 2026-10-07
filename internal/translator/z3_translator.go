// Package translator содержит реализацию транслятора в Z3
package translator

import (
	"fmt"
	"symbolic-execution-course/internal/symbolic"

	"github.com/ebukreev/go-z3/z3"
)

// Z3Translator транслирует символьные выражения в Z3 формулы
type Z3Translator struct {
	ctx    *z3.Context
	config *z3.Config
	vars   map[string]z3.Value // Кэш переменных
}

// NewZ3Translator создаёт новый экземпляр Z3 транслятора
func NewZ3Translator() *Z3Translator {
	config := &z3.Config{}
	ctx := z3.NewContext(config)

	return &Z3Translator{
		ctx:    ctx,
		config: config,
		vars:   make(map[string]z3.Value),
	}
}

// GetContext возвращает Z3 контекст
func (zt *Z3Translator) GetContext() interface{} {
	return zt.ctx
}

// Reset сбрасывает состояние транслятора
func (zt *Z3Translator) Reset() {
	zt.vars = make(map[string]z3.Value)
}

// Close освобождает ресурсы
func (zt *Z3Translator) Close() {
	// Z3 контекст закрывается автоматически
}

// TranslateExpression транслирует символьное выражение в Z3
func (zt *Z3Translator) TranslateExpression(expr symbolic.SymbolicExpression) (interface{}, error) {
	return expr.Accept(zt), nil
}

// TODO: Реализуйте следующие методы в рамках домашнего задания

// VisitVariable транслирует символьную переменную в Z3
func (zt *Z3Translator) VisitVariable(expr *symbolic.SymbolicVariable) interface{} {
	// TODO: Реализовать
	// Проверить, есть ли переменная в кэше
	// Если нет - создать новую Z3 переменную соответствующего типа
	// Добавить в кэш и вернуть

	if value, ok := zt.vars[expr.Name]; ok {
		return value
	}
	value := zt.createZ3Variable(expr.Name, expr.Type())
	zt.vars[expr.Name] = value
	return value
}

// VisitIntConstant транслирует целочисленную константу в Z3
func (zt *Z3Translator) VisitIntConstant(expr *symbolic.IntConstant) interface{} {
	return zt.ctx.FromInt(expr.Value, zt.ctx.IntSort()).(z3.Int)
}

// VisitBoolConstant транслирует булеву константу в Z3
func (zt *Z3Translator) VisitBoolConstant(expr *symbolic.BoolConstant) interface{} {
	return zt.ctx.FromBool(expr.Value)
}

// VisitBinaryOperation транслирует бинарную операцию в Z3
func (zt *Z3Translator) VisitBinaryOperation(expr *symbolic.BinaryOperation) interface{} {
	// TODO: Реализовать
	// 1. Транслировать левый и правый операнды
	// 2. В зависимости от оператора создать соответствующую Z3 операцию

	// Подсказки по операциям в Z3:
	// - Арифметические: left.Add(right), left.Sub(right), left.Mul(right), left.Div(right)
	// - Сравнения: left.Eq(right), left.LT(right), left.LE(right), etc.
	// - Приводите типы: left.(z3.Int), right.(z3.Int) для int операций

	left  :=  expr.Left.Accept(zt)
	right := expr.Right.Accept(zt)

	switch expr.Left.Type() {
	case symbolic.BoolType:
		leftBool, err  := zt.castToZ3BoolType(left)
		if err != nil {
			panic(fmt.Errorf("Error in binary operation: left operand: %w", err))
		}
		rightBool, err := zt.castToZ3BoolType(right)
		if err != nil {
			panic(fmt.Errorf("Error in binary operation: right operand: %w", err))
		}
		switch expr.Operator {
		case symbolic.EQ:
			return leftBool.Eq(rightBool)
		case symbolic.NE:
			return leftBool.NE(rightBool)
		default:
			panic(fmt.Errorf("Operation '%s' not supported for bool expressions", expr.Operator))
		}

	case symbolic.IntType:
		leftInt, err  := zt.castToZ3IntType(left)
		if err != nil {
			panic(fmt.Errorf("Error in binary operation: left operand: %w", err))
		}
		rightInt, err := zt.castToZ3IntType(right)
		if err != nil {
			panic(fmt.Errorf("Error in binary operation: right operand: %w", err))
		}

		switch expr.Operator {
		case symbolic.ADD:
			return leftInt.Add(rightInt)
		case symbolic.SUB:
			return leftInt.Sub(rightInt)
		case symbolic.MUL:
			return leftInt.Mul(rightInt)
		case symbolic.DIV:
			return leftInt.Div(rightInt)
		case symbolic.MOD:
			return leftInt.Mod(rightInt)
		case symbolic.EQ:
			return leftInt.Eq(rightInt)
		case symbolic.NE:
			return leftInt.NE(rightInt)
		case symbolic.LT:
			return leftInt.LT(rightInt)
		case symbolic.LE:
			return leftInt.LE(rightInt)
		case symbolic.GT: 
			return leftInt.GT(rightInt)
		case symbolic.GE:
			return leftInt.GE(rightInt)
		default:
			panic(fmt.Errorf("Operation '%s' not supported for int expressions", expr.Operator))
		}
	default:
		panic("не реализовано")
	}
}

// VisitLogicalOperation транслирует логическую операцию в Z3
func (zt *Z3Translator) VisitLogicalOperation(expr *symbolic.LogicalOperation) interface{} {
	// TODO: Реализовать
	// 1. Транслировать все операнды
	// 2. Применить соответствующую логическую операцию

	// Подсказки:
	// - AND: zt.ctx.And(operands...)
	// - OR: zt.ctx.Or(operands...)
	// - NOT: operand.Not() (для единственного операнда)
	// - IMPLIES: antecedent.Implies(consequent)

	boolOperands := make([]z3.Bool, len(expr.Operands))
	for i, operand := range expr.Operands {
		vBool, err  := zt.castToZ3BoolType(operand.Accept(zt))
		if err != nil {
			panic(fmt.Errorf("Error in logic operation: operand is not bool: %w", err))
		}
		boolOperands[i] = vBool
	}

	switch expr.Operator {
	case symbolic.AND:
		return boolOperands[0].And(boolOperands[1:]...)
	case symbolic.OR:
		return boolOperands[0].Or(boolOperands[1:]...)
	case symbolic.NOT:
		return boolOperands[0].Not()
	case symbolic.IMPLIES:
		return boolOperands[0].Implies(boolOperands[1])
	default:
		panic(fmt.Errorf("Logic operation '%s' not supported", expr.Operator))
	}
}

// Вспомогательные методы

// createZ3Variable создаёт Z3 переменную соответствующего типа
func (zt *Z3Translator) createZ3Variable(name string, exprType symbolic.ExpressionType) z3.Value {
	switch exprType {
	case symbolic.IntType:
		return zt.ctx.IntConst(name)
	case symbolic.BoolType:
		return zt.ctx.BoolConst(name)
	default:
		return nil
	}
}

// castToZ3Type приводит значение к нужному Z3 типу
func (zt *Z3Translator) castToZ3Type(value interface{}, targetType symbolic.ExpressionType) (z3.Value, error) {
	switch targetType {
	case symbolic.IntType:
		return zt.castToZ3IntType(value)
	case symbolic.BoolType:
		return zt.castToZ3BoolType(value)
	default:
		return nil, fmt.Errorf("Error: unsupported symbolic type: %s", targetType)
	}
}

func (zt *Z3Translator) castToZ3IntType(value interface{}) (z3.Int, error) {
	z3Value, ok := value.(z3.Int)
	if !ok {
		return z3Value, fmt.Errorf("Error: expected z3.Int, got %T", value)
	}
	return z3Value, nil
}

func (zt *Z3Translator) castToZ3BoolType(value interface{}) (z3.Bool, error) {
	z3Value, ok := value.(z3.Bool)
	if !ok {
		return z3Value, fmt.Errorf("Error: expected z3.Bool, got %T", value)
	}
	return z3Value, nil
}
