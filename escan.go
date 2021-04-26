package escan

import (
	"errors"
	"reflect"
)

// Scan saves data into a model easy
func Scan(Model interface{}, Data []string) error {
	items := reflect.ValueOf(Model).Elem()

	if !items.CanAddr() {
		return errors.New("it must be a pointer")
	}

	if items.Kind() != reflect.Struct {
		return errors.New("not a struct")
	}

	value := reflect.Indirect(items) // Pointer to item

	if value.NumField() != len(Data) {
		return errors.New("it should be the same length")
	}

	for i := 0; i < value.NumField(); i++ {
		fieldName := value.Type().Field(i).Name
		actual := value.FieldByName(fieldName)

		if !actual.CanSet() {
			return errors.New("cannot set to field")
		}

		// Passing data as string or full memory address?
		switch actual.Type().String() {
		case "string":
		case "int":
		case "int8":
		case "int16":
		case "int32":
		case "int64":
		case "uint":
		case "uint8":
		case "uint16":
		case "uint32":
		case "uint64":
		case "bool":
		default:
		}

	}

	return nil
}
