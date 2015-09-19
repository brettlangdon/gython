package ast

type TestChildNode interface {
	Node
	testChild()
}

type Test struct {
	ParentNode
}

func NewTest() *Test {
	node := &Test{}
	node.initBaseNode(TEST)
	return node
}

func (node *Test) testlistStarExpressionChild() {}
func (node *Test) SetChild(n TestChildNode)     { node.ParentNode.SetChild(n) }

type OrTestChildNode interface {
	Node
	orTestChild()
}

type OrTest struct {
	ListNode
}

func NewOrTest() *OrTest {
	node := &OrTest{}
	node.initBaseNode(OR_TEST)
	return node
}

func (node *OrTest) testChild()               {}
func (node *OrTest) Append(n OrTestChildNode) { node.ListNode.Append(n) }
