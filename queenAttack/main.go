package main

import (
	"fmt"
	"math"
)

func queensAttack(n int32, k int32, r_q int32, c_q int32, obstacles [][]int32) int32 {
	// Write your code here
	attacks := make([]int32, 8)
	if n - c_q > 0 {
		attacks[0] = n - c_q
	}
	if n - c_q > 0 && n - r_q > 0  {
		attacks[1] = int32(math.Min(float64(n - c_q), float64(n - r_q)))
	}
	if n - r_q > 0 {
		attacks[2] = n - r_q
	}
	if c_q - 1 > 0 && n - r_q > 0 {
		attacks[3] = int32(math.Min(float64(c_q - 1), float64(n - r_q)))
	}
	if c_q - 1 > 0 {
		attacks[4] = c_q - 1
	}
	if c_q - 1 > 0 && r_q - 1 > 0  {
		attacks[5] = int32(math.Min(float64(c_q - 1), float64(r_q - 1)))
	}
	if r_q - 1 > 0 {
		attacks[6] = r_q - 1
	}
	if n - c_q > 0 && r_q - 1 > 0 {
		attacks[7] = int32(math.Min(float64(n - c_q), float64(r_q - 1)))
	}

	for i:= int32(0); i<k; i++ {
		obstacle := obstacles[i]
		if obstacle[0] - c_q > 0 && obstacle[1] - r_q == 0 {
			attacks[0] = obstacle[0] - c_q - 1
		} else if obstacle[0] - c_q > 0 && obstacle[1] - r_q > 0 &&
			(obstacle[0] - c_q)/(obstacle[1] - r_q) == 1  {
			attacks[1] = int32(math.Min(float64(obstacle[0] - c_q), float64(obstacle[1] - r_q))) - 1
		} else if obstacle[0] - c_q == 0 && obstacle[1] - r_q > 0 {
			attacks[2] = obstacle[1] - r_q - 1
		} else if c_q - obstacle[0] > 0 && obstacle[1] - r_q > 0 &&
			(c_q - obstacle[0])/(obstacle[1] - r_q) == 1 {
			attacks[3] = int32(math.Min(float64(c_q - obstacle[0]), float64(obstacle[1] - r_q))) - 1
		} else if c_q - obstacle[0] > 0 && obstacle[1] - r_q == 0 {
			attacks[4] = c_q - obstacle[0] - 1
		} else if c_q - obstacle[0] > 0 && r_q - obstacle[1] > 0 &&
			(c_q - obstacle[0])/(r_q - obstacle[1]) == 1  {
			attacks[5] = int32(math.Min(float64(c_q - obstacle[0]), float64(r_q - obstacle[1]))) - 1
		} else if c_q - obstacle[0] == 0 && r_q - obstacle[1] > 0 {
			attacks[6] = r_q - obstacle[1] - 1
		} else if obstacle[0] - c_q > 0 && r_q - obstacle[1] > 0 &&
			(obstacle[0] - c_q)/(r_q - obstacle[1]) == 1 {
			attacks[7] = int32(math.Min(float64(obstacle[0] - c_q), float64(r_q - obstacle[1]))) - 1
		}
	}
	sum := int32(0)
	for i:=0; i<8; i++ {
		sum = sum + attacks[i]
	}
	return sum
}

func main()  {
	fmt.Println(queensAttack(5, 3, 3, 4, [][]int32{{5,5},{4,2},{2,3}}))
	//fmt.Println(queensAttack(5, 8, 3, 3, [][]int32{{1,1},{1,1},{1,1},{1,1},
	//	{1,1},{1,1},{1,1},{1,1}}))
}

