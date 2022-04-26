package main
import (
	"testing"
)


func TestLinkedList(t *testing.T) {
	mylist 	:= linkedList{}
	node1 	:= &node{data:48}
	node2 	:= &node{data:48}
	node3 	:= &node{data:48}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)

	if mylist.head.data != 48 {
		t.Errorf("Expected head to be 48, got %d", mylist.head.data)
	}
	if mylist.head.next.data != 48 {
		t.Errorf("Expected head.next to be 48, got %d", mylist.head.next.data)
	}
	if mylist.head.next.next.data != 48 {
		t.Errorf("Expected head.next.next to be 48, got %d", mylist.head.next.next.data)
	}
	if mylist.head.next.next.next != nil {
		t.Errorf("Expected head.next.next.next to be nil, got %d", mylist.head.next.next.next)
	}
}