package unclenelly 

import (
    "fmt"
    "testing"
    "reflect"
)

type gotResult struct {
    Effects     map[string]string
    Price       int32
    MixHistory  []string
    MixQueue    []string
}

type expectedResult struct {
    Effects     []string
    Price       int32
}

type testProduct struct {
    BaseIngredient  string
    MixQueue        []string
    ExpectedResults expectedResult
}

var testingMixProducts = []testProduct{
    {
        // OG Kush + Cuke + Mega Bean -> [Glowing, Cyclopean, Foggy]
        BaseIngredient: "OG Kush", 
        MixQueue: []string{"Cuke", "Mega Bean"}, 
        ExpectedResults: expectedResult{
            Effects: []string{"Glowing", "Cyclopean", "Foggy"},
            Price: 84,
        },
    },
    {
        BaseIngredient: "OG Kush", 
        MixQueue: []string{"Horse Semen"}, 
        ExpectedResults: expectedResult{
            Effects: []string{"Calming", "Long faced"},
            Price: 57,
        },
    },      
    {
        BaseIngredient: "OG Kush", 
        MixQueue: []string{"Horse Semen", "Addy"}, 
        ExpectedResults: expectedResult{
            Effects: []string{"Calming", "Thought-Provoking", "Electrifying"},
            Price: 71,
        },
    },
    {
        BaseIngredient: "OG Kush", 
        MixQueue: []string{"Horse Semen", "Addy", "Iodine"}, 
        ExpectedResults: expectedResult{
            Effects: []string{"Balding", "Thought-Provoking", "Jennerising", "Electrifying"},
            Price: 93,
        },
    },
    {
        BaseIngredient: "OG Kush", 
        MixQueue: []string{"Horse Semen", "Addy", "Iodine", "Battery"}, 
        ExpectedResults: expectedResult{
            Effects: []string{"Euphoric","Balding", "Jennerising", "Thought-Provoking", "Bright-Eyed"},
            Price: 96,
        },
    },
    {
        BaseIngredient: "OG Kush", 
        MixQueue: []string{"Horse Semen", "Addy", "Iodine", "Battery", "Chili"}, 
        ExpectedResults: expectedResult{
            Effects: []string{"Euphoric","Balding", "Spicy", "Bright-Eyed", "Jennerising", "Thought-Provoking"},
            Price: 109,
        },
    },
    {
        BaseIngredient: "OG Kush", 
        MixQueue: []string{"Horse Semen", "Addy", "Iodine", "Battery", "Chili", "Mega Bean"}, 
        ExpectedResults: expectedResult{
            Effects: []string{"Euphoric","Balding", "Spicy", "Bright-Eyed", "Paranoia", "Energizing", "Foggy"},
            Price: 99,
        },
    },
    {
        BaseIngredient: "OG Kush", 
        MixQueue: []string{"Horse Semen", "Addy", "Iodine", "Battery", "Chili", "Mega Bean", "Motor Oil"}, 
        ExpectedResults: expectedResult{
            Effects: []string{"Munchies", "Sedating", "Toxic", "Balding", "Slippery", "Spicy", "Bright-Eyed", "Anti-gravity"},
            Price: 117,
        },
    },
    {
        BaseIngredient: "OG Kush", 
        MixQueue: []string{"Horse Semen", "Addy", "Iodine", "Battery", "Chili", "Mega Bean", "Motor Oil", "Energy Drink", "Viagor"}, // No change from Viagor
        ExpectedResults: expectedResult{
            Effects: []string{"Munchies", "Sedating", "Toxic", "Balding", "Slippery", "Euphoric", "Bright-Eyed", "Anti-gravity"},
            Price: 110,
        },
    },
    {
        BaseIngredient: "OG Kush", 
        MixQueue: []string{"Horse Semen", "Addy", "Iodine", "Battery", "Chili", "Mega Bean", "Motor Oil", "Energy Drink", "Viagor", "Mouth Wash"}, // No change from Mouth Wash
        ExpectedResults: expectedResult{
            Effects: []string{"Munchies", "Sedating", "Toxic", "Balding", "Slippery", "Euphoric", "Bright-Eyed", "Anti-gravity"},
            Price: 110,
        },
    },
    {
        BaseIngredient: "OG Kush", 
        MixQueue: []string{"Horse Semen", "Addy", "Iodine", "Battery", "Chili", "Mega Bean", "Motor Oil", "Energy Drink",
            "Viagor", "Mouth Wash", "Cuke"},
        ExpectedResults: expectedResult{
            Effects: []string{"Euphoric", "Munchies", "Sedating", "Laxative", "Balding", "Athletic", "Bright-Eyed", "Anti-gravity"},
            Price: 109,
        },
    },
    {
        BaseIngredient: "OG Kush", 
        MixQueue: []string{"Horse Semen", "Addy", "Iodine", "Battery", "Chili", "Mega Bean", "Motor Oil", "Energy Drink",
            "Viagor", "Mouth Wash", "Cuke", "Donut"},
        ExpectedResults: expectedResult{
            Effects: []string{"Calming", "Euphoric", "Sedating", "Sneaky", "Laxative", "Athletic", "Slippery", "Bright-Eyed"},
            Price: 99,
        },
    },
    {
        BaseIngredient: "OG Kush", 
        MixQueue: []string{"Horse Semen", "Addy", "Iodine", "Battery", "Chili", "Mega Bean", "Motor Oil", "Energy Drink",
            "Viagor", "Mouth Wash", "Cuke", "Donut", "Cuke"},
        ExpectedResults: expectedResult{
            Effects: []string{"Calming", "Euphoric", "Paranoia", "Munchies", "Sedating", "Laxative", "Athletic", "Bright-Eyed"},
            Price: 83,
        },
    },
    {
        BaseIngredient: "OG Kush", 
        MixQueue: []string{"Horse Semen", "Addy", "Iodine", "Battery", "Chili", "Mega Bean", "Motor Oil", "Energy Drink",
            "Viagor", "Mouth Wash", "Cuke", "Donut", "Cuke", "Motor Oil"},
        ExpectedResults: expectedResult{
            Effects: []string{"Calming", "Euphoric", "Sedating", "Laxative", "Athletic", "Schizophrenic", "Bright-Eyed", "Anti-gravity"}, 
            Price: 98,
        },
    },
}

func fullyProcessOneProduct(tp *testProduct) Product {
    product := Product{}
    product.Initialize(tp.BaseIngredient)
    product.SetMixQueue(tp.MixQueue)
    product.MixAll()
    product.UpdatePrice()
    return product
}

func TestMixing(t *testing.T) {
    for _, tp := range testingMixProducts {
        t.Run(fmt.Sprintf("%s + %s", tp.BaseIngredient, tp.MixQueue), func(t *testing.T) {
            product := fullyProcessOneProduct(&tp)
            got := gotResult{
                Effects: product.EffectSet(),
                Price: product.Price(),
                MixQueue: product.MixQueue(),
                MixHistory: product.MixHistory(),
            }
            if len(product.MixQueue()) != 0 {
                t.Errorf("Expected MixQueue to be empty, got %v", got.MixQueue)
            }
            expectedHistory := []string{}
            for _, ingredient := range tp.MixQueue {
                expectedHistory = append(expectedHistory, ingredient)
            }
            if len(product.MixHistory()) != len(expectedHistory){
                t.Errorf("Expected %d effects in history, got %d", len(expectedHistory), len(got.MixHistory))
            }
            history_ok := true
            for i, eff := range got.MixHistory {
                if eff != expectedHistory[i]{
                    history_ok = false
                    break
                }
            }
            if !history_ok {
                t.Errorf("Expected history %v, got %v", expectedHistory, got.MixHistory)
            }
            if got.Price != tp.ExpectedResults.Price {
                t.Errorf("Expected price %v, got %v", tp.ExpectedResults.Price, got.Price)
            }
            if len(got.Effects) != len(tp.ExpectedResults.Effects) {
                t.Errorf("Expected %d effects, got %d", len(tp.ExpectedResults.Effects), len(got.Effects))
            }
            expectedEffects := make(map[string]string)
            for _, effect := range tp.ExpectedResults.Effects {
                expectedEffects[effect] = string( EffectName(effect).Valid() )
            }
            if !reflect.DeepEqual(expectedEffects, got.Effects) {
                t.Errorf("Wrong effects.\nExpected: \t%v\nGot :\t%v\n", expectedEffects, got.Effects)
            }
        })
    }
}

