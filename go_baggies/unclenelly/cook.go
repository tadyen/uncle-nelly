package unclenelly

import "fmt"

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
    product.UpdateCost()
    return product, nil
}
