package filters

type NotEqualToFilter struct {
	BaseFilter
}

func (f NotEqualToFilter) GetOperator() string {
	return "!="
}

func (f NotEqualToFilter) GetClause(parameterIndex *int) (clause string, value any) {
	return f.BaseFilter.GetClause(f.GetOperator(), parameterIndex)
}

func (f NotEqualToFilter) GetValue() any {
	return f.Value
}

func NotEqualTo(col string, value any) NotEqualToFilter {
	return NotEqualToFilter{
		BaseFilter: BaseFilter{
			Col:   col,
			Value: value,
		},
	}
}

var _ Filter = NotEqualToFilter{}
