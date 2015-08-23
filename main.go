package main

import (
	"fmt"
	"os"

	"github.com/brettlangdon/gython/parser"
	"github.com/brettlangdon/gython/token"
)

func main() {
	tokenizer, err := parser.TokenizerFromFileName(os.Args[1])
	if err != nil {
		panic(err)
	}
	for {
		tok := tokenizer.Next()
		tokenRange := fmt.Sprintf("%d,%d-%d,%d:", tok.LineStart, tok.ColumnStart, tok.LineEnd, tok.ColumnEnd)
		literalRep := fmt.Sprintf("%#v", tok.Literal)
		fmt.Printf("%-20s%-15s%15s\n", tokenRange, tok.String(), literalRep)
		if tok.ID == token.ENDMARKER || tok.ID == token.ERRORTOKEN {
			break
		}
	}
}
