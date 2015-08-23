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
		if tok.ID == token.ENDMARKER || tok.ID == token.ERRORTOKEN {
			break
		}
		fmt.Println(fmt.Sprintf("<%s> %s", tok, tok.Repr()))
	}
}
