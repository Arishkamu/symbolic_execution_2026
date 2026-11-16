// Package symbolic определяет базовые типы символьных выражений
package symbolic

// ExpressionType представляет тип символьного выражения
type ExpressionType int

const (
	IntType ExpressionType = iota
	Float64Type
	BoolType
	ArrayType
	// Добавьте другие типы по необходимости
)

// String возвращает строковое представление типа
func (et ExpressionType) String() string {
	switch et {
	case IntType:
		return "int"
	case Float64Type:
		return "float64"
	case BoolType:
		return "bool"
	case ArrayType:
		return "array"
	default:
		return "unknown"
	}
}
