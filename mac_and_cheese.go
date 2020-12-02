package main

import (
	"fmt"

	"github.com/lbordowitz/CookDex/ingredients"
)

func macAndCheese() {

	fmt.Println("Mac and Cheese is delicious, here is how to make it.")

	boil := &ingredients.CookingTechnique{Technique: "Boil"}
	bake := &ingredients.CookingTechnique{Technique: "Bake"}

	dryPasta := ingredients.NewIngredient("dry pasta", "1 package")
	cookedPasta := boil.Cook("boiled potatoes", []ingredients.Ingredient{dryPasta}, 10, "high")

	cheese := ingredients.NewIngredient("cheese", "8 oz")
	shreddedCheese := ingredients.MechanicalTechnique("Shred", "shredded cheese", []ingredients.Ingredient{cheese}, 1)

	mixedMac := ingredients.MechanicalTechnique("Mix", "mac and cheese", []ingredients.Ingredient{cookedPasta, shreddedCheese}, 1)
	unbakedMac := ingredients.MechanicalTechnique("Sprinkle", "unbaked mac", []ingredients.Ingredient{mixedMac, shreddedCheese}, 1)

	macAndCheese := bake.Cook("Mac and Cheese", []ingredients.Ingredient{unbakedMac}, 15, "350")

	fmt.Println("Ingredients:")
	displayIngredients(macAndCheese)

	fmt.Println()
	fmt.Println("Recipe:")

	displayRecipe(macAndCheese)
}
