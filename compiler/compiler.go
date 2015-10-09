package compiler

import (
	"fmt"

	"github.com/brettlangdon/gython/ast"
	"github.com/brettlangdon/gython/bytecode"
	"github.com/brettlangdon/gython/gython"
)

type Compiler struct {
	nestLevel    int
	currentScope *Scope
	scopeStack   []*Scope
}

func NewCompiler() *Compiler {
	return &Compiler{
		nestLevel:    0,
		currentScope: nil,
		scopeStack:   make([]*Scope, 0),
	}
}

func (compiler *Compiler) enterScope() *Scope {
	scope := NewScope()
	compiler.scopeStack = append(compiler.scopeStack, scope)
	compiler.nestLevel++
	compiler.currentScope = scope
	return scope
}

func (compiler *Compiler) exitScope() *Scope {
	last := len(compiler.scopeStack) - 1
	scope := compiler.scopeStack[last]
	compiler.scopeStack = compiler.scopeStack[:last]
	if last > 0 {
		compiler.currentScope = compiler.scopeStack[last-1]
	} else {
		compiler.currentScope = nil
	}
	compiler.nestLevel--
	return scope
}

func (compiler *Compiler) assemble(addNone bool) *gython.CodeObject {
	if addNone {
		// compiler.addOp(bytecode.LOAD_CONST)
		// compiler.addOp(bytecode.RETURN_VALUE)
	}

	codeobject := gython.NewCodeObject([]byte{}, []byte{}, 0)
	return codeobject
}

func (compiler *Compiler) addOp(op bytecode.Opcode, value gython.Object) bool {
	// TODO: add `value` object and get oparg
	oparg := 0
	instr := NewInstruction(op, oparg, true)
	compiler.currentScope.AddInstruction(instr)
	return true
}

func (compiler *Compiler) visitExpression(expr ast.Expression) bool {
	switch expr := expr.(type) {
	case *ast.Num:
		compiler.addOp(bytecode.LOAD_CONST, expr.Value)
	case *ast.Name:
		compiler.addOp(bytecode.STORE_NAME, expr.Identifier)
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
				// compiler.addOp(bytecode.DUP_TOP)
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

func (compiler *Compiler) CompileMod(root ast.Mod) *gython.CodeObject {
	addNone := true
	compiler.enterScope()
	var codeobject *gython.CodeObject
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
