package valo

import (
	"fmt"
	"reflect"
	"strconv"
)

func validateMin(param string, value reflect.Value, fieldName string) error {
	min, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return ValidationError{Field: fieldName, Message: "unexpected error in min validator"}
	}

	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if float64(value.Int()) < min {
			return ValidationError{Field: fieldName, Message: fmt.Sprintf("must be at least %v", min)}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if float64(value.Uint()) < min {
			return ValidationError{Field: fieldName, Message: fmt.Sprintf("must be at least %v", min)}
		}
	case reflect.Float32, reflect.Float64:
		if value.Float() < min {
			return ValidationError{Field: fieldName, Message: fmt.Sprintf("must be at least %v", min)}
		}
	default:
		return ValidationError{Field: fieldName, Message: "min validator can only be used on numeric fields"}
	}
	return nil
}

func validateMax(param string, value reflect.Value, fieldName string) error {
	max, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return ValidationError{Field: fieldName, Message: "unexpected error in max validator"}
	}

	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if float64(value.Int()) > max {
			return ValidationError{Field: fieldName, Message: fmt.Sprintf("must be at most %v", max)}
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if float64(value.Uint()) > max {
			return ValidationError{Field: fieldName, Message: fmt.Sprintf("must be at most %v", max)}
		}
	case reflect.Float32, reflect.Float64:
		if value.Float() > max {
			return ValidationError{Field: fieldName, Message: fmt.Sprintf("must be at most %v", max)}
		}
	default:
		return ValidationError{Field: fieldName, Message: "max validator can only be used on numeric fields"}
	}
	return nil
}

func validateNumeric(value reflect.Value, fieldName string) error {
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return nil
	case reflect.String:
		_, err := strconv.ParseFloat(value.String(), 64)
		if err != nil {
			return ValidationError{Field: fieldName, Message: "must be a valid numeric value"}
		}
		return nil
	default:
		return ValidationError{Field: fieldName, Message: "must be a numeric field"}
	}
}
