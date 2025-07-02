package db

import (
	"fmt"
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

func ParseDataType(typeName string) (DataType, error) {
	switch typeName {
	case "string":
		return String, nil
	case "uint", "int":
		return Integer, nil
	case "bool":
		return Boolean, nil
	case "Time":
		return Time, nil
	}

	return String, fmt.Errorf("invalid type name")
}
