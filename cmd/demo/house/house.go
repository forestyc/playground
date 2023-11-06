package main

import "github.com/forestyc/playground/cmd/demo/house/diff"

func main() {
	diff.FundMonthly(800000, []float64{0.031, 0.03575}, 360)
	diff.FundInterest(800000, []float64{0.031, 0.03575}, 132)
}
