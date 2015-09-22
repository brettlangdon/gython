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
	rule := &Test{}
	rule.initBaseNode(symbol.TEST)
	return rule
}

func (rule *Test) testlistStarExpressionChild() {}
func (rule *Test) testChild()                   {}
func (rule *Test) Append(n TestChild)           { rule.ListNode.Append(n) }

type OrTestChild interface {
	Node
	orTestChild()
}

type OrTest struct {
	ListNode
}

func NewOrTest() *OrTest {
	rule := &OrTest{}
	rule.initBaseNode(symbol.OR_TEST)
	return rule
}

func (rule *OrTest) testChild()           {}
func (rule *OrTest) Append(n OrTestChild) { rule.ListNode.Append(n) }

type AndTestChild interface {
	Node
	andTestChild()
}

type AndTest struct {
	ListNode
}

func NewAndTest() *AndTest {
	rule := &AndTest{}
	rule.initBaseNode(symbol.AND_TEST)
	return rule
}

func (rule *AndTest) orTestChild()          {}
func (rule *AndTest) Append(n AndTestChild) { rule.ListNode.Append(n) }

type NotTestChild interface {
	Node
	notTestChild()
}

type NotTest struct {
	ParentNode
}

func NewNotTest() *NotTest {
	rule := &NotTest{}
	rule.initBaseNode(symbol.NOT_TEST)
	return rule
}

func (rule *NotTest) notTestChild()           {}
func (rule *NotTest) andTestChild()           {}
func (rule *NotTest) SetChild(n NotTestChild) { rule.ParentNode.SetChild(n) }
