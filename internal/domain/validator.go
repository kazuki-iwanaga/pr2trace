package domain

import (
	"errors"
	"time"
)

var (
	ErrInvalidTimeOrder = errors.New("invalid time order")
	ErrEmptySlice       = errors.New("empty slice")
)

// ValidateTimeOrder checks if t1 < t2.
func ValidateTimeOrder(t1 time.Time, t2 time.Time) error {
	if t1.After(t2) {
		return ErrInvalidTimeOrder
	}

	return nil
}

// ValidateNotEmptySlice checks if the slice is not empty.
func ValidateNotEmptySlice[T any](s []T) error {
	if len(s) == 0 {
		return ErrEmptySlice
	}

	return nil
}
