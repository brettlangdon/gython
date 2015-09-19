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
func (node *FileInput) AppendNode(n StatementNode) {
	node.ListNode.AppendNode(n)
}
