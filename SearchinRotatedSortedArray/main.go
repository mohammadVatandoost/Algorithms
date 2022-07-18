package main

import "fmt"

func main() {
	fmt.Println(search([]int{4,5,6,7,0,1,2}, 0))
	fmt.Println(search([]int{1,3,5}, 5))
}

func search(nums []int, target int) int {
    start := 0
    end := len(nums)
    for {
        if end - start == 1 {
           if nums[start] == target {
              return start 
           } 
           return -1 
        } else if end - start == 2 {
            if nums[start] == target {
              return start 
           } else if nums[start+1] == target {
              return start + 1 
           }
           return -1
        } 
        half := ((end-start)/2)+start
        if (nums[start] <= target &&  target <= nums[half]) ||
        (nums[start] <= target &&  target >= nums[half] && nums[half] < nums[start]) || 
		(nums[start] >= target &&  target <= nums[half] && nums[half] < nums[start])  {
            end = half+1	
			fmt.Printf("First Half start: %d, half: %d, end: %d \n", start, half, end)
        } else {
            start = ((end-start)/2) + start
			fmt.Printf("Second Half start: %d, half: %d, end: %d \n", start, half, end)
        }
    }
}