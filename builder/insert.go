package builder

import (
	"fmt"
	"gorm/utils"
	"strconv"
)

func buildClausesAndVals[T any](q *QueryBuilder[T]) (colsClause string, valuesClause string, values []any) {
	for i, col := range q.table.Cols {
		if col.Name == "id" {
			continue
		}
		if v := q.insert[col.Name]; utils.IsEmpty(v) {
			continue
		}

		colsClause += col.Name
		valuesClause += "$" + strconv.Itoa(q.parameterIndex)
		if i+1 < len(q.table.Cols) {
			colsClause += ", "
			valuesClause += ", "
		}

		v := q.insert[col.Name]
		values = append(values, v)

		q.parameterIndex++
	}

	return colsClause, valuesClause, values
}

func (q *QueryBuilder[T]) buildInsert() (query string, values []any) {
	colsClause, valuesClause, values := buildClausesAndVals(q)
	query = fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", q.table.Name, colsClause, valuesClause)
	return query, values
}
