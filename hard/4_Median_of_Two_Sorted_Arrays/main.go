package main

import (
	"fmt"
	"sort"
)

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2, 4}
	fmt.Print(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	median := float64(nums1[(len(nums1)-1)>>1]+nums1[len(nums1)>>1]) / 2.0
	return median
}
