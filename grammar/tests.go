package grammar

import "github.com/brettlangdon/gython/symbol"

type TestChild interface {
	Rule
	testChild()
}

type Test struct {
	ListRule
}

func NewTest() *Test {
	rule := &Test{}
	rule.initBaseRule(symbol.TEST)
	return rule
}

func (rule *Test) testlistStarExpressionChild() {}
func (rule *Test) testChild()                   {}
func (rule *Test) Append(n TestChild)           { rule.ListRule.Append(n) }

type OrTestChild interface {
	Rule
	orTestChild()
}

type OrTest struct {
	ListRule
}

func NewOrTest() *OrTest {
	rule := &OrTest{}
	rule.initBaseRule(symbol.OR_TEST)
	return rule
}

func (rule *OrTest) testChild()           {}
func (rule *OrTest) Append(n OrTestChild) { rule.ListRule.Append(n) }

type AndTestChild interface {
	Rule
	andTestChild()
}

type AndTest struct {
	ListRule
}

func NewAndTest() *AndTest {
	rule := &AndTest{}
	rule.initBaseRule(symbol.AND_TEST)
	return rule
}

func (rule *AndTest) orTestChild()          {}
func (rule *AndTest) Append(n AndTestChild) { rule.ListRule.Append(n) }

type NotTestChild interface {
	Rule
	notTestChild()
}

type NotTest struct {
	ParentRule
}

func NewNotTest() *NotTest {
	rule := &NotTest{}
	rule.initBaseRule(symbol.NOT_TEST)
	return rule
}

func (rule *NotTest) notTestChild()           {}
func (rule *NotTest) andTestChild()           {}
func (rule *NotTest) SetChild(n NotTestChild) { rule.ParentRule.SetChild(n) }
