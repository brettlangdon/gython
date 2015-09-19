package ast

type FileInput struct {
	ListNode
}

func NewFileInput() *FileInput {
	node := &FileInput{}
	node.initBaseNode(FILE_INPUT)
	node.initListNode()
	return node
}

func (node *FileInput) Append(n *Statement) { node.ListNode.Append(n) }
