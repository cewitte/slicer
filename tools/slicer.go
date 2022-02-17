package tools

import (
	"errors"
	"fmt"
	"math"
)

// Slicer splits a slicer into smaller chunks.
func Slicer(chunkSize int, slice []interface{}) ([]interface{}, error) {
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
