package filters

import "fmt"

type IsNotNullFilter struct {
	BaseFilter
}

func (f IsNotNullFilter) GetOperator() string {
	return "IS NOT NULL"
}

func (f IsNotNullFilter) GetClause(parameterIndex *int) (clause string, value any) {
	return fmt.Sprintf("%s %s", f.Col, f.GetOperator()), f.Value
}

func (f IsNotNullFilter) GetValue() any {
	return f.Value
}

func IsNotNull(col string) IsNotNullFilter {
	return IsNotNullFilter{
		BaseFilter: BaseFilter{
			Col:   col,
			Value: nil,
		},
	}

}

var _ Filter = IsNotNullFilter{}
