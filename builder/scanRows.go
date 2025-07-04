package builder

import (
	"gorm/db"
	"log"
	"reflect"

	"github.com/jackc/pgx/v5"
)

func (q *QueryBuilder[T]) scanRows(rows pgx.Rows) (items []T) {
	for rows.Next() {
		v := reflect.New(reflect.TypeOf(q.model)).Elem()
		dest := db.PrepareScanDest(q.table, v)

		if err := rows.Scan(dest...); err != nil {
			log.Fatal("failed to scan row: ", err)
		}

		item := v.Interface().(T)
		items = append(items, item)
	}

	return items
}
