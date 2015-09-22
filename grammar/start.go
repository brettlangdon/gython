package grammar

import "github.com/brettlangdon/gython/symbol"

type FileInputChild interface {
	Node
	fileInputChild()
}

type FileInput struct {
	ListNode
}

func NewFileInput() *FileInput {
	node := &FileInput{}
	node.initBaseNode(symbol.FILE_INPUT)
	node.initListNode()
	return node
}

func (node *FileInput) Append(n FileInputChild) { node.ListNode.Append(n) }
