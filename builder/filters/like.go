package filters

type LikeFilter struct {
	BaseFilter
}

func (f LikeFilter) GetOperator() string {
	return "LIKE"
}

func (f LikeFilter) GetClause(parameterIndex *int) (clause string, value any) {
	return f.BaseFilter.GetClause(f.GetOperator(), parameterIndex)
}

func (f LikeFilter) GetValue() any {
	return f.Value
}

func Like(col string, value any) LikeFilter {
	return LikeFilter{
		BaseFilter: BaseFilter{
			Col:   col,
			Value: value,
		},
	}

}

var _ Filter = LikeFilter{}
