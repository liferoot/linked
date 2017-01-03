package linked

type Node struct {
	next, prev *Node

	// The List to which this Node belongs.
	list *List

	// The value stored with this Node.
	Value interface{}
}

func (n *Node) Attach(node *Node) *Node {
	if n == node {
		return n
	}
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
	if n.next != nil {
		n.next.prev = node
	}
	node.next, n.next = n.next, node
	node.prev = n

	if node.list != n.list {
		if node.list != nil {
			node.list.length--
		}
		if n.list != nil {
			n.list.length++
		}
		node.list = n.list
	}
	return node
}

func (n *Node) Detach() *Node {
	if n.prev != nil {
		n.prev.next = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	}
	n.next = nil
	n.prev = nil

	if n.list != nil {
		n.list.length--
		n.list = nil
	}
	return n
}

// List returns the pointer to the List that the Node belongs to.
func (n *Node) List() *List { return n.list }

// Next returns the pointer to the next Node or nil.
func (n *Node) Next() *Node {
	if n.list != nil && n.next == &n.list.root {
		return nil
	}
	return n.next
}

// Prev returns the pointer to the previous Node or nil.
func (n *Node) Prev() *Node {
	if n.list != nil && n.prev == &n.list.root {
		return nil
	}
	return n.prev
}

func NewNode(value interface{}) *Node { return &Node{Value: value} }
