package filters

type GreaterThanFilter struct {
	BaseFilter
}

func (f GreaterThanFilter) GetOperator() string {
	return ">"
}

func (f GreaterThanFilter) GetClause(parameterIndex *int) (clause string, value any) {
	return f.BaseFilter.GetClause(f.GetOperator(), parameterIndex)
}

func (f GreaterThanFilter) GetValue() any {
	return f.Value
}

func GreaterThan(col string, value any) GreaterThanFilter {
	return GreaterThanFilter{
		BaseFilter: BaseFilter{
			Col:   col,
			Value: value,
		},
	}

}

var _ Filter = GreaterThanFilter{}
