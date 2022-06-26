package median

import (
	"fmt"
	"math"
	"sort"
)

func Median(slice []float64) {
	sort.Float64s(slice)
	var median float64
	l := len(slice)
	if l%2 == 0 {
		median = float64((slice[(l-1)/2])+(slice[l/2])) / 2.0
	} else {
		median = float64(slice[l/2])
	}
	fmt.Printf("Median: %d\n", int(math.Round(median)))
}
