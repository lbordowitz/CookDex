package main

import (
	"fmt"

	"github.com/lbordowitz/CookDex/ingredients"
)

func main() {
	fmt.Println("Shepherd's pie is delicious, here is how to make it.")
	potatoes := ingredients.RawIngredient{Name: "Potatoes"}
	peeledPotatoes := ingredients.MechanicallyPreparedIngredient{Technique: "Peel", Name: "Peeled Potatoes", components: []ingredients.Ingredient{&potatoes}, time: 2}
	quarteredPotatoes := ingredients.MechanicallyPreparedIngredient{Technique: "Quarter", Name: "Quartered Potatoes", components: []ingredients.Ingredient{&peeledPotatoes}, time: 2}
	// boil potatoes
	butter := ingredients.RawIngredient{Name: "Butter"}
	// mash potatoes, with butter
	onions := ingredients.RawIngredient{Name: "Onions"}
	choppedOnions := ingredients.MechanicallyPreparedIngredient{Technique: "Chop", Name: "chopped Onions", components: []ingredients.Ingredient{&onions}, time: 2}
	// saute onions
	veggies := ingredients.RawIngredient{Name: "Veggies"}
	// cook veggies and onions
	rawBeef := ingredients.RawIngredient{Name: "Raw Ground Beef"}
	// cook raw beef vegetables
	beefBroth := ingredients.RawIngredient{Name: "Beef Broth"}
	// layer shepherds pie
	// bake shepherds pie
}
