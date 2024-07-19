package valo

import (
	"reflect"
	"strconv"
	"strings"
)

func validateNotBlank(value reflect.Value) error {
	if value.Kind() == reflect.String && strings.TrimSpace(value.String()) == "" {
		return ErrNotBlank
	}
	return nil
}

func validateSizeMin(param string, value reflect.Value) error {
	min, err := strconv.Atoi(param)
	if err != nil {
		return ErrUnexpected
	}

	switch value.Kind() {
	case reflect.String:
		if value.Len() < min {
			return ErrSizeMin
		}
	}
	return nil
}

func validateSizeMax(param string, value reflect.Value) error {
	max, err := strconv.Atoi(param)
	if err != nil {
		return ErrUnexpected
	}

	switch value.Kind() {
	case reflect.String:
		if value.Len() > max {
			return ErrSizeMax
		}
	}
	return nil
}
