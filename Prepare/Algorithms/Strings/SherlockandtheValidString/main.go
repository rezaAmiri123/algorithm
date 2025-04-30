package main

// medium

import (
	"bufio"
	"fmt"
	"io"
	"slices"
	"strings"
)

/*
 * Complete the 'isValid' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func isValid(s string) string {
	array := []rune(s)
	slices.Sort(array)
	fmt.Println(array)
	base := 0
	for i := 0; i < len(array); i++ {
		count := 0
		for j := i; j < len(array); j++ {
			if array[i] != array[j] {
				if base == 0 {
					base = count
					break
				}
				if count+1 == base || count == base || count-1 == base {
					break
				}else{
					return "NO"
				}
			}
		}
	}
	return "YES"
}
func isValid2(s string) string {
	array := []rune(s)
	slices.Sort(array)
	fmt.Println(array)
	baseCount := 0
	distance :=1
	index1 :=0
	index2 :=1
	for i:=0;i<len(array)&&index2<len(array);i++{
		fmt.Println("i:",i, "index2: ",index2)
		if array[index1] == array[index2]{
			index2++
			fmt.Println("distance: ",distance, "baseCount: ", baseCount)
			if distance -1 >baseCount{
				return "NO"
			}	
		}else{
			if baseCount == 0{
				baseCount = distance
				fmt.Println("baseCount: ", baseCount)
			}else{
				fmt.Printf("distance: %d, baseCount:%d\n",distance,baseCount)
				diff := distance - baseCount
				if diff<0 {
					diff *=-1
				}
				if diff>1{
					return "NO"
				}
			}
			// check
			index1 = index2
			distance = 0
			index2++
		}
		distance++
	}
	return "YES"
}

func main() {

	fmt.Println(isValid2("abccc"))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
