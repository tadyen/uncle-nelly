package main

import (
    "fmt"
)

func main(){
    effectsLookup := GetEffectsTable()
    fmt.Println("Effects Lookup: %v\n", effectsLookup)
    baseIngredientsLookup := GetBaseIngredientsTable()
    fmt.Println("Base Ingredients Lookup: %v\n", baseIngredientsLookup)
    mixIngredientsLookup := GetMixIngredientsTable()
    fmt.Println("Mix Ingredients Lookup: %v\n", mixIngredientsLookup)
}

