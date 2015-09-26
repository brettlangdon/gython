package ast

import (
	"fmt"
	"strings"
)

type Mod interface {
	Node
	mod()
}

type Module struct {
	Body []Statement
}

func NewModule() *Module {
	return &Module{
		Body: make([]Statement, 0),
	}
}

func (module *Module) node() {}
func (module *Module) mod()  {}
func (module *Module) Append(stmt Statement) {
	module.Body = append(module.Body, stmt)
}
func (module *Module) String() string {
	stmts := make([]string, 0)
	for _, stmt := range module.Body {
		stmts = append(stmts, stmt.String())
	}
	return fmt.Sprintf("Module(body=[%s])", strings.Join(stmts, ", "))
}
