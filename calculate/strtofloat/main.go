package strtofloat

import "strconv"

func StrToFloat(slice []string) ([]float64, error) {
	var arr []float64
	for i := 0; i < len(slice); i++ {
		num, err := strconv.ParseFloat(slice[i], 64)
		if err != nil {
			return []float64{}, err
		}
		arr = append(arr, num)
		num = 0
	}
	return arr, nil
}
