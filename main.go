package main

type list struct {
	head *Node
	tail *Node
}

// SINGLE LINKED LISTS
func (l list) First() *Node {
	return l.head
}

func (l list) Push(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
	}
	l.tail = node

}

type Node struct {
	value int
	next  *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func main() {
	l := &list{}
	l.Push(1)
	l.Push(2)
	l.Push(3)

	n := l.First()
	for {
		println(n.value)
		n = n.Next()
		if n == nil {
			break
		}
		
	}
}
