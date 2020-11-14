package main

import "fmt"

func main() {
	fmt.Println("Shepherd's pie is delicious, here is how to make it.")
	potatoes := RawIngredient{Name: "Potatoes"}
	peeledPotatoes := MechanicallyPreparedIngredient{Technique: "Peel", Name: "Peeled Potatoes", components: []Ingredient{&potatoes}, time: 2}
	butter := RawIngredient{Name: "Butter"}
	onions := RawIngredient{Name: "Onions"}
	veggies := RawIngredient{Name: "Veggies"}
	rawBeef := RawIngredient{Name: "Raw Ground Beef"}
	beefBroth := RawIngredient{Name: "Beef Broth"}
}
