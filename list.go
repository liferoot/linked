package linked

type List struct {
	root   Node
	length int
}

func (l *List) Len() int { return l.length }

func (l *List) First() *Node {
	if l.length == 0 {
		return nil
	}
	return l.root.next
}

func (l *List) Last() *Node {
	if l.length == 0 {
		return nil
	}
	return l.root.prev
}

func (l *List) Add(value ...interface{}) *Node {
	if len(value) == 0 {
		return nil
	}
	return l.insert(l.root.prev, value)
}

func (l *List) InsertAfter(mark *Node, value ...interface{}) *Node {
	if mark == nil && len(value) > 0 {
		mark = &l.root
	}
	return l.insert(mark, value)
}

func (l *List) InsertBefore(mark *Node, value ...interface{}) *Node {
	if mark == nil && len(value) > 0 {
		mark = &l.root
	}
	return l.insert(mark.prev, value)
}

func (l *List) Pop() interface{} {
	if l.length == 0 {
		return nil
	}
	l.length--
	return l.root.next.Detach().Value
}

func (l *List) Push(value interface{}) *Node {
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
