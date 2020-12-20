package utils

import (
	"reflect"
)

//Contains - Generic function to implement contains operation over arrays
func Contains(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	if isArrayType(arr.Kind()) {
		panic("Invalid data-type")
	}

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

func isArrayType(kind reflect.Kind) bool {
	return kind != reflect.Array && kind != reflect.Slice
}
