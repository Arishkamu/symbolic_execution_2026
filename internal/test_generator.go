package internal

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"strings"

	"github.com/ebukreev/go-z3/z3"
	"symbolic-execution-course/internal/symbolic"
)

func GenerateTestFile(sourceFile string) string {
    filePath, err := filepath.Abs(filename)
    if err != nil {
        panic(fmt.Sprintf("Generate tests: Error while getting filapath: %w", err))
	}
	ext := filepath.Ext(filePath)
	if ext != ".go" {
		panic(fmt.Sprintf("Generate tests: Error file needs to be '.go': %w", err))
	}
	source, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("Generate tests: Error while reading file: %w", err))
	}


	file, err := parser.ParseFile(token.NewFileSet(), filePath, source, 0)
	if err != nil {
		panic(fmt.Sprintf("Generate tests: Error while parsing file: %w", err))
	}

	var funcNames []string
	var body strings.Builder
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl);
        if !ok || fn.Body == nil {
            continue
        }
		testCases := analyzeAndTestFunc(source, fn.Name.Name)
		generateTestBody(&body, fn.Name.Name, testCases)
    }


	testsFile := strings.TrimSuffix(filePath, ext) + "_test.go"
	createTestFile(body.String(), testsFile, file.Name.Name)

    return testsFile
}


func analyzeAndTestFunc(source []byte, funcName string) *[]TestCase {
	testCases := make([]TestCase)
	states := Analyse(source, name)
	var dropped strings.Builder
	for _, state := range states {
		if state.Status.Status == Running {
			panic("Unexpected behavour. Analyse returned running interpreter")
		}
		if state.Status.Status == Dropped {
			dropped.WriteString(fmt.Sprintf("Dropped in func='%s' with message='%s', in pathCondition=%s\n", funcName, state.Status.Message, state.PathCondition))
		}
		
		if sat := state.Analyser.solverCheck(state.PathFormula); !sat {
			continue
		}
		model := state.Analyser.Solver.Model()
		
		args, res := getModelArgs()
		testCases = append(testCases, TestCase{
			Args: args,
			Predicted: res,
			Status: state.Status,
		})
	}
	fmt.Printf(dropped.String())
	return testCases
}


type TestCase struct {
    Args       []interface{}
    Predicted  interface{}
	Status     InterpreterState
}

func generateTestBody(body *strings.Builder, funcName string, cases []TestCase) {
	body.WriteString(fmt.Sprint("func TestGenerated_%s(t *testing.T) {\n", funcName))
	defer body.WriteString("}\n\n")

    for _, cs := range cases {
		strArgs := make([]string, len(parts))
		for i, arg := range cs.Args {
			strArgs[i] = fmt.Sprintf("%#v", arg)
		}
		argsString := strings.Join(strArgs, ", ")

		switch cs.Status.Status {
		case Returned:
            builder.WriteString(fmt.Sprintf(`
    if res := %s(%s); res != %#v {
		t.Errorf("%s() = %%v, but predicted %#v", res)
	}
`,
			funcName, argsString, cs.Predicted, funcName, cs.Predicted))

		case Panic:
			builder.WriteString(fmt.Sprintf(`
    if err := panicCheck(t, %#v, func() { %s(%s) }); err != nil {
		t.Error(err)
	}
`,
			cs.Status.Message, funcName, argsString))
		default:
			panic("Unexpected status while generating test body")
		}
    }
}

func createTestFile(body string, testPath filepath, fileName string) {
	var fileBody strings.Builder

	fileBody.WriteString(fmt.Sprintf("package %s\n\n", file.Name.Name))

	fileBody.WriteString(`import (
	"testing"
)
`)

	fileBody.WriteString(body)
	formatted, err := format.Source([]byte(fileBody.String()))
	if err != nil {
		panic(fmt.Sprintf("Generate tests: Error while formating generated file body: %w", err))
	}

	if err := os.WriteFile(testPath, formatted, 0644); err != nil {
		panic(fmt.Sprintf("Generate tests: Error while writing generated file: %w", err))
	}
}


