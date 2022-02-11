package main

import (
    "fmt"
	"math"
)

/*
 * Complete the 'formingMagicSquare' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY s as parameter.
 */

// func formingMagicSquare(s [][]int32) int32 {
//     // Write your code here
    
//     d := len(s)
//     magicNumber := (d*(d*d+1))-1
//     for i:=1; i<10; i++ {
        
//     }
// }


var allMatrixes [][]int

func isMagicalMatrix(matrix []int) bool {
	d := 3
	magicNumber := (d*(d*d+1))/2
	var temp [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			temp[i][j] = matrix[3 * i + j]
		}
	}

	// Checking if each row sum is same
	for i := 1; i < 3; i++ {
		tmp := 0
		for j := 0; j < 3; j++ {
			tmp += temp[i][j];
		}
		if tmp != magicNumber {
			return false
		}
	}

	// Checking if each column sum is same
	for  j := 0; j < 3; j++ {
		tmp := 0
		for i := 0; i < 3; i++ {
			tmp += temp[i][j]
		}
		if tmp != magicNumber {
			return false
		}
	}

	// Checking if diagonal 1 sum is same
	tmp := 0
	for i := 0; i < 3; i++ {
		tmp += temp[i][i];
	}
	if tmp != magicNumber {
		return false
	}

	// Checking if diagonal 2 sum is same
	tmp = 0
	for  i := 0; i < 3; i++ {
		tmp += temp[2 - i][i];
	}
	if tmp != magicNumber {
		return false
	}

	return true
}

func generateAllMatrixes(matrix []int, numbers []int) {
	if len(numbers) == 1 {
		matrix = append(matrix, numbers[0])
		if isMagicalMatrix(matrix) {
			allMatrixes = append(allMatrixes, matrix)
		}
		return
	}
	for i, v := range numbers {
		temp1 := make([]int, len(matrix))
		temp2 := make([]int, len(numbers))
		copy(temp1, matrix)
		copy(temp2, numbers)
		temp1 = append(temp1, v)
		temp2 = append(temp2[:i], temp2[i+1:]...)
		generateAllMatrixes(temp1, temp2)
	}
}

func calculateCost(matrix []int, magicalMatrix []int) int {
	sum := 0
	for i:=0; i<len(matrix); i++ {
		sum = sum + int(math.Abs(float64(matrix[i])-float64(magicalMatrix[i])))
	}
	return sum
}

func main() {
	generateAllMatrixes(make([]int,0), []int{1,2,3,4,5,6,7,8,9})
	matrix := []int{1,2,3,4,5,6,7,8,9}
	min := 1000
	for _, v := range allMatrixes {
		tmp := calculateCost(matrix,v)
		if tmp < min {
			min = tmp
		}
	}
	fmt.Println(min)
}

