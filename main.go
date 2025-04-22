package main

import (
    "fmt"

    UN "github.com/tadyen/uncle-nelly/unclenelly"
)

func main(){
    product := UN.Product{}
    product.Initialize("Meth")
    product.SetMixQueue([]string{"Cuke", "Addy", "Horse Semen", "Viagra", "Banana", "Motor Oil", "Battery", "Addy", "Mega Bean", "Paracetamol"})
    product.MixAll()
    result := product.Effects()
    mult := product.Multiplier()
    for _,v := range result {
        fmt.Println(v)
    }
    fmt.Printf("Multiplier: %f\n", mult)

    return
}
