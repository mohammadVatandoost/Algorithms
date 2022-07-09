func maxProduct(nums []int) int {
    max := math.MinInt32
    bufPositive := 1
    bufAll := 1
    for _, num := range nums {
        if num == 0 {
            bufPositive = 1
            bufAll = 1
            if max < 0 {
                max = 0
            }
        } else if num < 0 {
            bufAll = bufAll * num
            bufPositive = 1
            if max < bufAll {
                max = bufAll
                bufPositive = bufAll
                bufAll = num
            }
        } else {
            bufAll = bufAll * num
            bufPositive = bufPositive * num
            if max < bufPositive {
                max = bufPositive
            }
        }     
    }
    return max
}