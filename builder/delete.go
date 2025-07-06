package builder

import "fmt"

func (q *QueryBuilder[T]) buildDelete() (query string) {
	query = fmt.Sprintf("DELETE FROM %s", q.table.Name)
	return query
}
