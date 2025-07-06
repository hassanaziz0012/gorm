package types

type Bool string

const (
	TRUE  Bool = "TRUE"
	FALSE Bool = "FALSE"
)

type Defaults struct {
	TimeNow bool
	Text    string
	Int     int
	Bool    Bool
}
