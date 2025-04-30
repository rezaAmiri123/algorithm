package main

import "fmt"

// ✅ 1. Basic Binary Search
// Problem: Given a sorted array nums and a target value,
// return the index if the target exists. Otherwise, return -1.

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <=right{
		mid:=left+(right-left)/2

		if nums[mid]==target{
			return mid
		}else if nums[mid]<target{
			left = mid+1
		}else{
			right = mid -1
		}
	}

	return -1
}

func main(){
	nums := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(binarySearch(nums,3))
	fmt.Println(binarySearch(nums,6))
}