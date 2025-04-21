package main

import (
    "errors"
    "fmt"

    "gopkg.in/yaml.v3"
)

// Data stored as YAML string instead of a .yaml in order to hardcode and build it as wasm

var BaseIngredientsLookup = GetBaseIngredientsTable()
type BaseIngredientName string
func (b BaseIngredientName) BaseIngredientName() (BaseIngredientName, error) {
    if _, ok := BaseIngredientsLookup[string(b)]; ok {
        return b, nil
    }else{
        return "", errors.New(fmt.Sprintf("BaseIngredientName %s not found", b))
    }
}

type BaseIngredientsYaml struct {
    BaseTypePrice map[string]int    `yaml:"BaseTypePrice"`
    BaseIngredients map[string]struct{
        Type string         `yaml:"Type"`
        Effect []string     `yaml:"Effect"`
    } `yaml:"BaseIngredients"`
}

// Todo?: Type (eg Weed) is not statically checked. Treated as generic string field for now. Fix this.
type BaseIngredient struct{
    Name string
    Type string
    Effect []string
    Price int
}

// BaseIngredientRef is a reference to a BaseIngredient by name, providing a Lookup method
type BaseIngredientRef struct{
    Name BaseIngredientName
}
func (b BaseIngredientRef) Lookup() BaseIngredient {
    return BaseIngredientsLookup[string(b.Name)]
}

func GetBaseIngredientsTable() map[string]BaseIngredient{
    var baseIngredients BaseIngredientsYaml
    err := yaml.Unmarshal([]byte(BaseIngredientsRawYAML), &baseIngredients)
    if err != nil {
        panic(err)
    }
    baseIngredientsTable := make(map[string]BaseIngredient)
    for name, ingredient := range baseIngredients.BaseIngredients {
        baseIngredientsTable[name] = BaseIngredient{
            Name: name,
            Type: ingredient.Type,
            Effect: ingredient.Effect,
            Price: baseIngredients.BaseTypePrice[ingredient.Type],
        }
    }
    return baseIngredientsTable
}


var BaseIngredientsRawYAML = `
---
BaseTypePrice:  # <Base>:<Base Price>
    Weed: 45
    Meth: 70
    Cocaine: 150
BaseIngredients:
    "OG Kush":
        Type: Weed
        Effect: 
            - Calming
    "Sour Diesel":
        Type: Weed
        Effect:
            - Refreshing
    "Green Crack":
        Type: Weed
        Effect:
            - Energizing
    "Granddaddy Purple":
        Type: Weed
        Effect:
            - Sedating
    "Meth":
        Type: Meth
        Effect: []
    "Cocaine":
        Type: Cocaine
        Effect: []
`
