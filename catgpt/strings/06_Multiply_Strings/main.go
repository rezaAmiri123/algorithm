package main

import (
	"fmt"
	"strings"
)

// 🔹 Problem: Multiply Strings
// Description:
// Given two non-negative integers num1 and num2 represented as strings,
// return the product as a string.
// You must not use built-in big integer libraries or convert the inputs directly to integers.

// 🔸 Example
// Input:  num1 := "123", num2 := "456"
// Output: "56088"
// 🔹 Explanation
// This simulates how we do multiplication by hand:
// Multiply each digit of num1 with each digit of num2.
// Store intermediate results in an array (to simulate digit positions).
// Carry over values > 10.
// Skip leading zeros in the final result.
// For two numbers of lengths m and n, the maximum length of the result is m + n.

func multiply(num1, num2 string) string {
    if num1 == "0" || num2 == "0" {
        return "0"
    }

    m, n := len(num1), len(num2)
    result := make([]int, m+n)

    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            mul := (int(num1[i]-'0') * int(num2[j]-'0'))
            p1, p2 := i+j, i+j+1
            sum := mul + result[p2]

            result[p2] = sum % 10
            result[p1] += sum / 10
        }
    }

    var sb strings.Builder
    for _, digit := range result {
        if sb.Len() == 0 && digit == 0 {
            continue
        }
        sb.WriteByte(byte(digit + '0'))
    }

    return sb.String()
}

func main() {
    num1 := "123"
    num2 := "456"
    fmt.Println("Product:", multiply(num1, num2)) // Output: "56088"
}

// ✅ Why It’s Medium Level
// Requires low-level digit manipulation without converting to integers.

// Tests your ability to simulate algorithms you normally take for granted.

// Careful indexing and carry-over management.

// 🔹 Goal
// Multiply two large non-negative integers represented as strings 
// (e.g., "123" × "456") without using built-in integer conversion or big integer libraries.

// 🔸 Step-by-Step Explanation
// ✅ Step 1: Understand digit-by-digit multiplication
// We multiply numbers the way we do on paper:
//       123
//    ×  456
//   --------
//      738   ← 123 × 6 (ones place)
//     615    ← 123 × 5 (tens place, shift one left)
//    492     ← 123 × 4 (hundreds place, shift two left)
//   --------
//    56088   ← Final sum
// Each digit multiplication affects a different position in the result array.

// ✅ Step 2: Initialize result array
// If num1 has m digits and num2 has n digits, 
// then the maximum result length is m + n (e.g., 3 digits × 3 digits = max 6 digits).

// So we make an array result := make([]int, m+n) to store digits of the product.

// ✅ Step 3: Multiply digits in nested loops
// We go right-to-left (like manual multiplication):
// for i := m-1; i >= 0; i-- {
//     for j := n-1; j >= 0; j-- {
//         mul := (num1[i] - '0') * (num2[j] - '0')
//         p1 := i + j       // carry position
//         p2 := i + j + 1   // unit position

//         sum := mul + result[p2]
//         result[p2] = sum % 10
//         result[p1] += sum / 10
//     }
// }
// Here:

// p2 stores the current digit.

// p1 stores the carry (e.g., if 18 → 8 goes to p2, 1 carried to p1).

// ✅ Step 4: Convert result array to string
// Skip leading zeros (unless the number is actually zero).

// Convert digits in the result array into a string using a strings.Builder.

// for _, digit := range result {
//     if sb.Len() == 0 && digit == 0 {
//         continue // skip leading zeros
//     }
//     sb.WriteByte(byte(digit + '0'))
// }
// ✅ Step 5: Handle edge case for zero
// If either number is "0", return "0" immediately.

// 🧠 Example: Multiply "12" × "34"

//     12
// ×   34
// -----
//    48  ← 2×4
//   30   ← 2×3 (shift left)
//   40   ← 1×4 (shift left)
//  30    ← 1×3 (shift more left)
// -------
//  408
// Breakdown by position:

// 2 × 4 = 8 → position 2

// 2 × 3 = 6 → position 1 (add to existing)

// 1 × 4 = 4 → position 1 and 0

// 1 × 3 = 3 → position 0

// Result: [0, 4, 0, 8] → "408"

