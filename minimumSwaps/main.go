package main

import (
	"fmt"
)

type Element struct {
	value int32
	index int
}

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	count := int32(0)
	indexsArr := make([]int, len(arr))
	//var elements []Element
	var elements []int32
	for i, v := range arr {
		if int(v) == i + 1 {
			indexsArr[v-1] = i
			continue
		}
		if i+1 == int(arr[v-1]) {
			arr [i] = arr[v-1]
			arr[v-1] = v
			count++
			indexsArr[v-1] = i
			continue
		}
		indexsArr[v-1] = i
		elements = append(elements, v)
	}
	return minimumSwapsRecursive(elements, 0, indexsArr) + count
}

func minimumSwapsRecursive(arr []int32, index int, indexsArr []int) int32 {
	n := len(arr)
	min := int32(len(arr)*len(arr))
	//if index == n {
	//	return 0
	//}
	for i:= 0; i<n; i++ {
		v := arr[i]
		arrTemp := make([]int32, n)
		indexsTemp := make([]int, n)
		copy(arrTemp, arr)
		copy(indexsTemp, indexsArr)
		if indexsArr[int(v)-1] == int(v)-1 {
			arrTemp = append(arrTemp[:i], arrTemp[i:]...)
			tmp := minimumSwapsRecursive(arrTemp, i+1, indexsTemp)+1
			if tmp < min {
				min = tmp
			}
			continue
		}
		//arrTemp[i] = arrTemp[v-1]
		indexsTemp[v-1] = indexsTemp[v-1]
		//arrTemp[v-1] = v
		arrTemp = append(arrTemp[:v-1], arrTemp[v:]...)
		indexsTemp[v-1] = int(v-1)
        tmp := minimumSwapsRecursive(arrTemp, i+1, indexsTemp)+1
		if tmp < min {
			min = tmp
		}
	}
	return min
}


func main()  {
	//fmt.Println(minimumSwaps([]int32{7, 1, 3, 2, 4, 5, 6}) )
	fmt.Println(minimumSwaps([]int32{4, 3, 1, 2}) )
}
