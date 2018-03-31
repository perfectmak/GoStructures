package linkedlist

type Node struct {
	Item int
	Next *Node
	Prev *Node
}

type LinkedList struct {
	Head *Node
}

// append item (Node) into list
func (list *LinkedList) Append(item int) {
	if list.Head == nil {
		list.Head = &Node { Item: item, Next: nil, Prev: nil }
	} else {
		current := list.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = &Node { Item: item, Next: nil, Prev: current }
	}
}

// Remove node from linkedlist
func (list *LinkedList) Remove(node *Node) bool {
	current := list.Head

	// special case delete head
	if node == list.Head {
		list.Head = list.Head.Next
		list.Head.Prev = nil
		return true
	}

	// find the item
	for current != nil && current != node {
		current = current.Next
	}

	// delete item if exists
	if current != nil {
		current.Prev.Next = current.Next
		return true
	} else {
		return false
	}
}

// find node for item
func (list *LinkedList) Find(item int) *Node {
	current := list.Head
	for current != nil {
		if current.Item == item {
			return current
		} else {
			current = current.Next
		}
	}

	return nil
}

// traverse list and aggregate into a slice
func (list *LinkedList) ToSlice() []int {
	items := make([]int, 0)
	current := list.Head
	for current != nil {
		items = append(items, current.Item)
		current = current.Next
	}

	return items
}
