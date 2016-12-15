package linked

type List struct {
	root   Node
	length int
}

// Len returns the number of Node's actually contained in the List.
// The complexity is O(1).
func (l *List) Len() int { return l.length }

// First returns the first Node of the List or nil.
func (l *List) First() *Node {
	if l.length == 0 {
		return nil
	}
	return l.root.next
}

// Last returns the last Node of the List or nil.
func (l *List) Last() *Node {
	if l.length == 0 {
		return nil
	}
	return l.root.prev
}

// Add adds new Node's containing the specified valuees at the end of the List,
// returns the last Node added or nil.
func (l *List) Add(value ...interface{}) *Node {
	if len(value) == 0 {
		return nil
	}
	if l.root.next == nil && len(value) > 0 {
		l.Init()
	}
	return l.insert(l.root.prev, value)
}

// AddAfter adds a new Node's containing the specified values after
// the specified existing Node in the List.
func (l *List) AddAfter(mark *Node, value ...interface{}) *Node {
	if l.root.next == nil && len(value) > 0 {
		l.Init()
	}
	if mark == nil {
		if len(value) == 0 {
			return nil
		}
		mark = &l.root
	} else if mark.list != l {
		return nil
	}
	return l.insert(mark, value)
}

// AddBefore adds a new Node's containing the specified values before
// the specified existing Node in the List.
func (l *List) AddBefore(mark *Node, value ...interface{}) *Node {
	if l.root.next == nil && len(value) > 0 {
		l.Init()
	}
	if mark == nil {
		if len(value) == 0 {
			return nil
		}
		mark = &l.root
	} else if mark.list != l {
		return nil
	}
	return l.insert(mark.prev, value)
}

// Pop removes the first Node and returns its value.
func (l *List) Pop() interface{} {
	if l.length == 0 {
		return nil
	}
	l.length--
	return l.root.next.Detach().Value
}

// Push adds a new Node containing specified value at the top of the List.
func (l *List) Push(value interface{}) *Node {
	if l.root.next == nil {
		l.Init()
	}
	node, ok := value.(*Node)
	if ok {
		if node.list != nil && node.list != l {
			node.list.length--
		}
	} else {
		node = &Node{Value: value}
	}
	l.length++
	return l.root.Attach(node)
}

func (l *List) Remove(node ...*Node) (value interface{}) {
	if l.length == 0 {
		return nil
	}
	if len(node) == 0 {
		l.length--
		return l.First().Detach().Value
	}
	for _, n := range node {
		if n == nil || n.list != l {
			continue
		}
		l.length--
		value = n.Detach().Value
	}
	return
}

func (l *List) RemoveLast() interface{} {
	if l.length == 0 {
		return nil
	}
	l.length--
	return l.root.prev.Detach().Value
}

func (l *List) Clear() *List {
	var next *Node

	for n := l.First(); n != nil; n = next {
		next = n.Next()
		n.Detach()
	}
	return l.Init()
}

func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.root.list = l
	l.length = 0
	return l
}

func (l *List) insert(mark *Node, value []interface{}) *Node {
	var (
		node *Node
		ok   bool
		v    interface{}
	)
	for _, v = range value {
		if node, ok = v.(*Node); ok {
			if node == mark {
				continue
			}
			if node.list != mark.list {
				if node.list != nil {
					node.list.length--
				}
				l.length++
			}
		} else {
			node = &Node{Value: v}
			l.length++
		}
		mark = mark.Attach(node)
	}
	return mark
}

func NewList(value ...interface{}) *List {
	l := new(List).Init()
	l.insert(&l.root, value)
	return l
}
