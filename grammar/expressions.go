package grammar

import "github.com/brettlangdon/gython/symbol"

type TestlistStarExpressionChild interface {
	Rule
	testlistStarExpressionChild()
}

type TestlistStarExpression struct {
	ParentRule
}

func NewTestListStarExpression() *TestlistStarExpression {
	rule := &TestlistStarExpression{}
	rule.initBaseRule(symbol.TESTLIST_STAR_EXPR)
	return rule
}

func (rule *TestlistStarExpression) expressionStatementChild() {}
func (rule *TestlistStarExpression) SetChild(n TestlistStarExpressionChild) {
	rule.ParentRule.SetChild(n)
}

type ComparisonChild interface {
	Rule
	comparisonChild()
}

type Comparison struct {
	ListRule
}

func NewComparison() *Comparison {
	rule := &Comparison{}
	rule.initBaseRule(symbol.COMPARISON)
	rule.initListRule()
	return rule
}

func (rule *Comparison) notTestChild()            {}
func (rule *Comparison) Append(n ComparisonChild) { rule.ListRule.Append(n) }

type ExpressionChild interface {
	Rule
	expressionChild()
}

type Expression struct {
	ListRule
}

func NewExpression() *Expression {
	rule := &Expression{}
	rule.initBaseRule(symbol.EXPR)
	rule.initListRule()
	return rule
}

func (rule *Expression) comparisonChild()         {}
func (rule *Expression) Append(n ExpressionChild) { rule.ListRule.Append(n) }

type XorExpressionChild interface {
	Rule
	xorExpressionChild()
}

type XorExpression struct {
	ListRule
}

func NewXorExpression() *XorExpression {
	rule := &XorExpression{}
	rule.initBaseRule(symbol.XOR_EXPR)
	rule.initListRule()
	return rule
}

func (rule *XorExpression) expressionChild()            {}
func (rule *XorExpression) Append(n XorExpressionChild) { rule.ListRule.Append(n) }

type AndExpressionChild interface {
	Rule
	andExpressionChild()
}

type AndExpression struct {
	ListRule
}

func NewAndExpression() *AndExpression {
	rule := &AndExpression{}
	rule.initBaseRule(symbol.AND_EXPR)
	rule.initListRule()
	return rule
}

func (rule *AndExpression) xorExpressionChild()         {}
func (rule *AndExpression) Append(n AndExpressionChild) { rule.ListRule.Append(n) }

type ShiftExpressionChild interface {
	Rule
	shiftExpressionChild()
}

type ShiftExpression struct {
	ListRule
}

func NewShiftExpression() *ShiftExpression {
	rule := &ShiftExpression{}
	rule.initBaseRule(symbol.SHIFT_EXPR)
	rule.initListRule()
	return rule
}

func (rule *ShiftExpression) andExpressionChild()           {}
func (rule *ShiftExpression) Append(n ShiftExpressionChild) { rule.ListRule.Append(n) }

type ArithmeticExpressionChild interface {
	Rule
	arithmeticExpressionChild()
}

type ArithmeticExpression struct {
	ListRule
}

func NewArithmeticExpression() *ArithmeticExpression {
	rule := &ArithmeticExpression{}
	rule.initBaseRule(symbol.ARITH_EXPR)
	rule.initListRule()
	return rule
}

func (rule *ArithmeticExpression) shiftExpressionChild()              {}
func (rule *ArithmeticExpression) Append(n ArithmeticExpressionChild) { rule.ListRule.Append(n) }

type TermChild interface {
	Rule
	termChild()
}

type Term struct {
	ListRule
}

func NewTerm() *Term {
	rule := &Term{}
	rule.initBaseRule(symbol.TERM)
	rule.initListRule()
	return rule
}

func (rule *Term) arithmeticExpressionChild() {}
func (rule *Term) Append(n TermChild)         { rule.ListRule.Append(n) }

type FactorChild interface {
	Rule
	factorChild()
}

type Factor struct {
	ListRule
}

func NewFactor() *Factor {
	rule := &Factor{}
	rule.initBaseRule(symbol.FACTOR)
	rule.initListRule()
	return rule
}

func (rule *Factor) factorChild()         {}
func (rule *Factor) powerChild()          {}
func (rule *Factor) termChild()           {}
func (rule *Factor) Append(n FactorChild) { rule.ListRule.Append(n) }

type PowerChild interface {
	Rule
	powerChild()
}

type Power struct {
	ListRule
}

func NewPower() *Power {
	rule := &Power{}
	rule.initBaseRule(symbol.POWER)
	rule.initListRule()
	return rule
}

func (rule *Power) factorChild()        {}
func (rule *Power) Append(n PowerChild) { rule.ListRule.Append(n) }

type AtomExpressionChild interface {
	Rule
	atomExpressionChild()
}

type AtomExpression struct {
	ListRule
}

func NewAtomExpression() *AtomExpression {
	rule := &AtomExpression{}
	rule.initBaseRule(symbol.ATOM_EXPR)
	rule.initListRule()
	return rule
}

func (rule *AtomExpression) powerChild()                  {}
func (rule *AtomExpression) Append(n AtomExpressionChild) { rule.ListRule.Append(n) }

type AtomChild interface {
	Rule
	atomChild()
}

type Atom struct {
	ListRule
}

func NewAtom() *Atom {
	rule := &Atom{}
	rule.initBaseRule(symbol.ATOM)
	rule.initListRule()
	return rule
}

func (rule *Atom) atomExpressionChild() {}
func (rule *Atom) Append(n AtomChild)   { rule.ListRule.Append(n) }

type TrailerChild interface {
	Rule
	trailerChild()
}

type Trailer struct {
	ListRule
}

func NewTrailer() *Trailer {
	rule := &Trailer{}
	rule.initBaseRule(symbol.TRAILER)
	rule.initListRule()
	return rule
}

func (rule *Trailer) atomExpressionChild()  {}
func (rule *Trailer) Append(n TrailerChild) { rule.ListRule.Append(n) }
