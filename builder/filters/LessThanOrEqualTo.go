package filters

type LessThanOrEqualToFilter struct {
	BaseFilter
}

func (f LessThanOrEqualToFilter) GetOperator() string {
	return "<="
}

func (f LessThanOrEqualToFilter) GetClause(parameterIndex *int) (clause string, value any) {
	return f.BaseFilter.GetClause(f.GetOperator(), parameterIndex)
}

func (f LessThanOrEqualToFilter) GetValue() any {
	return f.Value
}

func LessThanOrEqualTo(col string, value any) LessThanOrEqualToFilter {
	return LessThanOrEqualToFilter{
		BaseFilter: BaseFilter{
			Col:   col,
			Value: value,
		},
	}
}

var _ Filter = LessThanOrEqualToFilter{}
