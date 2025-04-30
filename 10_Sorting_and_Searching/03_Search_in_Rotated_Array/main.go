package main

// Search in Rotated Array: Given a sorted array of n integers that has been rotated an unknown
// number of times, write code to find an element in the array. You may assume that the array was
// originally sorted in increasing order.
// EXAMPLE
// lnput:find5in{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}
// Output: 8 (the index of 5 in the array)

// SOLUTION
// If this problem smells like binary search to you, you're right!
// In classic binary search, we compare x with the midpoint to figure out if x belongs on the left or the right
// side. The complication here is that the array is rotated and may have an inflection point. Consider, for
// example, the following two arrays:
// Arrayl: {10, 15, 20, 0, 5}
// Array2: {50, 5, 20, 30, 40}

// Note that both arrays have a midpoint of 20, but5appears on the left side of one and on the right side of
// the other. Therefore, comparing x with the midpoint is insufficient.
// However, if we look a bit deeper, we can see that one half of the array must be ordered normally (in
// increasing order). We can therefore look at the normally ordered half to determine whether we should
// search the left or right half.
// For example, if we are searching for 5 in Array 1, we can look at the left element (1 O) and middle element
// (20). Since 10 < 20, the left half must be ordered normally. And, since 5 is not between those, we know that
// we must search the right half.

// In Array 2, we can see that since 50 > 20, the right half must be ordered normally. We turn to the middle
// (20) and right (40) element to check if 5 would fall between them. The value 5 would not; therefore, we
// search the left half.
// The tricky condition is if the left and the middle are identical, as in the example array { 2, 2, 2, 3, 4,
// 2}. In this case, we can check if the rightmost element is different. If it is, we can search just the right side.
// Otherwise, we have no choice but to search both halves.
func search(a []int, left, right, x int) int {
	mid := (left + right) / 2
	if x == a[mid] { //found element
		return mid
	}
	if right < left {
		return -1
	}

	// Either the left or right half must be normally ordered. Find out which side
	// is normally ordered, and then use the normally ordered half to figure out
	// which side to search to find x.
	if a[left] < a[mid] { // Left is normally ordered.
		if x > a[left] && x < a[mid] {
			return search(a, left, mid-1, x) //search left
		} else {
			return search(a, mid+1, right, x) //search right
		}
	}else if a[mid]<a[left]{ //Right is normally ordered.
		if x>a[mid]&&x<=a[right]{
			return search(a,mid+1,right,x) // search right
		}else{
			return search(a,left,mid-1,x)// search left
		}
	}else if a[left] == a[mid]{ // left or right half is all repeats
		if a[mid]!= a[right]{ //If right is different, search it
			return search(a,mid+1,right,x)//search right
		}else{ //Else, we have to search both halves
			result := search(a,left,mid-1,x) // search left
			if result == -1{
				return search(a,mid+1,right,x)//search right
			}else{
				return result
			}
		}
	}
	return -1
}

// This code will run in O(log n) if all the elements are unique. However, with many duplicates, the algo­
// rithm is actually O( n). This is because with many duplicates, we will often have to search both the left and
// right sides of the array (or subarrays).
// Note that while this problem is not conceptually very complex, it is actually very difficult to implement flaw­
// lessly. Don't feel bad if you had trouble implementing it without a few bugs. Because of the ease of making
// off-by-one and other minor errors, you should make sure to test your code very thoroughly.
