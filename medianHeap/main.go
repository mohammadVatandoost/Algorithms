package main

import "fmt"

type Node struct {
	Left *Node
	Right *Node
	Value float64
}

type MedianHeap struct {
	Elements []float64
}

func (MH *MedianHeap) Add(value float64)  {
	if len(MH.Elements) == 0 {
		MH.Elements = append(MH.Elements, value)
		return
	}

	index := int(len(MH.Elements)/2)
	for {
		if MH.Elements[index] == value {
			MH.insert(value, index)
			return
		} else if index == len(MH.Elements) - 1 {
			if MH.Elements[index] < value {
				MH.Elements = append(MH.Elements, value)
				return
			}
			MH.insert(value, index)
			return
		} else if index == 0 {
			if MH.Elements[index] > value {
				MH.Elements = append([]float64{value}, MH.Elements...)
				return
			}
			MH.insert(value, index)
			return
		} else if MH.Elements[index-1] < value && MH.Elements[index] > value {
			MH.insert(value, index)
			return
		} else if MH.Elements[index] > value {
			index = index/2
		} else {
			index = index + ((len(MH.Elements) - index)/2)
		}
	}
}

func (MH *MedianHeap) insert(value float64, index int)  {
	MH.Elements = append(MH.Elements[:index+1], MH.Elements[index:]...)
	MH.Elements[index] = value
}

func (MH *MedianHeap) Remove(value float64)  {
	index := findIndex(value, MH.Elements, 0, len(MH.Elements))
	MH.Elements = append(MH.Elements[:index], MH.Elements[index+1:]...)
}

func findIndex(value float64, Elements []float64, i int, j int) int {
	if len(Elements) == 0 {
		return 0
	}
	half := ((j-i)/2) + i
	if Elements[half] == value {
		return half
	} else if Elements[half] < value {
		return findIndex(value, Elements, half, j)
	}
	return findIndex(value, Elements, i, half)
}

func (MH *MedianHeap) Get() float64 {
	if len(MH.Elements) == 0 {
		return 0
	}
	if len(MH.Elements) == 1 {
		return MH.Elements[0]
	}
	if len(MH.Elements) % 2 == 1 {
		return MH.Elements[int(len(MH.Elements)/2)]
	}
	return (MH.Elements[len(MH.Elements)/2] + MH.Elements[len(MH.Elements)/2-1]) /2
}

func provideMedianHeap() *MedianHeap {
	return &MedianHeap{Elements: make([]float64,0)}
}

func activityNotifications(expenditure []int32, d int32) int32 {
	// Write your code here
	mh := provideMedianHeap()
	count := int32(0)
	for i, v := range expenditure {
		if len(mh.Elements) == int(d) {
			if int32(mh.Get()*2) <= v {
				count++
			}
			mh.Remove(float64(expenditure[i-int(d)]))
		}
		mh.Add(float64(v))
	}
	return count
}

func main()  {

	fmt.Println(activityNotifications([]int32{2, 3, 4, 2, 3, 6, 8, 4, 5}, 5))

 	//testData := []float64{2, 3, 4, 2, 3, 6, 8, 4, 5}
	//answer :=  []float64{2, 2.5, 3, 2.5, 3, 3, 3, 3.5, 4}
	//mh := provideMedianHeap()
	//for i,v:= range testData {
	//	mh.Add(v)
	//	if mh.Get() != answer[i] {
	//		fmt.Printf("Wrong Expected value: %v, result : %v \n", answer[i], mh.Get())
	//	}
	//}
	//for i:= len(testData)-1;  i>= 0; i-- {
	//	if mh.Get() != answer[i] {
	//		fmt.Printf("Remove Wrong Expected value: %v, result : %v \n", answer[i], mh.Get())
	//	}
	//	mh.Remove(testData[i])
	//}
	//
	//fmt.Println("Finished success")
}
