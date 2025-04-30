package main

// English Int: Given any integer, print an English phrase that describes the integer (e.g., "One
// Thousand, Two Hundred Thirty Four").

// SOLUTION
// This is not an especially challenging problem, but it is a somewhat tedious one. The key is to be organized
// in how you approach the problem-and to make sure you have good test cases.
// We can think about converting a number like 19,323,984 as converting each of three 3-digit segments of
// the number, and inserting "thousands" and "millions" in between as appropriate. That is,
// convert(19,323,984) = convert(19) + " million " + convert(323) + " thousand " + convert(984)
// The code below implements this algorithm.

var smalls = []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven",
	"Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen",
	"Sixteen", "Seventeen", "Eighteen", "Nineteen"}

var tens = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy",
	"Eighty", "Ninety"}
var bigs = []string{"", "Thousand", "Million", "Billion"}
var hundred = "Hundred"
var negative = "Negative"

type Node struct {
	next  *Node
	value string
}

type LinkedList struct {
	heat, tail *Node
	size       int
}

func (ll *LinkedList) push(v string) {
	node := &Node{value: v}
	if ll.size == 0 {
		ll.heat = node
		ll.tail = node
		ll.size++
		return
	}
	ll.tail.next = node
	ll.tail = node
	ll.size++
}

func (ll *LinkedList) pop() string {
	node := ll.heat
	ll.heat = ll.heat.next
	ll.size--
	return node.value
}

func (ll *LinkedList) isEmpty() bool {
	if ll.size <= 0 {
		return true
	}
	return false
}
// TODO: needs to be complement
func (ll *LinkedList) toString() string {
	panic("implelet me")
}
// Convert a linked list of strings to a string, dividing it up with spaces.
func listToString(parts *LinkedList) string {
	sb := &LinkedList{}
	for !parts.isEmpty(){
		sb.push(parts.pop())
		sb.push(" ")
	}
	sb.push(parts.pop())
	return sb.toString()  
}

func convertChunk(number int)string{
	parts := &LinkedList{}
	// Convert hundreds place
	if number >=100{
		parts.push(smalls[number/100])
		parts.push(hundred)
		number %=100
	}
	// Convert tens place
	if number>=10&&number<=19{
		parts.push(smalls[number])
	}else if number>=20{
		parts.push(tens[number/10])
		number %=10
	}
	//  Convert ones place
	if number>=1&&number<=9{
		parts.push(smalls[number])
	}
	return listToString(parts)
}

func convert(num int)string{
	if num == 0{
		return smalls[0]
	}else if num<0{
		return negative + " "+convert(-1*num)
	}

	parts := &LinkedList{}
	chunkCount :=0
	for num>0{
		if num%1000!= 0{
			chunk := convertChunk(num%1000)+" "+bigs[chunkCount]
			parts.push(chunk) // TODO it needs to be pushFirst ===> at firts of the queue
		}
		num/=1000 //shift chunk
		chunkCount++
	}
	return listToString(parts)
}

// The key in a problem like this is to make sure you consider all the special cases. There are a lot of them.
// page 489
