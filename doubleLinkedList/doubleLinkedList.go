package doubleLinkedList

type Data interface{}

type Node struct {
	Next *Node
	Prev *Node
	Data Data
}

type List struct {
	Head *Node
	Tail *Node
	Size int
}

func (l *List) InsertBeginning(newNode *Node) {
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		newNode.Next = nil
		newNode.Prev = nil
		l.Size = 1
	} else {
		l.InsertBefore(l.Head, newNode)
	}
}

func (l *List) InsertEnd(newNode *Node) {
	if l.Tail == nil {
		l.InsertBeginning(newNode)
	} else {
		l.InsertAfter(l.Tail, newNode)
	}
}

func (l *List) InsertAfter(node *Node, newNode *Node) {
	newNode.Prev = node
	if node.Next == nil {
		node.Next = newNode
		l.Tail = newNode
	} else {
		node.Next.Prev = newNode
		newNode.Next = node.Next
	}
	node.Next = newNode
	l.Size += 1
}

func (l *List) InsertBefore(node *Node, newNode *Node) {
	newNode.Next = node
	if node.Prev == nil {
		l.Head = newNode
	} else {
		newNode.Prev = node.Prev
		node.Prev.Next = newNode
	}
	node.Prev = newNode
	l.Size += 1
}

// Assumes the node exists in the list
func (l *List) Remove(node *Node) {
	if node.Prev == nil {
		l.Head = node.Next
	} else {
		node.Prev.Next = node.Next
	}
	if node.Next == nil {
		l.Tail = node.Prev
	} else {
		node.Next.Prev = node.Prev
	}
	l.Size -= 1
}
