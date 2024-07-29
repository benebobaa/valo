package valo

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func validateNotBlank(value reflect.Value, fieldName string) error {
	if value.Kind() != reflect.String {
		return ValidationError{Field: fieldName, Message: "notblank validator can only be used on string fields"}
	}
	if strings.TrimSpace(value.String()) == "" {
		return ValidationError{Field: fieldName, Message: "must not be blank"}
	}
	return nil
}

func validateSizeMin(param string, value reflect.Value, fieldName string) error {
	min, err := strconv.Atoi(param)
	if err != nil {
		return ValidationError{Field: fieldName, Message: "invalid sizemin parameter"}
	}

	switch value.Kind() {
	case reflect.String, reflect.Slice, reflect.Map, reflect.Array:
		if value.Len() < min {
			return ValidationError{Field: fieldName, Message: fmt.Sprintf("must have at least %d elements", min)}
		}
	default:
		return ValidationError{Field: fieldName, Message: "sizemin validator can only be used on strings, slices, maps, or arrays"}
	}
	return nil
}

func validateSizeMax(param string, value reflect.Value, fieldName string) error {
	max, err := strconv.Atoi(param)
	if err != nil {
		return ValidationError{Field: fieldName, Message: "invalid sizemax parameter"}
	}

	switch value.Kind() {
	case reflect.String, reflect.Slice, reflect.Map, reflect.Array:
		if value.Len() > max {
			return ValidationError{Field: fieldName, Message: fmt.Sprintf("must have at most %d elements", max)}
		}
	default:
		return ValidationError{Field: fieldName, Message: "sizemax validator can only be used on strings, slices, maps, or arrays"}
	}
	return nil
}
