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
	node := &TestlistStarExpression{}
	node.initBaseNode(symbol.TESTLIST_STAR_EXPR)
	return node
}

func (node *TestlistStarExpression) expressionStatementChild() {}
func (node *TestlistStarExpression) SetChild(n TestlistStarExpressionChild) {
	node.ParentNode.SetChild(n)
}

type ComparisonChild interface {
	Node
	comparisonChild()
}

type Comparison struct {
	ListNode
}

func NewComparison() *Comparison {
	node := &Comparison{}
	node.initBaseNode(symbol.COMPARISON)
	node.initListNode()
	return node
}

func (node *Comparison) notTestChild()            {}
func (node *Comparison) Append(n ComparisonChild) { node.ListNode.Append(n) }

type ExpressionChild interface {
	Node
	expressionChild()
}

type Expression struct {
	ListNode
}

func NewExpression() *Expression {
	node := &Expression{}
	node.initBaseNode(symbol.EXPR)
	node.initListNode()
	return node
}

func (node *Expression) comparisonChild()         {}
func (node *Expression) Append(n ExpressionChild) { node.ListNode.Append(n) }

type XorExpressionChild interface {
	Node
	xorExpressionChild()
}

type XorExpression struct {
	ListNode
}

func NewXorExpression() *XorExpression {
	node := &XorExpression{}
	node.initBaseNode(symbol.XOR_EXPR)
	node.initListNode()
	return node
}

func (node *XorExpression) expressionChild()            {}
func (node *XorExpression) Append(n XorExpressionChild) { node.ListNode.Append(n) }

type AndExpressionChild interface {
	Node
	andExpressionChild()
}

type AndExpression struct {
	ListNode
}

func NewAndExpression() *AndExpression {
	node := &AndExpression{}
	node.initBaseNode(symbol.AND_EXPR)
	node.initListNode()
	return node
}

func (node *AndExpression) xorExpressionChild()         {}
func (node *AndExpression) Append(n AndExpressionChild) { node.ListNode.Append(n) }

type ShiftExpressionChild interface {
	Node
	shiftExpressionChild()
}

type ShiftExpression struct {
	ListNode
}

func NewShiftExpression() *ShiftExpression {
	node := &ShiftExpression{}
	node.initBaseNode(symbol.SHIFT_EXPR)
	node.initListNode()
	return node
}

func (node *ShiftExpression) andExpressionChild()           {}
func (node *ShiftExpression) Append(n ShiftExpressionChild) { node.ListNode.Append(n) }

type ArithmeticExpressionChild interface {
	Node
	arithmeticExpressionChild()
}

type ArithmeticExpression struct {
	ListNode
}

func NewArithmeticExpression() *ArithmeticExpression {
	node := &ArithmeticExpression{}
	node.initBaseNode(symbol.ARITH_EXPR)
	node.initListNode()
	return node
}

func (node *ArithmeticExpression) shiftExpressionChild()              {}
func (node *ArithmeticExpression) Append(n ArithmeticExpressionChild) { node.ListNode.Append(n) }

type TermChild interface {
	Node
	termChild()
}

type Term struct {
	ListNode
}

func NewTerm() *Term {
	node := &Term{}
	node.initBaseNode(symbol.TERM)
	node.initListNode()
	return node
}

func (node *Term) arithmeticExpressionChild() {}
func (node *Term) Append(n TermChild)         { node.ListNode.Append(n) }

type FactorChild interface {
	Node
	factorChild()
}

type Factor struct {
	ListNode
}

func NewFactor() *Factor {
	node := &Factor{}
	node.initBaseNode(symbol.FACTOR)
	node.initListNode()
	return node
}

func (node *Factor) factorChild()         {}
func (node *Factor) powerChild()          {}
func (node *Factor) termChild()           {}
func (node *Factor) Append(n FactorChild) { node.ListNode.Append(n) }

type PowerChild interface {
	Node
	powerChild()
}

type Power struct {
	ListNode
}

func NewPower() *Power {
	node := &Power{}
	node.initBaseNode(symbol.POWER)
	node.initListNode()
	return node
}

func (node *Power) factorChild()        {}
func (node *Power) Append(n PowerChild) { node.ListNode.Append(n) }

type AtomExpressionChild interface {
	Node
	atomExpressionChild()
}

type AtomExpression struct {
	ListNode
}

func NewAtomExpression() *AtomExpression {
	node := &AtomExpression{}
	node.initBaseNode(symbol.ATOM_EXPR)
	node.initListNode()
	return node
}

func (node *AtomExpression) powerChild()                  {}
func (node *AtomExpression) Append(n AtomExpressionChild) { node.ListNode.Append(n) }

type AtomChild interface {
	Node
	atomChild()
}

type Atom struct {
	ListNode
}

func NewAtom() *Atom {
	node := &Atom{}
	node.initBaseNode(symbol.ATOM)
	node.initListNode()
	return node
}

func (node *Atom) atomExpressionChild() {}
func (node *Atom) Append(n AtomChild)   { node.ListNode.Append(n) }

type TrailerChild interface {
	Node
	trailerChild()
}

type Trailer struct {
	ListNode
}

func NewTrailer() *Trailer {
	node := &Trailer{}
	node.initBaseNode(symbol.TRAILER)
	node.initListNode()
	return node
}

func (node *Trailer) atomExpressionChild()  {}
func (node *Trailer) Append(n TrailerChild) { node.ListNode.Append(n) }
