package helpers

import (
	"reflect"
)

// ReMapStructToMapMap recursively flattens structs, and maps with structs into a json-line map.
// accepted concrete types are: map, struct, slice, array, and primitive types.
// not accepted concrete types are: func, chan, unsafe.Pointer, pointer, and complex.
// non-accepted types are converted to nil.
func ReMapStruct2MapMap(obj any) map[string]any {
    result := map[string]any{}
    v := reflect.ValueOf(obj)
    t := v.Type()

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
        for i := range v.NumField(){
            field := t.Field(i)
            value := v.Field(i)
            if field.IsExported() {
                result[field.Name] = handleInner(value.Interface())
            }
        }
        return result
    default:
        return nil
    }
}

func handleInner(value any) any {
    v := reflect.ValueOf(value)
    t := reflect.TypeOf(value)
    switch t.Kind() {
    case reflect.Map, reflect.Struct:
        return ReMapStruct2MapMap(v.Interface())
    case reflect.Array, reflect.Slice:
        var res = make([]any, v.Len())
        for i := range v.Len(){
            val := v.Index(i)
            res = append(res, handleInner(val.Interface()))
        }
        return res
    case reflect.Func, reflect.Chan, reflect.UnsafePointer, reflect.Ptr:
        return nil
    case reflect.Complex64, reflect.Complex128:
        return nil
    case reflect.String:
        return v.String()
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        return v.Int()
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
