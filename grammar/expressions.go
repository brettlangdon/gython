package grammar

import "github.com/brettlangdon/gython/symbol"

type TestlistStarExpressionChild interface {
	Node
	testlistStarExpressionChild()
}

type TestlistStarExpression struct {
	ParentNode
}

func NewTestListStarExpression() *TestlistStarExpression {
	rule := &TestlistStarExpression{}
	rule.initBaseNode(symbol.TESTLIST_STAR_EXPR)
	return rule
}

func (rule *TestlistStarExpression) expressionStatementChild() {}
func (rule *TestlistStarExpression) SetChild(n TestlistStarExpressionChild) {
	rule.ParentNode.SetChild(n)
}

type ComparisonChild interface {
	Node
	comparisonChild()
}

type Comparison struct {
	ListNode
}

func NewComparison() *Comparison {
	rule := &Comparison{}
	rule.initBaseNode(symbol.COMPARISON)
	rule.initListNode()
	return rule
}

func (rule *Comparison) notTestChild()            {}
func (rule *Comparison) Append(n ComparisonChild) { rule.ListNode.Append(n) }

type ExpressionChild interface {
	Node
	expressionChild()
}

type Expression struct {
	ListNode
}

func NewExpression() *Expression {
	rule := &Expression{}
	rule.initBaseNode(symbol.EXPR)
	rule.initListNode()
	return rule
}

func (rule *Expression) comparisonChild()         {}
func (rule *Expression) Append(n ExpressionChild) { rule.ListNode.Append(n) }

type XorExpressionChild interface {
	Node
	xorExpressionChild()
}

type XorExpression struct {
	ListNode
}

func NewXorExpression() *XorExpression {
	rule := &XorExpression{}
	rule.initBaseNode(symbol.XOR_EXPR)
	rule.initListNode()
	return rule
}

func (rule *XorExpression) expressionChild()            {}
func (rule *XorExpression) Append(n XorExpressionChild) { rule.ListNode.Append(n) }

type AndExpressionChild interface {
	Node
	andExpressionChild()
}

type AndExpression struct {
	ListNode
}

func NewAndExpression() *AndExpression {
	rule := &AndExpression{}
	rule.initBaseNode(symbol.AND_EXPR)
	rule.initListNode()
	return rule
}

func (rule *AndExpression) xorExpressionChild()         {}
func (rule *AndExpression) Append(n AndExpressionChild) { rule.ListNode.Append(n) }

type ShiftExpressionChild interface {
	Node
	shiftExpressionChild()
}

type ShiftExpression struct {
	ListNode
}

func NewShiftExpression() *ShiftExpression {
	rule := &ShiftExpression{}
	rule.initBaseNode(symbol.SHIFT_EXPR)
	rule.initListNode()
	return rule
}

func (rule *ShiftExpression) andExpressionChild()           {}
func (rule *ShiftExpression) Append(n ShiftExpressionChild) { rule.ListNode.Append(n) }

type ArithmeticExpressionChild interface {
	Node
	arithmeticExpressionChild()
}

type ArithmeticExpression struct {
	ListNode
}

func NewArithmeticExpression() *ArithmeticExpression {
	rule := &ArithmeticExpression{}
	rule.initBaseNode(symbol.ARITH_EXPR)
	rule.initListNode()
	return rule
}

func (rule *ArithmeticExpression) shiftExpressionChild()              {}
func (rule *ArithmeticExpression) Append(n ArithmeticExpressionChild) { rule.ListNode.Append(n) }

type TermChild interface {
	Node
	termChild()
}

type Term struct {
	ListNode
}

func NewTerm() *Term {
	rule := &Term{}
	rule.initBaseNode(symbol.TERM)
	rule.initListNode()
	return rule
}

func (rule *Term) arithmeticExpressionChild() {}
func (rule *Term) Append(n TermChild)         { rule.ListNode.Append(n) }

type FactorChild interface {
	Node
	factorChild()
}

type Factor struct {
	ListNode
}

func NewFactor() *Factor {
	rule := &Factor{}
	rule.initBaseNode(symbol.FACTOR)
	rule.initListNode()
	return rule
}

func (rule *Factor) factorChild()         {}
func (rule *Factor) powerChild()          {}
func (rule *Factor) termChild()           {}
func (rule *Factor) Append(n FactorChild) { rule.ListNode.Append(n) }

type PowerChild interface {
	Node
	powerChild()
}

type Power struct {
	ListNode
}

func NewPower() *Power {
	rule := &Power{}
	rule.initBaseNode(symbol.POWER)
	rule.initListNode()
	return rule
}

func (rule *Power) factorChild()        {}
func (rule *Power) Append(n PowerChild) { rule.ListNode.Append(n) }

type AtomExpressionChild interface {
	Node
	atomExpressionChild()
}

type AtomExpression struct {
	ListNode
}

func NewAtomExpression() *AtomExpression {
	rule := &AtomExpression{}
	rule.initBaseNode(symbol.ATOM_EXPR)
	rule.initListNode()
	return rule
}

func (rule *AtomExpression) powerChild()                  {}
func (rule *AtomExpression) Append(n AtomExpressionChild) { rule.ListNode.Append(n) }

type AtomChild interface {
	Node
	atomChild()
}

type Atom struct {
	ListNode
}

func NewAtom() *Atom {
	rule := &Atom{}
	rule.initBaseNode(symbol.ATOM)
	rule.initListNode()
	return rule
}

func (rule *Atom) atomExpressionChild() {}
func (rule *Atom) Append(n AtomChild)   { rule.ListNode.Append(n) }

type TrailerChild interface {
	Node
	trailerChild()
}

type Trailer struct {
	ListNode
}

func NewTrailer() *Trailer {
	rule := &Trailer{}
	rule.initBaseNode(symbol.TRAILER)
	rule.initListNode()
	return rule
}

func (rule *Trailer) atomExpressionChild()  {}
func (rule *Trailer) Append(n TrailerChild) { rule.ListNode.Append(n) }
