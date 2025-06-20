package unclenelly

import (
    "gopkg.in/yaml.v3"
)

// Data stored as YAML string instead of a .yaml in order to hardcode and build it as wasm


var MixIngredientsLookup = GetMixIngredientsTable()
type MixIngredientName string
func (m MixIngredientName) Valid() bool{
    if _, ok := MixIngredientsLookup[string(m)]; ok{
        return true
    }else{
        return false
    }
}

type MixIngredientsYAML map[string]struct {
    Effect string   `yaml:"Effect"`
    Price  int32    `yaml:"Price"`
}

type MixIngredient struct {
    Name    string
    Effect  string
    Price   int32
}

// MixIngredientRef is a reference to a mix ingredient, providing a lookup method
type MixIngredientRef struct {
    Name MixIngredientName
}
func (m MixIngredientRef) Lookup() MixIngredient {
    return MixIngredientsLookup[string(m.Name)]
}

func GetMixIngredientsTable() map[string]MixIngredient{
    var mixIngredients MixIngredientsYAML
    err := yaml.Unmarshal([]byte(MixIngredientsRawYAML), &mixIngredients)
    if err != nil {
        panic(err)
    }
    mixIngredientsTable := make(map[string]MixIngredient)
    for name, ingredient := range mixIngredients {
        mixIngredientsTable[name] = MixIngredient{
            Name: name,
            Effect: ingredient.Effect,
            Price: ingredient.Price,
        }
    }
    return mixIngredientsTable
}

var MixIngredientsRawYAML = `
# Kush cooking effects table
# https://docs.google.com/spreadsheets/d/1Swo-SuDGqPy5hSvRVM-Moix8RjlqQkql0nl1_8CREUM/edit?usp=sharing
---
# Ingredient: 
#   Effect: <Effect>
#   Price: <int>
Cuke:
    Effect: Energizing
    Price: 2
Banana: 
    Effect: Gingeritis
    Price: 2
Paracetamol:
    Effect: Sneaky
    Price: 3
Donut:
    Effect: Calorie-Dense
    Price: 3
Viagor:
    Effect: "Tropic Thunder"
    Price: 4
"Mouth Wash":
    Effect: Balding
    Price: 4
"Flu Medicine":
    Effect: Sedating
    Price: 5
Gasoline:
    Effect: Toxic
    Price: 5
"Energy Drink":
    Effect: Athletic
    Price: 6
"Motor Oil":
    Effect: Slippery
    Price: 6
"Mega Bean":
    Effect: Foggy
    Price: 7
Chili:
    Effect: Spicy
    Price: 7
Battery: 
    Effect: Bright-Eyed
    Price: 8
Iodine:
    Effect: Jennerising
    Price: 8
Addy:
    Effect: Thought-Provoking
    Price: 9
"Horse Semen":
    Effect: "Long faced"
    Price: 9
`
