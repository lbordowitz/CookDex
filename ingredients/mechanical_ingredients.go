package ingredients

import (
	"fmt"
	"strings"
)

type MechanicallyPreparedIngredient struct {
	Technique  string
	Name       string
	components []Ingredient
	time       int
}

func (ingredient *MechanicallyPreparedIngredient) Display() string {

	componentIngredientNames := make([]string, len(ingredient.components))
	for i, v := range ingredient.components {
		componentIngredientNames[i] = v.IngredientReferenceName()
	}

	return fmt.Sprintf("{} the {}", ingredient.Technique, strings.Join(componentIngredientNames, ", "))
}

func (ingredient *MechanicallyPreparedIngredient) IngredientReferenceName() string {
	return ingredient.Name
}

func (ingredient *MechanicallyPreparedIngredient) Components() []Ingredient {
	return ingredient.components
}

func (ingredient *MechanicallyPreparedIngredient) Time() int {
	return ingredient.time
}