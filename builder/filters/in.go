package filters

import (
	"fmt"
	"reflect"
	"strconv"
)

type InFilter struct {
	BaseFilter
}

func (f InFilter) GetOperator() string {
	return "IN"
}

func (f InFilter) GetClause(parameterIndex *int) (clause string, value any) {

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

func (f InFilter) GetValue() any {
	return f.Value
}

func In(col string, value []any) InFilter {
	return InFilter{
		BaseFilter: BaseFilter{
			Col:   col,
			Value: value,
		},
	}

}

var _ Filter = InFilter{}
