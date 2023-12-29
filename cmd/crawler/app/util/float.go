package util

import "math"

// Round 四舍五入，precision为小数点位数
func Round(f float64, precision int) float64 {
	pow := math.Pow10(precision)
	if f < 0 {
		return math.Trunc((f+(-0.5)/pow)*pow) / pow
	} else {
		return math.Trunc((f+(0.5)/pow)*pow) / pow
	}
}

// Float64Equal 判断f1,f2是否相等，precision为小数点位数
func Float64Equal(f1, f2 float64, precision int) bool {
	diff := math.SmallestNonzeroFloat64
	if precision != -1 {
		diff = math.Pow10(-1 * precision)
	}
	if math.Abs(f1-f2) < diff {
		return true
	}
	return false
}

// InvalidFloat 非法浮点数
func InvalidFloat(f float64) bool {
	return Float64Equal(f, math.MaxFloat64, 5)
}
