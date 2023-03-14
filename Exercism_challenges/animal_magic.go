package chance
import (
    "math/rand"
    "time"
)

// SeedWithTime seeds math/rand with the current computer time.
func SeedWithTime() {
	// return rand.Seed(time.Now().UnixNano())
    rand.Seed(int64(time.Now().Nanosecond()))
}

// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	return rand.Intn(19) + 1
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	return rand.Float64()*12
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
	rand.Shuffle(8, func(i, j int) { animals[i], animals[j] = animals[j], animals[i] })
	return animals
}

