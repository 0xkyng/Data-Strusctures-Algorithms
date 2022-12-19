package main


//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
// SINGLE & DOUBLE LINKED LISTS

// type list struct {
// 	head *Node
// 	tail *Node
// }

// // SINGLE LINKED LISTS
// func (l list) First() *Node {
// 	return l.head
// }

// func (l *list) Push(value int) {
// 	node := &Node{value: value}
// 	if l.head == nil {
// 		l.head = node
// 	} else {
// 		l.tail.next = node
// 	}
// 	l.tail = node

// }

// type Node struct {
// 	value int
// 	next  *Node
// }

// func (n *Node) Next() *Node {
// 	return n.next
// }

// func main() {
// 	l := &list{}
// 	l.Push(1)
// 	l.Push(2)
// 	l.Push(3)

// 	n := l.First()
// 	for {
// 		println(n.value)
// 		n = n.Next()
// 		if n == nil {
// 			break
// 		}
		
// 	}
// }

// DOUBLE LINKED LISTS
// type list struct {
// 	head *Node
// 	tail *Node
// }

// func (l *list) First() *Node {
// 	return l.head
// }

// func (l *list) Last() *Node {
// 	return l.tail
// }


// func (l *list) Push(value int) {
// 	node := &Node{value: value}
// 	if l.head == nil {
// 		l.head = node
// 	} else {
// 		l.tail.next = node
// 		node.prev = l.tail
// 	}
// 	l.tail = node

// }

// type Node struct {
// 	value int
// 	next  *Node
// 	prev *Node
// }

// func (n *Node) Next() *Node {
// 	return n.next
// }

// func (n *Node) Prev() *Node {
// 	return n.prev
// }

// func main() {
// 	l := &list{}
// 	l.Push(1)
// 	l.Push(2)
// 	l.Push(3)

// 	n := l.First()
// 	for {
// 		println(n.value)
// 		n = n.Next()
// 		if n == nil {
// 			break
// 		}	
// 	}

// 	n = l.Last()
// 	for {
// 		println(n.value)
// 		n = n.Prev()
// 		if n == nil {
// 			break
// 		}	
// 	}
// }
//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX


//-------------------------------------------------------------------------------------------------------------
// IMPLEMENTING STACKS
type Stack struct {
	items []string
}

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() string {
	item, items := s.items[len(s.items)-1], s.items[0:len(s.items)-1]
	s.items = items
	return item
}

func main() {
	s := Stack{}
	 s.Push("Read")
	 s.Push("this")
	 s.Push("slowly")


	 println(s.Pop())
	 println(s.Pop())
	 println(s.Pop())
}
//-------------------------------------------------------------------------------------------------------------