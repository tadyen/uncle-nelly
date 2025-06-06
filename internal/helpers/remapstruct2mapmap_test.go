package helpers_test

import (
    "testing"
    "fmt"

    "github.com/google/go-cmp/cmp"
    "github.com/tadyen/uncle-nelly/internal/helpers"
)

func Comparer[T comparable](a, b T) bool {
    return a == b
}

func Compare(a, b any) bool {
    return cmp.Equal(a,b)
}

func TestSanity(t *testing.T){
    tests := []struct{
        name string
        input any
        expected map[string]any
    }{
        {
            name: "empty map",
            input: map[string]any{},
            expected: map[string]any{},
        },
        {
            name: "empty struct",
            input: struct{}{},
            expected: map[string]any{},
        },
        {
            name: "nil",
            input: nil,
            expected: map[string]any{},
        },
    }
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := helpers.ReMapStruct2MapMap(test.input)
            if ! Compare(result, test.expected) {
                fmt.Println("diff", cmp.Diff(result, test.expected))
                t.Errorf("\nexpected %#v,\n got     %#v\n", test.expected, result)
            }
        })
    }
}


func TestSimple(t *testing.T) {
    tests := []struct{
        name string
        input any
        expected map[string]any
    }{
        {
            name: "simple map",
            input: map[string]any{"a": 1.0,"b": 2.0,"c": 3.0,},
            expected: map[string]any{"a": 1.0,"b": 2.0,"c": 3.0,},
        },
        {
            name: "simple struct",
            input: struct{ a int; b string; c float32 }{a: 1, b: "2", c: 3.0},
            expected: map[string]any{"a": 1, "b": "2", "c": 3.0},
        },
        {
            name: "mixed value map",
            input: map[string]any{"a": 1.0, "b": "2", "c": true},
            expected: map[string]any{"a": 1.0, "b": "2", "c": true},
        },
    }
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := helpers.ReMapStruct2MapMap(test.input)
            if ! Compare(result, test.expected) {
                fmt.Println("diff", cmp.Diff(result, test.expected))
                t.Errorf("\nexpected %#v,\n got     %#v\n", test.expected, result)
            }
        })
    }
}

func TestHarder(t *testing.T) {
    tests := []struct{
        name string
        input any
        expected map[string]any
    }{
        {
            name: "map of map",
            input: map[string]any{
                "asdf1": map[string]any{"a": 1.0,"b": 2.0,"c": 3.0,},
                "asdf2": map[int]any{1: 1.0, 2: 2.0, 3: 3.0,},
            },
            expected: map[string]any{
                "asdf1": map[string]any{"a": 1.0,"b": 2.0,"c": 3.0,},
                "asdf2": map[string]any{"1": 1.0, "2": 2.0, "3": 3.0,},     // Note: map[int]any is converted to map[string]any
            },
        },
        {
            name: "map of struct",
            input: map[string]any{
                "asdf1": struct{ a int; b string; c float32 }{a: 1, b: "2", c: 3.0},
            },
            expected: map[string]any{
                "asdf1": map[string]any{"a": 1, "b": "2", "c": 3.0},
            },
        },
        {
            name: "map of array",
            input:      map[string]any{"a": []string{"a", "b", "c"}, "b": []int{1, 2, 3},},
            expected:   map[string]any{"a": []any{"a", "b", "c"}, "b": []any{1, 2, 3},},    // Note: []int is converted to []any
        },
    }
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := helpers.ReMapStruct2MapMap(test.input)
            if ! Compare (result, test.expected) {
                fmt.Println("diff", cmp.Diff(result, test.expected))
                t.Errorf("\nexpected %#v,\n got     %#v\n", test.expected, result)
            }
        })
    }
}
