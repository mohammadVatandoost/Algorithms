package main
// https://fizzbuzzed.com/top-interview-questions-1/

func main() {
	
}

func threeSum(nums []int) [][]int {
    output := make([][]int, 0)
    sort.Ints(nums)
    for i:=0; i<len(nums); i++ {
        if i != 0 && nums[i] == nums[i - 1] { continue }
        set := make(map[int]bool)
        for j := i+1; j<len(nums); j++ {
            tmp := -(nums[i]+nums[j])
            if set[tmp] {
                output = append(output, []int{nums[i], nums[j], tmp })
                for {
                    if j+1< len(nums) && nums[j+1] == nums[j] {
                        j++
                        continue
                    }
                    break
                }
                continue
            } 
            set[nums[j]] = true
        }  
    }   
    return output
}