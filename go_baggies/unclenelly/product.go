package unclenelly

import (
    "math"
    "fmt"
    "errors"
)

const ProductMaxEffects = 8

// Responsibility of validating effect/ingredient names is on wherever string is being converted
// NoEffect is a string const.

// StatefulEffect is a wrapper around EffectRef to keep track of the previous effect to allow reverting mutations
type StatefulEffect struct {
    Current     *EffectRef
    Previous    *EffectRef
}
func (e *StatefulEffect) Reset() error {
    if !NoEffect.Valid(){
        return fmt.Errorf("Invalid effect %s", NoEffect)
    }
    e.Current = &EffectRef{NoEffect}
    e.Previous = &EffectRef{NoEffect}   // Previous refers to value of previous step, including updating if no-change
    return nil
}
func (e *StatefulEffect) SetCurrentEffect(effect EffectName){
    if e.Current == (&EffectRef{}){
        e.Reset()
    }
    e.Current = &EffectRef{effect}
}
func (e *StatefulEffect) MutateWith(input EffectName){
    if e.Current == (&EffectRef{}) {
        e.Reset()
    }
    prev := e.Current.Name
    e.Current.MutateWith(input)
    e.Previous = &EffectRef{prev}
}

func (e *StatefulEffect) Revert() (ok bool){
    // Revert is a one-way operation.
    if e.Current.Name == e.Previous.Name {
        return false
    }
    e.Current = e.Previous
    return true
}

// Product is the version that works with string type
// SafeProduct uses custom types instead which needs to be converted from string
// This makes Product easier to use directly
type Product struct {
    SafeProduct *SafeProduct
}
type SafeProduct struct {
    Base        BaseIngredientRef
    Multiplier  float64
    Price       int32
    MixQueue    []MixIngredientName
    MixHistory  []MixIngredientRef
    Effects     [ProductMaxEffects]StatefulEffect
}

// Not an object. Just a struct to hold info queried off product
type ProductStatus struct {
    Base        string
    Multiplier  float64
    Price       int32
    MixQueue    []string
    MixHistory  []string
    Effects     []string
}

func NewProduct(baseIngredient string) (*Product, error){
    if baseIngredient == "" {
        return nil, errors.New("base ingredient cannot be empty. Perhaps use BlankBaseIngredient instead?")}
    if ! BaseIngredientName(baseIngredient).Valid(){
        return nil, fmt.Errorf("base ingredient %s is not valid", baseIngredient)
    }
    p := &Product{}
    p.Initialize(baseIngredient)
    return p, nil
}
// Preferred way to cook a product
func Cook(product *Product, mixIngredients []string) (*Product, error){
    for _, ingredient := range mixIngredients {
        if !MixIngredientName(ingredient).Valid(){
            return product, fmt.Errorf("Mix ingredient %s not valid", ingredient)
        }
    }
    // Clear queue and replace with new ingredients
    err := product.SetMixQueue(mixIngredients)
    if err != nil {
        return product, err
    }
    product.MixAll()
    product.UpdateMultiplier()
    product.UpdatePrice()
    return product, nil
}

// Product Getters
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
    effects := []string{}
    for key := range p.EffectSet(){
        effects = append(effects, key)
    }
    return effects
}
func (p *Product) Status() ProductStatus{
    return ProductStatus{
        Base:       p.Base(),
        Multiplier: p.Multiplier(),
        Price:      p.Price(),
        MixQueue:   p.MixQueue(),
        MixHistory: p.MixHistory(),
        Effects:    p.Effects(),
    }
}

// Product Setters
func (p *Product) SetBase(baseIngredient string) error {
    if !BaseIngredientName(baseIngredient).Valid() {
        return fmt.Errorf("Base ingredient %s not valid", baseIngredient)
    }
    p.SafeProduct.Base = BaseIngredientRef{BaseIngredientName(baseIngredient)}
    return nil
}
func (p *Product) SetMixQueue(ingredients []string) error {
    safeIngredients := make([]MixIngredientName, len(ingredients))
    for i, ingredient := range ingredients {
        if !MixIngredientName(ingredient).Valid() {
            return fmt.Errorf("Mix ingredient %s not valid", ingredient)
        }
        safeIngredients[i] = MixIngredientName(ingredient) 
    }
    p.SafeProduct.MixQueue = safeIngredients
    return nil
}

// === Method Pairs ===
// Each method described onwards exists in pairs of Product and SafeProduct versions.
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
        if ref.Current.Name != NoEffect {
            effectSet[ref.Current.Name] = *ref.Current
        }
    }
    return effectSet
}

func (p *Product) Initialize(baseIngredient string) error{
    p.SafeProduct = &SafeProduct{}
    baseIngredientRef := BaseIngredientRef{BaseIngredientName(baseIngredient)}
    if ! baseIngredientRef.Name.Valid(){
        return fmt.Errorf("Base ingredient %s is not valid", baseIngredient)
    }
    err := p.SafeProduct.Initialize(baseIngredientRef)
    if err != nil {
         return err
    }
    return nil
}
func (p *SafeProduct) Initialize(baseIngredient BaseIngredientRef) error {
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
        if !EffectName(effect).Valid(){
            return fmt.Errorf("Base ingredient %s has invalid effect %s", baseIngredient, effect)
        }
        p.Effects[i].SetCurrentEffect( EffectName(effect) )
    }
    return nil
}

func (p *Product) QueueIngredients(ingredients []string) error {
    safeIngredients := make([]MixIngredientName, len(ingredients))
    for i, ingredient := range ingredients {
        if !MixIngredientName(ingredient).Valid(){
            return fmt.Errorf("Mix ingredient %s not valid", ingredient)
        }
        safeIngredients[i] = MixIngredientName(ingredient)
    }
    p.SafeProduct.QueueIngredients(safeIngredients)
    return nil
}
func (p *SafeProduct) QueueIngredients(ingredients []MixIngredientName){
    p.MixQueue = append(p.MixQueue, ingredients...)
}


func (p *Product) AddEffect(newEffect string) error {
    if !EffectName(newEffect).Valid(){
        return fmt.Errorf("Effect %s not valid", newEffect)
    }
    p.SafeProduct.AddEffect(EffectName(newEffect))
    return nil
}
func (p *SafeProduct) AddEffect(newEffect EffectName){
    for i, e := range p.Effects {
        // find empty slot
        if e.Current.Name == NoEffect {
            // avoid duplicating
            if e.Current.Name == newEffect{
                break
            }
            p.Effects[i].SetCurrentEffect(newEffect)
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
    // 1. Find duplicates of post-mutation.
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
            if p.Effects[i].Revert() {
                break dedupe
            }
        }
    }
    // Finally, add the new effect
    p.AddEffect(nextEffectName)

    return
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


