package ast

type FileInputChildNode interface {
	Node
	fileInputChild()
}

type FileInput struct {
	ListNode
}

func NewFileInput() *FileInput {
	node := &FileInput{}
	node.initBaseNode(FILE_INPUT)
	node.initListNode()
	return node
}

func (node *FileInput) Append(n FileInputChildNode) { node.ListNode.Append(n) }
