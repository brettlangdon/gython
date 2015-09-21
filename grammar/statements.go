package grammar

import "github.com/brettlangdon/gython/symbol"

type StatementChild interface {
	Rule
	stmtChild()
}

type Statement struct {
	ParentRule
}

func NewStatement() *Statement {
	rule := &Statement{}
	rule.initBaseRule(symbol.STMT)
	return rule
}

func (rule *Statement) fileInputChild()           {}
func (rule *Statement) SetChild(n StatementChild) { rule.ParentRule.SetChild(n) }

type SimpleStatementChild interface {
	Rule
	simpleStatementChild()
}

type SimpleStatement struct {
	ListRule
}

func NewSimpleStatement() *SimpleStatement {
	rule := &SimpleStatement{}
	rule.initBaseRule(symbol.SIMPLE_STMT)
	rule.initListRule()
	return rule
}

func (rule *SimpleStatement) stmtChild()                    {}
func (rule *SimpleStatement) Append(n SimpleStatementChild) { rule.ListRule.Append(n) }

type CompoundStatement struct {
	BaseRule
}

func NewCompoundStatement() *CompoundStatement {
	rule := &CompoundStatement{}
	rule.initBaseRule(symbol.COMPOUND_STMT)
	return rule
}
func (rule *CompoundStatement) stmtChild() {}

type SmallStatementChild interface {
	Rule
	smallStmtChild()
}

type SmallStatement struct {
	ParentRule
}

func NewSmallStatement() *SmallStatement {
	rule := &SmallStatement{}
	rule.initBaseRule(symbol.SMALL_STMT)
	return rule
}

func (rule *SmallStatement) simpleStatementChild()          {}
func (rule *SmallStatement) SetChild(n SmallStatementChild) { rule.ParentRule.SetChild(n) }

type ExpressionStatementChild interface {
	Rule
	expressionStatementChild()
}

type ExpressionStatement struct {
	ListRule
	Expression *TestlistStarExpression
}

func NewExpressionStatement() *ExpressionStatement {
	rule := &ExpressionStatement{}
	rule.initBaseRule(symbol.EXPR_STMT)
	rule.initListRule()
	return rule
}

func (rule *ExpressionStatement) smallStmtChild()                   {}
func (rule *ExpressionStatement) Append(n ExpressionStatementChild) { rule.ListRule.Append(n) }
