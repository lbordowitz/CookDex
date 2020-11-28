package ingredients

import "fmt"

// Ingredient has the basic functions an ingredient needs for being interacted with.
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

func NewIngredient(name string, amount string) *RawIngredient {
	return &RawIngredient{Name: name, Amount: amount}
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
