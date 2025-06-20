package unclenelly

import (
    "gopkg.in/yaml.v3"
)

// Data stored as YAML string instead of a .yaml in order to hardcode and build it as wasm

var EffectsLookup = GetEffectsTable()
type EffectName string

func (e EffectName) Valid() bool {
    if _, ok := EffectsLookup[string(e)]; ok{
        return true
    }else{
        return false
    }
}

type EffectsYAML map[string]struct{
    Multiplier float64 `yaml:"Multiplier"`
    Conversion []map[string]string `yaml:"Conversion"`
}
type Effect struct {
    Name string
    Multiplier float64
    Conversion []map[string]string
}

// EffectRef is a reference to an effect by name, providing a lookup method
// EffectRef can be used as a stateful object that mutates
type EffectRef struct {
    Name EffectName
}
func (e EffectRef) Lookup() Effect {
    return EffectsLookup[string(e.Name)]
}
func (e *EffectRef) MutateWith(effect EffectName) {
   conversion := e.Lookup().Conversion
    for _, entry := range conversion {
        for name, result := range entry{
            if name == string(effect) {
                e.Name = EffectName(result)
                break
            }
        }
    }
}

type EffectsTable map[string]Effect

func GetEffectsTable() EffectsTable{
    effects_yaml := EffectsYAML{}
    err := yaml.Unmarshal([]byte(EffectsRawYAML), &effects_yaml)
    if err != nil {
        panic(err)
    }
    table := map[string]Effect{}
    for name, effect := range effects_yaml {
        table[name] = Effect{
            Name: name,
            Multiplier: effect.Multiplier,
            Conversion: effect.Conversion,
        }
    }
    return table
}

const NoEffect = EffectName("None") // To reference the None effect
var EffectsRawYAML = `
# Kush cooking effects table
# https://docs.google.com/spreadsheets/d/1Swo-SuDGqPy5hSvRVM-Moix8RjlqQkql0nl1_8CREUM/edit?usp=sharing
---
# Effects table
# <Ingredient>:
#   Multiplier: <number>
#   Conversion:
#     - <mixed with>: <result>
None:
  Multiplier: 0.00
  Conversion: []
Energizing:
  Multiplier: 0.22
  Conversion:
    - Gingeritis: Thought-Provoking
    - Sneaky:     Paranoia
    - Toxic:      Euphoric
    - Slippery:   Munchies
    - Foggy:      Cyclopean
Gingeritis:
  Multiplier: 0.20
  Conversion:
    - Energizing:   Thought-Provoking
    - Toxic:        Smelly
    - "Long faced": Refreshing
Sneaky:
  Multiplier: 0.24
  Conversion:
    - Energizing: Paranoia
    - Toxic:      "Tropic Thunder"
    - Foggy:      Calming
    - Spicy:      Bright-Eyed
Calorie-Dense:
  Multiplier: 0.28
  Conversion:
    - Calorie-Dense:  Explosive
    - Balding:        Sneaky
    - Jennerising:    Gingeritis
"Tropic Thunder":
  Multiplier: 0.46
  Conversion:
    - Athletic: Sneaky
Balding:
  Multiplier: 0.30
  Conversion:
    - Calorie-Dense: Sneaky
Sedating:
  Multiplier: 0.26
  Conversion:
    - Athletic:           Munchies
    - Thought-Provoking:  Gingeritis
Toxic:
  Multiplier: 0.00
  Conversion:
    - Energizing:   Euphoric
    - Gingeritis:   Smelly
    - Sneaky:       "Tropic Thunder"
    - Jennerising:  Sneaky
Athletic:
  Multiplier: 0.32
  Conversion:
    - "Tropic Thunder": Sneaky
    - Sedating:         Munchies
    - Foggy:            Laxative
    - Spicy:            Euphoric
Slippery:
  Multiplier: 0.34
  Conversion:
    - Energizing: Munchies
    - Foggy:      Toxic
Foggy:
  Multiplier: 0.36
  Conversion:
    - Energizing:         Cyclopean
    - Sneaky:             Calming
    - Athletic:           Laxative
    - Slippery:           Toxic
    - Jennerising:        Paranoia
    - Thought-Provoking:  Energizing
Spicy:
  Multiplier: 0.38
  Conversion:
    - Sneaky: Bright-Eyed
    - Athletic: Euphoric
Bright-Eyed:
  Multiplier: 0.40
  Conversion: []
Jennerising:
  Multiplier: 0.42
  Conversion:
    - Calorie-Dense:  Gingeritis
    - Toxic:          Sneaky
    - Foggy:          Paranoia
Thought-Provoking:
  Multiplier: 0.44
  Conversion:
    - Sedating:     Gingeritis
    - Foggy:        Energizing
    - "Long faced": Electrifying
"Long faced":
  Multiplier: 0.52
  Conversion:
    - Gingeritis:         Refreshing
    - Thought-Provoking:  Electrifying
Calming:
  Multiplier: 0.10
  Conversion:
    - Gingeritis:   Sneaky
    - Sneaky:       Slippery
    - Balding:      Anti-gravity
    - Sedating:     Bright-Eyed
    - Foggy:        Glowing
    - Jennerising:  Balding
Refreshing:
  Multiplier: 0.14
  Conversion:
    - Jennerising: Thought-Provoking
Munchies:
  Multiplier: 0.12
  Conversion:
    - Energizing:     Athletic
    - Sneaky:         Anti-gravity
    - Calorie-Dense:  Calming
    - Sedating:       Slippery
    - Toxic:          Sedating
    - Slippery:       Schizophrenic
    - Spicy:          Toxic
    - Bright-Eyed:    "Tropic Thunder"
Euphoric:
  Multiplier: 0.18
  Conversion:
    - Energizing:       Laxative
    - "Tropic Thunder": Bright-Eyed
    - Sedating:         Toxic
    - Toxic:            Spicy
    - Athletic:         Energizing
    - Slippery:         Sedating
    - Bright-Eyed:      Zombifying
    - Jennerising:      Seizure-Inducing
Paranoia:
  Multiplier: 0.00
  Conversion:
    - Gingeritis: Jennerising
    - Sneaky:     Balding
    - Toxic:      Calming
    - Slippery:   Anti-gravity
Anti-gravity:
  Multiplier: 0.54
  Conversion:
    - Calorie-Dense:  Slippery
    - Spicy:          "Tropic Thunder"
    - "Long faced":   Calming
Glowing:
  Multiplier: 0.48
  Conversion:
    - Sneaky:             Toxic
    - Athletic:           Disorienting
    - Thought-Provoking:  Refreshing
Disorienting:
  Multiplier: 0.00
  Conversion:
    - Gingeritis:       Focused
    - "Tropic Thunder": Toxic
    - Toxic:            Glowing
    - Athletic:         Electrifying
Electrifying:
  Multiplier: 0.50
  Conversion:
    - Sneaky:       Athletic
    - Sedating:     Refreshing
    - Toxic:        Disorienting
    - Bright-Eyed:  Euphoric
Laxative:
  Multiplier: 0.00
  Conversion:
    - "Tropic Thunder": Calming
    - Sedating:         Euphoric
    - Toxic:            Foggy
    - Spicy:            "Long faced"
    - Bright-Eyed:      Calorie-Dense
Schizophrenic:
  Multiplier: 0.00
  Conversion:
    - Athletic: Balding
Zombifying:
  Multiplier: 0.58
  Conversion: []
Seizure-Inducing:
  Multiplier: 0.00
  Conversion:
    - Foggy:        Focused
    - "Long faced": Energizing
Explosive:
  Multiplier: 0.00
  Conversion:
    - Balding:            Sedating
    - Thought-Provoking:  Euphoric
Cyclopean:
  Multiplier: 0.56
  Conversion:
    - Gingeritis:   Energizing
    - Bright-Eyed:  Glowing
Smelly:
  Multiplier: 0.00
  Conversion:
    - Gingeritis: Anti-gravity
Focused:
  Multiplier: 0.16
  Conversion:
    - Gingeritis:     Seizure-Inducing
    - Sneaky:         Gingeritis
    - Calorie-Dense:  Euphoric
    - Balding:        Jennerising
    - Sedating:       Calming
    - Athletic:       Shrinking
    - Foggy:          Disorienting
Shrinking:
  Multiplier: 0.60
  Conversion:
    - Calorie-Dense:    Energizing
    - "Tropic Thunder": Gingeritis
    - Sedating:         Paranoia
    - Toxic:            Focused
    - Foggy:            Electrifying
    - Spicy:            Refreshing
    - Bright-Eyed:      Munchies
`
