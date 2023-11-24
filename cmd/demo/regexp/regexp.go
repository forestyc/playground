package main

import (
	"fmt"
	"regexp"
)

func main() {
	contract := "ZC401C1010"

	fmt.Println(GetVarietyFromContract(contract))
}

// GetVarietyFromContract 截取合约id中的品种id
func GetVarietyFromContract(contract string) string {
	re := regexp.MustCompile(`(\d{1})(0[1-9]|1[0-2])`)
	matches := re.FindAllStringIndex(contract, -1)
	if len(matches) > 0 {
		if len(matches[0]) > 0 {
			return contract[0:matches[0][0]]
		}
	}
	return ""
}
