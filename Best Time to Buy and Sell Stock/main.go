func maxProfit(prices []int) int {
    minTemp := prices[0]
    max_profit := 0
    for _, p := range prices {
        if p < minTemp {
            minTemp = p
        } else if (p - minTemp) > max_profit {
            max_profit = p - minTemp
        }  
    }
  
    return max_profit;
}
