func convert(s string, numRows int) string {
	res := make([][]string, numRows)
	i := 0
	isAscending := true
	for _, char := range s {
		res[i] = append(res[i], string(char))
		if isAscending {
			i = i + 1
		} else {
			i = i - 1
		}
		if i == numRows {
			isAscending = false
			i = numRows - 2
			if i < 0 {
				i = 0
			}
		} else if i == -1 {
			isAscending = true
			i = i + 2
			if i >= numRows {
				i = numRows - 1
			}
		}
	}
	out := ""
	for _, r := range res {
		out = out + strings.Join(r[:], "")
	}

	return out
}