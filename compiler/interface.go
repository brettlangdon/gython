package compiler

import (
	"github.com/brettlangdon/gython/ast"
	"github.com/brettlangdon/gython/bytecode"
)

func CompileAST(root ast.Mod) *bytecode.CodeObject {
	compiler := NewCompiler()
	return compiler.CompileMod(root)
}
