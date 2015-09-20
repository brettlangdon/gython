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
	return node
}

func (node *ShiftExpression) andExpressionChild()               {}
func (node *ShiftExpression) Append(n ShiftExpressionChildNode) { node.ListNode.Append(n) }
