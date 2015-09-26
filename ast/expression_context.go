package ast

type ExpressionContext interface {
	Node
	exprCtx()
}

type Store struct{}

func NewStore() *Store {
	return &Store{}
}

func (store *Store) node()          {}
func (store *Store) exprCtx()       {}
func (store *Store) String() string { return "Store()" }

type Load struct{}

func NewLoad() *Load {
	return &Load{}
}

func (load *Load) node()          {}
func (load *Load) exprCtx()       {}
func (load *Load) String() string { return "Load()" }
