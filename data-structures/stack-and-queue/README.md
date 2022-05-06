# Learn Go With Tests

### implementation LIFO Stack and FIFO Queue

* A stack is an ordered list in which all insertions and deletions are made at one end, called the top. A queue is an ordered list in which all insertions take place at one end, the rear, while all deletions take place at the other end, the front.Stacks are sometimes referred to as Last In First Out (LIFO) lists and queues are known as First In First Out (FIFO) lists.


### Example
```go
package main
 
import (
    "fmt"
)
 
type Node struct {
    Value int
}
 
func (n *Node) String() string {
    return fmt.Sprint(n.Value)
}
 
// NewStack returns a new stack.
func NewStack() *Stack {
    return &Stack{}
}
 
// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
    nodes []*Node
    count int
}
 
// Push adds a node to the stack.
func (s *Stack) Push(n *Node) {
    s.nodes = append(s.nodes[:s.count], n)
    s.count++
}
 
// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *Node {
    if s.count == 0 {
        return nil
    }
    s.count--
    return s.nodes[s.count]
}

 

func main() {
	s := NewStack()
    s.Push(&Node{3})
    s.Push(&Node{5})
    s.Push(&Node{7})
    s.Push(&Node{9})
    fmt.Println(s.Pop(), s.Pop(), s.Pop(), s.Pop())
}
```