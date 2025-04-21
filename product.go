package main

const ProductMaxEffects = 8

type Product struct {
    Base BaseIngredientRef
    Multiplier float64
    MixQueue []MixIngredientName
    MixHistory []MixIngredientRef
    Effects [ProductMaxEffects]EffectRef
}

// To assert that effect order in product does not matter
func EffectSet(erefs []EffectRef) map[EffectName]EffectRef {
    effectSet := make(map[EffectName]EffectRef)
    for _, eref := range erefs {
        if eref.Name.Valid() != NoEffect {
            effectSet[eref.Name] = eref
        }
    }
    return effectSet
}

func (p *Product) Initialize(baseIngredient BaseIngredientRef){
    p.Base = baseIngredient
    p.Multiplier = 1.0
    p.MixQueue = []MixIngredientName{}
    p.MixHistory = []MixIngredientRef{}
    for i := range len(p.Effects){
        p.Effects[i] = EffectRef{NoEffect}
    }
    for i, effect := range baseIngredient.Lookup().Effect {
        p.Effects[i] = EffectRef{EffectName(effect).Valid()}
    }
}

func (p *Product) QueueIngredients(ingredients []MixIngredientName){
    p.MixQueue = append(p.MixQueue, ingredients...)
    for _, val := range p.MixQueue{
        val.Valid()
    }
    return
}

func (p *Product) AddEffect(newEffect EffectName){
    for i, e := range p.Effects {
        if e.Name.Valid() == NoEffect {
            p.Effects[i] = EffectRef{newEffect.Valid()}
            break
        }
    }
    return
}

func (p *Product) MixNext() {
    if len(p.MixQueue) == 0 {
        return
    }
    nextIngredient := MixIngredientRef{p.MixQueue[0]}
    p.MixHistory = append(p.MixHistory, nextIngredient)
    p.MixQueue = p.MixQueue[1:]
    nextEffectName := EffectName(nextIngredient.Lookup().Effect)
    for i := range p.Effects {
        p.Effects[i].MutateWith(nextEffectName)
    }
    p.AddEffect(nextEffectName)
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
