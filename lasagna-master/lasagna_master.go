package lasagna

func PreparationTime(layers []string, prepTimePerLayer int) int {
	if prepTimePerLayer <= 0 {
		prepTimePerLayer = 2
	}

	return len(layers) * prepTimePerLayer
}


func Quantities(layers []string) (int, float64) {
	noodlesQuantity := 0
	waterQuantity := 0.0

	for _, layer := range layers {
		if layer == "noodles" {
			noodlesQuantity += 50
		}
		if layer == "sauce" {
			waterQuantity += 0.2
		}
	}
	
	return noodlesQuantity, waterQuantity
}

func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(quantities []float64, portionAmount int) []float64{
	amounts := make([]float64, len(quantities))

	for i, quantity := range quantities {
		amounts[i] = quantity * (float64(portionAmount) / 2.0)
	}

	return amounts
}
