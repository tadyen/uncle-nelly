package main

import (
    "fmt"
    "os"

    "gopkg.in/yaml.v3"
)

type effect map[string]interface{}

var effects_yaml = "effects.yaml"
var ingredients_yaml = "ingredients.yaml"

func main(){
    fs, err := os.ReadFile(effects_yaml)
    if err != nil {
        panic(err)
    }

    effects := make(effect)
    err = yaml.Unmarshal(fs, &effects)
    if err != nil {
        panic(err)
    }
    
    fmt.Println("Effects: %v", effects)

}

