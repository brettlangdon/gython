package main

import (
	"fmt"
	"os"

	"github.com/brettlangdon/gython/scanner"
	"github.com/brettlangdon/gython/token"
)

func main() {
	tokenizer := scanner.NewScanner(os.Stdin)
	for {
		tok := tokenizer.NextToken()
		tokenRange := fmt.Sprintf("%d,%d-%d,%d:", tok.LineStart, tok.ColumnStart, tok.LineEnd, tok.ColumnEnd)
		literalRep := fmt.Sprintf("%#v", tok.Literal)
		fmt.Printf("%-20s%-15s%15s\n", tokenRange, tok.String(), literalRep)
		if tok.ID == token.ENDMARKER || tok.ID == token.ERRORTOKEN {
			break
		}
	}
}
