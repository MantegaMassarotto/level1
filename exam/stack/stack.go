package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	Elements []int
}

func (stack *Stack) push(element int) {
	stack.Elements = append(stack.Elements, element)
}

func (stack *Stack) pop() int {
	lastElement := stack.Elements[len(stack.Elements)-1]
	stack.Elements = stack.Elements[:len(stack.Elements)-1]
	return lastElement
}

func (stack *Stack) peek(position int) (int, error) {
	if len(stack.Elements)-1 < position+1 {
		return 0, errors.New("Index not found")
	}
	element := stack.Elements[position+1]
	return element, nil
}

func main() {
	newStack := new(Stack)

	newStack.push(1)
	newStack.push(2)
	newStack.push(3)

	fmt.Println(newStack.peek(2))
}
