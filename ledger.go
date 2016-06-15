package ledgable

type LedgerQueryRunner func(filter *LedgerQuery) []*LedgerCommit
type Ledger struct {
	applyFuncs map[string]LedgerApplyFunc

	queryRunner LedgerQueryRunner
}

// Create a new ledger
func NewLedger(queryRunner LedgerQueryRunner) *Ledger {
	return &Ledger{
		applyFuncs:  map[string]LedgerApplyFunc{},
		queryRunner: queryRunner,
	}
}

type LedgerQuery struct {
	ledger *Ledger
}

func (l *Ledger) Query() *LedgerQuery {
	return &LedgerQuery{
		ledger: l,
	}
}

type LedgerCommit struct {
	Id   int
	Type string
}

type LedgerApplyFunc func(res interface{}, commit *LedgerCommit)

func (l *Ledger) RegisterApply(name string, applyFunc LedgerApplyFunc) {
	l.applyFuncs[name] = applyFunc
}

func (lq *LedgerQuery) WithType(_type string) *LedgerQuery {
	return lq
}

func (lq *LedgerQuery) Apply(name string, obj interface{}) {
	// Retrieve all entries from this query

	entries := lq.ledger.queryRunner(lq)

	applyFunc := lq.ledger.applyFuncs[name]
	for _, entry := range entries {
		applyFunc(obj, entry)
	}
}
