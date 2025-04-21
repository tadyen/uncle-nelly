package main

import (
    "fmt"

    UN "github.com/tadyen/uncle-nelly/unclenelly"
)

func main(){
    product := UN.Product{}
    product.Initialize("Meth")
    product.SetMixQueue([]string{"Cuke", "Addy", "Horse Semen"})
    product.MixAll()
    result := product.EffectSet()
    fmt.Printf("Effects: %v\n", result)

    return
}
