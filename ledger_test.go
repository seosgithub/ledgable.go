package models

import (
	"testing"

	. "dattch.com/hercall/kernutil"
	. "dattch.com/hercall/persist"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLedgerInit(t *testing.T) {
	Convey("Can create a new ledger", t, func() {
		SetSettingsString("persist_sql_mock", "TRUE")

		PersistInit()
		defer PersistTeardown()

		ModelsInit()
		defer ModelsTeardown()

		NewLedger("acl")
	})
	Convey("Can query a ledger", t, func() {
		SetSettingsString("persist_sql_mock", "TRUE")

		PersistInit()
		defer PersistTeardown()

		ModelsInit()
		defer ModelsTeardown()

		// Create our users table
		Sql.Exec(`
			CREATE TABLE acl_log (
				id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
				type TEXT NOT NULL,
				object_id INT NOT NULL DEFAULT(0),
				filter INT NOT NULL DEFAULT(0),
				exp_time INT NOT NULL DEFAULT(0),
				created_at INT NOT NULL DEFAULT(0)
			);

			INSERT INTO acl_log (type, object_id, filter) VALUES ('foo', 13, 1);
			INSERT INTO acl_log (type, object_id, filter) VALUES ('bar', 14, 2);
		`)

		aclLedger := NewLedger("acl_log")

		res, err := aclLedger.Query().FilterTypes("foo", "bar").Exec()
		if err != nil {
			Panic("%s", err)
		}

		So(len(res), ShouldEqual, 2)
	})
}
