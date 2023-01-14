package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scaner := bufio.NewScanner(os.Stdin)
	if scaner.Scan() {
		str := scaner.Text()
		fmt.Printf("%d", lengthOfLongestSubstring(str))
	}
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	var right, result int
	for left, v := range s {
		if n, ok := m[v]; ok {
			right = max(right, n+1)
		}
		result = max(result, left-right+1)
		m[v] = left
	}
	return result
}

func max(right, left int) int {
	if right > left {
		return right
	}
	return left
}
