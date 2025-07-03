package types

type Column struct {
	// The struct name in Go
	FieldName string
	// The SQL table column name
	Name    string
	Coltype DataType
	Constraints
	Validators
}
