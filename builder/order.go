package builder

import (
	"fmt"
	"strings"
)

func (q *QueryBuilder[T]) buildOrderClause() (query string) {
	dir := strings.ToUpper(string(q.orderDirection))
	query = fmt.Sprintf(" ORDER BY %s %s ", q.orderBy, dir)
	return query
}
