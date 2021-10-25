package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averagePrepTime int) int {
	if averagePrepTime == 0 {
		averagePrepTime = 2
	}
	return len(layers) * averagePrepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	noodles = 0
	sauce = float64(0)

	for _, e := range layers {
		switch e {
		case "sauce":
			sauce += .2
		case "noodles":
			noodles += 50
		}
	}

	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, myList []string) (comboList []string) {
	comboList = make([]string, len(myList)+1)
	for i := range comboList {
		if i < len(myList) {
			comboList[i] = myList[i]
		} else {
			comboList[i] = friendsList[len(friendsList)-1]
		}
	}

	return
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) (scaled []float64) {
	scaled = make([]float64, len(quantities))
	scaleFactor := float64(portions) / 2
	for i := range quantities {
		scaled[i] = quantities[i] * scaleFactor
	}

	return
}
