package cyrillicfilter

import (
	"errors"
	"reflect"
	"strings"
	"unsafe"
)

var (
	NilInterfaceError = errors.New("CyrillicFilter nil argument error")
	NotStructError    = errors.New("CyrillicFilter argument is not struct")
)

func CyrillicFilter(v interface{}) error {
	if v == nil {
		return NilInterfaceError
	}

	myStruct := reflect.ValueOf(v)

	if myStruct.Elem().Kind() != reflect.Struct {
		return NotStructError
	}

	for i := 0; i < myStruct.Elem().NumField(); i++ {
		field := myStruct.Elem().Field(i)
		field = reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem()

		if field.Type().Kind() == reflect.String && field.CanSet() {
			field.SetString(cleanCyrillic(field.Interface().(string)))
		}

		if field.Type().Kind() == reflect.Ptr {
			pointerField := field.Elem()

			if pointerField.Kind() == reflect.String && pointerField.CanSet() {
				pointerField.SetString(cleanCyrillic(pointerField.Interface().(string)))
			} else if pointerField.Kind() == reflect.Struct {
				if err := CyrillicFilter(field.Interface()); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func cleanCyrillic(s string) string {
	var result strings.Builder

	for _, v := range []rune(s) {
		if (v < 'А' || v > 'я') && v != 'Ё' && v != 'ё' {
			result.WriteString(string(v))
		}
	}

	return result.String()
}
