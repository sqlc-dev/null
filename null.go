package nullable

import (
	"database/sql/driver"
)

// An Nullable represents a value that may or may not exist.
//
// The zero value represents a non-existent value.
type Nullable[T any] struct {
		// FIXME: Valid and Val are too similar
	Val  T // What should this be called?
	Valid bool
}

// New creates a new Nullable without a value.
func Empty[T any]() Nullable[T] {
	var empty T
	return Nullable[T]{empty, false}
}

// Of creates a new Nullable with a value.
func Of[T any](val T) Nullable[T] {
	return Nullable[T]{val, true}
}

// It's invalid to use the returned value if the bool is false.
func (n Nullable[T]) Get() (T, bool) {
	return n.Val, n.Valid
}

// Scan implements the Scanner interface.
func (n *Nullable[T]) Scan(value any) error {
	if value == nil {
		var empty T
		n.Valid = false
		n.Val = empty
		return nil
	}
	n.Valid = true
	return convertAssign(&n.Val, value)

	// TODO: Figure out why NullByte uses a different patter for setting Valid
	// https://github.com/golang/go/blob/master/src/database/sql/sql.go#L304
	//
	//   err := convertAssign(&n.Val, value)
	//   n.Valid = err == nil
	//   return err
}

// Value implements the Valuer interface.
func (n Nullable[T]) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return driver.DefaultParameterConverter.ConvertValue(any(n.Val))
}
