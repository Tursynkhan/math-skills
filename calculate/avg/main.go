package avg

import (
	"fmt"
	"math"
)

func Average(slice []float64) float64 {
	var avg float64
	var sum float64 = 0
	for i := 0; i < len(slice); i++ {
		sum += slice[i]
	}
	avg = float64(sum) / float64(len(slice))
	fmt.Printf("Average: %d\n", int(math.Round(avg)))
	return avg
}
