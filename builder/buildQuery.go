package builder

import (
	"fmt"
	"gorm/utils"
)

func (q *QueryBuilder[T]) Build() *QueryBuilder[T] {
	var query string
	var values []any

	switch q.queryType {
	case SELECT:
		query = q.buildSelect()
	case INSERT:
		query, values = q.buildInsert()
	case UPDATE:
		query, values = q.buildUpdate()
	case DELETE:
		query = q.buildDelete()
	default:
		panic("query type is not set")
	}

	if len(q.where) > 0 {
		whereClause, whereValues := q.buildWhereClause()
		query += whereClause
		values = append(values, whereValues...)
	}

	if q.orderBy != "" {
		orderByClause := q.buildOrderClause()
		query += orderByClause
	}

	if q.limit > 0 {
		query += fmt.Sprintf(" LIMIT %d ", q.limit)
	}

	if q.offset > 0 {
		query += fmt.Sprintf(" OFFSET %d ", q.offset)
	}

	query = utils.RemoveExtraSpaces(query)
	query += ";"

	q.finalQuery = query
	q.finalValues = values
	return q
}
