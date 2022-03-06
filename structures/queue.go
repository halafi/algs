package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	Elements []int
}

func (q Queue) Front() int {
	return q.Elements[len(q.Elements) - 1]
}

// Enqueue - adds an element to the front of the queue
func (q *Queue) Enqueue(elem int) {
	q.Elements = append(q.Elements, elem)
}

// Dequeue - removes the first element from a queue
func (q Queue) Dequeue() int {
	removed := q.Elements[0]
	q.Elements = q.Elements[1:]
	return removed
}

// Peek - returns the first element from our queue without updating queue
func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty queue")
	}
	return q.Elements[0], nil
}

// GetLength - returns the length of the queue
func (q Queue) GetLength() int {
	return len(q.Elements)
}

// LastElem - returns the last element of the queue
func (q Queue) IsEmpty() bool {
	return len(q.Elements) == 0
}

func main() {
	fmt.Println("Queues Section")

	queue := Queue{}
	queue.Enqueue(1)
	fmt.Println(queue)
	queue.Enqueue(2)
	fmt.Println(queue)
	elem := queue.Dequeue()
	fmt.Println(elem)
	fmt.Println(queue)

}
