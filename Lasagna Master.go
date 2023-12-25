package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
    if time == 0 {
        time = 2
    }
	return len(layers) * time
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
    noodles = 0
    sauce = 0
	for _, v := range layers {
        switch {
            case v == "noodles":
        			noodles += 50
            case v == "sauce":
        			sauce += 0.2
        }
    }
    return
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList, myList []string) {
    myList[len(myList)-1] = friendList[len(friendList)-1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, numPort int) []float64 {
    ans := []float64{}	
    for _, v := range quantities {
        ans = append(ans, (v * float64(numPort)) / 2)
    }
	return ans
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
