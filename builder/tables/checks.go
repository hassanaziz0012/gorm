package tables

import (
	"fmt"
	"gorm/builder/filters"
)

type Check struct {
	Col    string
	Filter filters.Filter
}

func (c Check) BuildClause() string {
	return fmt.Sprintf("CHECK (%s %s %v)", c.Col, c.Filter.GetOperator(), c.Filter.GetValue())
}
