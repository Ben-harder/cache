package doubleLinkedList

type Data interface{}

type Node struct {
	Next *Node
	Prev *Node
	Data Data
}

func New(data Data) *Node {
	return &Node{Data: data}
}

type List struct {
	Head *Node
	Tail *Node
}
