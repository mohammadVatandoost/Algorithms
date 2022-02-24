package main

import (
	"fmt"
	"math"
)

func nonDivisibleSubset(k int32, s []int32) int32 {
	// Write your code here
	modes := make([]int32, k)
	for _, v := range s {
		modes[v%k] = modes[v%k] + 1
	}

	if k%2 == 0 {
		modes[k/2] = int32(math.Min(float64(modes[k/2]), 1))
	}

	sum := int32(math.Min(float64(modes[0]), 1))

	for i := int32(1); i<=k/2; i++ {
		sum = sum + int32(math.Max(float64(modes[i]), float64(modes[k-i])))
	}

	return sum
}

func main()  {
	fmt.Println(nonDivisibleSubset(4, []int32{19,10,12,10,24,25,22}) )
	fmt.Print(nonDivisibleSubset(7,
		[]int32{278, 576, 496, 727, 410, 124, 338, 149, 209, 702, 282, 718, 771, 575, 436}) )

}
