package builder

import "fmt"

func (q *QueryBuilder[T]) buildSelect() (query string) {
	selectColsQuery := q.buildSelectCols()
	query = fmt.Sprintf("SELECT %s FROM %s", selectColsQuery, q.table.Name)
	return query
}

func (q *QueryBuilder[T]) buildSelectCols() string {
	if len(q.selectCols) == 0 {
		return "*"
	} else {
		var query string
		for i, col := range q.selectCols {
			query += col
			if i+1 != len(q.selectCols) {
				query += ", "
			}
		}
		return query
	}
}
