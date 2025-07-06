package filters

import (
	"fmt"
	"strconv"
)

type BetweenFilter struct {
	BaseFilter
}

func (f BetweenFilter) GetOperator() string {
	return "BETWEEN"
}

func (f BetweenFilter) GetClause(parameterIndex *int) (clause string, value any) {
	p1 := "$" + strconv.Itoa(*parameterIndex)
	*parameterIndex++
	p2 := "$" + strconv.Itoa(*parameterIndex)

	clause = fmt.Sprintf("%s %s %s AND %s", f.Col, f.GetOperator(), p1, p2)
	return clause, f.Value
}

func (f BetweenFilter) GetValue() any {
	return f.Value
}

func Between(col string, value []any) BetweenFilter {
	return BetweenFilter{
		BaseFilter: BaseFilter{
			Col:   col,
			Value: value,
		},
	}

}

var _ Filter = BetweenFilter{}
