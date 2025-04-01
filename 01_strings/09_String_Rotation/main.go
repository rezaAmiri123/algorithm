package main

import (
	"fmt"
	"strings"
)

// String Rotation: Assume you have a method i5Su b5tring which checks if one word is a substring
// of another. Given two strings, 51 and 52, write code to check if 52 is a rotation of 51 using only one
// call to i5Sub5tring (e.g., "waterbottle" is a rotation of" erbottlewat").

// SOLUTION
// If we imagine that 52 is a rotation of 51, then we can ask what the rotation point is. For example, if you
// rotate waterbottle after wat. you get erbottlewat. In a rotation, we cut 51 into two parts, x and y,
// and rearrange them to get 52.
// s1= xy = waterbottle
// x = wat
// y = erbottle
// s2 = yx = erbottlewat

// So, we need to check if there's a way to split s1 into x and y such that xy = s1 andyx = s2. Regardless of
// where the division between x andy is, we can see thatyx will always be a substring of xyxy.That is, s2 will
// always be a substring of s1s1.
// And this is precisely how we solve the problem: simply do isSubstring(slsl, s2).
// The code below implements this algorithm.
func isRotation(str1,str2 string)bool{
	// Check that sl and s2 are equal length and not empty
	if len(str1) != len(str2) || str1 == ""{
		return false
	}
	str1str1 := str1 + str1
	return strings.Contains(str1str1, str2)
}

func main(){
	fmt.Println(isRotation("waterbottle","erbottlewat"))
	fmt.Println(isRotation("waterbotttle","erbottlewat"))
}

// The runtime of this varies based on the runtime of isSubstring. But if you assume that isSubstring
// runs inO(A+B) time (on strings of length A and B), then the runtime of is Rotation isO(N).
