package main

import "testing"

func check(t *testing.T, method string, actual, expected interface{}) {
	if actual != expected {
		t.Errorf("LinkedList.%s: expected: %v, actual: %v", method, expected, actual)
	}

}

func TestSize(t *testing.T) {
	list := NewLinkedList()

	check(t, "Size()", list.Size(), 0)
}

func TestIsEmpty(t *testing.T) {
	list := NewLinkedList()

	check(t, "IsEmpty()", list.IsEmpty(), true)
}

func TestAdd(t *testing.T) {
	list := NewLinkedList()

	list.Add(1)
	list.Add(2)
	check(t, "Size()", list.Size(), 2)
	check(t, "Get(0)", list.Get(0), 1)
	check(t, "Get(1)", list.Get(1), 2)
}

func TestRemove(t *testing.T) {
	list := NewLinkedList()

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Remove(1)

	check(t, "Size()", list.Size(), 2)
	check(t, "Get(0)", list.Get(0), 1)
	check(t, "Get(1)", list.Get(1), 3)
}
