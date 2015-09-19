package ast

type TestChildNode interface {
	Node
	testChild()
}

type Test struct {
	ListNode
}

func NewTest() *Test {
	node := &Test{}
	node.initBaseNode(TEST)
	return node
}

func (node *Test) testlistStarExpressionChild() {}
func (node *Test) testChild()                   {}
func (node *Test) Append(n TestChildNode)       { node.ListNode.Append(n) }

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

type AndTestChildNode interface {
	Node
	andTestChild()
}

type AndTest struct {
	ListNode
}

func NewAndTest() *AndTest {
	node := &AndTest{}
	node.initBaseNode(AND_TEST)
	return node
}

func (node *AndTest) orTestChild()              {}
func (node *AndTest) Append(n AndTestChildNode) { node.ListNode.Append(n) }

type NotTestChild interface {
	Node
	notTestChild()
}

type NotTest struct {
	ParentNode
}

func NewNotTest() *NotTest {
	node := &NotTest{}
	node.initBaseNode(NOT_TEST)
	return node
}

func (node *NotTest) notTestChild()           {}
func (node *NotTest) andTestChild()           {}
func (node *NotTest) SetChild(n NotTestChild) { node.ParentNode.SetChild(n) }
