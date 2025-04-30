package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

/*
 * Complete the 'funnyString' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func funnyString(s string) string {
	array := []rune(s)
	arrayDiff := []rune{}
	for i := len(array) - 1; i >= 0; i-- {
		arrayDiff = append(arrayDiff, array[i])
	}
	for i := 0; i < len(array)-2; i++ {
		adjust1 := array[i] - array[i+1]
		if adjust1 < 0 {
			adjust1 *= -1
		}
		adjust2 := arrayDiff[i] - arrayDiff[i+1]
		if adjust2 < 0 {
			adjust2 *= -1
		}
		if adjust1-adjust2 != 0{
			return "Not Funny"
		}
	}
	return "Funny"
	// Write your code here

}

func main() {
	fmt.Println(funnyString("lmnop"))
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
