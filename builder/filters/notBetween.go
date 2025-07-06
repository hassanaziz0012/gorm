package filters

import (
	"fmt"
	"strconv"
)

type NotBetweenFilter struct {
	BaseFilter
}

func (f NotBetweenFilter) GetOperator() string {
	return "NOT BETWEEN"
}

func (f NotBetweenFilter) GetClause(parameterIndex *int) (clause string, value any) {
	p1 := "$" + strconv.Itoa(*parameterIndex)
	*parameterIndex++
	p2 := "$" + strconv.Itoa(*parameterIndex)

	clause = fmt.Sprintf("%s %s %s AND %s", f.Col, f.GetOperator(), p1, p2)
	return clause, f.Value
}

func (f NotBetweenFilter) GetValue() any {
	return f.Value
}

func NotBetween(col string, value []int) NotBetweenFilter {
	return NotBetweenFilter{
		BaseFilter: BaseFilter{
			Col:   col,
			Value: value,
		},
	}

}

var _ Filter = NotBetweenFilter{}
