package utils

import "math"

func EqualFloat(x, y, epsilon float64) bool {
	diff := math.Abs(x - y)
	return diff < epsilon
}

func Round(number float64, decimalPlaces int) float64 {
	return math.Round(number*math.Pow(10, float64(decimalPlaces))) / math.Pow(10, float64(decimalPlaces))
}
