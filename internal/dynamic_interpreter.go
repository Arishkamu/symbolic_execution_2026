package internal

import (
	"fmt"
	"go/constant"
	"go/token"
	"go/types"
	"symbolic-execution-course/internal/memory"
	"symbolic-execution-course/internal/symbolic"

	"github.com/ebukreev/go-z3/z3"
	"golang.org/x/tools/go/ssa"
)

type Interpreter struct {
	CallStack     []CallStackFrame
	Analyser      *Analyser
	PathCondition symbolic.SymbolicExpression
	PathFormula   z3.Bool
	Heap          memory.Memory

	LoopUnroll *StepTracker
	JumpDepth  *StepTracker
	Status     InterpreterStatus
}

type CallStackFrame struct {
	Function         *ssa.Function
	LocalMemory      map[string]symbolic.SymbolicExpression
	ReturnValue      symbolic.SymbolicExpression
	Block            *ssa.BasicBlock
	PreviousBlock    *ssa.BasicBlock
	InstructionIndex int
}

type InterpreterStatus int

const (
	Running InterpreterStatus = iota
	Returned
	Terminated
	Panic
)

// Interpreter utils
func (interpreter *Interpreter) curFrame() *CallStackFrame {
	if len(interpreter.CallStack) == 0 {
		panic("Call stack is empty. No current frame available.")
	}
	return &interpreter.CallStack[len(interpreter.CallStack)-1]
}

func (interpreter *Interpreter) curInstruction() ssa.Instruction {
	curFrame := interpreter.curFrame()
	return curFrame.Block.Instrs[curFrame.InstructionIndex]
}

func (interpreter *Interpreter) incrInstructionIndex() {
	interpreter.curFrame().InstructionIndex++
}

func (interpreter *Interpreter) clone() Interpreter {
	clone := *interpreter

	clone.CallStack = append([]CallStackFrame(nil), interpreter.CallStack...)
	return clone
}

func (interpreter *Interpreter) moveBlock(newBlock *ssa.BasicBlock) {
	curFrame := interpreter.curFrame()
	curFrame.PreviousBlock = curFrame.Block
	curFrame.Block = newBlock
	curFrame.InstructionIndex = 0
}

func (interpreter *Interpreter) feasible(cond symbolic.SymbolicExpression) bool {
	if c, ok := cond.(*symbolic.BoolConstant); ok {
		return c.Value
	}
	interpreter.PathCondition = symbolic.NewLogicalOperation(
		[]symbolic.SymbolicExpression{interpreter.PathCondition, cond},
		symbolic.AND,
	)
	value, err := interpreter.Analyser.Z3Translator.TranslateExpression(cond)
	if err != nil {
		panic(fmt.Sprintf("analysis: translate path: %v", err))
	}
	interpreter.PathFormula = interpreter.PathFormula.And(value.(z3.Bool))
	return interpreter.Analyser.solverCheck(interpreter.PathFormula)
}

// interpretDynamically block
func (interpreter *Interpreter) interpretDynamically(element ssa.Instruction) []Interpreter {
	switch instr := element.(type) {
	case *ssa.UnOp:
		return interpreter.interpretUnOp(instr)
	case *ssa.BinOp:
		return interpreter.interpretBinOp(instr)
	case *ssa.If:
		return interpreter.interpretIf(instr)
	case *ssa.Jump:
		return interpreter.interpretJump(instr)
	case *ssa.Return:
		return interpreter.interpretReturn(instr)
	case *ssa.Phi:
		return interpreter.interpretPhi(instr)
	default:
		// Unsupported instructions must not silently succeed.
		panic("interpretDynamically: unsupported instruction type")
	}
}

func (interpreter *Interpreter) interpretUnOp(instruction *ssa.UnOp) []Interpreter {
	result := interpreter.resolveExpression(instruction)
	interpreter.curFrame().LocalMemory[instruction.Name()] = result

	interpreter.incrInstructionIndex()
	return []Interpreter{*interpreter}
}

func (interpreter *Interpreter) interpretBinOp(instruction *ssa.BinOp) []Interpreter {
	result := interpreter.resolveExpression(instruction)
	interpreter.curFrame().LocalMemory[instruction.Name()] = result

	interpreter.incrInstructionIndex()
	return []Interpreter{*interpreter}
}

func (interpreter *Interpreter) interpretIf(instruction *ssa.If) []Interpreter {
	condition := interpreter.resolveExpression(instruction.Cond)
	notCondition := symbolic.UnNot(condition)

	var succsessors []Interpreter
	for i, cond := range []symbolic.SymbolicExpression{condition, notCondition} {
		interpr := interpreter.clone()
		if interpr.feasible(cond) {
			interpr.moveBlock(instruction.Block().Succs[i])
			succsessors = append(succsessors, interpr)
		}
	}
	return succsessors
}

func (interpreter *Interpreter) interpretJump(instruction *ssa.Jump) []Interpreter {
	panic("unimplemented")
}

func (interpreter *Interpreter) interpretReturn(instruction *ssa.Return) []Interpreter {
	curFrame := interpreter.curFrame()

	if len(instruction.Results) > 0 {
		result := interpreter.resolveExpression(instruction.Results[0])
		curFrame.ReturnValue = result
	}

	curFrame.Block = nil
	interpreter.Status = Returned
	return []Interpreter{*interpreter}
}

func (interpreter *Interpreter) interpretPhi(instruction *ssa.Phi) []Interpreter {
	panic("unimplemented")
}

// Resolve expressions block
func (interpreter *Interpreter) resolveExpression(value ssa.Value) symbolic.SymbolicExpression {

	switch v := value.(type) {
	case *ssa.Const:
		return interpreter.resolveConst(v)
	case *ssa.UnOp:
		return interpreter.resolveUnOp(v)
	case *ssa.BinOp:
		return interpreter.resolveBinOp(v)
	// case *ssa.Parameter:
	// 	return interpreter.resolveParameter(v)
	case *ssa.Phi:
		panic("resolveExpression: Phi not implemented yet")
		// return interpreter.resolvePhi(v)
	default:
		if expr, ok := interpreter.curFrame().LocalMemory[value.Name()]; ok {
			return expr
		}
		panic(fmt.Sprintf("unresolved value: %s, with type: %T", value.Name(), value))
	}
}

func (interpreter *Interpreter) resolveConst(v *ssa.Const) symbolic.SymbolicExpression {
	if v.IsNil() {
		panic("resolveConst: nil constant not implemented yet")
	}

	constVal := v.Value
	switch constVal.Kind() {
	case constant.Bool:
		return symbolic.NewBoolConstant(constant.BoolVal(constVal))
	case constant.Int:
		return symbolic.NewIntConstant(v.Int64())
	default:
		panic(fmt.Errorf("unknown type for const %s", v.String()))
	}
}

func (interpreter *Interpreter) resolveUnOp(v *ssa.UnOp) symbolic.SymbolicExpression {
	value := interpreter.resolveExpression(v.X)

	switch v.Op {
	case token.SUB:
		return symbolic.NewUnaryOperation(value, symbolic.UN_NEG)
	case token.XOR:
		return symbolic.UnXor(value)
	case token.NOT:
		return symbolic.UnNot(value)
	default:
		panic(fmt.Errorf("unknown or not implemented type for unop %s", v.Op))
	}
}

func (interpreter *Interpreter) resolveBinOp(v *ssa.BinOp) symbolic.SymbolicExpression {
	left := interpreter.resolveExpression(v.X)
	right := interpreter.resolveExpression(v.Y)

	switch v.Op {
	case token.ADD:
		return symbolic.NewBinaryOperation(left, right, symbolic.ADD)
	case token.SUB:
		return symbolic.NewBinaryOperation(left, right, symbolic.SUB)
	case token.MUL:
		return symbolic.NewBinaryOperation(left, right, symbolic.MUL)
	case token.QUO:
		return symbolic.NewBinaryOperation(left, right, symbolic.DIV)
	case token.REM:
		return symbolic.NewBinaryOperation(left, right, symbolic.MOD)

	case token.EQL:
		return symbolic.NewBinaryOperation(left, right, symbolic.EQ)
	case token.NEQ:
		return symbolic.NewBinaryOperation(left, right, symbolic.NE)
	case token.LSS:
		return symbolic.NewBinaryOperation(left, right, symbolic.LT)
	case token.GTR:
		return symbolic.NewBinaryOperation(left, right, symbolic.GT)
	case token.LEQ:
		return symbolic.NewBinaryOperation(left, right, symbolic.LE)
	case token.GEQ:
		return symbolic.NewBinaryOperation(left, right, symbolic.GE)
	default:
		// Comparisons and logical operators need to map to the
		// corresponding expression types in your symbolic system.
		panic(fmt.Errorf("unsupported binary operator %s", v.Op))
	}
}

// func (interpreter *Interpreter) resolveParameter(v *ssa.Parameter) symbolic.SymbolicExpression {
// 	panic("Ssa.Parameter resolution not implemented yet. I want to see when it happends.")
// }

func convertType(t types.Type) symbolic.ExpressionType {
	if t == nil {
		panic("cannot convert expression type")
	}
	if t, ok := t.Underlying().(*types.Basic); ok {
		switch {
		case t.Kind() == types.Bool || t.Kind() == types.UntypedBool:
			return symbolic.BoolType
		case t.Kind() == types.Int || t.Kind() == types.UntypedInt:
			return symbolic.IntType
		}
	}
	panic(fmt.Errorf("unsupported basic type: %s", t.String()))
}

func (interpreter *Interpreter) String() string {
	result := fmt.Sprintf("Interpreter:\n")
	result += fmt.Sprintf("PathCondition: %s\n", interpreter.PathCondition.String())
	result += fmt.Sprintf("Status: %d\n", interpreter.Status)

	curFrame := interpreter.curFrame()
	if len(interpreter.CallStack) > 0 {
		result += fmt.Sprintf("Current Frame:\n")
		result += fmt.Sprintf("Function: %s\n", curFrame.Function.Name())

		if len(curFrame.LocalMemory) > 0 {
			result += fmt.Sprintf("LocalMemory:\n")
			for k, v := range curFrame.LocalMemory {
				result += fmt.Sprintf("%s: %s\n", k, v.String())
			}
		}

		if curFrame.ReturnValue != nil {
			result += fmt.Sprintf("ReturnValue: %s\n", curFrame.ReturnValue.String())
		}
	}

	if interpreter.curFrame().Block != nil {
		result += fmt.Sprintf("CurrentBlock: %s\n", interpreter.curFrame().Block.String())
	}
	result += fmt.Sprintf("InstrIndex: %d\n", curFrame.InstructionIndex)
	return result
}
