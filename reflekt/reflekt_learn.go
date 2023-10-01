package reflekt

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)
	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}
	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		// numberOfValues = val.NumField()
		// getFields = val.Field
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		// numberOfValues = val.Len()
		// getFields = val.Index
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walkValue(v)
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}
	// fn("I still can't believe south korea beat germany 2-0 to put last in their group")
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	return val
}

// func walk(x interface{}, fn func(input string)) {
// 	val := getValue(x)
// 	if val.Kind() == reflect.Slice {
// 		for i := 0; i < val.Len(); i++ {
// 			walk(val.Index(i).Interface(), fn)
// 		}
// 		return
// 	}
//
// 	for l := 0; l < val.NumField(); l++ {
// 		field := val.Field(l)
// 		switch field.Kind() {
// 		case reflect.String:
// 			fn(field.String())
// 		case reflect.Struct:
// 			walk(field.Interface(), fn)
// 		}
// 	}
// 	// fn("I still can't believe south korea beat germany 2-0 to put last in their group")
// }

// func walk(x interface{}, fn func(input string)) {
// 	val := reflect.ValueOf(x)
// 	for l := 0; l < val.NumField(); l++ {
// 		field := val.Field(l)
// 		if field.Kind() == reflect.String {
// 			fn(field.String())
// 		}
//
// 		if field.Kind() == reflect.Struct {
// 			walk(field.Interface(), fn)
// 		}
// 	}
// 	// fn("I still can't believe south korea beat germany 2-0 to put last in their group")
// }
