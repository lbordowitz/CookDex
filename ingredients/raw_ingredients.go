package ingredients

type Ingredient interface {
	Display() string
	IngredientReferenceName() string
	Components() []Ingredient
	Time() int
}

type RawIngredient struct {
	Name string
}

func (ingredient *RawIngredient) Display() string {
	return ingredient.Name
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
