//go:build wasm
// +build wasm
package main

import (
    "fmt"
    "syscall/js"

    UN "github.com/tadyen/uncle-nelly/go_baggies/unclenelly"
)

// wrapper for JS objects
type jsUncleNelly struct {
    unclenelly *UN.Job
}


func NewUncleNelly(this js.Value, args []js.Value) any {
    if len(args) != 1 {
        panic("NewJob: expected 1 argument")
    }
    var jobName UN.JobName
    jobOptions := []UN.JobName{UN.CookingSim, UN.ReverseCook, UN.Optimise}
    jobSelectOk := true
    jobSelect: for _, option := range jobOptions {
        if args[0].String() == option.String() {
            jobName = option
            break jobSelect
        }
        jobSelectOk = false
    }
    if !jobSelectOk {
        panic(fmt.Sprintf("NewUncleNelly: expected one of %v, got %s", jobOptions, args[0].String()))
    }
    job, err := UN.NewJob(jobName.String())
    if err != nil {
        panic(err)
    }
    jsUncleNelly := &jsUncleNelly{
        unclenelly: job,
    }
    return js.ValueOf(map[string]any{
        "init_product": js.FuncOf(jsUncleNelly.unclenelly.Product.Initialize),
    })

}

func main(){
    js.Global().Set("InitUncleNelly", js.FuncOf(NewUncleNelly))

    <- make(chan struct{}) // Block forever so that Go does not terminate execution
}
