package lasagna

// PreparationTime returns the preparation time for the layers
func PreparationTime(layers []string, avg int) int {
	if avg == 0 {
		avg = 2
	}

	return len(layers) * avg
}

// Quantities returns the sauce and noodles nedded for the layers
func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauce float64
	for _, kind := range layers {
		if kind == "sauce" {
			sauce += .2
		}

		if kind == "noodles" {
			noodles += 50
		}
	}

	return noodles, sauce
}

// AddSecretIngredient adds the last ingredient from friendsList into myList
func AddSecretIngredient(friendsList []string, myList []string) {
	secret := friendsList[len(friendsList)-1]
	myList[len(myList)-1] = secret
}

// ScaleRecipe returns a slice of the quantities scaled up by the portions multiplier 
func ScaleRecipe(quantities []float64, portions int) []float64 {
	var scaled []float64
	for _, qt := range quantities {
		scaled = append(scaled, qt/2*float64(portions))
	}

	return scaled
}