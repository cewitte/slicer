package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
)

func slicer(chunkSize int, slice []interface{}) ([]interface{}, error) {
	if len(slice) == 0 || chunkSize > len(slice) {
		return nil, errors.New("nothing to be done here, check your chunk and content sizes")
	}

	numOfSlices := int(math.Ceil(float64(len(slice)) / float64(chunkSize)))

	var slices []interface{}
	begin, end := 0, chunkSize

	for i := 0; i < numOfSlices; i++ {
		slices = append(slices, slice[begin:end])
		fmt.Printf("i = %d, begin = %d, end = %d\n", i, begin, end)
		if end+chunkSize > len(slice) {
			begin, end = end, len(slice)
		} else {
			begin, end = end, end+chunkSize
		}
	}

	return slices, nil
}

func main() {
	// Test case with small slice
	slice := []interface{}{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	slices, err := slicer(3, slice)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Final result returned:", slices)

	// Test case 2 with large slice containing 150.000 entries
	var lgSlice []interface{}

	for i := 0; i <= 150000; i++ {
		lgSlice = append(lgSlice, rand.Intn(150000))
	}

	slices, err = slicer(1000, lgSlice)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Final result returned:", slices)

}
