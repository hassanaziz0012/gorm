package builder

import (
	"fmt"
	"strconv"
)

func (q *QueryBuilder[T]) buildWhereClause() (query string, values []any) {
	query = " WHERE "

	for k, v := range q.where {
		parameterIndex := "$" + strconv.Itoa(q.parameterIndex)
		query += fmt.Sprintf("%s = %s", k, parameterIndex)

		if q.parameterIndex < len(q.where) {
			query += " AND "
		}

		values = append(values, v)
		q.parameterIndex++
	}
	return query, values
}
