package main

import "fmt"

// write a method to replace all space in a string with %20
// example
// input: "Mr John Smith     "
// output: "Mr%20John%20Smith"

// Solution: start from the end. we have two scan approach
// first: count the number of space. so we can copute how many extra charachter
// we will have in the end
// second: in reverse order we edit the string and replace

func replaceSpaces(str string)string{
	lastSpaceCount := 0
	for i:=len(str) -1;i>0 ;i--{
		if str[i] == ' '{
			lastSpaceCount++
		}else{
			break
		}
	}

	str = str[:len(str) - lastSpaceCount]
	newStr := ""
	for i := 0;i<len(str);i++{
		if str[i] == ' '{
			newStr += "%20"
		} else{
			newStr += string(str[i])
		}
	}
	return newStr
}

func main(){
	fmt.Println(replaceSpaces("Hi Mr. Hamedani    "))
}