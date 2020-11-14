package main

import (
	"fmt"

	"github.com/lbordowitz/CookDex/ingredients"
)

func main() {
	fmt.Println("Shepherd's pie is delicious, here is how to make it.")
	potatoes := &ingredients.RawIngredient{Name: "Potatoes"}
	peeledPotatoes := ingredients.MechanicalTechnique("Peel", "Peeled Potatoes", []ingredients.Ingredient{potatoes}, 2)
	quarteredPotatoes := ingredients.MechanicalTechnique("Quarter", "Quartered Potatoes", []ingredients.Ingredient{peeledPotatoes}, 2)
	// boil potatoes
	butter := &ingredients.RawIngredient{Name: "Butter"}
	// mash potatoes, with butter
	onions := &ingredients.RawIngredient{Name: "Onions"}
	choppedOnions := ingredients.MechanicalTechnique("Chop", "chopped Onions", []ingredients.Ingredient{onions}, 2)
	// saute onions
	veggies := ingredients.RawIngredient{Name: "Veggies"}
	// cook veggies and onions
	rawBeef := ingredients.RawIngredient{Name: "Raw Ground Beef"}
	// cook raw beef vegetables
	beefBroth := ingredients.RawIngredient{Name: "Beef Broth"}
	// layer shepherds pie
	// bake shepherds pie
}
