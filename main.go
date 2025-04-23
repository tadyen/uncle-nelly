//go:build wasm
// +build wasm
package main

import (
    "fmt"
    "syscall/js"

    UN "github.com/tadyen/uncle-nelly/go_baggies/unclenelly"
)

type WrappedProduct struct {
    product     *UN.Product
}

func NewWrappedProduct(this js.Value, 
func main(){
    js.Global().Set("UncleNelly", js.FuncOf(myFunc))

    <- make(chan struct{}) // Block forever so that Go does not terminate execution
}
