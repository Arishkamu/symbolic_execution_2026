package internal

import (
	"fmt"
	"log"
	"symbolic-execution-course/internal/translator"

	"golang.org/x/tools/go/ssa"
	intssa "symbolic-execution-course/internal/ssa"
)

type Analyser struct {
	Package      *ssa.Package
	StatesQueue  PriorityQueue
	PathSelector PathSelector
	Results      []Interpreter
	Z3Translator *translator.Z3Translator
	Stepper    	 *StepTracker
}

type AnalyseOptions struct {
    MaxSteps     int
    MaxJumpDepth int
}

func DefaultOptions() AnalyseOptions {
    return AnalyseOptions{
        MaxSteps:     10000,
        MaxJumpDepth: 50,
    }
}

func Analyse(source string, functionName string) []Interpreter {
	return AnalyzeWithOptions(source, functionName, DefaultOptions())
}

func AnalyzeWithOptions(source, functionName string, analyseOptions AnalyseOptions) []Interpreter {
	panic("unimplemented")
	ssaBuilder := intssa.NewBuilder()
	ssaSource, err := ssaBuilder.ParseAndBuildSSAPkg(source, functionName)
	if err != nil {
		log.Fatalf("Ошибка построения SSA: %v", err)
	}

	analyser := &Analyser{
		Package:      ssaSource.Pkg(),
		StatesQueue:  make(PriorityQueue, 0),
		PathSelector: selector,
		Results:      make([]Interpreter, 0),
		Z3Translator: translator.NewZ3Translator(),
		Stepper:      newStepTracker(maxSteps),
		stepsCounter: 0,
	}
}



func NewInterpreter(function *ssa.Function) any {
	panic("unimplemented")
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