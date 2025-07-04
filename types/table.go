package types

type Table[T any] struct {
	Model T
	Name  string
	Cols  []Column
}
