package main

import "fmt"

// ✅ 2. First and Last Position of Element in Sorted Array
// Problem: Return the starting and ending position of a given target in a sorted array.
// If not found, return [-1, -1].

func searchRange(nums []int, target int) []int {
    return []int{findBound(nums, target, true), findBound(nums, target, false)}
}

func findBound(nums []int, target int, isFirst bool) int {
    left, right := 0, len(nums)-1
    result := -1

    for left <= right {
        mid := left + (right-left)/2

        if nums[mid] == target {
            result = mid
            if isFirst {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return result
}

func main(){
	nums := []int{1,2,3,4,5,36,7,8,9}
	fmt.Println(searchRange(nums,4))
}