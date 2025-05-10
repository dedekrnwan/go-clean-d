package ref

import "reflect"

func IsEmpty(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0.0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Slice, reflect.Map, reflect.Array:
		return value.Len() == 0
	case reflect.Ptr:
		return value.IsNil()
	case reflect.Interface:
		return value.IsNil()
	case reflect.Struct:
		// For structs, we can check if all fields are empty
		for i := 0; i < value.NumField(); i++ {
			if !IsEmpty(value.Field(i)) {
				return false
			}
		}
		return true
	}
	return false
}
