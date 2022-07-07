func productExceptSelf(nums []int) []int {
    lenNums := len(nums)
    ascending := make([]int, lenNums)
    deAscending := make([]int, lenNums)
    ascending[0] = nums [0]
    deAscending[lenNums-1] = nums[lenNums-1]
    for i :=1 ; i< lenNums; i++ {
        ascending[i] = ascending[i-1]*nums[i]
        deAscending[lenNums-i-1] = deAscending[lenNums-i]*nums[lenNums-i-1]
    } 
    nums[0] = deAscending[1]
    nums[lenNums-1] = ascending[lenNums-2]
    for i :=1 ; i< lenNums-1; i++ {
        nums[i] = deAscending[i+1]*ascending[i-1]
    }
    return nums
}
