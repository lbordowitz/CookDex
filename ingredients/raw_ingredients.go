package ingredients

import "fmt"

type Ingredient interface {
	Display() string
	IngredientReferenceName() string
	Components() []Ingredient
	Time() int
}

type RawIngredient struct {
	Name   string
	Amount string
}

func (ingredient *RawIngredient) Display() string {
	return fmt.Sprintf("%s %s", ingredient.Amount, ingredient.Name)
}

func (ingredient *RawIngredient) IngredientReferenceName() string {
	return ingredient.Name
}

func (ingredient *RawIngredient) Components() []Ingredient {
	return nil
}

func (ingredient *RawIngredient) Time() int {
	return 0
}
