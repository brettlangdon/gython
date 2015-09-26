package compiler

import (
	"fmt"

	"github.com/brettlangdon/gython/ast"
	"github.com/brettlangdon/gython/bytecode"
)

type Compiler struct {
	nestLevel  int
	scopeStack []*Scope
}

func NewCompiler() *Compiler {
	return &Compiler{
		nestLevel:  0,
		scopeStack: make([]*Scope, 0),
	}
}

func (compiler *Compiler) currentScope() *Scope {
	length := len(compiler.scopeStack)
	if length > 0 {
		return compiler.scopeStack[length-1]
	}
	return nil
}

func (compiler *Compiler) enterScope() *Scope {
	scope := NewScope()
	compiler.scopeStack = append(compiler.scopeStack, scope)
	compiler.nestLevel++
	return scope
}

func (compiler *Compiler) exitScope() *Scope {
	last := len(compiler.scopeStack) - 1
	scope := compiler.scopeStack[last]
	compiler.scopeStack = compiler.scopeStack[:last]
	compiler.nestLevel--
	return scope
}

func (compiler *Compiler) assemble(addNone bool) *bytecode.CodeObject {
	if addNone {
		compiler.addOp(bytecode.LOAD_CONST)
		compiler.addOp(bytecode.RETURN_VALUE)
	}

	codeobject := bytecode.NewCodeObject()
	return codeobject
}

func (compiler *Compiler) addOp(op bytecode.Opcode) bool {
	fmt.Println(bytecode.Opnames[op])
	return true
}

func (compiler *Compiler) visitExpression(expr ast.Expression) bool {
	switch expr := expr.(type) {
	case *ast.Num:
		compiler.addOp(bytecode.LOAD_CONST)
	case *ast.Name:
		compiler.addOp(bytecode.STORE_NAME)
	default:
		fmt.Println(expr)
	}

	return true
}

func (compiler *Compiler) visitStatement(stmt ast.Statement) bool {
	switch stmt := stmt.(type) {
	case *ast.Assign:
		compiler.visitExpression(stmt.Value)
		length := len(stmt.Targets)
		for i := 0; i < length; i++ {
			if i < length-1 {
				compiler.addOp(bytecode.DUP_TOP)
			}
			compiler.visitExpression(stmt.Targets[i])
		}
	}
	return true
}

func (compiler *Compiler) compileBody(stmts []ast.Statement) bool {
	// TODO: Check for docstring
	for _, stmt := range stmts {
		compiler.visitStatement(stmt)
	}
	return true
}

func (compiler *Compiler) CompileMod(root ast.Mod) *bytecode.CodeObject {
	addNone := true
	compiler.enterScope()
	var codeobject *bytecode.CodeObject
	switch root := root.(type) {
	case *ast.Module:
		if !compiler.compileBody(root.Body) {
			compiler.exitScope()
			return nil
		}
	}

	codeobject = compiler.assemble(addNone)
	compiler.exitScope()
	return codeobject
}
