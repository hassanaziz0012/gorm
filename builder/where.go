package builder

import (
	"fmt"
	"reflect"
	"strings"
)

func (q *QueryBuilder[T]) buildWhereClause() (query string, values []any) {
	query = " WHERE "

	var finalCombine string
	for _, conditionGroup := range q.where {
		finalCombine = string(conditionGroup.Combine)

		for _, f := range conditionGroup.Filters {
			filterQuery, value := f.GetClause(&q.parameterIndex)
			query += filterQuery

			query += fmt.Sprintf(" %s ", conditionGroup.Combine)

			if value != nil {
				t := reflect.TypeOf(value)
				switch t.Kind() {
				case reflect.Array, reflect.Slice:
					v := reflect.ValueOf(value)
					for i := 0; i < v.Len(); i++ {
						val := v.Index(i)
						var valueToAppend any = val.Interface()
						values = append(values, valueToAppend)
					}
				default:
					values = append(values, value)
				}
			}
		}
	}

	query = strings.TrimSuffix(query, " "+finalCombine+" ")
	return query, values
}
