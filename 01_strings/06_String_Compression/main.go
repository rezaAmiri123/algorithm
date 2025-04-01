package main

import "fmt"

// String Compression: Implement a method to perform basic string compression using the counts
// of repeated characters. For example, the string aabcccccaaa would become a2blc5a3. If the
// "compressed" string would not become smaller than the original string, your method should return
// the original string. You can assume the string has only uppercase and lowercase letters (a - z).

// SOLUTION
// At first glance, implementing this method seems fairly straightforward, but perhaps a bit tedious. We iterate
// through the string, copying characters to a new string and counting the repeats. At each iteration, check
// if the current character is the same as the next character. If not, add its compressed version to the result.
// How hard could it be?
func comparessBad(str string)string{
	comparessString :=""
	countConsecutive:=0
	for i:=0;i<len(str);i++{
		countConsecutive++

		// If next character is different than current, append this char to result.
		if i+1>=len(str)||str[i]!=str[i+1]{
			comparessString+=fmt.Sprintf("%s%d",string(str[i]),countConsecutive)
			countConsecutive = 0
		}
	}

	if len(comparessString)<len(str){
		return comparessString
	}else{
		return str
	}
}

// This works. ls it efficient, though?Take a look at the runtime of this code.
// The runtime is O(p + k2 ), where p is the size of the original string and k is thelnumber of character
// sequences. For example, if the string is aabccdeeaa, then there are six characte sequences. It's slow
// because string concatenation operates in O(n 2 ) time (see StringBuilder on pg 8 ).
// We can fix this by using a StringBuilder.


func main(){
	fmt.Println("comparessBad: ",comparessBad("aabcccccaaa"))
}