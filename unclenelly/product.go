package unclenelly

import (
    "math"
    "fmt"
)

const ProductMaxEffects = 8

type StatefulEffect struct {
    Current     EffectRef
    Previous    EffectRef
}
func (e *StatefulEffect) Reset() {
    e.Current = EffectRef{NoEffect.Valid()}
    e.Previous = EffectRef{NoEffect.Valid()}
}
func (e *StatefulEffect) SetCurrentEffect(effect EffectName){
    if e.Current == (EffectRef{}){
        e.Reset()
    }
    e.Current = EffectRef{effect.Valid()}
}

func (e *StatefulEffect) MutateWith(input EffectName){
    if e.Current == (EffectRef{}) {
        e.Reset()
        return
    }
    if (e.Current.Name == NoEffect) || (e.Current.Name == input) {
        return
    }
    prev := e.Current.Name
    e.Current.MutateWith(input)
    if e.Current.Name == prev {
        // no change
        return
    }
    e.Previous = EffectRef{prev}
}

// Revert is a one-way operation.
func (e *StatefulEffect) Revert() (ok bool) {
    if e.Current.Name == e.Previous.Name {
        return false
    }
    e.Current = e.Previous
    return true
}

// Product is the exposed version of the product struct for simpler usage
// Assertations are done in the safe version
type Product struct {
    SafeProduct *SafeProduct
}

// Safe Product is for type safety. The 'unsafe' version wraps the safe version
// avoiding the need to convert string types
type SafeProduct struct {
    Base        BaseIngredientRef
    Multiplier  float64
    Price       int32
    MixQueue    []MixIngredientName
    MixHistory  []MixIngredientRef
    Effects     [ProductMaxEffects]StatefulEffect
}


// Getters
func (p *Product) Base() string{
    return string(p.SafeProduct.Base.Name)
}
func (p *Product) Multiplier() float64{
    return p.SafeProduct.Multiplier
}
func (p *Product) Price() int32{
    return p.SafeProduct.Price
}
func (p *Product) MixQueue() []string{
    mixQueue := make([]string, len(p.SafeProduct.MixQueue))
    for i, ingredient := range p.SafeProduct.MixQueue {
        mixQueue[i] = string(ingredient)
    }
    return mixQueue
}
func (p *Product) MixHistory() []string{
    mixHist:= make([]string, len(p.SafeProduct.MixHistory))
    for i, ref := range p.SafeProduct.MixHistory{
        mixHist[i] = string(ref.Name)
    }
    return mixHist
}
func (p *Product) Effects() []string{
    effects := make([]string, ProductMaxEffects)
    for i, effect := range p.SafeProduct.Effects {
        effects[i] = string(effect.Current.Name)
    }
    return effects
}
  
// Setters
func (p *Product) SetBase(baseIngredient string){
    p.SafeProduct.Base = BaseIngredientRef{BaseIngredientName(baseIngredient).Valid()}
}
func (p *Product) SetMultiplier(multiplier float64){
    p.SafeProduct.Multiplier = multiplier
}
func (p *Product) SetPrice(price int32){
    p.SafeProduct.Price = price
}
func (p *Product) SetMixQueue(ingredients []string){
    safeIngredients := make([]MixIngredientName, len(ingredients))
    for i, ingredient := range ingredients {
        safeIngredients[i] = MixIngredientName(ingredient).Valid()
    }
    p.SafeProduct.MixQueue = safeIngredients
}
func (p *Product) SetMixHistory(ingredients []string){
    safeRefs := make([]MixIngredientRef, len(ingredients))
    for i, ingredient := range ingredients {
        safeRefs[i] = MixIngredientRef{MixIngredientName(ingredient).Valid()}
    }
    p.SafeProduct.MixHistory = safeRefs
}

// Each method exists as a pair of safe and unsafe methods. The unsafe wraps the safe

// Do not allow setting effects directly. Effects must be cleared and added one-by-one via AddEffect()
// TODO: Add a wrapper for adding multiple effects 
func (p *Product) ClearEffects(){
    p.SafeProduct.ClearEffects()
}
func (p *SafeProduct) ClearEffects(){
    for i := range p.Effects {
        p.Effects[i] = StatefulEffect{}
        p.Effects[i].Reset()
    }
}

// Retrieve Effects as a set instead of a list
func (p *Product) EffectSet() map[string]string{
    effectSet := make(map[string]string)
    safeEffectSet := p.SafeProduct.EffectSet()
    for k, v := range safeEffectSet {
        if v.Name != NoEffect{
            effectSet[string(k)] = string(v.Name)
        }
    }
    return effectSet
}
func (p *SafeProduct) EffectSet() map[EffectName]EffectRef {
    effectSet := make(map[EffectName]EffectRef)
    for _, ref := range p.Effects {
        if ref.Current.Name.Valid() != NoEffect {
            effectSet[ref.Current.Name] = ref.Current
        }
    }
    return effectSet
}

func (p *Product) Initialize(baseIngredient string){
    p.SafeProduct = &SafeProduct{}
    p.SafeProduct.Initialize(BaseIngredientRef{BaseIngredientName(baseIngredient).Valid()}) 
}
func (p *SafeProduct) Initialize(baseIngredient BaseIngredientRef){
    p.Base = baseIngredient
    p.Multiplier = 1.0
    p.Price = baseIngredient.Lookup().Price
    p.MixQueue = []MixIngredientName{}
    p.MixHistory = []MixIngredientRef{}
    for i := range len(p.Effects){
        p.Effects[i] = StatefulEffect{}
        p.Effects[i].Reset()
    }
    for i, effect := range baseIngredient.Lookup().Effect {
        p.Effects[i].SetCurrentEffect( EffectName(effect).Valid() )
    }
}

func (p *Product) QueueIngredients(ingredients []string){
    safeIngredients := make([]MixIngredientName, len(ingredients))
    for i, ingredient := range ingredients {
        safeIngredients[i] = MixIngredientName(ingredient).Valid()
    }
    p.SafeProduct.QueueIngredients(safeIngredients)
}
func (p *SafeProduct) QueueIngredients(ingredients []MixIngredientName){
    p.MixQueue = append(p.MixQueue, ingredients...)
}


func (p *Product) AddEffect(newEffect string){
    p.SafeProduct.AddEffect(EffectName(newEffect).Valid())
}
func (p *SafeProduct) AddEffect(newEffect EffectName){
    for i, e := range p.Effects {
        if e.Current.Name == NoEffect {
            // avoid duplicates
            if e.Current.Name == newEffect{
                break
            }
            p.Effects[i].SetCurrentEffect(newEffect.Valid())
            break
        }
    }
    return
}

// Mixing mechanics:
//      Mutate all effefcts with the next ingredient. Prevent/avoid duplicates. Add new effect without duplicating it.
// Special case: 
//      If an effect were to mutates to an already existing effect, and that existing effect does not mutate, 
//      the former should not be mutated. If the latter mutates, then the former is allowed to mutate too
//      Unknown behaviour: 2 effects mutating into the same effect. However no such example exists in the table.
func (p *Product) MixNext(){
    p.SafeProduct.MixNext()
}   
func (p *SafeProduct) MixNext() {
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

    // Revert effects that are blocked. 
    // 1. Find duplicates of-post mutation.
    indexes := make(map[EffectName][]int)
    for i, e := range p.Effects {
        indexes[e.Current.Name] = append(indexes[e.Current.Name], i)
    }
    dupes := map[EffectName][]int{}
    for effect, indexList := range indexes {
        if effect != NoEffect && len(indexList) > 1 {
            dupes[effect] = indexList
        }
    }
    // 2. Find an effect that can be reverted
    // break if revert is successful. This assumes the "unknown behaviour" case is not possible
    for _, indexList := range dupes {
        dedupe: for _, i := range indexList {
            curr := p.Effects[i].Current.Name
            if p.Effects[i].Revert() {
                fmt.Printf("Reverted %s to %s\n",curr, p.Effects[i].Current.Name)
                break dedupe
            }
        }
    }
    // Finally, add the new effect
    p.AddEffect(nextEffectName)

    return
}

// Effects shouldnt actually be duplicated.
/*
func (p *Product) DedupeEffects(){
    p.SafeProduct.DedupeEffects()
}
func (p *SafeProduct) DedupeEffects(){
    effectCheck := make(map[EffectName]bool)
    for i, e := range p.Effects {
        _, exists := effectCheck[e.Name]
        if !exists {
            effectCheck[e.Name] = true
        } else {
            p.Effects[i] = EffectRef{NoEffect.Valid()}
        }
    }
}
*/

func (p *Product) UpdateMultiplier(){
    p.SafeProduct.UpdateMultiplier()
}
func (p *SafeProduct) UpdateMultiplier(){
    p.Multiplier = float64(1.0)
    for _, e := range p.Effects {
        if e.Current.Name != NoEffect {
            p.Multiplier += e.Current.Lookup().Multiplier
        }
    }
    return
}

func (p *Product) UpdatePrice(){
    p.SafeProduct.UpdatePrice()
}
func (p *SafeProduct) UpdatePrice(){
    basePrice := p.Base.Lookup().Price
    p.UpdateMultiplier()
    p.Price = int32( math.Round(float64(basePrice) * p.Multiplier) )
}

func (p *Product) MixAll(){
    p.SafeProduct.MixAll()
}
func (p *SafeProduct) MixAll(){
    for len(p.MixQueue) > 0 {
        p.MixNext()
    }
    return
}
