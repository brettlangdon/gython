package ast

import "github.com/brettlangdon/gython/token"

type FileInput struct {
	BaseNode
	children []interface{}
}

func NewFileInput() *FileInput {
	node := &FileInput{
		children: make([]interface{}, 0),
	}
	node.initBaseNode(FILE_INPUT)
	return node
}

func (node *FileInput) AppendToken(t *token.Token) {
	node.children = append(node.children, t)
}

func (node *FileInput) AppendNode(n StatementNode) {
	node.children = append(node.children, n)
}

func (node *FileInput) Children() []interface{} {
	return node.children
}
