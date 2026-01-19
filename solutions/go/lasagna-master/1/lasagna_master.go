package lasagna

func PreparationTime(layers []string, avgTime int) int {
	if avgTime == 0 {
		avgTime = 2
	}
	return len(layers) * avgTime
}

func Quantities(layers []string) (gNoodles int, lSauce float64) {
	for _, layer := range layers {
		if layer == "noodles" {
			gNoodles += 50
		} else if layer == "sauce" {
			lSauce += 0.2
		}
	}
	return
}

func AddSecretIngredient(friendIngs, myIngs []string) {
	myIngs[len(myIngs)-1] = friendIngs[len(friendIngs)-1]
}

func ScaleRecipe(ingsForTwo []float64, numPortions int) []float64 {
	scale := float64(numPortions) / 2.0
	ingsNeeded := make([]float64, len(ingsForTwo))
	for i, amt := range ingsForTwo {
		ingsNeeded[i] = amt * scale
	}
	return ingsNeeded
}
