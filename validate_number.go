package valo

import (
	"reflect"
	"strconv"
)

func validateMin(param string, value reflect.Value) error {
	min, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return ErrUnexpected
	}

	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if float64(value.Int()) < min {
			return ErrMin
		}

	case reflect.Float32, reflect.Float64:
		if value.Float() < min {
			return ErrMin
		}
	}
	return nil
}

func validateMax(param string, value reflect.Value) error {
	max, err := strconv.ParseFloat(param, 64)
	if err != nil {
		return ErrUnexpected
	}

	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if float64(value.Int()) > max {
			return ErrMax
		}
	case reflect.Float32, reflect.Float64:
		if value.Float() > max {
			return ErrMax
		}
	}
	return nil
}

func validateNumeric(value reflect.Value) error {
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return nil
	}
	return ErrNumeric
}
