package main

type Ingredient interface {
	Display() string
	Components() []Ingredient
	Time() int
}

type RawIngredient struct {
	Name string
}

func (ingredient *RawIngredient) Display() string {
	return ingredient.Name
}

func (ingredient *RawIngredient) Components() []Ingredient {
	return nil
}

func (ingredient *RawIngredient) Time() int {
	return 0
}
