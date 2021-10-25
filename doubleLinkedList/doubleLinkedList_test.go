package doubleLinkedList

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListOperations(t *testing.T) {
	testList := new(List)
	node1 := &Node{
		Data: "1",
	}
	testList.InsertBeginning(node1)
	assert.Equal(t, 1, testList.Size)

	testList.Remove(node1)
	assert.Zero(t, testList.Size)

	node2 := &Node{
		Data: "2",
	}
	node3 := &Node{
		Data: "3",
	}
	node4 := &Node{
		Data: "4",
	}
	// Testing multiple inserts
	testList.InsertBeginning(node2)
	testList.InsertBeginning(node3)
	testList.InsertBeginning(node4)
	assert.Equal(t, node2, testList.Tail)
	assert.Equal(t, node4, testList.Head)
	assert.Equal(t, 3, testList.Size)

	// Testing removing middle node
	testList.Remove(node3)
	assert.Equal(t, 2, testList.Size)
	assert.Equal(t, node2.Prev, node4)
	assert.Equal(t, node4.Next, node2)
}
