package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(twoSum([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	arr := make([]int, len(nums))
	copy(arr, nums)
	sort.Ints(nums)
	i := 0
	j := len(nums) - 1
	for i < j {
		if nums[i]+nums[j] == target {
			return []int{index(arr, nums[i]), index(arr, nums[j])}
		}
		if nums[i]+nums[j] < target {
			i++
		} else {
			j--
		}
	}
	return nums
}

func index(arr []int, item int) int {
	for k, v := range arr {
		if v == item {
			i := k
			arr[k] = v - 1 // for case [4, 4]
			return i
		}
	}
	return -1
}

// ---- OR ----

//func twoSum(nums []int, target int) []int {
//    m := make(map[int]int, len(nums))
//
//    for i, num := range nums {
//        if idx, ok := m[target - num]; ok {
//            return []int{idx, i}
//        }
//        m[num] = i
//    }
//    return []int{0, 0}
//}
