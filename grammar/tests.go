package grammar

import "github.com/brettlangdon/gython/symbol"

type TestChild interface {
	Node
	testChild()
}

type Test struct {
	ListNode
}

func NewTest() *Test {
	node := &Test{}
	node.initBaseNode(symbol.TEST)
	return node
}

func (node *Test) testlistStarExpressionChild() {}
func (node *Test) testChild()                   {}
func (node *Test) Append(n TestChild)           { node.ListNode.Append(n) }

type OrTestChild interface {
	Node
	orTestChild()
}

type OrTest struct {
	ListNode
}

func NewOrTest() *OrTest {
	node := &OrTest{}
	node.initBaseNode(symbol.OR_TEST)
	return node
}

func (node *OrTest) testChild()           {}
func (node *OrTest) Append(n OrTestChild) { node.ListNode.Append(n) }

type AndTestChild interface {
	Node
	andTestChild()
}

type AndTest struct {
	ListNode
}

func NewAndTest() *AndTest {
	node := &AndTest{}
	node.initBaseNode(symbol.AND_TEST)
	return node
}

func (node *AndTest) orTestChild()          {}
func (node *AndTest) Append(n AndTestChild) { node.ListNode.Append(n) }

type NotTestChild interface {
	Node
	notTestChild()
}

type NotTest struct {
	ParentNode
}

func NewNotTest() *NotTest {
	node := &NotTest{}
	node.initBaseNode(symbol.NOT_TEST)
	return node
}

func (node *NotTest) notTestChild()           {}
func (node *NotTest) andTestChild()           {}
func (node *NotTest) SetChild(n NotTestChild) { node.ParentNode.SetChild(n) }
