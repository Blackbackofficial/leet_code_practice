package main

import (
	"fmt"
	"math"
)

func main() {
	var x int
	_, err := fmt.Scan(&x)
	if err != nil {
		return
	}
	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}

	length := int(math.Log10(float64(x))) + 1
	arrNum := make([]byte, length)

	for k, _ := range arrNum {
		arrNum[k] = byte(x % 10)
		x = x / 10
	}

	length = length - 1
	for i := 0; i <= length/2; i++ {
		if !(arrNum[i] == arrNum[length-i]) {
			return false
		}
	}

	return true
}