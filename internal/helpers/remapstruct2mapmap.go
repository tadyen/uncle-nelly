package helpers

import (
	"reflect"
	"unsafe"
)

// ReMapStructToMapMap recursively flattens structs, and maps with structs into a json-line map.
// accepted concrete types are: map, struct, slice, array, and primitive types.
// not accepted concrete types are: func, chan, unsafe.Pointer, pointer, and complex.
// non-accepted types are converted to nil.
func ReMapStruct2MapMap(obj any) map[string]any {
    result := map[string]any{}
    v := reflect.ValueOf(obj)

    switch v.Kind() {
    case reflect.Map:
        iter := reflect.ValueOf(obj).MapRange()
        for iter.Next() {
            key := iter.Key()
            val := iter.Value()
            result[key.String()] = handleInner(val.Interface())
        }
        return result
    case reflect.Struct:
        mapped := Struct2Map(obj)
        for k, v := range mapped {
            result[k] = handleInner(v)
        }
        return result
    default:
        return nil
    }
}

func handleInner(in any) any {
    v := reflect.ValueOf(in)
    t := v.Type()
    switch t.Kind() {
    case reflect.Map, reflect.Struct:
        return ReMapStruct2MapMap(v.Interface())
    case reflect.Array, reflect.Slice:
        res := []any{}
        for i := range v.Len() {
            res = append(res, handleInner(v.Index(i).Interface()))
        }
        return res
    case reflect.Func, reflect.Chan, reflect.UnsafePointer, reflect.Ptr:
        return nil
    case reflect.Complex64, reflect.Complex128:
        return nil
    case reflect.String:
        return v.String()
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        return int(v.Int())
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
        return v.Uint()
    case reflect.Float32, reflect.Float64:
        return v.Float()
    case reflect.Bool:
        return v.Bool()
    default:
        return nil
    }
}

func Struct2Map(obj any) map[string]any {
	result := map[string]any{}
	v := reflect.ValueOf(obj)
	t := v.Type()
	// Check if the object is a pointer and dereference it
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		t = v.Type()
	}

	if v.Kind() != reflect.Struct {
		return nil // Handle only structs
	}
	for i := range v.NumField(){
		field := t.Field(i)
		fieldValue := v.Field(i)
		// Check if the field is exported
		if fieldValue.CanInterface() {
			result[field.Name] = fieldValue.Interface()
		} else {
            rs := reflect.ValueOf(obj)
            rs2 := reflect.New(rs.Type()).Elem()
            rs2.Set(rs)
            rf := rs2.Field(i)
            rf = reflect.NewAt(rf.Type(), unsafe.Pointer(rf.UnsafeAddr())).Elem()
			result[field.Name] = rf.Interface()
		}
	}
	return result
}


