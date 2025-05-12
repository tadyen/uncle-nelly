package helpers_test

import (
    "testing"
    "reflect"
    "github.com/tadyen/uncle-nelly/internal/helpers"
)


func TestSanity(t *testing.T) {
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
            input: map[string]any{"a": 1.0, "b": "2", "c": 3.0},
            expected: map[string]any{"a": 1.0, "b": "2", "c": 3.0},
        },
    }
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := helpers.ReMapStruct2MapMap(test.input)
            if !reflect.DeepEqual(result, test.expected) {
                t.Errorf("expected %v, got %v", test.expected, result)
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
            name: "map of map",
            input: map[string]any{
                "asdf1": map[string]any{"a": 1.0,"b": 2.0,"c": 3.0,},
                "asdf2": map[string]any{"a": 1.0,"b": 2.0,"c": 3.0,},
            },
            expected: map[string]any{
                "asdf1": map[string]any{"a": 1.0,"b": 2.0,"c": 3.0,},
                "asdf2": map[string]any{"a": 1.0,"b": 2.0,"c": 3.0,},
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
            expected:   map[string]any{"a": []string{"a", "b", "c"}, "b": []int{1, 2, 3},},
        },
    }
    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := helpers.ReMapStruct2MapMap(test.input)
            if !reflect.DeepEqual(result, test.expected) {
                t.Errorf("expected %v, got %v", test.expected, result)
            }
        })
    }
}
