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

func MechanicalTechnique(Technique string, ResultingIngredient string, SubIngredients []Ingredient, Time int) *MechanicallyPreparedIngredient {
	return &MechanicallyPreparedIngredient{
		Technique:  Technique,
		Name:       ResultingIngredient,
		components: SubIngredients,
		time:       Time,
	}
}

func conjoin(conj string, items []string) string {
	if len(items) == 0 {
		return ""
	}
	if len(items) == 1 {
		return items[0]
	}
	if len(items) == 2 { // "a and b" not "a, and b"
		return items[0] + " " + conj + " " + items[1]
	}

	sep := ", "
	pieces := []string{items[0]}
	for _, item := range items[1 : len(items)-1] {
		pieces = append(pieces, sep, item)
	}
	pieces = append(pieces, sep, conj, " ", items[len(items)-1])

	return strings.Join(pieces, "")
}

func (ingredient *MechanicallyPreparedIngredient) Display() string {

	componentIngredientNames := make([]string, len(ingredient.components))
	for i, v := range ingredient.components {
		componentIngredientNames[i] = v.IngredientReferenceName()
	}

	return fmt.Sprintf("%s the %s", ingredient.Technique, conjoin("and", componentIngredientNames))
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
