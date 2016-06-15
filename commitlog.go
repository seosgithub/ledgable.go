package models

import "github.com/Masterminds/squirrel"

// A ledger manages a sequence of events
type Ledger struct {
	Name string // Should match SQL table name
}

func NewLedger(name string) Ledger {
	SqlMapped.AddTableWithName(LedgerEntry{}, name)
	return Ledger{Name: name}
}

type LedgerEntry struct {
	/* Managed by Gorp */
	M_Id        int64  `db:"id"`
	M_Type      string `db:"type"`
	M_ObjectId  int64  `db:"object_id"`
	M_Filter    int64  `db:"filter"`
	M_ExpTime   int64  `db:"exp_time"`
	M_CreatedAt int64  `db:"created_at"`
}

/* Call this function first to get a generic query */
func (l *Ledger) Query() LedgerQuery {
	return LedgerQuery{Sql: Sql.Select("*").From(l.Name)}
}

/* Filter for only certain types of entries in the ledger */
func (lq LedgerQuery) FilterTypes(types ...string) LedgerQuery {
	lq.Sql = lq.Sql.Where(squirrel.Eq{"type": types})

	return lq
}

func (lq LedgerQuery) Exec() ([]LedgerEntry, error) {

	//sqlQuery, args, err := lq.Sql.ToSql()
	//if err != nil {
	//return nil, err
	//}

	entries := []LedgerEntry{}
	_, err = SqlMapped.Select(&entries, lq.Sql)
	return entries, err
}
