package main

// Missing Int: Given an input file with four billion non-negative integers, provide an algorithm to
// generate an integer that is not contained in the file. Assume you have 1 GB of memory available for
// this task.
// FOLLOW UP
// What if you have only 1O MB of memory? Assume that all the values are distinct and we now have
// no more than one billion non-negative integers.

// SOLUTION
// There are a total of 232, or 4 billion, distinct integers possible and 231 non-negative integers. Therefore, we
// know the input file (assuming it is ints rather than longs) contains some duplicates.
// We have 1 GB of memory, or 8 billion bits. Thus, with 8 billion bits, we can map all possible integers to a
// distinct bit with the available memory. The logic is as follows:
// 1. Create a bit vector (BV) with 4 billion bits. Recall that a bit vector is an array that compactly stores
// boolean values by using an array of ints (or another data type). Each int represents 32 boolean values.
// 2. Initialize BV with all Os.
// 3. Scan all numbers (num) from the file and call BV. set ( num, 1) .
// 4. Now scan again BV from the 0th index.
// 5. Return the first index which has a value of 0.
// The following code demonstrates our algorithm.

// to be continued at page 414
