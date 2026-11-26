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
        panic(fmt.Errorf("Generate tests: Error while getting filapath: %w", err))
	}
	ext := filepath.Ext(filePath)
	if ext != ".go" {
		panic(fmt.Errorf("Generate tests: Error file needs to be '.go': %w", err))
	}
	source, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Errorf("Generate tests: Error while reading file: %w", err))
	}


	file, err := parser.ParseFile(token.NewFileSet(), filePath, source, 0)
	if err != nil {
		panic(fmt.Errorf("Generate tests: Error while parsing file: %w", err))
	}

	var funcNames []string
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl);
        if !ok || fn.Body == nil {
            continue
        }
		analyzeAndTestFunc(source, fn.Name.Name)
    }


	testsFile := strings.TrimSuffix(filePath, ext) + "_test.go"

    return testsFile
}


func analyzeAndTestFunc(source []byte, funcName string) {
	states := Analyse(source, name)
	var body strings.Builder
	initTestBody()
	fmt.Fprintf(&body, "func TestGenerated_%s(t *testing.T) {\n", name)
	var dropped strings.Builder
	for _, state := range states {
		if state.Status.Status == Running {
			panic("Unexpected behavour. Analyse returned running interpreter")
		}
		if state.Status.Status == Dropped {
			dropped.WriteString(fmt.Sprintf("Dropped in func='%s' with message='%s', in pathCondition=%s", funcName, state.Status.Message, state.PathCondition))
		}
		
		if sat := state.Analyser.solverCheck(state.PathFormula); !sat {
			continue
		}
		model := state.Analyser.Solver.Model()
		
		...
		generateTestBody()

		frame := state.frame()
		arguments := make([]string, len(frame.Function.Params))

		for i, param := range frame.Function.Params {
			expression := symbolic.NewSymbolicVariable(param.Name(), expressionType(param.Type()))
			value, err := translator.TranslateExpression(expression)
			if err != nil {
				panic(err)
			}
			literal, floating := modelLiteral(model, value.(z3.Value), expression.Type())
			usesMath = usesMath || floating
			typ := types.TypeString(param.Type(), func(p *types.Package) string {
				if p == frame.Function.Pkg.Pkg {
					return ""
				}
				return p.Name()
			})
			arguments[i] = fmt.Sprintf("%s(%s)", typ, literal)
		}
		fmt.Fprintf(&body, "t.Run(\"path_%d\", func(t *testing.T) {\n", cases)
		call := fmt.Sprintf("%s(%s)", name, strings.Join(arguments, ", "))
		if state.Status == Panicked {
			body.WriteString("completed := false\ndefer func() { value := recover(); if completed { t.Fatal(\"expected panic\") }; _ = value }()\n")
			body.WriteString(call + "\ncompleted = true\n")
		} else {
}


type TestCase struct {
    Args       []interface{}
    Predicted  interface{}
	Status     InterpreterState
}

func generateTestBody(body *strings.Builder funcName string, cases []TestCase) {
	body.WriteString(fmt.Sprintf("func Test%s(t *testing.T) {\n", funcName))
	defer builder.WriteString("}\n\n")

	wasPanicCase := false
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

func createTestFile(filename string) (*bufio.Writer, *os.File, error) {
	tmpl, err := template.New("tests").Parse(testTemplate)
    if err != nil {
        return "", err
    }

    var builder strings.Builder

    builder.WriteString("package main\n\n")
    builder.WriteString(`import (
	"testing"
)
`)
    file, err := os.Create(filename)
    if err != nil {
        return nil, nil, err
    }

	builder.WriteString(`
func panicCheck(t *testing.T, panicExpect string, fn func()) {
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
	`)

    writer := bufio.NewWriter(file)

    return writer, file, nil
}


