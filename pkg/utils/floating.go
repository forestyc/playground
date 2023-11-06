package utils

import "math"

func EqualFloat(x, y, epsilon float64) bool {
	diff := math.Abs(x - y)
	return diff < epsilon
}
