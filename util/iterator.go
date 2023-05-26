package util

import (
	"fmt"
	"reflect"
)

// Check if an item is in an array/slice
func ItemExists(arrayType any, item any) (bool, error) {
	arr := reflect.ValueOf(arrayType)

	if arr.Kind() != reflect.Slice {
		return false, fmt.Errorf("must be a reflect.Array")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true, nil
		}
	}

	return false, nil
}
