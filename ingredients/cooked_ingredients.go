package ingredients

import (
	"fmt"
	"strings"
)

type CookedIngredient struct {
	Technique  string
	Name       string
	components []Ingredient
	time       int
	heat       string
	// indexOfIngredientInWhichTheRestOfTheIngredientsCook int
}

func CookingTechnique(Technique string, ResultingIngredient string, SubIngredients []Ingredient, Time int, Heat string) *CookedIngredient {
	return &CookedIngredient{
		Technique:  Technique,
		Name:       ResultingIngredient,
		components: SubIngredients,
		time:       Time,
		heat:       Heat,
	}
}
func (ingredient *CookedIngredient) Display() string {

	componentIngredientNames := make([]string, len(ingredient.components))
	for i, v := range ingredient.components {
		componentIngredientNames[i] = v.IngredientReferenceName()
	}
	// TODO exclude "in" ingredient, add it later as "in x" or something.

	return fmt.Sprintf("%s the %s, at %s heat, for %d minutes", ingredient.Technique, strings.Join(componentIngredientNames, ", "), ingredient.heat, ingredient.time)
}

func (ingredient *CookedIngredient) IngredientReferenceName() string {
	return ingredient.Name
}

func (ingredient *CookedIngredient) Components() []Ingredient {
	return ingredient.components
}

func (ingredient *CookedIngredient) Time() int {
	return ingredient.time
}
