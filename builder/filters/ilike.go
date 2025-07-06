package filters

type ILikeFilter struct {
	BaseFilter
}

func (f ILikeFilter) GetOperator() string {
	return "ILIKE"
}

func (f ILikeFilter) GetClause(parameterIndex *int) (clause string, value any) {
	return f.BaseFilter.GetClause(f.GetOperator(), parameterIndex)
}

func (f ILikeFilter) GetValue() any {
	return f.Value
}

func ILike(col string, value any) ILikeFilter {
	return ILikeFilter{
		BaseFilter: BaseFilter{
			Col:   col,
			Value: value,
		},
	}

}

var _ Filter = ILikeFilter{}
