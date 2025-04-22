package main

import (
    "fmt"

    UN "github.com/tadyen/uncle-nelly/unclenelly"
)

func main(){
    product := UN.Product{}
    product.Initialize("Meth")
    product.SetMixQueue([]string{"Cuke", "Addy", "Horse Semen", "Viagor", "Banana", "Motor Oil", "Battery", "Addy", "Mega Bean", "Paracetamol"})
    product.MixAll()
    fmt.Println("Effects:")
    result := product.Effects()
    for _, v := range result{
        fmt.Println("\t", v)
    }
    product.UpdatePrice()
    fmt.Println("Price: ", product.Price())

    return
}
