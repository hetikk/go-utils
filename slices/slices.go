package slices

import (
	"reflect"
)

// Проверяет вхождение элемента в слайс
func Contains(src interface{}, val interface{}) bool {
	slice := reflect.ValueOf(src)

	if slice.Kind() == reflect.Slice {
		for i := 0; i < slice.Len(); i++ {
			if slice.Index(i).Interface() == val {
				return true
			}
		}
	}

	return false
}
