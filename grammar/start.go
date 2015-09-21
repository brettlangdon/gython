package grammar

import "github.com/brettlangdon/gython/symbol"

type FileInputChild interface {
	Rule
	fileInputChild()
}

type FileInput struct {
	ListRule
}

func NewFileInput() *FileInput {
	rule := &FileInput{}
	rule.initBaseRule(symbol.FILE_INPUT)
	rule.initListRule()
	return rule
}

func (rule *FileInput) Append(n FileInputChild) { rule.ListRule.Append(n) }
