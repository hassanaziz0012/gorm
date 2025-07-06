package filters

import (
	"fmt"
	"gorm/utils"
	"reflect"
	"strconv"
)

type CombineType string

const (
	AND  CombineType = "AND"
	OR   CombineType = "OR"
	XAND CombineType = "AND NOT"
	XOR  CombineType = "OR NOT"
)

type ConditionGroup struct {
	Filters []Filter
	Combine CombineType
}

type Filter interface {
	GetClause(parameterIndex *int) (clause string, value any)
	GetOperator() string
	GetValue() any
}

type BaseFilter struct {
	Col   string
	Value any
}

func (f BaseFilter) GetClause(op string, parameterIndex *int) (clause string, value any) {
	switch utils.IsEmpty(parameterIndex) {
	case true:
		return fmt.Sprintf("%s %s %s", f.Col, op, f.Value), f.Value
	case false:
		i := int(reflect.ValueOf(parameterIndex).Elem().Int())
		parameter := "$" + strconv.Itoa(i)
		*parameterIndex++
		return fmt.Sprintf("%s %s %s", f.Col, op, parameter), f.Value
	}
	return clause, value
}
