package dsa

import (
	"testing"
)

func Test_singlyLinkedList_search(t *testing.T) {
	ll := &singlyLinkedList{}
	_, err := ll.search(1)
	if err == nil {
		t.Error("expected searching empty list to return error but did not")
	}
	n := &node{value: 1}
	ll.head = n
	ll.append(5)
	ll.append(4)
	ll.append(8)
	ll.append(6)
	_, err = ll.search(8)
	if err != nil {
		t.Errorf("expected successful search not to return error but did: %s", err.Error())
	}
	_, err = ll.search(3)
	if err == nil {
		t.Error("expected searching for value not in list to return error but did not")
	}
}

func Test_singlyLinkedList_insertAfter(t *testing.T) {
	ll := &singlyLinkedList{}
	ll.head = &node{value: 1}
	ll.append(5)
	ll.append(4)
	ll.append(8)
	ll.append(6)
	foundNode, _ := ll.search(8)
	ll.insertAfter(foundNode, 10)
	foundNode, _ = ll.search(10)
	if foundNode.next.value != 6 {
		t.Error("list not in expected order after inserting in middle")
	}
	foundNode, _ = ll.search(8)
	if foundNode.next.value != 10 {
		t.Error("list not in expected order after inserting in middle")
	}
}
