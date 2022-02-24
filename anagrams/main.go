package main

import (
	"fmt"
	"math"
	"strconv"
)

func sherlockAndAnagrams(s string) int32 {
	// Write your code here
	set := make(map[string]bool)
	findAnagrams(0, len(s)-1, s, set)
	return int32(len(set)/2)
}

func findLocalMaxAnagrams(i int, j int, s string) (string, string) {
	tmp1 := ""
	tmp2 := ""
	min := int(math.Min(float64(len(s)-j-1), float64(i)))
	for t:=0; t <= min; t++ {
		if s[i-t] == s[j+t] {
			tmp1 = string(s[i-t]) + strconv.Itoa(i-t) + tmp1
			tmp2 = tmp2 + string(s[j+t]) + strconv.Itoa(j+t)
		} else {
			break
		}
	}
	return tmp1, tmp2
}

func findAnagrams(i int, j int, s string, set map[string]bool) {
	if i==j {
		tmp1, tmp2 := findLocalMaxAnagrams(i,j,s)
		if len(tmp1) > 4 && len(tmp1) != len(s)*2 {
			set[tmp1] = true
			set[tmp2] = true
		}
	} else if s[i] == s[j] {
		set[string(s[i])+strconv.Itoa(i)+string(s[j])+strconv.Itoa(j)] = true
		set[string(s[j])+strconv.Itoa(j)+string(s[i])+strconv.Itoa(i)] = true
		tmp1, tmp2 := findLocalMaxAnagrams(i,j,s)
		if len(tmp1) > 2  && len(tmp1) != len(s)*2 {
			set[tmp1] = true
			set[tmp2] = true
		}
	}
	if i == len(s)-1 && j == 0 {
		return
	} else if i == len(s)-1 {
		findAnagrams(i, j-1, s, set)
	} else if j == 0 {
		findAnagrams(i+1, j, s, set)
	} else {
		findAnagrams(i, j-1, s, set)
		findAnagrams(i+1, j, s, set)
		findAnagrams(i+1, j-1, s, set)
	}

}

func main()  {
	//fmt.Println(sherlockAndAnagrams("abba"))
	//fmt.Println(sherlockAndAnagrams("abcd"))
	fmt.Println(sherlockAndAnagrams("ifailuhkqq"))
	fmt.Println(sherlockAndAnagrams("kkkk"))
}
