package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string,avg_time int) int{
    switch{
        case avg_time > 0: return int(len(layers) * avg_time)
        default: return int(len(layers) * 2) 
    }
    
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int,float64){
    num_noodle := 0
    num_sauce := 0
  
    for _,ele := range layers{
        
        if ele == "noodles"{
        num_noodle++
    }
    
      if ele == "sauce" { 
          num_sauce++
      } 
    
    }
    return int(50 * num_noodle) , 0.2 * float64( num_sauce)
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string,myList []string){
    myList = append(myList[:len(myList)-1],friendsList[len(friendsList)-1])
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64,amounts int) []float64{
    
	var scaledAmounts []float64
    
	for _, v := range quantities {
		scaledAmounts = append(scaledAmounts, v * float64(amounts)/2.0)
	}
	return scaledAmounts
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
