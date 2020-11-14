package main

import "fmt"

func main() {
	fmt.Println("Shepherd's pie is delicious, here is how to make it.")
	potatoes := RawIngredient{Name: "Potatoes"}
	peeledPotatoes := MechanicallyPreparedIngredient{Technique: "Peel", Name: "Peeled Potatoes", components: []Ingredient{&potatoes}, time: 2}
	quarteredPotatoes := MechanicallyPreparedIngredient{Technique: "Quarter", Name: "Quartered Potatoes", components: []Ingredient{&peeledPotatoes}, time: 2}
	// boil potatoes
	butter := RawIngredient{Name: "Butter"}
	// mash potatoes, with butter
	onions := RawIngredient{Name: "Onions"}
	choppedOnions := MechanicallyPreparedIngredient{Technique: "Chop", Name: "chopped Onions", components: []Ingredient{&onions}, time: 2}
	// saute onions
	veggies := RawIngredient{Name: "Veggies"}
	// cook veggies and onions
	rawBeef := RawIngredient{Name: "Raw Ground Beef"}
	// cook raw beef vegetables
	beefBroth := RawIngredient{Name: "Beef Broth"}
	// layer shepherds pie
	// bake shepherds pie
}
