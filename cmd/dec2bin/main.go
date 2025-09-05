package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var errBadDec = errors.New("bad dec")

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Please enter a decimal number to convert (type q to quit)")

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

		inp := tokens[0]
		if inp == "q" {
			break
		}
		dec, err := validateDec(inp)
		if err != nil {
			if errors.Is(err, errBadDec) {
				fmt.Println("Please enter a valid decimal number")
				continue
			}
			log.Fatal(err)
		}

		fmt.Printf("Your binary number for %d is %s\n\n", dec, DecimalToBinary(dec))
	}
}

func DecimalToBinary(num uint64) string {
	if num == 0 {
		return "0"
	}
	var res []byte
	for num > 0 {
		fmt.Println(num)
		if num%2 == 0 {
			res = append(res, '0')
		} else {
			res = append(res, '1')
		}
		num /= 2
	}

	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}

	return string(res)
}

// Validates decimal number and returns the number.
func validateDec(dec string) (uint64, error) {
	num, err := strconv.ParseUint(dec, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("error during parsing dec: %w", errBadDec)
	}

	return num, nil
}
