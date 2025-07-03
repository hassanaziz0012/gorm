package types

import (
	"log"
)

type DataType int

const (
	String DataType = iota
	Integer
	Boolean
	Time
)

func (t DataType) String() string {
	switch t {
	case String:
		return "TEXT"
	case Integer:
		return "INT"
	case Boolean:
		return "BOOLEAN"
	case Time:
		return "TIMESTAMPTZ"
	default:
		log.Fatal("invalid data type")
	}
	return ""
}
