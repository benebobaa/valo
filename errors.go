package valo

import "errors"

var (
	ErrNotBlank   = errors.New("field must not be blank")
	ErrMin        = errors.New("field must be greater than or equal to minimum value")
	ErrMax        = errors.New("field must be less than or equal to maximum value")
	ErrNumeric    = errors.New("field must be numeric")
	ErrSizeMin    = errors.New("field size must be greater than or equal to minimum size")
	ErrSizeMax    = errors.New("field size must be less than or equal to maximum size")
	ErrUnexpected = errors.New("unexpected error occurred during validation")
	ErrNotNil     = errors.New("value cannot be nil")
)
