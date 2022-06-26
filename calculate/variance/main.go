package variance

import (
	"fmt"
	"math"
)

func Variance(slice []float64, avg float64) float64 {
	var variance float64
	total := 0.0

	for j := 0; j < len(slice); j++ {
		total += math.Pow(slice[j]-avg, 2)
	}
	variance = float64(total) / float64(len(slice))
	fmt.Printf("Variance: %d\n", int(math.Round(variance)))
	return variance
}
