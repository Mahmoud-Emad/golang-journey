package main
import "testing"


func TestStack(t *testing.T) {
	// Test 1:
	// Here we want to define a new stack
	stack := NewStack()

	// Push some values onto the stack
	prepareStack(stack)

	actual := stack.Pop().String()
	expected := "9"

	if expected != actual {
		t.Errorf("Test 1: Expected %v, but got %v", expected, actual)
	}
	
}

func prepareStack(stack *Stack) *Stack {
	// Here we want to push some values onto the stack
	stack.Push(&Node{3})
	stack.Push(&Node{5})
	stack.Push(&Node{7})
	stack.Push(&Node{9})

	return stack
}