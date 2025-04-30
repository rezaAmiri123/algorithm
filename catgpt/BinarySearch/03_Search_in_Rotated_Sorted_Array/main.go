package main

// ✅ 3. Search in Rotated Sorted Array
// Problem: Array is rotated at some unknown pivot. Search for the target.

func searchRotated(nums []int, target int) int {
	left,right := 0, len(nums)-1
	for left<=right{
		mid := left+(right-left)/2

		if nums[mid]==target{
			return mid
		}

		// Left side is sorted
		if nums[left]<=nums[mid]{
			if nums[left]<=target&&target <=nums[mid]{
				right = mid-1
			}else{
				left = mid+1
			}
		}else { // Right side is sorted
			if nums[mid]<target&&target<=nums[right]{
				left=mid+1
			}else{
				right = mid-1
			}
		}
	}
	return -1
}