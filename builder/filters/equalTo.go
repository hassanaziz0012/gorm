package filters

type EqualToFilter struct {
	BaseFilter
}

func (f EqualToFilter) GetOperator() string {
	return "="
}

func (f EqualToFilter) GetClause(parameterIndex *int) (clause string, value any) {
	return f.BaseFilter.GetClause(f.GetOperator(), parameterIndex)
}

func (f EqualToFilter) GetValue() any {
	return f.Value
}

func EqualTo(col string, value any) EqualToFilter {
	return EqualToFilter{
		BaseFilter: BaseFilter{
			Col:   col,
			Value: value,
		},
	}
}

var _ Filter = EqualToFilter{}
