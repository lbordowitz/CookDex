package main

import (
	"fmt"

	"github.com/lbordowitz/CookDex/ingredients"
)

func displayIngredients(ingredient ingredients.Ingredient) {
	for _, v := range ingredient.Components() {
		displayIngredients(v)
	}
	if ingredient.Components() == nil {
		fmt.Println(ingredient.Display())
	}
}

func displayRecipe(ingredient ingredients.Ingredient) {
	if ingredient.Components() == nil {
		return
	}
	for _, v := range ingredient.Components() {
		displayRecipe(v)
	}
	fmt.Println(ingredient.Display())
}

func main() {
	fmt.Println("Shepherd's pie is delicious, here is how to make it.")

	boil := &ingredients.CookingTechnique{Technique: "Boil"}
	saute := &ingredients.CookingTechnique{Technique: "Saute"}
	simmer := &ingredients.CookingTechnique{Technique: "Simmer"}
	bake := &ingredients.CookingTechnique{Technique: "Bake"}
	// TODO: Really? "Cook"? Isn't there something more descriptive?
	cook := &ingredients.CookingTechnique{Technique: "Cook"}

	potatoes := ingredients.NewIngredient("potatoes", "3 large")
	// TODO: Should these two mechanical steps be combined? Even though, technically, they're separate?
	peeledPotatoes := ingredients.MechanicalTechnique("Peel", "peeled potatoes", []ingredients.Ingredient{potatoes}, 2)
	quarteredPotatoes := ingredients.MechanicalTechnique("Quarter", "quartered potatoes", []ingredients.Ingredient{peeledPotatoes}, 2)
	boiledPotatoes := boil.Cook("boiled potatoes", []ingredients.Ingredient{quarteredPotatoes}, 20, "high")
	butter := ingredients.NewIngredient("butter", "4 tbsp")
	mashedPotatoes := ingredients.MechanicalTechnique("Mash", "mashed potatoes", []ingredients.Ingredient{boiledPotatoes, butter}, 1)

	onions := ingredients.NewIngredient("onion", "1 medium")
	choppedOnions := ingredients.MechanicalTechnique("Chop", "chopped onions", []ingredients.Ingredient{onions}, 2)
	sautedOnions := saute.Cook("sauted onions", []ingredients.Ingredient{choppedOnions, butter}, 8, "medium")
	veggies := ingredients.NewIngredient("veggies", "1-2 cups")
	onionsAndVeggies := cook.Cook("onions and veggies", []ingredients.Ingredient{sautedOnions, veggies}, 8, "medium")
	rawBeef := ingredients.NewIngredient("raw ground beef", "1 1/2 pound")
	// TODO: "until beef is no longer pink" is the precise verbiage used. NOT an integer time in minutes!
	beefAndVegetables := cook.Cook("beef and veggies", []ingredients.Ingredient{rawBeef, onionsAndVeggies}, 4, "medium")
	seasonedBeef := ingredients.MechanicalTechnique("Season", "seasoned beef and veggies", []ingredients.Ingredient{beefAndVegetables}, 1)
	beefBroth := ingredients.NewIngredient("beef broth", "1/2 cup")
	shepherdsPieFilling := simmer.Cook("shepherds pie filling", []ingredients.Ingredient{seasonedBeef, beefBroth}, 10, "medium")

	uncookedShepherdsPie := ingredients.MechanicalTechnique("Layer", "unbaked shepherds pie", []ingredients.Ingredient{shepherdsPieFilling, mashedPotatoes}, 1)
	shepherdsPie := bake.Cook("Shepherds Pie", []ingredients.Ingredient{uncookedShepherdsPie}, 30, "400")

	fmt.Println("Ingredients:")
	displayIngredients(shepherdsPie)

	fmt.Println()
	fmt.Println("Recipe:")

	displayRecipe(shepherdsPie)
}
