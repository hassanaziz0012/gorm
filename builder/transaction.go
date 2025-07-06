package builder

import (
	"context"
	"gorm/db"

	"github.com/jackc/pgx/v5"
)

func (q *QueryBuilder[T]) UseTx(tx pgx.Tx) *QueryBuilder[T] {
	q.tx = tx
	return q
}

func BeginTx() (pgx.Tx, error) {
	return db.DB.Begin(context.Background())
}

func RollbackTx(tx pgx.Tx) error {
	return tx.Rollback(context.Background())
}

func CommitTx(tx pgx.Tx) error {
	return tx.Commit(context.Background())
}
