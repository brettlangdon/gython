package ast

type FileInput struct {
	ListNode
}

func NewFileInput() *FileInput {
	node := &FileInput{}
	node.initChildren()
	return node
}

func (node *FileInput) ID() NodeID {
	return FILE_INPUT
}

func (node *FileInput) Name() string {
	return NodeNames[FILE_INPUT]
}
