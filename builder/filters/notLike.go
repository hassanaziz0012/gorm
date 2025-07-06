package filters

type NotLikeFilter struct {
	BaseFilter
}

func (f NotLikeFilter) GetOperator() string {
	return "NOT LIKE"
}

func (f NotLikeFilter) GetClause(parameterIndex *int) (clause string, value any) {
	return f.BaseFilter.GetClause(f.GetOperator(), parameterIndex)
}

func (f NotLikeFilter) GetValue() any {
	return f.Value
}

func NotLike(col string, value any) NotLikeFilter {
	return NotLikeFilter{
		BaseFilter: BaseFilter{
			Col:   col,
			Value: value,
		},
	}

}

var _ Filter = NotLikeFilter{}
