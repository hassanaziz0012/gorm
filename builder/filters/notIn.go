package filters

import (
	"fmt"
	"reflect"
	"strconv"
)

type NotInFilter struct {
	BaseFilter
}

func (f NotInFilter) GetOperator() string {
	return "NOT IN"
}

func (f NotInFilter) GetClause(parameterIndex *int) (clause string, value any) {

	var valuesClause string
	values := reflect.ValueOf(f.Value)
	switch values.Kind() {
	case reflect.Array, reflect.Slice:
		for i := 0; i < values.Len(); i++ {
			param := "$" + strconv.Itoa(*parameterIndex)

			if i+1 != values.Len() {
				param += ", "
				*parameterIndex++
			}

			valuesClause += param
		}
	}

	clause = fmt.Sprintf("%s %s (%s)", f.Col, f.GetOperator(), valuesClause)
	return clause, f.Value
}

func (f NotInFilter) GetValue() any {
	return f.Value
}

func NotIn(col string, value []any) NotInFilter {
	return NotInFilter{
		BaseFilter: BaseFilter{
			Col:   col,
			Value: value,
		},
	}

}

var _ Filter = NotInFilter{}
