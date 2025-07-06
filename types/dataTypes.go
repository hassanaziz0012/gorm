package types

import (
	"fmt"
)

type DataType int

const (
	String DataType = iota
	Integer
	Boolean
	Time
	ForeignKey
)

func (t DataType) String() (string, error) {
	switch t {
	case String:
		return "TEXT", nil
	case Integer:
		return "INT", nil
	case Boolean:
		return "BOOLEAN", nil
	case Time:
		return "TIMESTAMPTZ", nil
	case ForeignKey:
		return "FOREIGN KEY", nil
	}
	return "", fmt.Errorf("invalid data type")
}
