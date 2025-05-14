//go:build wasm
// +build wasm

package main

import (
	"errors"
	"fmt"
	"syscall/js"

	UN "github.com/tadyen/uncle-nelly/go_baggies/unclenelly"
    helpers "github.com/tadyen/uncle-nelly/internal/helpers"
)

// wrapper for JS objects
type jsUncleNelly struct {
	job *UN.Job
}

const resOK = "OK"
func flatten (v any) (result any) {
    return helpers.ReMapStruct2MapMap(v)
    // return helpers.Jsonify(v)
}

func jsRes(response any, err error) any {
	if err != nil {
		return js.ValueOf(map[string]any{
			"response": nil,
			"error":    err.Error(),
		})
	}
	return js.ValueOf(map[string]any{
		"response": response,
		"error":    nil,
	})
}

// JS bindings
func NewUncleNelly(this js.Value, args []js.Value) any {
	jsUN := &jsUncleNelly{}
	return js.ValueOf(map[string]any{
		"init_job":         js.FuncOf(jsUN.InitJob),
		"reset_product":    js.FuncOf(jsUN.ResetProduct),
		"get_tables":       js.FuncOf(jsUN.GetTables),
		"set_product_base": js.FuncOf(jsUN.SetProductBase),
		"cook_with":        js.FuncOf(jsUN.CookWith),
		"product_info":     js.FuncOf(jsUN.ProductInfo),
        "test_func":        js.FuncOf(jsUN.TestFunc),
	})
}

func (jsUN *jsUncleNelly) InitJob(this js.Value, args []js.Value) any {
	// invalid nargs
	if len(args) > 1 {
		return jsRes(nil, errors.New("NewJob: expected 0 or 1 argument only for jobname"))
	}
	// Default
	if len(args) == 0 {
		job, err := UN.NewJob("")
		if err != nil {
			return jsRes(nil, err)
		} else {
			jsUN.job = job
			return jsRes(resOK, nil)
		}
	}
	// Validate jobname and set
	var jobName UN.JobName
	jobOptions := []UN.JobName{UN.CookingSim, UN.ReverseCook, UN.Optimise}
	jobSelectOk := true
jobSelect:
	for _, option := range jobOptions {
		if args[0].String() == option.String() {
			jobName = option
			break jobSelect
		}
		jobSelectOk = false
	}
	if !jobSelectOk {
		return jsRes(nil, fmt.Errorf("NewUncleNelly: expected one of %v, got %s", jobOptions, args[0].String()))
	}
	job, err := UN.NewJob(jobName.String())
	if err != nil {
		return jsRes(nil, err)
	}
	jsUN.job = job
	return jsRes(resOK, nil)
}

// will look to maybe refactor or move these around later
func (jsUN *jsUncleNelly) ResetProduct(this js.Value, args []js.Value) any {
	if len(args) != 0 {
		return jsRes(nil, fmt.Errorf("ResetProduct: expected 0 args, got %d", len(args)))
	}
	newProduct, err := UN.NewProduct("")
	if err != nil {
		return jsRes(nil, err)
	}
	jsUN.job.Product = newProduct
	return jsRes(resOK, nil)
}

func (jsUN *jsUncleNelly) GetTables(this js.Value, args []js.Value) any {
	if len(args) != 0 {
		return jsRes(nil, fmt.Errorf("GetTables: expected 0 args, got %d", len(args)))
	}
	tables := map[string]any{
		"effects":          flatten(UN.GetEffectsTable()),
		"mix_ingredients":  flatten(UN.GetMixIngredientsTable()),
		"base_ingredients": flatten(UN.GetBaseIngredientsTable()),
	}
	return jsRes(tables, nil)
}

func (jsUN *jsUncleNelly) SetProductBase(this js.Value, args []js.Value) any {
	if len(args) != 1 {
		return jsRes(nil, fmt.Errorf("SetProductBase: expected 1 arg, got %d", len(args)))
	}
	err := jsUN.job.Product.SetBase(args[0].String())
	if err != nil {
		return jsRes(nil, err)
	}
	return jsRes(resOK, nil)
}

func (jsUN *jsUncleNelly) CookWith(this js.Value, args []js.Value) any {
	if len(args) < 1 {
		return jsRes(nil, errors.New("CookWith: expected >= 1 arg, got none"))
	}
	mix_ingredients := []string{}
	for _, v := range args {
		mix_ingredients = append(mix_ingredients, v.String())
	}
	cooked, err := UN.Cook(jsUN.job.Product, mix_ingredients)
	if err != nil {
		return jsRes(nil, err)
	}
	jsUN.job.Product = cooked
	return jsRes(resOK, nil)
}

func (jsUN *jsUncleNelly) ProductInfo(this js.Value, args []js.Value) any {
	if len(args) != 0 {
		return jsRes(nil, fmt.Errorf("ProductInfo: expected 0 args, got %d", len(args)))
	}
    return jsRes(flatten(jsUN.job.Product.Status()), nil)
}

func (jsUN *jsUncleNelly) TestFunc(this js.Value, args []js.Value) any {
    if len(args) != 0 {
        return jsRes(nil, fmt.Errorf("TestFunc: expected 0 args, got %d", len(args)))
    }
    payload := map[string]any{
        "asdstr": "123",
        "qwenum": 123,
        "qwebool": true,
        "qwemap": map[string]any{
            "key1": "value1",
            "key2": 123,
            "key3": true,
        },
        "qwearr": []any{
            1,2,3,
        },
    }
    return jsRes(helpers.ReMapStruct2MapMap(payload), nil)
}

func main() {
	js.Global().Set("InitUncleNelly", js.FuncOf(NewUncleNelly))

	<-make(chan struct{}) // Block forever so that Go does not terminate execution
}
