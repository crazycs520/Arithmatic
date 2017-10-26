package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var nums2 []int
	for i := 0; i < 1000000; i++ {
		nums2 = append(nums2, rand.Intn(100000))
	}
	start := time.Now()
	fmt.Println(twoSum1(nums2, 1333))
	fmt.Println(time.Since(start).Nanoseconds())

	start = time.Now()
	fmt.Println(twoSum2(nums2, 1333))
	fmt.Println(time.Since(start).Nanoseconds())

	start = time.Now()
	fmt.Println(twoSum3(nums2, 1333))
	fmt.Println(time.Since(start).Nanoseconds())
}

//Approach #1 (Brute Force) [Accepted]
func twoSum1(nums []int, target int) []int {

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}

//Two-pass Hash Table
func twoSum2(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, v := range nums {
		numMap[v] = i
	}
	for i := 0; i < len(nums); i++ {
		value := target - nums[i]
		if mv, ok := numMap[value]; ok {
			return []int{i, mv}
		}
	}
	return []int{0, 0}
}

//one-pass Hash Table
func twoSum3(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, v := range nums {
		numMap[v] = i
	}
	for i := 0; i < len(nums); i++ {
		value := target - nums[i]
		if mv, ok := numMap[value]; ok {
			return []int{i, mv}
		}
		numMap[nums[i]] = i
	}
	return []int{0, 0}
}
