package valo

import "reflect"

func validateNotNil(value reflect.Value, fieldName string) error {
	switch value.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Interface, reflect.Map, reflect.Chan, reflect.Func:
		if value.IsNil() {
			return ValidationError{Field: fieldName, Message: "must not be nil"}
		}
	}
	return nil
}
