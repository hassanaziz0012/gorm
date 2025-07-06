package builder

import (
	"context"
	"fmt"
	"gorm/db"
	"gorm/utils"

	"github.com/jackc/pgx/v5"
)

func (q *QueryBuilder[T]) Execute() (items []T, err error) {
	fmt.Println(q.finalQuery, q.finalValues)
	if q.finalQuery == "" {
		return nil, fmt.Errorf("please Build() the query first")
	}

	var rows pgx.Rows
	if q.shouldReturn {
		if !utils.IsEmpty(q.tx) {
			rows, err = q.tx.Query(context.Background(), q.finalQuery, q.finalValues...)
		} else {
			rows, err = db.DB.Query(context.Background(), q.finalQuery, q.finalValues...)
		}

		if err != nil {
			return nil, err
		}
		defer rows.Close()
		if rows == nil {
			return items, fmt.Errorf("no rows returned")
		}
		items = q.scanRows(rows)
		return items, err

	} else {
		if !utils.IsEmpty(q.tx) {
			_, err = q.tx.Exec(context.Background(), q.finalQuery, q.finalValues...)
		} else {
			_, err = db.DB.Exec(context.Background(), q.finalQuery, q.finalValues...)
		}
		if err != nil {
			return nil, err
		}
	}

	return items, err
}
