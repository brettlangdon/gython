package ast

type TestlistStarExpressionChildNode interface {
	Node
	testlistStarExpressionChild()
}

type TestlistStarExpression struct {
	ParentNode
}

func NewTestListStarExpression() *TestlistStarExpression {
	node := &TestlistStarExpression{}
	node.initBaseNode(TESTLIST_STAR_EXPR)
	return node
}

func (node *TestlistStarExpression) expressionStatementChild() {}
func (node *TestlistStarExpression) SetChild(n TestlistStarExpressionChildNode) {
	node.ParentNode.SetChild(n)
}

type ComparisonChildNode interface {
	Node
	comparisonChild()
}

type Comparison struct {
	ListNode
}

func NewComparison() *Comparison {
	node := &Comparison{}
	node.initBaseNode(COMPARISON)
	node.initListNode()
	return node
}

func (node *Comparison) notTestChild()                {}
func (node *Comparison) Append(n ComparisonChildNode) { node.ListNode.Append(n) }

type ExpressionChildNode interface {
	Node
	expressionChild()
}

type Expression struct {
	ListNode
}

func NewExpression() *Expression {
	node := &Expression{}
	node.initBaseNode(EXPR)
	node.initListNode()
	return node
}

func (node *Expression) comparisonChild()             {}
func (node *Expression) Append(n ExpressionChildNode) { node.ListNode.Append(n) }

type XorExpressionChildNode interface {
	Node
	xorExpressionChild()
}

type XorExpression struct {
	ListNode
}

func NewXorExpression() *XorExpression {
	node := &XorExpression{}
	node.initBaseNode(XOR_EXPR)
	node.initListNode()
	return node
}

func (node *XorExpression) expressionChild()                {}
func (node *XorExpression) Append(n XorExpressionChildNode) { node.ListNode.Append(n) }

type AndExpressionChildNode interface {
	Node
	andExpressionChild()
}

type AndExpression struct {
	ListNode
}

func NewAndExpression() *AndExpression {
	node := &AndExpression{}
	node.initBaseNode(AND_EXPR)
	node.initListNode()
	return node
}

func (node *AndExpression) xorExpressionChild()             {}
func (node *AndExpression) Append(n AndExpressionChildNode) { node.ListNode.Append(n) }

type ShiftExpressionChildNode interface {
	Node
	shiftExpressionChild()
}

type ShiftExpression struct {
	ListNode
}

func NewShiftExpression() *ShiftExpression {
	node := &ShiftExpression{}
	node.initBaseNode(SHIFT_EXPR)
	node.initListNode()
	return node
}

func (node *ShiftExpression) andExpressionChild()               {}
func (node *ShiftExpression) Append(n ShiftExpressionChildNode) { node.ListNode.Append(n) }

type ArithmeticExpressionChildNode interface {
	Node
	arithmeticExpressionChild()
}

type ArithmeticExpression struct {
	ListNode
}

func NewArithmeticExpression() *ArithmeticExpression {
	node := &ArithmeticExpression{}
	node.initBaseNode(ARITH_EXPR)
	node.initListNode()
	return node
}

func (node *ArithmeticExpression) shiftExpressionChild()                  {}
func (node *ArithmeticExpression) Append(n ArithmeticExpressionChildNode) { node.ListNode.Append(n) }

type TermChildNode interface {
	Node
	termChild()
}

type Term struct {
	ListNode
}

func NewTerm() *Term {
	node := &Term{}
	node.initBaseNode(TERM)
	node.initListNode()
	return node
}

func (node *Term) arithmeticExpressionChild() {}
func (node *Term) Append(n TermChildNode)     { node.ListNode.Append(n) }

type FactorChildNode interface {
	Node
	factorChild()
}

type Factor struct {
	ListNode
}

func NewFactor() *Factor {
	node := &Factor{}
	node.initBaseNode(FACTOR)
	node.initListNode()
	return node
}

func (node *Factor) factorChild()             {}
func (node *Factor) powerChild()              {}
func (node *Factor) termChild()               {}
func (node *Factor) Append(n FactorChildNode) { node.ListNode.Append(n) }

type PowerChildNode interface {
	Node
	powerChild()
}

type Power struct {
	ListNode
}

func NewPower() *Power {
	node := &Power{}
	node.initBaseNode(POWER)
	node.initListNode()
	return node
}

func (node *Power) factorChild()            {}
func (node *Power) Append(n PowerChildNode) { node.ListNode.Append(n) }

type AtomExpressionChildNode interface {
	Node
	atomExpressionChild()
}

type AtomExpression struct {
	ListNode
}

func NewAtomExpression() *AtomExpression {
	node := &AtomExpression{}
	node.initBaseNode(ATOM_EXPR)
	node.initListNode()
	return node
}

func (node *AtomExpression) powerChild()                      {}
func (node *AtomExpression) Append(n AtomExpressionChildNode) { node.ListNode.Append(n) }

type AtomChildNode interface {
	Node
	atomChild()
}

type Atom struct {
	ListNode
}

func NewAtom() *Atom {
	node := &Atom{}
	node.initBaseNode(ATOM)
	node.initListNode()
	return node
}

func (node *Atom) atomExpressionChild()   {}
func (node *Atom) Append(n AtomChildNode) { node.ListNode.Append(n) }

type TrailerChildNode interface {
	Node
	trailerChild()
}

type Trailer struct {
	ListNode
}

func NewTrailer() *Trailer {
	node := &Trailer{}
	node.initBaseNode(TRAILER)
	node.initListNode()
	return node
}

func (node *Trailer) atomExpressionChild()      {}
func (node *Trailer) Append(n TrailerChildNode) { node.ListNode.Append(n) }
