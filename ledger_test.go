package ledgable

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type AApplyTargetObject struct {
	types []string
}

func TestLedgerInit(t *testing.T) {
	Convey("Can create a new ledger", t, func() {
		queryRunner := func(lq *LedgerQuery) []*LedgerCommit {
			return []*LedgerCommit{&LedgerCommit{Type: "lol"}}
		}
		l := NewLedger(queryRunner)
		l.RegisterApply("test", func(res interface{}, commit *LedgerCommit) {
			target := res.(*AApplyTargetObject)

			target.types = append(target.types, commit.Type)
		})

		q := l.Query()

		obj := &AApplyTargetObject{}
		q.WithType("foo").Apply("test", obj)

		So(obj.types[0], ShouldEqual, "lol")
	})
}
