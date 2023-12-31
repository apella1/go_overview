package exercises

import (
	"fmt"
	"math/rand"
)

func PopulateIntSlice() {
	randomNums := []int{}
	for i := 0; i < 100; i++ {
		randNum := rand.Intn(100)
		randomNums = append(randomNums, randNum)
	}

	for _, num := range randomNums {
		switch {
		case num%2 == 0 && num%3 == 0:
			fmt.Println("Six.")
		case num%2 == 0:
			fmt.Println("Two.")
		case num%3 == 0:
			fmt.Println("Three.")
		default:
			fmt.Println("Never mind!")
		}
	}
}
