package main

import "fmt"

type Node struct {
	data interface{}
	prev *Node
	next *Node
}

func (this *Node) String() string {
	if this.next == nil {
		return fmt.Sprintf("%v", this.data)
	} else {
		return fmt.Sprintf("%v %s", this.data, this.next)
	}
}

type LinkedList struct {
	first *Node
	last  *Node
	size  int
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (this *LinkedList) String() string {
	return fmt.Sprintf("LinkedList: [%s]", this.first.String())
}

func (this *LinkedList) Size() int {
	return this.size
}

func (this *LinkedList) IsEmpty() bool {
	return this.Size() == 0
}

func (this *LinkedList) Add(data interface{}) {
	node := &Node{data: data}

	if this.first == nil {
		this.first = node
		this.last = node
	} else {
		this.last.next = node
		node.prev = this.last
		this.last = node
	}

	this.size++
}

func (this *LinkedList) getNode(index int) *Node {
	node := this.first

	for i := 0; i < index; i++ {
		node = node.next
	}

	return node
}

func (this *LinkedList) Get(index int) interface{} {
	return this.getNode(index).data
}

func (this *LinkedList) Remove(index int) interface{} {
	node := this.getNode(index)

	node.prev.next = node.next
	node.next.prev = node.prev
	this.size--

	return node.data
}
