package internal

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"go/types"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"symbolic-execution-course/internal/symbolic"

	"github.com/ebukreev/go-z3/z3"
)

func GenerateTestFile(sourceFile string) string {
    filePath, err := filepath.Abs(sourceFile)
    if err != nil {
        panic(fmt.Sprintf("Generate tests: Error while getting filapath: %v", err))
	}
	ext := filepath.Ext(filePath)
	if ext != ".go" {
		panic(fmt.Sprintf("Generate tests: Error file needs to be '.go': %v", err))
	}
	source, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("Generate tests: Error while reading file: %v", err))
	}


	file, err := parser.ParseFile(token.NewFileSet(), filePath, source, 0)
	if err != nil {
		panic(fmt.Sprintf("Generate tests: Error while parsing file: %v", err))
	}

	var body strings.Builder
	imports := make(Imports)
	imports["testing"] = true
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl);
        if !ok || fn.Body == nil {
            continue
        }
		testCases := analyzeAndTestFunc(source, fn.Name.Name, &imports)
		generateTestBody(&body, fn.Name.Name, testCases)
    }


	testsFile := strings.TrimSuffix(filePath, ext) + "_test.go"
	createTestFile(body.String(), testsFile, file.Name.Name, imports)

    return testsFile
}

func createTestFile(body string, testPath string, fileName string, imports Imports) {
	var fileBody strings.Builder

	fileBody.WriteString(fmt.Sprintf("package %s\n\n", fileName))

	fileBody.WriteString("import (\n")
	for k, v := range imports {
		if v {
			fileBody.WriteString(fmt.Sprintf("\t%#v\n", k))
		}
	}
	
	fileBody.WriteString(")\n")

	fileBody.WriteString(fmt.Sprintf(`
func panicCheck_%s(t *testing.T, panicExpect string, fn func()) {
	defer func() {
        r := recover()
        if r == nil {
			t.Errorf("Expected panic, but no panic occurred")
		}
		if r != panicExpect {
			t.Errorf("Expected panic with message='%%s', but no panic occurred", panicExpect)
		}
    }()

    fn()
}
`, 
	fileName))

	fileBody.WriteString(body)
	// fmt.Printf("======\n\n%s\n=======\n", fileBody.String())
	formatted, err := format.Source([]byte(fileBody.String()))
	if err != nil {
		panic(fmt.Sprintf("Generate tests: Error while formating generated file body: %v", err))
	}

	if err := os.WriteFile(testPath, formatted, 0644); err != nil {
		panic(fmt.Sprintf("Generate tests: Error while writing generated file: %v", err))
	}
}


func analyzeAndTestFunc(source []byte, funcName string, imports *Imports) []TestCase {
	states := Analyse(string(source), funcName)
	testCases := make([]TestCase, 0)
	dropped := 0
	for _, state := range states {
		if state.Status.Status == Running {
			panic("Unexpected behavour. Analyse returned running interpreter")
		}
		if state.Status.Status == Dropped {
			// fmt.Printf("In func='%s' Dropped with message='%s', in pathCondition=%s\n", funcName, state.Status.Message, state.PathCondition)
			dropped++
			continue
		}

		solver := state.Analyser.Solver
		solver.Push()
		solver.Assert(state.PathFormula)
		sat, err := solver.Check()
		if err != nil {
			panic(fmt.Sprintf("analysis: solver could not decide path: %v", err))
		}
		if !sat {
			continue
		}
		model := state.Analyser.Solver.Model()
		solver.Pop()
		curFrame := state.curFrame()
		
		testCase := TestCase{
			PathCond: state.PathCondition,
			Status: state.Status,
		}
		for _, param := range curFrame.Function.Params {
			argSymbolic := curFrame.LocalMemory[param.Name()]
			argZ3, err := simplifyAndTranslateExpression(state.Analyser.Z3Translator, argSymbolic)
			if err != nil {
				panic(fmt.Sprintf("Generate tests: Unexpected behavior. Can't translate argument expression %s: %v", argSymbolic, err))
			}
			arg := model.Eval(argZ3.(z3.Value), true)
			testCase.Args = append(testCase.Args, getGoLiteral(arg, param.Type(), argSymbolic.Type(), imports))
		}
		if state.Status.Status == Returned {
			resTypes := curFrame.Function.Signature.Results()
			for i, retVal := range curFrame.ReturnValues {
				retZ3, err := simplifyAndTranslateExpression(state.Analyser.Z3Translator, retVal)
				if err != nil {
					panic(fmt.Sprintf("Generate tests: Unexpected behavior. Can't translate argument expression %s: %v", retVal, err))
				}
				ret := model.Eval(retZ3.(z3.Value), true)
				testCase.Predicted = append(testCase.Predicted, getGoLiteral(ret, resTypes.At(i).Type(), retVal.Type(), imports))
			}
		}
		testCases = append(testCases, testCase)
	}
	// fmt.Printf("In func='%s': Dropped=%d\n", funcName, dropped)
	return testCases
}

type Imports map[string]bool
type TestCase struct {
    Args       []string
    Predicted  []string
	PathCond   symbolic.SymbolicExpression
	Status     InterpreterState
}


func generateTestBody(body *strings.Builder, funcName string, cases []TestCase) {
	body.WriteString(fmt.Sprintf("func TestGenerated_%s(t *testing.T) {\n", funcName))

    for iCS, cs := range cases {
		argsString := strings.Join(cs.Args, ", ")
		if len(cs.Args) == 0 {
			argsString = ""
		}

		switch cs.Status.Status {
		case Returned:
			resStr := make([]string, len(cs.Predicted))
			for i := 0; i < len(cs.Predicted); i++ {
				resStr[i] = fmt.Sprintf("res%dTC%d", i, iCS)
			}
			resiStr := strings.Join(resStr, ", ")
            // body.WriteString(fmt.Sprintf("\n\t// TestCase_PathCondition: %s\n", cs.PathCond))
			body.WriteString(fmt.Sprintf("\t%s := %s(%s)\n", resiStr, funcName, argsString))
			for i, predict := range cs.Predicted {
				body.WriteString(fmt.Sprintf(`
	if %s != %s {
		t.Errorf("%s = %%v, but predicted %s", %s)
	}
`,
				resStr[i], predict, resStr[i], predict, resStr[i]))
			}
		case Panic:
			body.WriteString(fmt.Sprintf("\tpanicCheck(t, %#v, func() { %s(%s) })\n", 
				cs.Status.Message, funcName, argsString))
		default:
			panic("Unexpected status while generating test body")
		}
    }
	body.WriteString("}\n\n")
}



func getGoLiteral(value z3.Value, ty types.Type, smbType symbolic.ExpressionType, imports *Imports) string {
    switch smbType {
    case symbolic.BoolType:
        boolVal, isLiteral := value.(z3.Bool).AsBool()
        if !isLiteral {
            panic("Generate tests: Expected concreet bool value")
        }
		return fmt.Sprintf("%t", boolVal)
    case symbolic.IntType:
		numVal, isLiteral, ok := value.(z3.BV).AsInt64()
		if !isLiteral || !ok{
			panic("Generate tests: Expected concreet int64 value")
		}
		return fmt.Sprintf("%d", numVal)
	case symbolic.Float64Type:
		floatVal, ok := value.(z3.Float).AsBigFloat()
		if !ok {
			panic("Generate tests: Expected concreet float value")
		}
		(*imports)["math"] = true
		if floatVal == nil {
			return "math.NaN()"
		}

		flt, _:= floatVal.Float64()
		switch {
		case math.IsNaN(flt):
			return "math.NaN()"
		case math.IsInf(flt, 1):
			return "math.Inf(1)"
		case math.IsInf(flt, -1):
			return "math.Inf(-1)"
		default:
			(*imports)["math"] = false
			return strconv.FormatFloat(flt, 'g', -1, 64)
		}
    default:
        panic("Generate tests: Unsupported Go type")
    }
}



