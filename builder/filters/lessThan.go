package filters

type LessThanFilter struct {
	BaseFilter
}

func (f LessThanFilter) GetOperator() string {
	return "<"
}

func (f LessThanFilter) GetClause(parameterIndex *int) (clause string, value any) {
	return f.BaseFilter.GetClause(f.GetOperator(), parameterIndex)
}

func (f LessThanFilter) GetValue() any {
	return f.Value
}

func LessThan(col string, value any) LessThanFilter {
	return LessThanFilter{
		BaseFilter: BaseFilter{
			Col:   col,
			Value: value,
		},
	}

}

var _ Filter = LessThanFilter{}
