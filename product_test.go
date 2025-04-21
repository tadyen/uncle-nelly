package main

import (
    "fmt"
    "testing"
    "reflect"
)

type productResult struct {
    Multiplier          float64
    ResultEffectSet     map[EffectName]EffectRef
}

type testProduct struct {
    BaseIngredient  BaseIngredientName
    MixQueue        []MixIngredientName
    ExpectedResults productResult
}

var testingMixProducts = []testProduct{
    {
        // OG Kush + Cuke + Mega Bean -> [Glowing, Cyclopean, Foggy]
        BaseIngredient: BaseIngredientName("OG Kush"), 
        MixQueue: []MixIngredientName{"Cuke", "Mega Bean"}, 
        ExpectedResults: productResult{
            ResultEffectSet: map[EffectName]EffectRef{
                "Glowing": {"Glowing"},
                "Cyclopean": {"Cyclopean"},
                "Foggy": {"Foggy"},
            },
            Multiplier: 1.0 + 0.48 + 0.56 + 0.36,
        },
    },
}

func fullyProcessOneProduct(tp *testProduct) Product {
    product := Product{}
    product.Initialize(BaseIngredientRef{tp.BaseIngredient})
    product.MixQueue = tp.MixQueue
    product.MixAll()
    product.UpdateMultiplier()
    return product
}

func TestMixing(t *testing.T) {
    for _, tp := range testingMixProducts {
        t.Run(fmt.Sprintf("%s + %s", tp.BaseIngredient, tp.MixQueue), func(t *testing.T) {
            product := fullyProcessOneProduct(&tp)
            got := productResult{
                Multiplier: product.Multiplier,
                ResultEffectSet: EffectSet(product.Effects[:]),
            }
            if len(product.MixQueue) != 0 {
                t.Errorf("Expected MixQueue to be empty, got %v", product.MixQueue)
            }
            expectedHistory := []MixIngredientRef{}
            for _, ingredient := range tp.MixQueue {
                expectedHistory = append(expectedHistory, MixIngredientRef{ingredient})
            }
            if len(product.MixHistory) != len(expectedHistory){
                t.Errorf("Expected %d effects in history, got %d", len(expectedHistory), len(product.MixHistory))
            }
            history_ok := true
            for i, eff := range product.MixHistory {
                if eff != expectedHistory[i]{
                    history_ok = false
                    break
                }
            }
            if !history_ok {
                t.Errorf("Expected history %v, got %v", expectedHistory, product.MixHistory)
            }
            if got.Multiplier != tp.ExpectedResults.Multiplier {
                t.Errorf("Expected multiplier %v, got %v", tp.ExpectedResults.Multiplier, got.Multiplier)
            }
            if len(got.ResultEffectSet) != len(tp.ExpectedResults.ResultEffectSet) {
                t.Errorf("Expected %d effects, got %d", len(tp.ExpectedResults.ResultEffectSet), len(got.ResultEffectSet))
            }
            if !reflect.DeepEqual(tp.ExpectedResults.ResultEffectSet, got.ResultEffectSet) {
                t.Errorf("Expected effects %v, got %v", tp.ExpectedResults.ResultEffectSet, got.ResultEffectSet)
            }
        })
    }
}

