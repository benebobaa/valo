package valo

import (
	"errors"
	"reflect"
	"strings"
)

func Validate(s interface{}) error {
	return validateStruct(reflect.ValueOf(s))
}

func validateStruct(v reflect.Value) error {
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		tag := field.Tag.Get("valo")
		if tag == "" {
			continue
		}

		validators := strings.Split(tag, ",")
		for _, validator := range validators {
			parts := strings.Split(validator, "=")
			validatorName := parts[0]

			var validatorValue string
			if len(parts) > 1 {
				validatorValue = parts[1]
			}

			// Handle nested struct,
			// Check tag if its 'valid', iterate field value and type recursively
			if validatorName == "valid" {
				if value.Kind() == reflect.Struct {
					if err := validateStruct(value); err != nil {
						return errors.New(field.Name + ": " + err.Error())
					}
				}
			} else {
				if err := runValidator(validatorName, validatorValue, value); err != nil {
					return errors.New(field.Name + ": " + err.Error())
				}
			}
		}
	}

	return nil
}

func runValidator(name, param string, value reflect.Value) error {
	switch name {
	case "notnil":
		return validateNotNil(value)
	case "notblank":
		return validateNotBlank(value)
	case "min":
		return validateMin(param, value)
	case "max":
		return validateMax(param, value)
	case "numeric":
		return validateNumeric(value)
	case "sizeMin":
		return validateSizeMin(param, value)
	case "sizeMax":
		return validateSizeMax(param, value)
	}
	return nil
}
