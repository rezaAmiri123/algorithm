package main

import "fmt"

// Animal Shelter: An animal shelter, which holds only dogs and cats, operates on a strictly"first in, first
// out" basis. People must adopt either the "oldest" (based on arrival time) of all animals at the shelter,
// or they can select whether they would prefer a dog or a cat (and will receive the oldest animal of
// that type). They cannot select which specific animal they would like. Create the data structures to
// maintain this system and implement operations such as enqueue, dequeueAny, dequeueDog,
// and dequeueCat. You may use the built-in Linkedlist data structure.

// SOLUTION
// We could explore a variety of solutions to this problem. For instance, we could maintain a single queue.
// This would make dequeueAny easy, but dequeueDog and dequeueCat would require iteration through
// the queue to find the first dog or cat. This would increase the complexity of the solution and decrease the
// efficiency.
// An alternative approach that is simple, clean and efficient is to simply use separate queues for dogs and
// cats, and to place them within a wrapper class called An imalQueue. We then store some sort of timestamp
// to mark when each animal was enqueued. When we call dequeueAny, we peek at the heads of both the
// dog and cat queue and return the oldest.
type Animal struct {
	name  string
	order int
}

func (a *Animal) isOlderThan(o *Animal) bool {
	return a.order < a.order
}

type Node struct {
	next *Node
	data *Animal
}

type AnimalQueue struct {
	dogs  *Node
	cats  *Node
	order int // acts as timestamp
}

func (q *AnimalQueue) enqueue(a *Animal) {
	// Order is used as a sort of timestamp, so that we can compare the insertion
	// order of a dog to a cat.
	a.order = q.order
	q.order++
	n := &Node{data: a}
	switch a.name {
	case "dogs":
		if q.dogs == nil {
			q.dogs = n
		} else {
			n.next = q.dogs
			q.dogs = n
		}
	case "cats":
		if q.cats == nil {
			q.cats = n
		} else {
			n.next = q.cats
			q.cats = n
		}
	}
}

func (q *AnimalQueue) dequeueAny() *Animal {
	var node *Node
	if q.cats != nil && q.dogs != nil {
		if q.cats.data.isOlderThan(q.dogs.data) {
			node = q.cats
			q.cats = q.cats.next
		} else {
			node = q.dogs
			q.dogs = q.dogs.next
		}
	} else {
		if q.cats != nil {
			node = q.cats
			q.cats = q.cats.next
		} else {
			node = q.dogs
			q.dogs = q.dogs.next
		}
	}
	q.order--
	return node.data
}

func (q *AnimalQueue) isEmpty() bool {
	return q.order <= 0
}
func (q *AnimalQueue) print() {
	cur := q.cats
	fmt.Println("cats")
	for cur != nil {
		fmt.Printf("%d-->", cur.data.order)
		cur = cur.next
	}

	fmt.Println()
	cur = q.dogs
	fmt.Println("dogs")
	for cur != nil {
		fmt.Printf("%d-->", cur.data.order)
		cur = cur.next
	}
	fmt.Println()
}
// It is important that Dog and Cat both inherit from an Animal class since dequeueAny() needs to be able
// to support returning both Dog and Cat objects.
// If we wanted, order could be a true timestamp with the actual date and time. The advantage of this is that
// we wouldn't have to set and maintain the numerical order. If we somehow wound up with two animals with
// the same timestamp, then (by definition) we don't have an older animal and we could return either one

func main() {
	queue := &AnimalQueue{}
	for i := 0; i < 10; i++ {
		a := &Animal{}
		if i%2 == 0 {
			a.name = "dogs"
		} else {
			a.name = "cats"
		}
		queue.enqueue(a)
	}
	queue.print()

	for !queue.isEmpty(){
		a :=queue.dequeueAny()
		fmt.Println(a.name, a.order)
	}
}

