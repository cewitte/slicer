package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/cewitte/slicer/tools"
)

func main() {
	// Test case with small slice
	slice := []interface{}{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	slices, err := tools.Slicer(3, slice)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Final result returned:", slices)

	// Test case 2 with large slice containing 150.000 entries
	var lgSlice []interface{}

	for i := 0; i <= 150000; i++ {
		lgSlice = append(lgSlice, rand.Intn(150000))
	}

	slices, err = tools.Slicer(1000, lgSlice)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Final result returned:", slices)
}
