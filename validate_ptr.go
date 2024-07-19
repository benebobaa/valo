package valo

import "reflect"

func validateNotNil(value reflect.Value) error {
	switch value.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Interface:
		if value.IsNil() {
			return ErrNotNil
		}
	}
	return nil
}
