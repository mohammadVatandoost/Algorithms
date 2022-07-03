package main

import (
	"fmt"
	"math"
	"unicode"
)

func main() {
	fmt.Printf("=== %d \n", myAtoi("   -42"))
	fmt.Printf("=== %d \n", myAtoi("   +0 123"))
	fmt.Printf("=== %d \n", myAtoi("20000000000000000000"))

}

func myAtoi(s string) int {
	numbers := ""
	sign := int64(1)
	isSigned := false
	result := int64(0)
	for _, char := range s {
		if char == ' ' && len(numbers) > 0 {
			break
		} else if char == ' ' {
			continue
		} else if len(numbers) == 0 && char == '-' {
			if isSigned {
				return 0
			}
			sign = -1
			isSigned = true
		} else if len(numbers) == 0 && char == '+' {
			if isSigned {
				return 0
			}
			sign = 1
			isSigned = true
		} else if unicode.IsNumber(char) {
			numbers = numbers + string(char)
		} else {
			break
		}
	}
	for i, char := range numbers {
		result = result + int64(char-'0')*int64(math.Pow(10, float64(len(numbers)-i-1)))
	}
	if result < 0 {
		return math.MaxInt32
	}
	if sign*result > math.MaxInt32 {
		return math.MaxInt32
	} else if sign*result < math.MinInt32 {
		return math.MinInt32
	}
	return int(sign * result)
}
