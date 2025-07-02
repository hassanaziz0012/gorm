package db

// Validators are different from Constraints. Constraints are
// applied at the DB-level and instantiated at table creation.
// Validators are custom functions executed when you create a
// new object.
type Validators struct {
	IsEmail     bool
	IsURL       bool
	IsMinLength int
	IsMaxLength int
}
