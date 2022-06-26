package standarddev

import (
	"fmt"
	"math"
)

func StandardDeviation(slice []float64, variance float64) {
	var deviation float64 = math.Sqrt(variance)
	fmt.Printf("Standard Deviation: %d\n", int(math.Round(deviation)))
}
