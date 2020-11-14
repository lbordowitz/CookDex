package main

import (
	"fmt"

	"github.com/lbordowitz/CookDex/ingredients"
)

func main() {
	fmt.Println("Shepherd's pie is delicious, here is how to make it.")

	boil := &ingredients.CookingTechnique{Technique: "Boil"}
	saute := &ingredients.CookingTechnique{Technique: "Saute"}
	simmer := &ingredients.CookingTechnique{Technique: "Simmer"}
	bake := &ingredients.CookingTechnique{Technique: "Bake"}
	// TODO: Really? "Cook"? Isn't there something more descriptive?
	cook := &ingredients.CookingTechnique{Technique: "Cook"}

	potatoes := &ingredients.RawIngredient{Name: "Potatoes"}
	peeledPotatoes := ingredients.MechanicalTechnique("Peel", "Peeled Potatoes", []ingredients.Ingredient{potatoes}, 2)
	quarteredPotatoes := ingredients.MechanicalTechnique("Quarter", "Quartered Potatoes", []ingredients.Ingredient{peeledPotatoes}, 2)
	boiledPotatoes := boil.Cook("Boiled Potatoes", []ingredients.Ingredient{quarteredPotatoes}, 20, "high")
	butter := &ingredients.RawIngredient{Name: "Butter"}
	mashedPotatoes := ingredients.MechanicalTechnique("Mash", "MashedPotatoes", []ingredients.Ingredient{boiledPotatoes, butter}, 1)

	onions := &ingredients.RawIngredient{Name: "Onions"}
	choppedOnions := ingredients.MechanicalTechnique("Chop", "chopped Onions", []ingredients.Ingredient{onions}, 2)
	sautedOnions := saute.Cook("Sauted Onions", []ingredients.Ingredient{choppedOnions, butter}, 8, "medium")
	veggies := &ingredients.RawIngredient{Name: "Veggies"}
	onionsAndVeggies := cook.Cook("Onions and Veggies", []ingredients.Ingredient{sautedOnions, veggies}, 8, "medium")
	rawBeef := &ingredients.RawIngredient{Name: "Raw Ground Beef"}
	// "until beef is no longer pink" is the precise verbiage used. NOT an integer time in minutes!
	beefAndVegetables := cook.Cook("Beef and Veggies", []ingredients.Ingredient{rawBeef, onionsAndVeggies}, 4, "medium")
	seasonedBeef := ingredients.MechanicalTechnique("Season", "Seasoned Beef and Veggies", []ingredients.Ingredient{beefAndVegetables}, 1)
	beefBroth := &ingredients.RawIngredient{Name: "Beef Broth"}
	shepherdsPieFilling := simmer.Cook("Shepherds Pie Filling", []ingredients.Ingredient{seasonedBeef, beefBroth}, 10, "medium")

	uncookedShepherdsPie := ingredients.MechanicalTechnique("Layer", "Uncooked Shepherds Pie", []ingredients.Ingredient{shepherdsPieFilling, mashedPotatoes}, 1)
	shepherdsPie := bake.Cook("Shepherds Pie", []ingredients.Ingredient{uncookedShepherdsPie}, 30, "400")

	fmt.Println(shepherdsPie.Display())
}
