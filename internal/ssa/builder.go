// Package ssa предоставляет функции для построения SSA представления
package ssa

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/ssa"
)

// Builder отвечает за построение SSA из исходного кода Go
type Builder struct {
	fset *token.FileSet
}

// NewBuilder создаёт новый экземпляр Builder
func NewBuilder() *Builder {
	return &Builder{
		fset: token.NewFileSet(),
	}
}

// TODO: Реализуйте следующие методы в рамках домашнего задания

// ParseAndBuildSSA парсит исходный код Go и создаёт SSA представление
// Возвращает SSA программу и функцию по имени
func (b *Builder) ParseAndBuildSSA(source string, funcName string) (*ssa.Function, error) {
	// TODO: Реализовать
	// Шаги:
	// 1. Парсинг исходного кода с помощью go/parser
	// 2. Создание SSA программы
	// 3. Поиск нужной функции по имени

	// Подсказки:
	// - Используйте parser.ParseFile для парсинга
	// - Создайте packages.Config и загрузите пакет
	// - Используйте ssautil.CreateProgram для создания SSA
	// - Найдите функцию в SSA программе

	file, err := parser.ParseFile(b.fset, "source.go", source, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("Error in parsing: %w", err)
	}

	info := &types.Info{
		Types:        make(map[ast.Expr]types.TypeAndValue),
		Instances:    make(map[*ast.Ident]types.Instance),
		Defs:         make(map[*ast.Ident]types.Object),
		Uses:         make(map[*ast.Ident]types.Object),
		Implicits:    make(map[ast.Node]types.Object),
		Selections:   make(map[*ast.SelectorExpr]*types.Selection),
		Scopes:       make(map[ast.Node]*types.Scope),
		FileVersions: make(map[*ast.File]string),
	}

	cnf := types.Config{}
	pkg, err := cnf.Check(file.Name.Name, b.fset, []*ast.File{file}, info)
	if err != nil {
		return nil, fmt.Errorf("Error type-checking package: %w", err)
	}

	prog := ssa.NewProgram(b.fset, ssa.SanityCheckFunctions)
	ssaPkg := prog.CreatePackage(pkg, []*ast.File{file}, info, true)
	prog.Build()

	function := ssaPkg.Func(funcName)
	if function == nil {
		return nil, fmt.Errorf("Error function '%s'not found", funcName)
	}

	return function, nil
}
