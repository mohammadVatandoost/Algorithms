func twoSum(nums []int, target int) []int {
	mTable := make(map[int]int)
	var pairs []int
	for _, num := range nums {
		mTable[num] = mTable[num] + 1
	}
	for i, num := range nums {
		compliment := target - num
		if compliment == num && mTable[compliment] == 1 {
			continue
		}
		if mTable[compliment] > 0 {
			pairs = append(pairs, i)
		}
	}
	return pairs
}