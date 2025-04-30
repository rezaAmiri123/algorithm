package main
// Living People: Given a list of people with their birth and death years, implement a method to
// compute the year with the most number of people alive. You may assume that all people were born
// between 1900 and 2000 (inclusive). If a person was alive during any portion of that year, they should
// be included in that year's count. For example, Person (birth= 1908, death= 1909) is included in the
// counts for both 1908 and 1909.

// SOLUTION
// The first thing we should do is outline what this solution will look like. The interview question hasn't speci­
// fied the exact form of input. In a real interview, we could ask the interviewer how the input is structured.
// Alternatively, you can explicitly state your (reasonable) assumptions.
// Here, we'll need to make our own assumptions. We will assume that we have an array of simple Person
// objects:

type Person struct{
	birth,death int
}

// We could have also given Person a getBirthVear() and getDeathYear() objects. Some would
// argue that's better style, but for compactness and clarity, we'll just keep the variables public.
// The important thing here is to actually use a Person object. This shows better style than, say, having an
// integer array for birth years and an integer array for death years (with an implicit association of bir ths [ i]
// and deaths [ i] being associated with the same person). You don't get a lot of chances to demonstrate
// great coding style, so it's valuable to take the ones you get.
// With that in mind, let's start with a brute force algorithm.

// Brute Force
// The brute force algorithm falls directly out from the wording of the problem. We need to find the year with
// the most number of people alive. Therefore, we go through each year and check how many people are alive
// in that year.
func MaxAliveYear(people []*Person, max,min int)int{
	maxAlive := 0
	maxAliveYear := min
	for year := min;year<=max;year ++{
		alive := 0
		for _,person := range people{
			if person.birth<=year&&year <=person.death{
				alive++
			}
		}
		if alive>maxAlive{
			maxAlive=alive
			maxAliveYear = year
		}
	}
	return maxAliveYear
}
// Note that we have passed in the values for the min year(l 900) and max year (2000). We shouldn't hard code
// these values.
// The runtime of this is O (RP), where Ris the range of years (100 in this case) and P is the number of people
// to be ontinued at page 494