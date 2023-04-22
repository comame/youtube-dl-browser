package server

import "reflect"

func validateIsFilled(obj interface{}) bool {
	v := reflect.ValueOf(obj)
	switch v.Kind() {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i += 1 {
			child := v.FieldByIndex([]int{i})
			return validateIsFilled(child.Interface())
		}
	case reflect.String:
		return !v.IsZero()
	default:
		panic("unimplemented")
	}
	panic("unreachable")
}
