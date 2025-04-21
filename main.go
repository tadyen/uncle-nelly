package main

import (
    "fmt"
)

const NoEffect = EffectName("None")

type EffectState struct{
    Name EffectName
}
func (e *EffectState) Lookup() Effect{
    return EffectsLookup[string(e.Name)]
}

func (e *EffectState) MutateWith(effect EffectName){
    conversion := e.Lookup().Conversion
    for _, entry := range conversion {
        for name, result := range entry{
            if name == string(effect) {
                e.Name = EffectName(result)
                break
            }
        }
    }
    return
}
const ProductMaxEffects = 8
type Product struct {
    Base BaseIngredientName
    Multiplier float64
    MixQueue []MixIngredientName
    MixHistory []MixIngredientName
    Effects [ProductMaxEffects]EffectState
}

func (p *Product) Initialize(baseIngredient BaseIngredientName){
    p.Base = baseIngredient
    p.Multiplier = 1.0
    p.MixQueue = []MixIngredientName{}
    p.MixHistory = []MixIngredientName{}
    for i := range len(p.Effects){
        p.Effects[i] = EffectState{NoEffect}
    }
    for i, effect := range BaseIngredientsLookup[string(baseIngredient)].Effect {
        p.Effects[i] = EffectState{EffectName(effect)}
    }
}

func (p *Product) QueueIngredients(ingredients []MixIngredientName){
    p.MixQueue = append(p.MixQueue, ingredients...)
    return
}

func (p *Product) AddEffect(newEffect EffectName){
    for i, e := range p.Effects {
        if e.Name == NoEffect {
            p.Effects[i] = EffectState{newEffect}
            break
        }
    }
    return
}

func (p *Product) MixNext() {
    if len(p.MixQueue) == 0 {
        return
    }
    nextIngredient := p.MixQueue[0]
    p.MixHistory = append(p.MixHistory, nextIngredient)
    p.MixQueue = p.MixQueue[1:]
    nextEffect := EffectName( MixIngredientsLookup[string(nextIngredient)].Effect )
    for i := range p.Effects {
        p.Effects[i].MutateWith(nextEffect)
    }
    p.AddEffect(nextEffect)
    return
}

func (p *Product) UpdateMultiplier(){
    p.Multiplier = 1.0
    for _, e := range p.Effects {
        if e.Name != NoEffect {
            p.Multiplier += e.Lookup().Multiplier
        }
    }
    return
}

func (p *Product) MixAll(){
    for len(p.MixQueue) > 0 {
        p.MixNext()
    }
    return
}

func main(){
    // Test example product:
    // OG Kush + Cuke + Mega Bean -> [Glowing, Cyclopean, Foggy]
    product := Product{}
    product.Initialize("OG Kush")
    product.MixQueue = []MixIngredientName{"Cuke", "Mega Bean"}
    product.Base = "OG Kush"
    product.MixAll()
    product.UpdateMultiplier()
    fmt.Printf("Product: %s\n", product)

}



