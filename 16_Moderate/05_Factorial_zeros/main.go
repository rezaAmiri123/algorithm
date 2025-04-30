package main

// Factorial zeros: Write an algorithm which computes the number of trailing zeros in n factorial.

// SOLUTION
// A simple approach is to compute the factorial, and then count the number of trailing zeros by continu­
// ously dividing by ten. The problem with this though is that the bounds of an int would be exceeded very
// quickly. To avoid this issue, we can look at this problem mathematically.
// Consider a factorial like 19 ! :
// 19! = 1*2*3*4* 5*6*7*8*9*10 *11*12*13*14*15*16*17*18*19
// A trailing zero is created with multiples of 10, and multiples of 10 are created with pairs of 5-multiples and
// 2-multiples.
// For example, in 19!, the following terms create the trailing zeros:
// 19! = 2 * ... * 5 * ... * 10 * ... * 15 * 16 * ...
// Therefore, to count the number of zeros, we only need to count the pairs of multiples of 5 and 2. There will
// always be more multiples of 2 than 5, though, so simply counting the number of multiples of 5 is sufficient.
// One "gotcha" here is 15 contributes a multiple of 5 (and therefore one trailing zero), while 25 contributes
// two (because 25 = 5 * 5).
// There are two different ways to write this code.
// The first way is to iterate through all the numbers from 2 through n, counting the number of times that 5
// goes into each number.

/* If the number is a 5 of five, return which power of 5. For example: 5 -> 1,
* 25-> 2, etc. */
func factorsOf5(i int) int {
	count := 0
	for i%5 == 0 {
		count++
		i /= 5
	}
	return count
}

func countFactZeros(num int) int {
	count := 0
	for i := 0; i <= num; i++ {
		count += factorsOf5(i)
	}
	return count
}

// This isn't bad, but we can make it a little more efficient by directly counting the factors of 5. Using this
// approach, we would first count the number of multiples of 5 between 1 and n (which is Ys ), then the
// number of multiples of 25 (/'is), then 125, and so on
// To count how many multiples of mare in n, we can just divide n by m.
func countFactZerosImproved(num int) int {
	count := 0
	if num < 0 {
		return -1
	}
	for i := 5; num/i > 0; i *= 5 {
		count += num / i
	}
	return count
}

// This problem is a bit of a brainteaser, but it can be approached logically (as shown above). By thinking
// through what exactly will contribute a zero, you can come up with a solution. You should be very clear in
// your rules upfront so that you can implement it correctly.
