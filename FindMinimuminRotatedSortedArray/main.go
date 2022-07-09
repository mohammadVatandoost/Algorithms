func findMin(nums []int) int {
	lenNums := len(nums)
	if lenNums == 1 {
		return nums[0]
	} else if lenNums == 2 {
		if nums[0] < nums[1] {
			return nums[0]
		}
		return nums[1]
	}

	if nums[lenNums-1] < nums[lenNums/2] {
		return findMin(nums[(lenNums / 2):])
	}
	return findMin(nums[:(lenNums/2)+1])
}