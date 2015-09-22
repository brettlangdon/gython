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
	rule := &FileInput{}
	rule.initBaseNode(symbol.FILE_INPUT)
	rule.initListNode()
	return rule
}

func (rule *FileInput) Append(n FileInputChild) { rule.ListNode.Append(n) }
