package filters

import "fmt"

type IsNullFilter struct {
	BaseFilter
}

func (f IsNullFilter) GetOperator() string {
	return "IS NULL"
}

func (f IsNullFilter) GetClause(parameterIndex *int) (clause string, value any) {
	return fmt.Sprintf("%s %s", f.Col, f.GetOperator()), f.Value
}

func (f IsNullFilter) GetValue() any {
	return f.Value
}

func IsNull(col string) IsNullFilter {
	return IsNullFilter{
		BaseFilter: BaseFilter{
			Col:   col,
			Value: nil,
		},
	}

}

var _ Filter = IsNullFilter{}
