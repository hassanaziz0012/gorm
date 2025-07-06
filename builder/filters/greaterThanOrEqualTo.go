package filters

type GreaterThanOrEqualToFilter struct {
	BaseFilter
}

func (f GreaterThanOrEqualToFilter) GetOperator() string {
	return ">="
}

func (f GreaterThanOrEqualToFilter) GetClause(parameterIndex *int) (clause string, value any) {
	return f.BaseFilter.GetClause(f.GetOperator(), parameterIndex)
}

func (f GreaterThanOrEqualToFilter) GetValue() any {
	return f.Value
}

func GreaterThanOrEqualTo(col string, value any) GreaterThanOrEqualToFilter {
	return GreaterThanOrEqualToFilter{
		BaseFilter: BaseFilter{
			Col:   col,
			Value: value,
		},
	}
}

var _ Filter = GreaterThanOrEqualToFilter{}
