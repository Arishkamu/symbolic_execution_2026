package internal

import (
	"container/heap"
	"fmt"
	"log"
	"symbolic-execution-course/internal/symbolic"
	"symbolic-execution-course/internal/translator"

	intssa "symbolic-execution-course/internal/ssa"

	"github.com/ebukreev/go-z3/z3"
	"golang.org/x/tools/go/ssa"
)

type Analyser struct {
	Package      *ssa.Package
	StatesQueue  PriorityQueue
	PathSelector PathSelector
	Results      []Interpreter
	Z3Translator *translator.Z3Translator
	Stepper    	 *StepTracker
	Solver       *z3.Solver
}

type AnalyseOptions struct {
	PathSelector PathSelector
    MaxSteps      int
    MaxLoopUnroll int
	MaxJumpDepth  int
}

func DefaultOptions() AnalyseOptions {
    return AnalyseOptions{
        PathSelector:  &BfsPathSelector{},
        MaxSteps:      10000,
        MaxLoopUnroll: 5,
        MaxJumpDepth:  50,
    }
}

func Analyse(source string, functionName string) []Interpreter {
	return AnalyzeWithOptions(source, functionName, DefaultOptions())
}

func AnalyzeWithOptions(source, functionName string, options AnalyseOptions) []Interpreter {
	ssaBuilder := intssa.NewBuilder()
	ssaSource, err := ssaBuilder.ParseAndBuildSSAPkg(source, functionName)
	if err != nil {
		log.Fatalf("Ошибка построения SSA: %v", err)
	}

	translator := translator.NewZ3Translator()
	ctx := translator.GetContext().(*z3.Context)
	analyser := &Analyser{
		Package:      ssaSource.Pkg,
		StatesQueue:  make(PriorityQueue, 0),
		PathSelector: options.PathSelector,
		Results:      make([]Interpreter, 0),
		Z3Translator: translator,
		Stepper:      newStepTracker(options.MaxSteps),
		Solver:       z3.NewSolver(ctx),
	}

	initialState := Interpreter{
		CallStack: []CallStackFrame{{
			Function:    ssaSource.Fn,
			Block:       ssaSource.Fn.Blocks[0],
			LocalMemory: make(map[string]symbolic.SymbolicExpression),
		}},
		Analyser:      analyser,
		LoopUnroll:    newStepTracker(options.MaxLoopUnroll),
		JumpDepth:     newStepTracker(options.MaxJumpDepth),
		PathCondition: symbolic.NewBoolConstant(true),
		PathFormula:   ctx.FromBool(true),
		Status:        Running,
	}
	
	for _, param := range ssaSource.Fn.Params {
		initialState.curFrame().LocalMemory[param.Name()] = symbolic.NewSymbolicVariable(param.Name(), convertType(param.Type()))
	}

    analyser.push(initialState)
	analyser.run(initialState)
	return analyser.Results
}

func (analyser *Analyser) run(initialState Interpreter) {
	for analyser.StatesQueue.Len() > 0 && !analyser.Stepper.finished() {
		interpreter := heap.Pop(&analyser.StatesQueue).(*Item).value
		analyser.Stepper.incr()
		for _, newInterpreter := range interpreter.interpretDynamically(interpreter.curInstruction()) {
			if newInterpreter.Status == Running {
				analyser.push(newInterpreter)
			} else {
				analyser.Results = append(analyser.Results, newInterpreter)
			}
		}
	}
}

func (analyser *Analyser) push(interpreter Interpreter) {
	heap.Push(&analyser.StatesQueue, &Item{
		value:    interpreter,
		priority: analyser.PathSelector.CalculatePriority(interpreter),
	})
}

func (analyser *Analyser) solverCheck(condition z3.Bool) bool {
	analyser.Solver.Push()
	defer analyser.Solver.Pop()
	
	analyser.Solver.Assert(condition)
	sat, err := analyser.Solver.Check()
	if err != nil {
		panic(fmt.Sprintf("analysis: solver could not decide path: %v", err))
	}
	fmt.Printf("Condi: %v\n",condition)
	fmt.Printf("Model: %v\n",analyser.Solver.Model())
	return sat
}
	


type StepTracker struct {
    MaxSteps  int
    StepCount int
}

func (st *StepTracker) incr() bool {
	st.StepCount++
	return st.StepCount <= st.MaxSteps
}

func (st *StepTracker) finished() bool {
	return st.StepCount >= st.MaxSteps
}

func newStepTracker(maxSteps int) *StepTracker {
	return &StepTracker{
		MaxSteps:  maxSteps,
		StepCount: 0,
	}
}