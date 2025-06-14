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

func fullyProcessOneProduct(tp *testProduct) Product {
    product, err := NewProduct(tp.BaseIngredient)
    if err != nil {
        panic(err)
    }
    product, err = Cook(product, tp.MixQueue)
    if err != nil {
        panic(err)
    }
    return *product
}

func mixingTest(t *testing.T, productsList *[]testProduct){
    for _, tp := range *productsList {
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
                if !EffectName(effect).Valid(){
                    panic(fmt.Sprintf("Effect %s is not valid", effect))
                }
                expectedEffects[effect] = effect
            }
            if !reflect.DeepEqual(expectedEffects, got.Effects) {
                t.Errorf("Wrong effects.\nExpected: \t%v\nGot :\t%v\n", expectedEffects, got.Effects)
            }
            fmt.Printf("Test %s + %s complete\n", tp.BaseIngredient, tp.MixQueue)
        })
    }
}


func TestNoIngredients(t *testing.T) {
    var productList = []testProduct{
        {
            BaseIngredient: "OG Kush",
            MixQueue: []string{},
            ExpectedResults: expectedResult{
                Effects: []string{"Calming"},
                Price: 39,
            },
        },
        {
            BaseIngredient: "Sour Diesel",
            MixQueue: []string{},
            ExpectedResults: expectedResult{
                Effects: []string{"Refreshing"},
                Price: 40,
            },
        },
        {
            BaseIngredient: "Green Crack",
            MixQueue: []string{},
            ExpectedResults: expectedResult{
                Effects: []string{"Energizing"},
                Price: 43,
            },
        },
        {
            BaseIngredient: "Granddaddy Purple",
            MixQueue: []string{},
            ExpectedResults: expectedResult{
                Effects: []string{"Sedating"},
                Price: 44,
            },
        },
        {
            BaseIngredient: "Meth",
            MixQueue: []string{},
            ExpectedResults: expectedResult{
                Effects: []string{},
                Price: 70,
            },
        },
        {
            BaseIngredient: "Cocaine",
            MixQueue: []string{},
            ExpectedResults: expectedResult{
                Effects: []string{},
                Price: 150,
            },
        },
    }
    mixingTest(t, &productList)
}

func TestSimpleMix(t *testing.T) {
    var productList = []testProduct{
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
            MixQueue: []string{"Mega Bean"},
            ExpectedResults: expectedResult{
                Effects: []string{"Glowing", "Foggy"},
                Price: 64,
            },
        },
        {
            BaseIngredient: "OG Kush",
            MixQueue: []string{"Mega Bean", "Paracetamol"},
            ExpectedResults: expectedResult{
                Effects: []string{"Calming", "Toxic", "Sneaky"},
                Price: 47,
            },
        },
    }
    mixingTest(t, &productList)
}

func TestBatonMix(t *testing.T) {
    // eg. {Blank + Paracetamol + Cuke + Paracetamol} -> { [Paranoia + Energizing] + Paracetamol } -> {Paranoia + Balding + Sneaky}.
    // This may look like Paranoia was not modified, when actually {Paranoia->Balding, Energizing->Paranoia, Sneaky added}
    var productList = []testProduct{
        {
            BaseIngredient: "Meth",
            MixQueue: []string{"Paracetamol"},
            ExpectedResults: expectedResult{
                Effects: []string{"Sneaky"},
                Price: 87,
            },
        },
        {
            BaseIngredient: "Meth",
            MixQueue: []string{"Paracetamol", "Cuke"},
            ExpectedResults: expectedResult{
                Effects: []string{"Paranoia", "Energizing"},
                Price: 85,
            },
        },
        {
            BaseIngredient: "Meth",
            MixQueue: []string{"Paracetamol", "Cuke", "Paracetamol"},
            ExpectedResults: expectedResult{
                Effects: []string{"Paranoia", "Sneaky", "Balding"},
                Price: 108,
            },
        },
    }
    mixingTest(t, &productList)
}

func TestBlockingMix(t *testing.T) {
    // {Sedating + Munchies} + Athletic -> {Munchies + Sedating + Athletic}
    // Sedating should've become Munchies, but because Munchies does not mutate, it blocks Sedating from mutating
    var productList = []testProduct{
        {
            BaseIngredient: "OG Kush",
            MixQueue: []string{"Energy Drink"},
            ExpectedResults: expectedResult{
                Effects: []string{"Calming", "Athletic"},
                Price: 50,
            },
        },
        {
            BaseIngredient: "OG Kush",
            MixQueue: []string{"Energy Drink", "Flu Medicine"},
            ExpectedResults: expectedResult{
                Effects: []string{"Munchies", "Sedating", "Bright-Eyed"},
                Price: 62,
            },
        },
        {
            BaseIngredient: "OG Kush",
            MixQueue: []string{"Energy Drink", "Flu Medicine", "Energy Drink"},
            ExpectedResults: expectedResult{
                Effects: []string{"Munchies", "Sedating", "Athletic", "Bright-Eyed"},
                Price: 74,
            },
        },
    }
    mixingTest(t, &productList)
}

func TestChainMixing(t *testing.T) {
    var productList = []testProduct{
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
            MixQueue: []string{"Horse Semen", "Addy", "Iodine", "Battery", "Chili", "Mega Bean", "Motor Oil", "Energy Drink"},
            ExpectedResults: expectedResult{
                Effects: []string{"Munchies", "Sedating", "Toxic", "Balding", "Slippery", "Euphoric", "Bright-Eyed", "Anti-gravity"},
                Price: 110,
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
    mixingTest(t, &productList)
}

