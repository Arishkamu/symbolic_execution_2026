package internal

import (
	"fmt"
	"go/constant"
	"go/token"
	"go/types"
	"symbolic-execution-course/internal/memory"
	"symbolic-execution-course/internal/symbolic"
	"symbolic-execution-course/internal/translator"

	"github.com/ebukreev/go-z3/z3"
	"golang.org/x/tools/go/ssa"
)

type Interpreter struct {
	CallStack     []CallStackFrame
	Analyser      *Analyser
	PathCondition symbolic.SymbolicExpression
	PathFormula   z3.Bool
	Heap          memory.Memory

	BlockVisits map[*ssa.BasicBlock]int
	LoopUnroll  StepTracker
	JumpDepth   StepTracker
	Status      InterpreterState
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
	Dropped
	Panic
)

func (s InterpreterStatus) String() string {
    switch s {
    case Running:
        return "Running"
    case Returned:
        return "Returned"
    case Dropped:
        return "Dropped"
    case Panic:
        return "Panic"
    default:
        return "Unknown"
    }
}

type InterpreterState struct {
    Status  InterpreterStatus
    Message string
}

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

	clone.CallStack = make([]CallStackFrame, len(interpreter.CallStack))

    for i, frame := range interpreter.CallStack {
        clone.CallStack[i] = frame
		frame.LocalMemory = make(map[string]symbolic.SymbolicExpression, len(interpreter.CallStack[i].LocalMemory))
		for k, v := range interpreter.CallStack[i].LocalMemory {
			frame.LocalMemory[k] = v
		}
	}

	clone.BlockVisits = make(map[*ssa.BasicBlock]int, len(interpreter.BlockVisits))

    for block, visits := range interpreter.BlockVisits {
        clone.BlockVisits[block] = visits
    }
	return clone
}

func (interpreter *Interpreter) moveBlock(newBlock *ssa.BasicBlock) {
	curFrame := interpreter.curFrame()
	interpreter.BlockVisits[curFrame.Block]++
	fmt.Printf("Interpreter: %p: curFrame.Block: %s = %s\n", interpreter, curFrame.Block, interpreter.BlockVisits[curFrame.Block])
	curFrame.PreviousBlock = curFrame.Block
	curFrame.Block = newBlock
	curFrame.InstructionIndex = 0
}

func (interpreter *Interpreter) finishInterpreter(status InterpreterStatus, msg string) []Interpreter {
	interpreter.curFrame().Block = nil
	interpreter.Status = InterpreterState{status, msg}
	return []Interpreter{*interpreter}
}

func (interpreter *Interpreter) feasible(cond symbolic.SymbolicExpression) bool {
	if c, ok := cond.(*symbolic.BoolConstant); ok {
		return c.Value
	}
	interpreter.PathCondition = symbolic.NewLogicalOperation(
		[]symbolic.SymbolicExpression{interpreter.PathCondition, cond},
		symbolic.AND,
	)
	value, err := simplifyAndTranslateExpression(interpreter.Analyser.Z3Translator, cond)
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
	case *ssa.Panic:
		return interpreter.interpretPanic(instr)
	case *ssa.MakeInterface:
		// Ignore for now. We think that this is only part of panic
		interpreter.incrInstructionIndex()
		return []Interpreter{*interpreter}
	default:
		// Unsupported instructions must not silently succeed.
		panic(fmt.Sprintf("interpretDynamically: unsupported instruction type: %s", instr))
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
	notCondition := symbolic.NewUnaryOperation(condition, symbolic.UN_NOT)

	var succsessors []Interpreter
	for i, cond := range []symbolic.SymbolicExpression{condition, notCondition} {
		interpr := interpreter.clone()
		if interpr.feasible(cond) {
			// TODO maybe check if maxBrancesCount reached
			interpr.moveBlock(instruction.Block().Succs[i])
			succsessors = append(succsessors, interpr)
		}
	}
	return succsessors
}

func (interpreter *Interpreter) interpretJump(instruction *ssa.Jump) []Interpreter {
	if interpreter.JumpDepth.finished() {
		return interpreter.finishInterpreter(Dropped, "JumpDepth maxStep reached.")
	}
	nextBlock := instruction.Block().Succs[0]
	// I probably can fix this if correctly clone (push and pop) interpreter and 
	// count loops and make zero when exit
	if (interpreter.BlockVisits[nextBlock] >= interpreter.LoopUnroll.MaxSteps) {
		return interpreter.finishInterpreter(Dropped, "LoopUnroll maxStep reached.")
	}

	interpreter.moveBlock(nextBlock)
	interpreter.JumpDepth.incr()
	return []Interpreter{*interpreter}
}

func (interpreter *Interpreter) interpretReturn(instruction *ssa.Return) []Interpreter {
	curFrame := interpreter.curFrame()

	if len(instruction.Results) > 0 {
		result := interpreter.resolveExpression(instruction.Results[0])
		curFrame.ReturnValue = result
	}

	return interpreter.finishInterpreter(Returned, "")
}

func (interpreter *Interpreter) interpretPanic(instruction *ssa.Panic) []Interpreter {
	// some conversion of make interface to string. Crutch
	panicMsg := instruction.X.String()
	if constSsa, ok := instruction.X.(*ssa.MakeInterface).X.(*ssa.Const); ok {
		constVal := constSsa.Value
		switch constVal.Kind() {
		case constant.String:
			panicMsg = constant.StringVal(constVal)
		default:
			panicMsg = constVal.String()
		}
	}
	return interpreter.finishInterpreter(Panic, panicMsg)
}

func (interpreter *Interpreter) interpretPhi(instruction *ssa.Phi) []Interpreter {
	curFrame := interpreter.curFrame()
    if curFrame.PreviousBlock == nil {
        panic("Error. Cannot resolve Phi: previous block is nil")
    }

	var value ssa.Value
    for i, pred := range instruction.Block().Preds {
		if pred == curFrame.PreviousBlock && i < len(instruction.Edges) {
			value = instruction.Edges[i]
			break
		}
	}
	if value == nil {
		panic("Error. Cannot find corresponding value in Phi")
	}
	
	result := interpreter.resolveExpression(value)
	interpreter.curFrame().LocalMemory[instruction.Name()] = result

	interpreter.incrInstructionIndex()
	return []Interpreter{*interpreter}
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
	// case *ssa.Phi:
	// 	fmt.Printf("BLOCK: %s\n", 7)
	// 	for _, instr := range interpreter.curFrame().Block.Instrs {
	// 		fmt.Printf("    %T: %s\n", instr, instr)
	// 	}
	// 	panic("resolveExpression: Phi not implemented yet")
		// return interpreter.resolvePhi(v)
	default:
		if expr, ok := interpreter.curFrame().LocalMemory[value.Name()]; ok {
			return expr
		} else {
			panic(fmt.Sprintf("unresolved value: %s, with type: %T", value.Name(), value))
		}
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
	case constant.Float:
		return symbolic.NewFloatConstant(v.Float64())
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
		return symbolic.NewUnaryOperation(value, symbolic.UN_XOR)
	case token.NOT:
		return symbolic.NewUnaryOperation(value, symbolic.UN_NOT)
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
	
	case token.AND:
		return symbolic.NewBinaryOperation(left, right, symbolic.BIT_AND)
	case token.OR:
		return symbolic.NewBinaryOperation(left, right, symbolic.BIT_OR)
	case token.XOR:
		return symbolic.NewBinaryOperation(left, right, symbolic.BIT_XOR)
	case token.SHL:
		return symbolic.NewBinaryOperation(left, right, symbolic.LSH)
	case token.SHR:
		return symbolic.NewBinaryOperation(left, right, symbolic.RSH)
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
		case t.Kind() == types.Float64 || t.Kind() == types.UntypedFloat:
			return symbolic.Float64Type
		}
	}
	panic(fmt.Errorf("unsupported basic type: %s", t.String()))
}

func (interpreter *Interpreter) String() string {
	result := fmt.Sprintf("Interpreter:\n")
	result += fmt.Sprintf("PathCondition: %s\n", interpreter.PathCondition.String())
	result += fmt.Sprintf("Status: %s: %s\n", interpreter.Status.Status, interpreter.Status.Message)

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

func simplifyAndTranslateExpression(zt *translator.Z3Translator, expr symbolic.SymbolicExpression) (interface{}, error) {
	simpifiedExpr := (&symbolic.Simplifier{}).Simplify(expr)
	return zt.TranslateExpression(simpifiedExpr)
}