package main

import (
	"fmt"
	"os"

	"github.com/brettlangdon/gython/ast"
	"github.com/brettlangdon/gython/compiler"
	"github.com/brettlangdon/gython/grammar"
	"github.com/brettlangdon/gython/scanner"
	"github.com/brettlangdon/gython/token"
)

func tokenize() {
	tokenizer := scanner.NewScanner(os.Stdin)
	for {
		tok := tokenizer.NextToken()
		tokenRange := fmt.Sprintf("%d,%d-%d,%d:", tok.LineStart, tok.ColumnStart, tok.LineEnd, tok.ColumnEnd)
		literalRep := fmt.Sprintf("%#v", tok.Literal)
		fmt.Printf("%-20s%-15s%-15s\n", tokenRange, tok.String(), literalRep)
		if tok.ID == token.ENDMARKER || tok.ID == token.ERRORTOKEN {
			break
		}
	}
}

func parseGrammar() *grammar.FileInput {
	tokenizer := scanner.NewScanner(os.Stdin)
	gp := grammar.NewGrammarParser(tokenizer)
	return gp.Parse()
}

func parseAST() ast.Mod {
	start := parseGrammar()
	mod, err := ast.ASTFromGrammar(start)
	if err != nil {
		panic(err)
	}
	return mod
}

func compile() {
	root := parseAST()
	codeobject := compiler.CompileAST(root)
	fmt.Println(codeobject)
}

func main() {
	// start := parseGrammar()
	// fmt.Println(start.Repr())
	// root := parseAST()
	// fmt.Println(root)
	compile()
}
