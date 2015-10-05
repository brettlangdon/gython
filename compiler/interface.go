package compiler

import (
	"github.com/brettlangdon/gython/ast"
	"github.com/brettlangdon/gython/gython"
)

func CompileAST(root ast.Mod) *gython.CodeObject {
	compiler := NewCompiler()
	return compiler.CompileMod(root)
}
