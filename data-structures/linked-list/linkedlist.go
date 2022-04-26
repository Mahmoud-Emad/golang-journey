package main
import (
	"fmt"
)

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
	length int
}


func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func main(){
	mylist 	:= linkedList{}
	node1 	:= &node{data:48}
	node2 	:= &node{data:48}
	node3 	:= &node{data:48}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)

	fmt.Println(mylist)
}