package builder

import (
	"fmt"
	"strconv"
)

func (q *QueryBuilder[T]) buildSetClause() (query string, values []any) {
	for k, v := range q.update {
		parameterIndex := "$" + strconv.Itoa(q.parameterIndex)
		query += fmt.Sprintf(" %s = %s", k, parameterIndex)
		values = append(values, v)
		q.parameterIndex++
	}

	return query, values
}

func (q *QueryBuilder[T]) buildUpdate() (string, []any) {
	setClause, setValues := q.buildSetClause()
	query := fmt.Sprintf("UPDATE %s SET %s", q.table.Name, setClause)

	return query, setValues
}
