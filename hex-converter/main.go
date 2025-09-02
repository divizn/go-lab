package main

import (
	"fmt"
	"math"
)

func main() {
	hex := "DCBA4255"

	fmt.Println(HexToDecimal(hex))

}

func HexToDecimal(num string) int {
	sum := 0

	hexMap := map[string]int{
		"F": 15,
		"E": 14,
		"D": 13,
		"C": 12,
		"B": 11,
		"A": 10,
		"9": 9,
		"8": 8,
		"7": 7,
		"6": 6,
		"5": 5,
		"4": 4,
		"3": 3,
		"2": 2,
		"1": 1,
	}

	for i := len(num) - 1; i >= 0; i-- {
		number := float64(hexMap[string(num[len(num)-i-1])])
		sum += int(number * math.Pow(16, float64(i)))
	}

	return sum
}
