package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var errBadHex = errors.New("bad hex")

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter a hex number to convert (type q to quit)")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		tokens := strings.Fields(line)

		if len(tokens) != 1 {
			fmt.Println("Please enter exactly one value")
			continue
		}

		hex := tokens[0]
		if hex == "q" {
			break
		}
		hex, err := validateHex(hex)
		if err != nil {
			if errors.Is(err, errBadHex) {
				fmt.Println("Please enter a valid hexadecimal number")
				continue
			}
			log.Fatal(err)
		}

		fmt.Printf("Your decimal number for %s is %d\n\n", hex, HexToDecimal(hex))
	}
}

// Converts hex number to decimal
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
		"0": 0,
	}

	for i := len(num) - 1; i >= 0; i-- {
		number := float64(hexMap[string(num[len(num)-i-1])])
		sum += int(number * math.Pow(16, float64(i)))
	}

	return sum
}

// Validates hex number and returns the parsed hex number.
func validateHex(hex string) (string, error) {
	hex = strings.ToUpper(hex)

	if _, err := strconv.ParseUint(hex, 16, 64); err != nil {
		return "", fmt.Errorf("error during parsing hex: %w", errBadHex)
	}

	return hex, nil
}
