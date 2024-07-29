package valo

import (
	"fmt"
	"reflect"
	"strings"
)

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// Validate performs validation on the given struct
func Validate(s interface{}) error {
	return validateValue(reflect.ValueOf(s), "")
}

func validateValue(v reflect.Value, prefix string) error {
	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			return nil
		}
		return validateValue(v.Elem(), prefix)
	case reflect.Struct:
		return validateStruct(v, prefix)
	case reflect.Slice, reflect.Array:
		return validateSlice(v, prefix)
	}
	return nil
}

func validateStruct(v reflect.Value, prefix string) error {
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		tag := field.Tag.Get("valo")
		if tag == "" {
			continue
		}

		fieldName := field.Name
		if prefix != "" {
			fieldName = prefix + "." + fieldName
		}

		if err := validateField(value, fieldName, tag); err != nil {
			return err
		}
	}
	return nil
}

func validateSlice(v reflect.Value, prefix string) error {
	for i := 0; i < v.Len(); i++ {
		fieldName := fmt.Sprintf("%s[%d]", prefix, i)
		if err := validateValue(v.Index(i), fieldName); err != nil {
			return err
		}
	}
	return nil
}

func validateField(value reflect.Value, fieldName, tag string) error {
	validators := strings.Split(tag, ",")
	for _, validator := range validators {
		parts := strings.SplitN(validator, "=", 2)
		validatorName := parts[0]
		var validatorValue string
		if len(parts) > 1 {
			validatorValue = parts[1]
		}

		if validatorName == "valid" {
			if err := validateValue(value, fieldName); err != nil {
				return err
			}
		} else {
			if err := runValidator(validatorName, validatorValue, value, fieldName); err != nil {
				return err
			}
		}
	}
	return nil
}

func runValidator(name, param string, value reflect.Value, fieldName string) error {
	switch name {
	case "notnil":
		return validateNotNil(value, fieldName)
	case "notblank":
		return validateNotBlank(value, fieldName)
	case "min":
		return validateMin(param, value, fieldName)
	case "max":
		return validateMax(param, value, fieldName)
	case "numeric":
		return validateNumeric(value, fieldName)
	case "sizeMin":
		return validateSizeMin(param, value, fieldName)
	case "sizeMax":
		return validateSizeMax(param, value, fieldName)
	}
	return ValidationError{Field: fieldName, Message: fmt.Sprintf("unknown validator: %s", name)}
}
