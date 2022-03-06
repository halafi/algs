// First Go program
package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	top        int
	capacity   int
	stackArray []interface{}
}

//init intialises a new stack object and returns a pointer for the same
func (stack *Stack) init(capacity int) *Stack {
	stack.top = -1
	stack.capacity = capacity
	stack.stackArray = make([]interface{}, capacity)

	return stack
}

//NewStack creates a new stack object
func NewStack(capacity int) *Stack {
	return new(Stack).init(capacity)
}

//IsFull checks whether the stack has reached it's capacity or not
func (stack *Stack) IsFull() bool {
	return stack.top == int(stack.capacity)-1
}

//IsEmpty checks whether the stack is empty
func (stack *Stack) IsEmpty() bool {
	return stack.top == -1
}

//Size returns the current size of the Stack
func (stack *Stack) Size() uint {
	return uint(stack.top + 1)
}

//Push inserts the data in the top of the stack
func (stack *Stack) Push(data interface{}) error {
	if stack.IsFull() {
		return errors.New("stack overflow")
	}
	stack.top++
	stack.stackArray[stack.top] = data
	return nil
}

//Pop removes the data from the top of the stack
func (stack *Stack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack underflow")
	}
	temp := stack.stackArray[stack.top]
	stack.top--
	return temp, nil
}

//Peek returns the data without removing from the top of stack
func (stack *Stack) Peek() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack underflow")
	}
	temp := stack.stackArray[stack.top]
	return temp, nil
}

//Print prints all the stack element
func (stack *Stack) Print() {
	fmt.Print("[")
	for i := 0; i <= stack.top; i++ {
		fmt.Print(stack.stackArray[i])
		if i != stack.top {
			fmt.Print(",")
		}
	}
	fmt.Print("]\n")
}

func main() {
	stack := NewStack(10)
	stack.Peek()
	stack.Push(1)
	stack.Print()
	stack.Push(2)
	stack.Print()
	stack.Push(3)
	stack.Print()
	stack.Pop()
	stack.Print()
	fmt.Println("Size := ", stack.Size())
}
