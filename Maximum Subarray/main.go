func maxSubArray(nums []int) int {
	max := math.MinInt32
	sum := math.MinInt32
	for _, num := range nums {
		if sum+num < num {
			sum = num
		} else {
			sum = sum + num
		}
		if max < sum {
			max = sum
		}
	}

	return max
}