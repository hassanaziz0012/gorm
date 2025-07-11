package types

type Column struct {
	// The struct name in Go
	FieldName string
	// The SQL table column name
	Name    string
	Coltype DataType
	Constraints
	Defaults
	Validators
	FKR
}

type FKR struct {
	ConstraintName string
	Name           string
	FKTable        string
	FKField        string
	FKStructField  string
	OnDelete       FKRDelete
}

type FKRDelete int

const (
	CASCADE FKRDelete = iota
)
